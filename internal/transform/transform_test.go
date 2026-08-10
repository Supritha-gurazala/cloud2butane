package transform

import (
	"testing"

	"github.com/Supritha-gurazala/cloud2butane/internal/cloudconfig"
)

func TestToButane(t *testing.T) {
	in := cloudconfig.Config{
		WriteFiles: []cloudconfig.WriteFile{
			{
				Path:    "/etc/example.conf",
				Content: "hello world",
			},
		},
	}

	out, err := ToButane(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out.Variant != "flatcar" {
		t.Fatalf("expected variant flatcar, got %q", out.Variant)
	}

	if out.Version != "1.0.0" {
		t.Fatalf("expected version 1.0.0, got %q", out.Version)
	}

	if len(out.Storage.Files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(out.Storage.Files))
	}

	file := out.Storage.Files[0]

	if file.Path != "/etc/example.conf" {
		t.Fatalf("unexpected path: %q", file.Path)
	}

	if file.Contents.Inline != "hello world" {
		t.Fatalf("unexpected content: %q", file.Contents.Inline)
	}
}

func TestToButaneEmptyWriteFiles(t *testing.T) {
	in := cloudconfig.Config{}

	_, err := ToButane(in)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestToButaneMissingPath(t *testing.T) {
	in := cloudconfig.Config{
		WriteFiles: []cloudconfig.WriteFile{
			{
				Content: "hello",
			},
		},
	}

	_, err := ToButane(in)
	if err == nil {
		t.Fatal("expected an error for missing path, got nil")
	}
}

func TestToButaneMissingContent(t *testing.T) {
	in := cloudconfig.Config{
		WriteFiles: []cloudconfig.WriteFile{
			{
				Path: "/etc/example.conf",
			},
		},
	}

	_, err := ToButane(in)
	if err == nil {
		t.Fatal("expected an error for missing content, got nil")
	}
}
