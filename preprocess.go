package main

import (
	"bufio"
	"os"
	"strings"
)

func preprocess(path string) error {
	var rawPath = path + "_raw"
	must(os.Rename(path, rawPath))

	var rawFile, err = os.Open(rawPath)
	must(err)

	outFile, err := os.Create(path)
	must(err)

	var scanner = bufio.NewScanner(rawFile)
	for scanner.Scan() {
		var token = scanner.Text()
		if strings.HasPrefix(token, "DROP SCHEMA") {
			continue
		}

		if strings.HasPrefix(token, "CREATE SCHEMA") {
			continue
		}

		if strings.HasPrefix(token, "COMMENT ON SCHEMA") {
			continue
		}

		var _, err = outFile.WriteString(token + "\n")
		must(err)
	}

	must(os.Remove(rawPath))
	return nil
}
