package main

import (
    "fmt"
    "math"
)

/*
    Let us assume the following formula for displacement s as a function of time t, acceleration a,
    initial velocity vo, and initial displacement so.

    s = ½ a t2 + vot + so

    Write a program which first prompts the user
    to enter values for acceleration, initial velocity, and initial displacement.
    Then the program should prompt the user to enter a value for time and the
    program should compute the displacement after the entered time.

    You will need to define and use a function called GenDisplaceFn() which takes three float64
    arguments, acceleration a, initial velocity vo, and initial displacement so.

    GenDisplaceFn() should return a function which computes displacement as a function of time,
    assuming the given values acceleration, initial velocity, and initial displacement.

    The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
    float64 argument which is the displacement travelled after time t.
*/

func GenDisplaceFn(a, vo, so float64) func (t float64) float64 {
    return func (t float64) float64 {
        return ((1/2.0) * a * math.Pow(t, 2)) + (vo * t) + so
    }
}

/*
    For example, let’s say that I want to assume the following values for acceleration,
    initial velocity, and initial displacement: a = 10, vo = 2, so = 1.

    I can use the following statement to call GenDisplaceFn() to
    generate a function fn which will compute displacement as a function of time.
    fn := GenDisplaceFn(10, 2, 1)

    Then I can use the following statement to print the displacement after 3 seconds.
    fmt.Println(fn(3))

    And I can use the following statement to print the displacement after 5 seconds.
    fmt.Println(fn(5))
*/

func main() {
    var a, vo, so float64

    fmt.Print("Enter value for acceleration: ")
    fmt.Scan(&a)

    fmt.Print("Enter value for initial velocity: ")
    fmt.Scan(&vo)

    fmt.Print("Enter value for initial displacement: ")
    fmt.Scan(&so)

    fn := GenDisplaceFn(a, vo, so)

    fmt.Println("Displacement after 3 seconds is", fn(3))
    fmt.Println("Displacement after 5 seconds is", fn(5))
}