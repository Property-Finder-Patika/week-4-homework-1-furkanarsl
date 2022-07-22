package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

var rootdir = flag.String("d", ".", "root directory to walk and list files")

// WalkFunc implementation for WalkDir
func walkFunc(p string, _ fs.DirEntry, err error) error {
	// Using path.Ext to check the file extension
	if path.Ext(p) == ".go" {
		fmt.Println(p)
	}
	return err
}

func main() {
	flag.Parse()
	// Get the root directory using flags. By default root dir is current directory
	// WalkDir walks all the subdirectories and calls the walkFunc for each entry(file or directory)
	err := filepath.WalkDir(*rootdir, walkFunc)

	if err != nil {
		fmt.Println(err)
	}
}
