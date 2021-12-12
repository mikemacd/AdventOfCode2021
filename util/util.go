package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadInput() []string {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return []string{}
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	output := []string{}
	for _, line := range bytes.Split(data, []byte("\n")) {
		output = append(output, strings.TrimSpace(string(line)))
	}

	return output
}
