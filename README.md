# QuEq

## Simple Quadratic Equation Solver

An application written on Go programming language to solve quadratic equations ( _a * x^2 + b * x + c = 0_ ).

### Build and run application

- Install Go 1.24+ : [go.dev](https://go.dev/dl/)

- Clone this repository:

```
git clone https://github.com/5aradise/queq.git
```

- Build app:

```
go build .
```

#### Interactive Mode

To start in interactive mode simply run

```
./queq
```

#### Non-interactive Mode

In non-interactive mode, you need to have a .txt file with lines containing three numbers.
Example of file structure:

```
4 2 0
2 1 -3
6 0 9
```

To start the application run the same command as in interactive mode, but add a file name to the end:

```
./queq file.txt
```

### Revert commit

[fbc8129f25bba1a4c8fbf47dfcb447b243d7c1f6](https://github.com/5aradise/queq/commit/fbc8129f25bba1a4c8fbf47dfcb447b243d7c1f6)
