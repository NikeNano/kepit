package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// ParseFile parses a golang file and removes all the omitempty structs
func ParseFile(filename string, tmpFile string) error {

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("File not found error")
	}

	output := string(input)
	output = strings.ReplaceAll(output, ",omitempty", "")
	// Write back to file
	err = ioutil.WriteFile(tmpFile, []byte(output), 0644)
	if err != nil {
		return errors.New("Unabla to write to tmp file")
	}

	return nil
}

// MakePlugin takes an input file and compiles a plugin of it
func MakePlugin(filePath string, pluginPath string) error {
	// Time to make the plugin
	//out, err := exec.Command("/bin/bash", "-c", "echo 'hello world'").Output()
	//command := fmt.Sprintf("Variable string %d content", data)
	command := fmt.Sprintf("go build -buildmode=plugin -o %s %s ", pluginPath, filePath)
	out, err := exec.Command("/bin/bash", "-c", command).Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%s", out)
	return nil
}
