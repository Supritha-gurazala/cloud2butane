package butane

// Config represents the top-level structure of a Butane config file
type Config struct {
	Variant string  `yaml:"variant"`
	Version string  `yaml:"version"`
	Storage Storage `yaml:"storage"`
}

// Storage represents the storage section of butane.
type Storage struct {
	Files []File `yaml:"files"`
}

// File represents one entry in storage.files.
type File struct {
	Path     string   `yaml:"path"`
	Contents Contents `yaml:"contents"`
}

// Contents represents the contents of a file in storage.files.
type Contents struct {
	Inline string `yaml:"inline"`
}
