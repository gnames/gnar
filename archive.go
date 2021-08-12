package gnar

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gnames/gnar/ent/archive"
	"github.com/gnames/gnar/io/gntar"
	"github.com/gnames/gnar/io/gnzip"
)

func NewArchive(path string) (archive.Archive, error) {
	var err error
	f := filepath.Base(path)
	parts := strings.Split(f, ".")
	l := len(parts)

	switch l {
	case 0, 1:
		err = fmt.Errorf("no extention in '%s'", f)
		return nil, err
	case 2:
		return getArch(path, parts[l-1])
	default:
		if parts[l-2] == "tar" {
			return gntar.New(path, parts[l-1])
		}
	}

	err = fmt.Errorf("unknown archive '%s'", f)
	return nil, err
}

func getArch(path, ext string) (archive.Archive, error) {
	switch ext {
	case "zip":
		return gnzip.New(path, ext)
	}

	err := fmt.Errorf("unknown archive '%s'", filepath.Base(path))
	return nil, err
}
