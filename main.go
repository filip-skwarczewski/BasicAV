package main

import (
	"basicAv/scanner"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filePath string
	args := os.Args[1:]
	for arg := 0; arg < len(args); arg++ {
		if strings.HasPrefix(args[arg], "--file=") {
			filePath = strings.Replace(args[arg], "--file=", "", 1)
			fmt.Println(filePath)
		}
	}
	if filePath == "" {
		fmt.Println("You haven't set a file! Use this: go run main.go --file=<file>")
		return
	}
	fmt.Println("-----------[AV]-----------")
	fmt.Println("Made by Skwaraa")
	fmt.Println("Initializing scanner...")
	scanner.Init()
	fmt.Println("File to check: ", filePath)
	fmt.Println("Scanning...")
	ifVirus, name := scanner.ScanFile(filePath)
	if ifVirus == true {
		fmt.Println("Virus found! Virus:", name)
		return
	}
	fmt.Println("File hasn't been detected as a virus!")

}
