/*
Copyright 2024 The oai Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oai

import (
	"bufio"
	"encoding/json"
	"os"
)

// Reader reads a file line by line. handler big file maybe >= 1GB file.[scanner.Buffer(buf, 2048*1024)]
// https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func Reader(filename string, fn func(le string) error) error {
	// Open the file.
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new Scanner for the file.
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		// Process the line here.
		if err := fn(line); err != nil {
			return err
		}
	}
	return nil
}

// Write writes a file line by line.
func Write(filename string, account interface{}) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	jsonStr, err := json.Marshal(account)
	if err != nil {
		return err
	}

	_, err = writer.WriteString(string(jsonStr) + "\n")
	if err != nil {
		return err
	}

	return writer.Flush()
}
