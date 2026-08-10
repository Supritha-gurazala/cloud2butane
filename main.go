package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Supritha-gurazala/cloud2butane/internal/parser"
	"github.com/Supritha-gurazala/cloud2butane/internal/transform"
)

func main() {
	// Expect exactly one argument: the input YAML file.
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: cloud2butane <input.yaml>")
		os.Exit(1)
	}

	inputPath := os.Args[1]

	// Read the input file.
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse cloud-config YAML.
	cfg, err := parser.Parse(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid YAML: %v\n", err)
		os.Exit(1)
	}

	// Transform to Butane.
	out, err := transform.ToButane(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "transform error: %v\n", err)
		os.Exit(1)
	}

	// Marshal the Butane config back to YAML.
	yamlData, err := yaml.Marshal(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "output error: %v\n", err)
		os.Exit(1)
	}

	// Print the generated YAML.
	fmt.Print(string(yamlData))
}
