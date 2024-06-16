package tool

import (
	"encoding/json"
	"log"
)

// Node represents a node in the hierarchy
type Node struct {
	UuId        string      `json:"uuid"`
	Data        interface{} `json:"data,omitempty"`
	ChildOfUuid string      `json:"childofuuid"`
	Children    []*Node     `json:"children,omitempty"`
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

		// Create a new node
		node := &Node{
			UuId:        uuid,
			Data:        item["data"],
			ChildOfUuid: item["childofuuid"].(string), // assuming childofuuid is always present
			Children:    []*Node{},
		}

		// Store node in map
		nodeMap[uuid] = node
	}

	// Step 2: Link child nodes to their parent nodes
	for _, node := range nodeMap {
		parentUUID := node.ChildOfUuid
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

	return rootNodes
}
