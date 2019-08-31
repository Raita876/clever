/*
Main package contains important processing of clever (task runner).

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

var version string

// UsageTxt how to use clever.
const UsageTxt string = `This is a task runner tool that can manage tasks in Yaml format.

Usage: 
	clever <task>...
Options:
	task: Required Arguments → Specify the task you want to execute.`

// YamlFile clever (task runner) configuration file.
const YamlFile string = "clefile.yaml"

// Runner Interface that performs the given process.
type Runner interface {
	Run()
}

// CleverFile clever (task runner) configuration file struct.
type CleverFile struct {
	Tasks        Tasks        `yaml:"tasks"`
	Environments Environments `yaml:"environments"`
}

// Task This is a struct that owns the command.
type Task struct {
	Command string `yaml:"command"`
}

// Tasks List of tasks.
type Tasks map[string]Task

// Environments struct that holds environment variables as a list.
type Environments map[string]string

// Usage function to display usage.
func Usage() {
	fmt.Printf("%s\n", UsageTxt)
}

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

// Run This function executes the command defined in Task.
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
	flag.Usage = Usage
	args := Args()
	cf, _ := Parse(YamlFile)

	switch {
	case len(args) == 0:
		flag.Usage()
	case args[0] == "help":
		flag.Usage()
	case args[0] == "version":
		fmt.Print(version)
	default:
		for _, a := range args {
			if _, ok := cf.Tasks[a]; ok {
				t := cf.Tasks[a]
				out, _ := t.Run()
				fmt.Print(out)
			}
		}
	}
}
