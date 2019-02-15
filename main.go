package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	wd, err := os.Getwd()
	check(err)

	flag.Parse()

	// Update Version File

	var version string

	err = filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "version" {
			read, e := ioutil.ReadFile(path)

			if e != nil {
				return e
			}

			oldVersion := string(read)

			version, e = UpdateVersion(oldVersion)

			if e != nil {
				return e
			}

			return ioutil.WriteFile(path, []byte(version), 0)
		}

		return nil
	})

	check(err)

	// Add Tag to Git Repository

	err = Tag(version)

	check(err)
}
