package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	taskFilePath := "tasks.yaml"
	tasks, err := Parse(taskFilePath)

	if err != nil {
		t.Error(err.Error())
	}

	expectedTasks := Tasks{
		"hello": Task{
			Command: "echo HelloWorld",
		},
		"ls": Task{
			Command: "ls -a",
		},
		"shell": Task{
			Command: "echo $SHELL",
		},
	}

	if diff := cmp.Diff(tasks, expectedTasks); diff != "" {
		t.Errorf("Tasks mismatch (-tasks +expectedTasks):\n%s", diff)
	}

}

func TestRun(t *testing.T) {
	task := Task{
		Command: "echo TestRun.",
	}

	out, err := task.Run()

	if err != nil {
		t.Error(err.Error())
	}

	expectedOut := "TestRun.\n"

	if diff := cmp.Diff(out, expectedOut); diff != "" {
		t.Errorf("Task Run mismatch (-out +expectedOut):\n%s", diff)
	}

}
