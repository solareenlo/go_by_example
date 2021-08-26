package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filepath")
	fmt.Println("p:", p)
	fmt.Println()

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	fmt.Println()

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	fmt.Println()

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	fmt.Println()

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext))
	fmt.Println()

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	fmt.Println()

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
