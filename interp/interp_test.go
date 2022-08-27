package interp

import (
	"fmt"
	"testing"

	"github.com/linkxzhou/gongx/interp"
	"github.com/linkxzhou/gongx/loader"
)

func TestNewInterp(t *testing.T) {
	sources := `
	package test

	func fib(i int) int {
		if i < 2 {
			return i
		}
		return fib(i - 1) + fib(i - 2)
	}
	func test(i int) int {
		return fib(i)
	}
	
	func main() {
		test(37)
	}`
	c := loader.NewContext(loader.EnableDumpImports)
	p, err := c.LoadFile("__main__.go", sources)
	fmt.Println("p: ", p, ", err: ", err)
	iv1, err1 := interp.NewInterp(c, p)
	fmt.Println("NewInterp err: ", err1, ", iv1: ", iv1)
	iv2, err2 := iv1.RunMain()
	fmt.Println("RunMain err: ", err2, ", iv2: ", iv2)
}