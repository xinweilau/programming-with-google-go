package main

import (
    "fmt"
    "time"
)

var counter = 0

func increment() {
    v := counter
    v++
    time.Sleep(1 * time.Second)
    counter = v
}

func decrement() {
    v := counter
    v--
    time.Sleep(1 * time.Second)
    counter = v
}


/*
    Write two goroutines which have a race condition when executed concurrently.
    Explain what the race condition is and how it can occur.

    The race condition can be detected by go race detector tool by running the command:
    <go run race.go -race> or simply running the code using: <go run race.go>

    Race condition is when two processes/threads tries to access and modify the same piece of data at the same time.

    The race condition in this code comes from the accessing the shared global variable counter.
    Both functions increment() and decrement() first stores the copy of counter, followed by
    incrementing/decrementing the value. At this point, the goroutine will execute the code time.Sleep().
    which is a blocking operation.

    When Go encounters a blocking operation such as time.Sleep(), it will replace the current running goroutine with
    another goroutine. (either increment or decrement as it is not deterministic). The function that is executed next
    will also store a copy of the counter and go to sleep and return control to the initial goroutine.

    The 2 functions will then override the global variable counter with their own local copy, resulting in a
    race condition (final value will either be -1 or 1 instead of 0).

    An example of execution sequence:
    Function    |   local   |   global   | Remark
    main()      |   -       |   0        |
    increment() |   1       |   0        | time.Sleep() result in switching to next goroutine
    decrement() |   -1      |   0        | time.Sleep() result in switching to next goroutine
    increment() |   1       |   1        | local variable value is restored and control is passed to decrement()
    decrement() |   -1      |   -1       | local variable value is restored and control is return to main()
*/

func main() {
    fmt.Println("Initial value of counter is", counter)
    go increment()
    go decrement()
    time.Sleep(1 * time.Second)
    fmt.Println("Final value of counter is", counter)
}