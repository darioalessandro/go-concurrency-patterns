# go-concurrency-patterns

**Note**: All code based on 
William Kennedy. “Go in Action.”

Goroutine:

Abstraction on top of go runtime to execute functions in a concurrent way.

1. The number of goroutines that are executed at the same time is limited by the number of logical processors that are 
allocated by the program.

2. Example that allocates 2 virtual processors: 

```go
runtime.GOMAXPROCS(2)
```

# Description of programs:

## 1logicalprocessor.go

Program that creates two goroutines that display the English alphabet with lower and uppercase letters in a concurrent
fashion.
Excerpt From: William Kennedy. “Go in Action.”.

## 2hellosyncchannels

Shows inter goroutine communications using an unbufferred channel.

```go
functionA -> hello -> functionB

after functionB receives hello, the program will terminate.
```

## 3hellobufferedchannels

Shows how to use buffered channels using a statically defined number of
go routines and tasks which are configurable.

```go
const (
	numberGoroutines 	= 4
	taskLoad 		= 10
)
```

## 4supervisedbufferedchannels

Introduces the concept of a supervisor, that waits for
all workers to end, before each worker goroutine ends, they send 
a dead letter to the supervisor.

