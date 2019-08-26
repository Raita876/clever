# Clerver
Task runner tool by Go

# How to use
tasks.yaml
```
hello:
  command: "echo HelloWorld"
ls:
  command: "ls -a"
```

run
```
$ ./clever hello ls
HelloWorld
.
..
Makefile
README.md
clever
main.go
tasks.yaml
```

# Todo
- [ ] Add test.
- [ ] Make the variable available.
- [ ] Make environment variables usable。
- [ ] Edit comment.
