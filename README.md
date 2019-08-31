# Clerver

This is a task runner tool that can manage tasks in Yaml format.

You can execute a command by specifying a task written in Yaml as an argument. Of course, multiple executions are possible.

You can also specify environment variables that are required at runtime.

# Install
OSX
```
curl -LO https://github.com/Raita876/clever/releases/download/0.1.0/clever && chmod +x clever && sudo mv clever /usr/local/bin/clever
```

Linux
```
curl -LO https://github.com/Raita876/clever/releases/download/0.1.0/clever-linux && chmod +x clever-linux && sudo mv clever-linux /usr/local/bin/clever
```

# How to use

tasks.yaml(sample)

```
tasks:
  echo:
    command: "echo $FOO"
  shell:
    command: "echo $SHELL"
environments:
  - name: "FOO"
    value: "bar"
```

run

```
$ ./clever echo shell
bar
/bin/bash
```

help
```
$ ./clever help
This is a task runner tool that can manage tasks in Yaml format.

Usage: 
        clever <task>...
Options:
        task: Required Arguments → Specify the task you want to execute.
```

version
```
$ ./clever version
0.1.0
```
