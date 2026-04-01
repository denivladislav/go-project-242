# Pathsize
Utility that calculates the size of a file or directory

### Hexlet tests and linter status:
[![Actions Status](https://github.com/denivladislav/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/denivladislav/go-project-242/actions)
### CI
[![Actions Status](https://github.com/denivladislav/go-project-242/actions/workflows/CI.yml/badge.svg)](https://github.com/denivladislav/go-project-242/actions)

### How to use
```bash
$ make build

# example: make run ARGS="-a ./testfolder"
$ make run ARGS="{PASS_FLAGS_AND_PATH}"

$ make test
```

### Flags
``` bash
-H (--human) #human-readable sizes (auto-select unit)

–a (--all)  #include hidden files and directories

-r (--recursive)  #recursive size of directories
```

### Demo
[![asciicast](https://asciinema.org/a/cjCIKIAZRNAco6X6.svg)](https://asciinema.org/a/cjCIKIAZRNAco6X6)
