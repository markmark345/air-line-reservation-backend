package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func addImportAndFixPackage(fileName string, filePath string) {
	var result string
	newPath := "./internal/infrastucture/postgres/"
	newPackage := "package postgres"
	newImport := ". \"air-line-reservation-backend/internal/infrastucture/postgres/model\""

	if strings.Compare(fileName, "models.go") == 0 {
		newPath += "model/"
		newPackage = "package model"
	}

	newPath += fileName

	b, err := os.ReadFile(filePath)
	if err != nil {
        panic(err)
    }

	data := string(b)
	data = strings.Replace(data, "package sqlc", newPackage, -1)

	buf := bytes.NewBufferString(data)

	for {
		line, err := buf.ReadString('\n')
		lineWithOutSpace := strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Print(line)
				break
			}
			fmt.Println(err)
			break
		}

		if strings.Compare(lineWithOutSpace, "\"context\"") == 0 && !(fileName == "models.go" || fileName == "db.go")  {
			line += "\n\n" + "	" + newImport
		}

		result += line
	}
	
	err = os.WriteFile(newPath, []byte(result), 0644)
    if err != nil {
    	log.Fatal(err)
    }
}

func main() {
	var files []string

	rootDir := "./db/sqlc"

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
            fmt.Println(err)
            return err
        }
		
		if !info.IsDir() && filepath.Ext(path) == ".go" {
            files = append(files, path)
        }

        return nil
	})

	if err := os.MkdirAll("./internal/infrastucture/postgres/model", os.ModePerm); err != nil {
        log.Fatal(err)
    }

	if err != nil {
        fmt.Println(err)
    }

	for _, file := range files {
		addImportAndFixPackage(filepath.Base(file), file)
	}
}