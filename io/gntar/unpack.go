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

func (t gntar) Unpack(destPath string) error {
	var err error
	if !gnsys.IsFile(t.path) {
		return fmt.Errorf("no such file '%s'", t.path)
	}

	f, err := os.Open(t.path)
	if err != nil {
		return err
	}
	defer f.Close()

	reader, err := t.uncompress(f)
	if err != nil {
		return err
	}

	return t.untar(reader, destPath)
}

func (t gntar) uncompress(f io.Reader) (io.Reader, error) {
	switch t.ext {
	case "gz":
		return gzip.NewReader(f)
	case "xz":
		return xz.NewReader(f)
	}
	return nil, fmt.Errorf("unknown compression '%s'", t.ext)
}

func (t gntar) untar(z io.Reader, destPath string) error {
	err := gnsys.MakeDir(destPath)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(z)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.Mkdir(filepath.Join(destPath, header.Name), 0755)
			if err != nil {
				perr := err.(*os.PathError)
				if perr.Err.Error() != "file exists" {
					return err
				}
			}
		case tar.TypeReg:
			outFile, err := os.Create(filepath.Join(destPath, header.Name))
			if err != nil {
				return err
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		default:
			return fmt.Errorf(
				"ExtractTarGz: uknown type: '%v' in '%s'",
				header.Typeflag,
				header.Name,
			)
		}
	}

	return nil
}
