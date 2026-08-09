# Upskill Go!!

## Six main points about Go

1. Statically typed language - type should be defined beside variable declaration or a value should be provided so it can be inferred.
2. Strongly typed language - for example, an integer and a string cannot be added together.
3. Go is a compiled language
4. Fast compile time
5. Built In concurrency
6. Simplicity

## Module vs Package

Package is just a collection of `.go` files and a Module is a collection of packages.

## Data types

below data types can be used with `const` and `var`

1. `bool`
2. `float32`, `float64`
3. `int`, `int16`, `int32`, `int64`
4. `rune`
5. `string`
6. `uint`, `uint8`, `uint16`, `uint32`, `uint64`

## Functions & Control statements

`func` keyword is used to define functions in Go. `main` function defined in package `main` is always automatically called when running the program.

For control statements, `else` block needs be on the same line where `if` block's closing `}` ends.

Below statements is the right way

```go
if a == b {
    // do something
} else {
    // do something else
}
```

the wrong way of doing is

```go
if a == b {
    // do something
}
else {
    // do something else
}
```
