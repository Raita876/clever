package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	filePath := "clefile.yaml"
	cf, err := Parse(filePath)

	if err != nil {
		t.Error(err.Error())
	}

	expectedCleverFile := CleverFile{
		Tasks: Tasks{
			"hello": Task{
				Command: "echo HelloWorld",
			},
			"ls": Task{
				Command: "ls -a",
			},
			"shell": Task{
				Command: "echo $SHELL",
			},
		},
		Environments: map[string]string{
			"FOO": "bar",
		},
	}

	if diff := cmp.Diff(cf, expectedCleverFile); diff != "" {
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
