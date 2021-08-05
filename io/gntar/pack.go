package gntar

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gnames/gnsys"
	"github.com/ulikunitz/xz"
)

func (t gntar) Pack(srcPath string) error {
	var err error
	dir, _ := filepath.Split(srcPath)
	if !gnsys.IsDir(dir) {
		err = gnsys.MakeDir(dir)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(t.path)
	if err != nil {
		return err
	}
	defer f.Close()

	writer, err := t.compress(f)
	if err != nil {
		return err
	}
	defer writer.Close()

	return t.pack(writer, srcPath)
}

func (t gntar) compress(f io.Writer) (io.WriteCloser, error) {
	switch t.ext {
	case "gz":
		return gzip.NewWriter(f), nil
	case "xz":
		return xz.NewWriter(f)
	}

	return nil, fmt.Errorf("unknown compression '%s'", t.ext)
}

func (t gntar) pack(w io.WriteCloser, srcPath string) error {
	defer w.Close()

	tw := tar.NewWriter(w)
	defer tw.Close()

	err := filepath.Walk(srcPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if gnsys.IsDir(path) {
				return nil
			}

			fmt.Println(path, info.Size())
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			info, err = file.Stat()
			if err != nil {
				return err
			}

			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}
			header.Name = path

			err = tw.WriteHeader(header)
			if err != nil {
				return err
			}

			_, err = io.Copy(tw, file)
			return err
		})

	return err
}
