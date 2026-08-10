package transform

import (
	"fmt"

	"github.com/Supritha-gurazala/cloud2butane/internal/butane"
	"github.com/Supritha-gurazala/cloud2butane/internal/cloudconfig"
)

// ToButane converts a cloud-config to a butane config.
func ToButane(cfg cloudconfig.Config) (butane.Config, error) {
	out := butane.Config{
		Variant: "flatcar",
		Version: "1.0.0",
	}
	//write_files should not be empty
	if len(cfg.WriteFiles) == 0 {
		return butane.Config{}, fmt.Errorf("write_files is empty")
	}

	//map each cloud-config file to a butane file.
	for i, wf := range cfg.WriteFiles {
		if wf.Path == "" {
			return butane.Config{}, fmt.Errorf("write_files[%d]: missing path", i)
		}

		if wf.Content == "" {
			return butane.Config{}, fmt.Errorf("write_files[%d]: missing content", i)
		}

		bf := butane.File{
			Path: wf.Path,
			Contents: butane.Contents{
				Inline: wf.Content,
			},
		}
		out.Storage.Files = append(out.Storage.Files, bf)
	}
	return out, nil

}
