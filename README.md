# Pathsize
CLI-utility that calculates the size of a file or a directory.  
Supports recursive traversal, hidden files, and human-readable output.

### Hexlet tests and linter status:
[![Actions Status](https://github.com/denivladislav/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/denivladislav/go-project-242/actions)
### CI
[![Actions Status](https://github.com/denivladislav/go-project-242/actions/workflows/CI.yml/badge.svg)](https://github.com/denivladislav/go-project-242/actions)

### Requirements
Go 1.26+

### How to use

```bash
# Install
make build

# Single file
bin/hexlet-path-size ./testdata/file.txt

# Folder, with flags
bin/hexlet-path-size -a -r ./testdata
```

See [Flags](#flags) and [Demo](#demo) for details.

### Development
See [Makefile](./Makefile) for tests, lint, etc.

### Flags
``` bash
-H (--human) #human-readable sizes (auto-select unit)

-a (--all)  #include hidden files and directories

-r (--recursive)  #recursive size of directories
```

### Demo
[![asciicast](https://asciinema.org/a/cjCIKIAZRNAco6X6.svg)](https://asciinema.org/a/cjCIKIAZRNAco6X6)
