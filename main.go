package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	// extension to filter
	ext string
	// min file size
	size int64
	// list files
	list bool
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	// Action options
	list := flag.Bool("list", false, "list files only")
	// Filter options
	ext := flag.String("ext", "", "file ext to filter out")
	size := flag.Int64("size", 0, "minimun file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	err := run(*root, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			// If list was explicitly set, don't do anything else
			if cfg.list {
				return listFile(path, out)
			}

			// List is the default option if nothing else was set
			return listFile(path, out)
		})
}
