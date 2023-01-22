package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analysed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO and ERROR")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')

		if err != nil {
			break
		}

		if strings.Contains(s, *level) {
			fmt.Println(s)
		}

	}
}

// go run BasicCLIProgram.go
// go run BasicCLIProgram.go -level DEBUG
// go run BasicCLIProgram.go -level DEBUG -path test.log
// go run BasicCLIProgram.go -level DEBUG -path myapp.log
