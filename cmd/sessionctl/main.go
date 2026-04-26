package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	ptr_string := flag.String("dir", "", "Path to the PHP sessions directory")
	flag.Parse()
	value := *ptr_string

	if value == "" {
		flag.Usage()
		os.Exit(1)
	}
	info, err := os.Stat(value)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintln(os.Stderr, "Not a PHP session directory")
		os.Exit(1)
	}
	obj_in_dir, err := os.ReadDir(value)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, obj := range obj_in_dir {
		obj_name := obj.Name()
		if !obj.IsDir() && strings.HasPrefix(obj_name, "sess_") {
			fullpath := filepath.Join(value, obj_name)
			file_text, read_file_err := os.ReadFile(fullpath)
			if read_file_err != nil {
				fmt.Fprintln(os.Stderr, read_file_err)
				os.Exit(1)
			}
			fmt.Println(string(file_text))
		}
	}
}
