package ramsom_test

import (
	"testing"

	"github.com/GLUD/RansomwareGLUD"
)

func TestRecorrerArchivos(t *testing.T) {
	// carpeta en el directorio actual
	basePath := "TreeDirExample/"

	c := ramsom.RecorrerArchivos(basePath+"f1", basePath+"f 3")
	expectedFiles := map[string]bool{
		basePath + "f1/b1/ejem.doc":   true,
		basePath + "f1/data.doc":      true,
		basePath + "f 3/test.txt":     true,
		basePath + "f 3/test.bin":     true,
		basePath + "f 3/a1/test2.bin": true,
	}
	for file := range c {
		if ok := expectedFiles[file.Name()]; !ok {
			t.Fatalf("%s is not valid", file.Name())
		}
		t.Log(file.Name())
	}

}
