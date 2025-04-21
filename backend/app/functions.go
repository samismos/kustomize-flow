package app

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"os"
)

var application_name string
var kustomization_overlay string

func TraverseKustomizations(basePath string, workingDirectory string, graph *Graph) error {
	// Check if the supplied workingDirectory exists
	if application_name != "" && kustomization_overlay != "" {
		workingDirectory = strings.Replace(workingDirectory, "${application_name}", application_name, -1)
	workingDirectory = strings.Replace(workingDirectory, "${kustomization_overlay}", kustomization_overlay, -1)
	}
	
	fmt.Println(Blue + workingDirectory + Reset)
	if _, err := os.Stat(workingDirectory); os.IsNotExist(err) {
        return fmt.Errorf("working directory %s does not exist", workingDirectory)
    }

	kustomization := ReadAndPrintKustomization(workingDirectory)
	currentWorkingDirectory := workingDirectory
	workingDirectory = filepath.Dir(workingDirectory)

	switch k := kustomization.(type) {
	case KustomizationConfig:
		for _, resource := range k.Resources {
			// Traverse resources
			workingDirectory := filepath.Join(workingDirectory, resource)
			newWorkingDirectory := SearchForKustomizationInPath(workingDirectory)
			if newWorkingDirectory != workingDirectory {
				graph.AddEdge(currentWorkingDirectory, newWorkingDirectory)
			}
			TraverseKustomizations(basePath, newWorkingDirectory, graph)
		}
		for _, component := range k.Components {
			// Traverse components
			workingDirectory := filepath.Join(workingDirectory, component)
			newWorkingDirectory := SearchForKustomizationInPath(workingDirectory)
			if newWorkingDirectory != workingDirectory {
				graph.AddEdge(currentWorkingDirectory, newWorkingDirectory)
			}
			TraverseKustomizations(basePath, newWorkingDirectory, graph)
		}
	case KustomizationToolkit:
		for _, patch := range k.Spec.Patches {
			// Traverse patch.path if it exists
			if len(strings.TrimSpace(patch.Path)) != 0 {
				// fmt.Println("Traversing patch path: ", patch.Path)
				workingDirectory := filepath.Join(workingDirectory, patch.Path)
				newWorkingDirectory := SearchForKustomizationInPath(workingDirectory)
				if newWorkingDirectory != workingDirectory {
					graph.AddEdge(currentWorkingDirectory, newWorkingDirectory)
				}
				TraverseKustomizations(basePath, newWorkingDirectory, graph)
			}
		}
		// Traverse k.Spec.Path
		// Paths under spec.path are relative to the base of the flux-repo
		if k.Spec.Path != "" {
			workingDirectory := filepath.Join(basePath, k.Spec.Path)
			newWorkingDirectory := SearchForKustomizationInPath(workingDirectory)
			if newWorkingDirectory != workingDirectory {
				graph.AddEdge(currentWorkingDirectory, newWorkingDirectory)
			}
			TraverseKustomizations(basePath, newWorkingDirectory, graph)
		}
	case BaseKustomization:
		// Traverse resources
		workingDirectory := filepath.Join(workingDirectory)
		graph.AddEdge(workingDirectory, currentWorkingDirectory)
	}
	return nil
}

func SearchForKustomizationInPath(path string) string {
	// Searches for kustomization.yaml if provided path is a directory
	if strings.HasSuffix(path, ".yaml") {
		return path
	}
	return filepath.Join(path, "kustomization.yaml")
}

func IsFluxVariableValid(v string) bool {
	return (!strings.Contains(v, "${") && len(strings.TrimSpace(v)) != 0)
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string]*Node), edges: []*Edge{}}
}

func (g *Graph) AddNode(name string) *Node {
	if _, exists := g.nodes[name]; !exists {
		g.nodes[name] = &Node{name: name}
	}
	return g.nodes[name]
}

func (g *Graph) AddEdge(from, to string) {
	fromNode := g.AddNode(from)
	toNode := g.AddNode(to)
	edge := &Edge{from: fromNode, to: toNode}
	g.edges = append(g.edges, edge)
}

func (g *Graph) PrintGraph() {
	for i, edge := range g.edges {
		fmt.Printf("Edge %d: %s --> %s\n", i, edge.from.name, edge.to.name)
	}
}

func GraphToJSON(g *Graph) (string, error) {
	var jsonGraph JSONGraph

	for _, node := range g.nodes {
		jsonGraph.Nodes = append(jsonGraph.Nodes, node.name)
	}

	for _, edge := range g.edges {
		jsonGraph.Edges = append(jsonGraph.Edges, JSONEdge{From: edge.from.name, To: edge.to.name})
	}

	jsonData, err := json.Marshal(jsonGraph)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func GetBasePathFromEntrypoint(entrypoint string) string {
	// Split the entrypoint into parts
	parts := strings.Split(entrypoint, "/")

	// Initialize basePath
	basePath := ""

	// Iterate over the parts
	for _, part := range parts {
		// Add the part to basePath
		basePath = basePath + part + "/"
		// Check if the part contains 'deployment'
		if strings.Contains(part, "deployment") {
			break
		}
	}

	// Remove the trailing slash
	basePath = strings.TrimSuffix(basePath, "/")

	return basePath
}
