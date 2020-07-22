package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Traverses along the files tree and fills the map
// Use it like:
// m := make(map[string][]string // map of ([directory_name] [ array of files if any ])
// path := "/home/user"
// if err := utils.Collect(path, m); err != nil {
// // handle an error
// }
// use the resulting map as you wish
func Collect(basePath string, m map[string][]string) error {
	if _, err := ioutil.ReadDir(basePath); err != nil {
		return err
	}
	err := traverse(&m, basePath)
	if err != nil {
		return err
	}
	return nil
}

// Traverse file tree and append to the map
func traverse(mapPtr *map[string][]string, path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	for _, file := range files {
        // I don't want any dotfiles
		if !strings.HasPrefix(file.Name(), ".") {
			if !file.IsDir() {
				(*mapPtr)[path] = append((*mapPtr)[path], file.Name())
			} else {
                // File type is directory so let's call traverse recursively
				_ = traverse(mapPtr,
					fmt.Sprintf("%s/%s", path, file.Name()))
			}
		}
	}
	return nil
}
