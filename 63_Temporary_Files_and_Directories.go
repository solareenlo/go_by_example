package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := ioutil.TempFile("", "sample")
	check(err)
	fmt.Println("Temp file name:", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3})
	check(err)

	dName, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name: ", dName)
	defer os.Remove(dName)

	fName := filepath.Join(dName, "file1")
	err = ioutil.WriteFile(fName, []byte{1, 2}, 0666)
	check(err)
}
