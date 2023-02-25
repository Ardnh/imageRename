package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	folder := "jambu"
	prefix := "PS"
	mimeType := "jpg"
	files, _ := ioutil.ReadDir(fmt.Sprintf("./data/%s", folder))
	count := 0
	for i, f := range files {
		oldName := fmt.Sprintf("./data/%s/%s", folder, f.Name())
		newName := fmt.Sprintf("./data/%s/%s-%d.%s", folder, prefix, i, mimeType)
		err := os.Rename(oldName, newName)
		if err != nil {
			panic(err)
		}

		count++
	}
}
