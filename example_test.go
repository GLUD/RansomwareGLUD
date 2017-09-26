package ramsom_test

import (
	"fmt"

	"github.com/GLUD/RansomwareGLUD"
)

func ExampleRecorrerArchivos() {
	c := ramsom.RecorrerArchivos("./foo", "./bar", "/usr", `C:\Windows`)
	for f := range c {
		fmt.Println(f)
	}
}
