package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	Word  string `json:"word"`
	Level string `json:"level"`
	Pos   string `json:"pos"`
}

func replaceByWord(levels map[string][]Data, key string, characters string) string {
	for _, value := range levels[key] {
		if value.Word == characters {
			keyReplace := "\\?n\\"
			str := strings.Replace(keyReplace, "n", strconv.Itoa(len(characters)), 1)
			return str
		}
	}
	return ""
}
func randomIndex(lengthParagraph int, list_container map[int]bool) int {
	for true {
		key := rand.Intn(lengthParagraph)
		if !list_container[key] {
			list_container[key] = true
			return key
		}
	}
	return -1
}
func removeSpaceAfterPunctuation(input string) string {
	re := regexp.MustCompile(`\s*([.,])`)
	result := re.ReplaceAllString(input, "$1")
	return result
}

// random index word in english paragraph
// check word (random index) in map list language
// if word in map list language then replace word by key
// if word not in map list language then continue
// if word in map list language then add amountKey
// No find word in map list language then return -1
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
	levels := make(map[string][]Data)
	for _, entry := range data {
		levels[entry.Level] = append(levels[entry.Level], entry)
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
	re := regexp.MustCompile(`[a-zA-Z0-9]+|\S+`)
	words := re.FindAllString(value, -1)
	max := 30
	min := 10
	amount_words_level_a1 := rand.Intn(max-min) + min
	amount_words_level_a2 := max - amount_words_level_a1
	easy := map[string]int{
		"a1": amount_words_level_a1,
		"a2": amount_words_level_a2,
	}
	container_duplicate := make(map[int]bool)
	answer := make(map[string]int)
	for key, value := range easy {
		for easy[key] > 0 {
			indexWord := randomIndex(len(words), container_duplicate)
			is_valid, _ := regexp.MatchString(`\S`, words[indexWord])
			if is_valid && err == nil {
				s1 := replaceByWord(levels, key, words[indexWord])
				if s1 != "" {
					answer[words[indexWord]] = indexWord
					words[indexWord] = s1
					value--
					easy[key] = value
				}
			}
		}
	}
	fmt.Println(removeSpaceAfterPunctuation(strings.Join(words, " ")))
}
