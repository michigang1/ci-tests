# Postfix to infix converter using Unit-test and GitHub Actions

This is lab 2 for software architecture subject.

## Installation

### 1. Clone repository

```bash
git clone https://github.com/michigang1/ci-tests.git
```

### 2.  Run with help

```bash
go run cmd/example/main.go -h
```

## Testing 
```bash
go test
```

## Using
### Specify input by flags:

* -e "expression"*
* -f file_path.txt

### Specify output by:

* -o file_path.txt

```bash

  -e string
        Expression to compute
  -f string
        Reading from file
  -o string
        Saving to file

```
## GitHub Actions:
* [Check is passed](https://github.com/michigang1/ci-tests/actions/runs/4559085587)
* [Check isn't passed](https://github.com/michigang1/ci-tests/actions/runs/4557343360)

--------------
* All checks of this project - [here](https://github.com/michigang1/ci-tests/actions)
* All pull-requests of this project - [here](https://github.com/michigang1/ci-tests/pulls?q=is%3Apr+is%3Aclosed)

## Contributors:

- Mykhailo Chirozidi

- Ivan Lotokhin

- Lev Pavelko
