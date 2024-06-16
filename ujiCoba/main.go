package main

import (
	"encoding/json"
	"fmt"
	"log"
	payload "ujicoba/payload"
	tool "ujicoba/tool"
)

func main() {
	// Convert JSON data string to byte slice
	dataBytes := []byte(payload.DataString1)

	// Convert JSON data to list of maps
	items := tool.ConvertByteToListMap(dataBytes)

	// Build hierarchy of nodes
	hierarchy := tool.BuildHierarchy(items)

	// Convert hierarchy to JSON
	hierarchyJSON, err := json.MarshalIndent(hierarchy, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling hierarchy to JSON: %v", err)
	}

	// Print JSON representation of hierarchy
	fmt.Println("Hierarchy1 JSON:")
	fmt.Println(string(hierarchyJSON))

}
