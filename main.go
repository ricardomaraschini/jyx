package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"sigs.k8s.io/yaml"
)

func main() {
	src := os.Stdin
	if len(os.Args) != 1 {
		var err error
		if src, err = os.Open(os.Args[1]); err != nil {
			fmt.Println("failed to open file:", err)
			os.Exit(1)
		}
		defer src.Close()
	}

	data, err := io.ReadAll(src)
	if err != nil {
		fmt.Println("failed to read data:", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &map[string]interface{}{}); err == nil {
		result, err := yaml.JSONToYAML(data)
		if err != nil {
			fmt.Println("failed to convert JSON to YAML:", err)
			os.Exit(1)
		}
		fmt.Println(string(result))
		os.Exit(0)
	}

	if err := yaml.Unmarshal(data, &map[string]interface{}{}); err != nil {
		fmt.Println("failed to validate input type:", err)
		os.Exit(1)
	}

	result, err := yaml.YAMLToJSON(data)
	if err != nil {
		fmt.Println("failed to convert YAML to JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(result))
}
