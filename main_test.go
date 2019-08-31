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
			"echo": Task{
				Command: "echo $FOO",
			},
			"shell": Task{
				Command: "echo $SHELL",
			},
		},
		Environments: Environments{
			Env{
				Name:  "FOO",
				Value: "bar",
			},
		},
	}

	if diff := cmp.Diff(cf, expectedCleverFile); diff != "" {
		t.Errorf("Tasks mismatch (-tasks +expectedTasks):\n%s", diff)
	}

}

func TestRun(t *testing.T) {
	task := Task{
		Command: "echo $FOO",
	}
	environments := Environments{
		Env{
			Name:  "FOO",
			Value: "bar",
		},
	}
	environments.Set()

	out, err := task.Run()

	if err != nil {
		t.Error(err.Error())
	}

	expectedOut := "bar\n"

	if diff := cmp.Diff(out, expectedOut); diff != "" {
		t.Errorf("Task Run mismatch (-out +expectedOut):\n%s", diff)
	}
}
