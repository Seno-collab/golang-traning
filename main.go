package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type Writer interface {
	Writer(p []byte) (int, error)
}

type Reader interface {
	Reader(p []byte) (int, error)
}
type Data struct {
	Word  string `json:"word"`
	Level string `json:"level"`
	Pos   string `json:"pos"`
}

func main() {
	file, err := os.Open("level.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	defer file.Close()
	var data []Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		panic(fmt.Sprintf("Failed to decode JSON: %v", err))
	}
	level := make(map[string][]Data)
	for _, entry := range data {
		level[entry.Level] = append(level[entry.Level], entry)
	}
	fileEnglish, err := os.Open("english.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	defer fileEnglish.Close()
	content, err := io.ReadAll(fileEnglish)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	value := string(content)
	modifiedWords := []string{}
	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(value, -1)
	modifiedWords = append(modifiedWords, words...)
	fmt.Println(modifiedWords)
}
