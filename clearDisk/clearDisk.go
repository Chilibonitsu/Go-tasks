package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type File struct {
	Name string
	Size int64
}

// type File struct {
// 	Info fs.FileInfo
// }
//filesArr[0].Info.Name()

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	downloadsDir := filepath.Join(homeDir, "Downloads")

	files, err := os.ReadDir(downloadsDir)

	if err != nil {
		log.Fatal(err)
	}

	var filesArr []File
	for _, file := range files {
		info, err := file.Info()

		if err != nil {
			log.Printf("cant read info of %s: %v", file.Name(), err)
			continue
		}
		filesArr = append(filesArr, File{
			Name: file.Name(),
			Size: info.Size(),
		})

	}

	sort.Slice(filesArr, func(i, j int) bool {
		return filesArr[i].Size > filesArr[j].Size
	})

	for i := range filesArr {
		fmt.Println("", filesArr[i].Name, filesArr[i].Size)
	}

	if len(filesArr) < 5 {
		err := errors.New("there are less than 5 files")
		fmt.Println(err.Error())
		return
	}

	var input string
	fmt.Println("Delete 5 largets files Y/N?")
	fmt.Scanln(&input)

	if input == "y" || input == "Y" {

		for i := 0; i < 5; i++ {
			name := filesArr[i].Name

			fullPath := filepath.Join(downloadsDir, name)
			err := os.Remove(fullPath)
			if err != nil {
				log.Printf("cant delete file %s: %v", name, err)
			} else {
				fmt.Printf("file %s deleted \n", name)
			}
		}

	}
}
