package scanner

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Signature struct {
	algorithm string
	hash      string
	virus     string
}

var signatures = []Signature{}

func Init() {
	loadHashes()
}

func loadHashes() {
	data, err := os.ReadFile("misc/signatures.txt")
	if err != nil {
		log.Fatal(err)
	}
	var converted string = string(data)
	for _, line := range strings.Split(converted, "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		if len(line) == 1 {
			continue
		}
		x := strings.Split(line, "|")
		fmt.Println(len(x))
		signatures = append(signatures, Signature{x[0], x[1], x[2]})
	}
}


func hashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Failed to open file: %w", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("Failed to close file: %w", cerr)
		}
	}()

	hash := sha256.New()

	if _, err = io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("Failed to copy file content to hash: %w", err)
	}
	hashInBytes := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashInBytes)

	return hashString, err
}

func ScanFile(filePath string) (bool,string){
	hash, err := hashFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	for _, signature := range signatures {
		if(signature.hash == hash) {
			return true,signature.virus
		}
	}
	return false,""
}
