package main

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	files := arguments

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	var results = list.New()

	for _, directory := range pathSplit {
		for _, file := range files {
			fullPath := filepath.Join(directory, file)
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err != nil {
				continue
			}
			mode := fileInfo.Mode()
			// Is it a regular file?
			if !mode.IsRegular() {
				continue
			}

			// Is it executable?
			if mode&0111 != 0 {
				//fmt.Println(fullPath)
				results.PushFront(fullPath)
			}
		}
	}
	if results.Len() > 0 {
		for element := results.Front(); element != nil; element = element.Next() {
			fmt.Println(element.Value)
		}
	}
}
