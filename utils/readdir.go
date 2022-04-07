package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/TeamIndex/Backend/schema"
)

func readdir(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return file
}

func ParseJokes() []schema.Joke {
	// jsonFile := readdir("./jokes/index.json")
	byteValue, _ := ioutil.ReadFile("./jokes/index.json")
	var jokes []schema.Joke
	err := json.Unmarshal(byteValue, &jokes)
	if err != nil {
		panic(err)
	}
	return jokes

}


func ParseFacts() []schema.Fact {
	byteValue, _ := ioutil.ReadFile("./facts/index.txt")
	var facts []schema.Fact
	err := json.Unmarshal(byteValue, &facts)
	if err != nil {
		panic(err)
	}
	return facts
}

