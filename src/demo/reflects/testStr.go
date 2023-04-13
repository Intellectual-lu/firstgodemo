package reflects

import (
	"bufio"
	"os"
)

func GetStrings(file string) ([]string, error) {
	var contant []string
	open, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		text := scanner.Text()
		contant = append(contant, text)
	}
	err = open.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return contant, nil
}
