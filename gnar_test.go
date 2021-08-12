package gnar_test

import (
	"os"
	"testing"

	"github.com/gnames/gnar"
	"github.com/gnames/gnsys"
	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	testDir := "testdata/test"
	tests := []struct {
		msg, path, file string
	}{
		{"tar.gz", "testdata/test.tar.gz", "testdata/test/a.txt"},
		{"tar.xz", "testdata/test.tar.xz", "testdata/test/ab/a.txt"},
		{"zip nodir", "testdata/test-0l.zip", "testdata/test/a.txt"},
		{"zip 1l dir", "testdata/test-1l.zip", "testdata/test/test-1l/a.txt"},
		{"zip 2l dir", "testdata/test-2l.zip", "testdata/test/test-2l/ab/a.txt"},
	}
	for _, v := range tests {
		assert.False(t, gnsys.IsDir(testDir))
		ar, err := gnar.NewArchive(v.path)
		assert.Nil(t, err)
		g := gnar.New(ar)
		err = g.Unpack(testDir)
		assert.Nil(t, err)
		assert.True(t, gnsys.IsDir(testDir))
		assert.True(t, gnsys.IsFile(v.file), v.msg)
		err = os.RemoveAll(testDir)
		assert.Nil(t, err)
	}
}

func TestPack(t *testing.T) {
	tests := []struct {
		msg, dir, arch string
	}{
		{"tar.gz 1l", "testdata/test-1l", "testdata/test1l.tar.gz"},
		{"tar.xz 1l", "testdata/test-1l", "testdata/test1l.tar.xz"},
		{"tar.gz 2l", "testdata/test-2l", "testdata/tests2l.tar.gz"},
		{"tar.xz 2l", "testdata/test-2l", "testdata/tests2l.tar.xz"},
		{"zip 1l", "testdata/test-1l", "testdata/tests1l.zip"},
		{"zip 2l", "testdata/test-2l", "testdata/tests2l.zip"},
	}

	for _, v := range tests {
		assert.True(t, gnsys.IsDir(v.dir))
		ar, err := gnar.NewArchive(v.arch)
		assert.Nil(t, err)
		g := gnar.New(ar)
		err = g.Pack(v.dir)
		assert.Nil(t, err)
		assert.True(t, gnsys.IsFile(v.arch))
		err = os.RemoveAll(v.arch)
		assert.Nil(t, err)
	}
}
