package tool

import (
	"encoding/json"
	"fmt"
	pl "ujicoba/payload"
)

// Node represents a node in the hierarchy
type Node struct {
	UuId     string  `json:"uuid"`
	Name     string  `json:"name,omitempty"`
	Children []*Node `json:"children,omitempty"`
}

func BuildHierarchy(data []byte, dataIner interface{}) map[string]*Node {
	var items []pl.Data1 // or Data2, depending on input

	err := json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	nodeMap := make(map[string]*Node)

	for _, item := range items {
		// Check if node already exists in nodeMap
		if _, exists := nodeMap[item.UuId]; !exists {
			node := &Node{
				UuId: item.UuId,
				Name: item.Name,
			}
			nodeMap[item.UuId] = node
		}

		// Check if child node already exists in nodeMap
		if item.Childofuuid != "" {
			if _, exists := nodeMap[item.Childofuuid]; !exists {
				childNode := &Node{
					UuId: item.Childofuuid,
				}
				nodeMap[item.Childofuuid] = childNode
			}
		}
	}

	// Link children
	for _, item := range items {
		if item.Childofuuid != "" {
			parentNode, parentExists := nodeMap[item.Childofuuid]
			childNode, childExists := nodeMap[item.UuId]

			if parentExists && childExists {
				parentNode.Children = append(parentNode.Children, childNode)
			}
		}
	}

	// Remove nodes without valid parents from nodeMap
	for _, item := range items {
		if item.Childofuuid == "" {
			delete(nodeMap, item.UuId)
		}
	}

	// Remove nodes that have no parent
	for uuid, node := range nodeMap {
		if node.Children == nil {
			delete(nodeMap, uuid)
		}
	}

	return nodeMap
}
