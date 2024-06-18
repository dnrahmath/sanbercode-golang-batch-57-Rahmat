package tool

import (
	"encoding/json"
	"log"
	"sort"
	"strconv"
)

// Node represents a node in the hierarchy
type Node struct {
	UuId     string      `json:"uuid"`
	Data     interface{} `json:"data,omitempty"`
	Conf     interface{} `json:"conf,omitempty"`
	Children []*Node     `json:"children,omitempty"`
}

// ConvertByteToListMap converts JSON byte data to a list of maps
func ConvertByteToListMap(data []byte) []map[string]interface{} {
	var items []map[string]interface{}

	err := json.Unmarshal(data, &items)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return nil
	}

	return items
}

// BuildHierarchy builds a hierarchy of nodes based on input items
func BuildHierarchy(items []map[string]interface{}) []*Node {
	nodeMap := make(map[string]*Node)
	var rootNodes []*Node

	// Step 1: Create nodes for each item and store them in the node map
	for _, item := range items {
		uuid, ok := item["uuid"].(string)
		if !ok {
			log.Println("Missing or invalid uuid in item:", item)
			continue
		}

		// Parse conf data dynamically
		conf := item["conf"]

		// Check if conf is a map[string]interface{} type
		if confMap, ok := conf.(map[string]interface{}); ok {
			if idxStr, ok := confMap["index"].(string); ok {
				_, err := strconv.Atoi(idxStr)
				if err != nil {
					log.Println("Invalid index value:", idxStr, err)
				}
			}
		}

		// Create a new node
		node := &Node{
			UuId:     uuid,
			Data:     item["data"],
			Conf:     conf,
			Children: []*Node{},
		}

		// Store node in map
		nodeMap[uuid] = node
	}

	// Step 2: Link child nodes to their parent nodes
	for _, node := range nodeMap {
		confMap, ok := node.Conf.(map[string]interface{})
		if !ok {
			log.Println("Invalid conf format for node:", node.UuId)
			continue
		}

		parentUUID := confMap["childofuuid"].(string)

		if parentUUID == node.UuId {
			// Skip if node is its own parent (to prevent circular reference)
			rootNodes = append(rootNodes, node)
		} else if parentNode, exists := nodeMap[parentUUID]; exists {
			parentNode.Children = append(parentNode.Children, node)
		} else {
			// If no parent node exists (root node), add to rootNodes
			rootNodes = append(rootNodes, node)
		}
	}

	// Sort children of each parent node based on index
	for _, node := range nodeMap {
		sort.Slice(node.Children, func(i, j int) bool {
			return getIndex(node.Children[i]) < getIndex(node.Children[j])
		})
	}

	// Sort root nodes based on index
	sort.Slice(rootNodes, func(i, j int) bool {
		return getIndex(rootNodes[i]) < getIndex(rootNodes[j])
	})

	return rootNodes
}

// getIndex extracts the index from node's Conf if possible, or returns 0
func getIndex(node *Node) int {
	if confMap, ok := node.Conf.(map[string]interface{}); ok {
		if idxStr, ok := confMap["index"].(string); ok {
			idx, err := strconv.Atoi(idxStr)
			if err == nil {
				return idx
			}
		}
	}
	return 0
}
