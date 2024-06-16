package main

import (
	"encoding/json"
	"fmt"
	tools "ujicoba/tool"
)

func main() {
	dataMap1 := []map[string]interface{}{
		{
			"uuId":        "2",
			"name":        "agus",
			"childofuuid": "0",
		},
		{
			"uuId":        "3",
			"name":        "agus",
			"childofuuid": "0",
		},
	}
	dataMap2 := []map[string]interface{}{
		{
			"uuId":        "2",
			"name":        "agus",
			"childofuuid": "0",
		},
		{
			"uuId":        "3",
			"name":        "agus",
			"childofuuid": "0",
		},
	}

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
	hierarchy1 := tools.BuildHierarchy([]byte(data1), dataMap1)

	// Panggil buildHierarchy dengan data2
	fmt.Println("\nHierarchy based on data2:")
	hierarchy2 := tools.BuildHierarchy([]byte(data2), dataMap2)

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
