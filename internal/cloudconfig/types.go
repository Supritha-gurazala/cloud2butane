package cloudconfig

// for this project ignoring unknown fields in the YAML file.

// Config represents the whole cloud config YAML structure
type Config struct {
	WriteFiles []WriteFile `yaml:"write_files"`
}

// WriteFile represents one entry inside the write_files list
type WriteFile struct {
	Path    string `yaml:"path"`
	Content string `yaml:"content"`
}
