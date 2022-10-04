package main

import (
	"encoding/json"
	"fmt"
)

func navFilter() string {
	return "{\"dimension\":\"department\",\"node_id\":\"7130074483159090732\",\"is_strand_id_in\":true}"
}

type request struct {
	Dimension    string `json:"dimension"`
	NodeId       string `json:"node_id"`
	IsStrandIdIn bool   `json:"is_strand_id_in"`
}

func myEnumOption(req string) {
	nav := request{}

	err := json.Unmarshal([]byte(req), &nav)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(nav)
}

func main() {
	myEnumOption(navFilter())
}
