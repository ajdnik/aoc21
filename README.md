# Advent of Code 2021

Solutions for [Advent of Code 2021](https://adventofcode.com/2021) in Go.

## Project Structure

```
cmd/dayXX/         - Entry points for each day
internal/dayXX/    - Solution logic and tests
input/             - Puzzle input files
utils/             - Shared utilities (parsing, math helpers)
```

## Prerequisites

- [Go](https://go.dev) 1.22+
- [golangci-lint](https://golangci-lint.run) (for linting)

## Usage

### Run a solution

```bash
make run DAY=01
```

### Run tests

```bash
make test
```

### Format and lint

```bash
make fmt
make lint
```

### Run all checks

```bash
make all
```
