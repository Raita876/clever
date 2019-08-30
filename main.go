/*
Main package contains important processing of clever (task runner).

Usage:
	./clever <task>...

Options:
	task: Required Arguments → Specify the task you want to execute.

*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/go-yaml/yaml"
	"github.com/mattn/go-shellwords"
)

// YamlFile TaskFile path to load.
const YamlFile string = "clefile.yaml"

// Runner interface
type Runner interface {
	Run()
}

// CleverFile struct
type CleverFile struct {
	Tasks        Tasks        `yaml:"tasks"`
	Environments Environments `yaml:"environments"`
}

// Tasks list
type Tasks map[string]Task

// Task struct
type Task struct {
	Command string `yaml:"command"`
}

// Environments list
type Environments map[string]string

// Args Returns command line arguments in []string format.
func Args() []string {
	flag.Parse()
	args := flag.Args()

	return args
}

// Parse Generate struct based on the specified yaml file.
func Parse(filePath string) (CleverFile, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}

	var cf CleverFile
	err = yaml.Unmarshal(buf, &cf)
	if err != nil {
		fmt.Println(err)
		return cf, err
	}

	return cf, nil
}

// Run method
func (task *Task) Run() (string, error) {
	p := shellwords.NewParser()
	p.ParseEnv = true
	c, err := p.Parse(task.Command)
	if err != nil {
		panic(err.Error())
	}

	var out []byte
	switch len(c) {
	case 1:
		out, _ = exec.Command(c[0]).CombinedOutput()
	default:
		out, _ = exec.Command(c[0], c[1:]...).CombinedOutput()
	}

	return string(out), nil
}

func main() {
	args := Args()
	cf, _ := Parse(YamlFile)

	for _, a := range args {
		if _, ok := cf.Tasks[a]; ok {
			t := cf.Tasks[a]
			out, _ := t.Run()
			fmt.Print(out)
		}
	}
}
