package app

type Graph struct {
	nodes map[string]*Node
	edges []*Edge
}

type Node struct {
	name string
}

type Edge struct {
	from *Node
	to   *Node
}

type JSONEdge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type JSONGraph struct {
	Nodes []string   `json:"nodes"`
	Edges []JSONEdge `json:"edges"`
}

type BaseKustomization struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

type Patch struct {
	Path   string `yaml:"path"`
	Target struct {
		Group string `yaml:"group"`
		Kind  string `yaml:"kind"`
	} `yaml:"target"`
	InlinePatch string `yaml:"patch"`
}

type KustomizationConfig struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Resources  []string `yaml:"resources"`
	Components []string `yaml:"components"`
	Patches    []Patch  `yaml:"patches"`
}

type KustomizationToolkitSpec struct {
	Components []string `yaml:"components"`
	Path       string   `yaml:"path"`
	Patches    []Patch  `yaml:"patches"`
	PostBuild  struct {
		Substitute struct {
			ApplicationName      string `yaml:"application_name"`
			KustomizationOverlay string `yaml:"kustomization_overlay"`
		} `yaml:"substitute"`
	} `yaml:"postBuild"`
}

type KustomizationToolkit struct {
	APIVersion string                   `yaml:"apiVersion"`
	Kind       string                   `yaml:"kind"`
	Spec       KustomizationToolkitSpec `yaml:"spec"`
}
