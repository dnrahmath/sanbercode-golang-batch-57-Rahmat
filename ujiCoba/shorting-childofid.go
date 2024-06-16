package main

import (
	"encoding/json"
	"fmt"
)

type Data1 struct {
	UuId        string `form:"uuid" json:"uuid" bson:"uuid"`
	Name        string `form:"kodename" json:"kodename" bson:"kodename"`
	Childofuuid string `form:"childofuuid" json:"childofuuid" bson:"childofuuid"`
}

type Data2 struct {
	UuId        string `form:"uuid" json:"uuid" bson:"uuid"`
	Kodename    string `form:"kodename" json:"kodename" bson:"kodename"`
	Childofuuid string `form:"childofuuid" json:"childofuuid" bson:"childofuuid"`
}

// Node represents a node in the hierarchy
type Node struct {
	UuId     string
	Children []*Node
}

func buildHierarchy(data []byte) map[string]*Node {
	var items []Data1 // or Data2, depending on input

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

func main() {
	data1 := `[{"uuid": "2", "name": "satu", "childofuuid": "0"},
	          {"uuid": "5", "name": "dua", "childofuuid": "0"},
	          {"uuid": "3", "name": "agus", "childofuuid": "2"},
	          {"uuid": "4", "name": "sutis", "childofuuid": "2"},
	          {"uuid": "6", "name": "sutarman", "childofuuid": "5"},
	          {"uuid": "7", "name": "sukija", "childofuuid": "5"},
	          {"uuid": "8", "name": "sukiyam", "childofuuid": "3"},
	          {"uuid": "9", "name": "amir", "childofuuid": "3"}]`

	data2 := `[{"uuid": "2", "kodename": "satu", "childofuuid": "0"},
	          {"uuid": "5", "kodename": "dua", "childofuuid": "0"},
	          {"uuid": "3", "kodename": "agus", "childofuuid": "2"},
	          {"uuid": "4", "kodename": "sutis", "childofuuid": "2"},
	          {"uuid": "6", "kodename": "sutarman", "childofuuid": "5"},
	          {"uuid": "7", "kodename": "sukija", "childofuuid": "5"},
	          {"uuid": "8", "kodename": "sukiyam", "childofuuid": "3"},
	          {"uuid": "9", "kodename": "amir", "childofuuid": "3"}]`

	// Panggil buildHierarchy dengan data1
	fmt.Println("Hierarchy based on data1:")
	hierarchy1 := buildHierarchy([]byte(data1))

	// Panggil buildHierarchy dengan data2
	fmt.Println("\nHierarchy based on data2:")
	hierarchy2 := buildHierarchy([]byte(data2))

	// Mencetak hasil hierarki dalam format JSON yang terstruktur
	hierarchyJSON1, err := json.MarshalIndent(hierarchy1, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling hierarchy1 to JSON:", err)
	} else {
		fmt.Println("\nHierarchy1 JSON:")
		fmt.Println(string(hierarchyJSON1))
	}

	hierarchyJSON2, err := json.MarshalIndent(hierarchy2, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling hierarchy2 to JSON:", err)
	} else {
		fmt.Println("\nHierarchy2 JSON:")
		fmt.Println(string(hierarchyJSON2))
	}
}
