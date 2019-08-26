package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/go-yaml/yaml"
	"github.com/mattn/go-shellwords"
)

// TasksYamlFile TaskFile path to load.
const TasksYamlFile string = "tasks.yaml"

// Runner interface
type Runner interface {
	Run()
}

// Tasks list
type Tasks map[string]Task

// Task struct
type Task struct {
	Command string
}

// Args method
func Args() []string {
	flag.Parse()
	args := flag.Args()

	return args
}

// Parse method
func Parse() (Tasks, error) {
	buf, err := ioutil.ReadFile(TasksYamlFile)
	if err != nil {
		panic(err.Error())
	}

	data := make(Tasks, 10)

	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

// Run method
func (task *Task) Run() {
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

	fmt.Print(string(out))
}

func main() {
	args := Args()
	tasks, _ := Parse()

	for _, a := range args {
		if _, ok := tasks[a]; ok {
			t := tasks[a]
			t.Run()
		}
	}
}
