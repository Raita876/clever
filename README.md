# Clerver

Task runner tool by Go

# How to use

tasks.yaml(sample)

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

- [x] Add test.
- [x] Make the variable available.
- [x] Make environment variables usable。
- [x] Edit comment.
- [ ] error handling.
