package main

import "fmt"

/*
    Write a program which allows the user to create a set of animals and to get information about those animals.
    Each animal has a name and can be either a cow, bird, or snake.

    With each command, the user can either create a new animal of one of the three types,
    or the user can request information about an animal that he/she has already created.

    Each animal has a unique name, defined by the user.
    Note that the user can define animals of a chosen type,
    but the types of animals are restricted to either cow, bird, or snake.

    Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
    Your program should accept one command at a time from the user, print out a response,
    and print out a new prompt on a new line. Your program should continue in this loop forever.

    Every command from the user must be either a “newanimal” command or a “query” command.

    Each “newanimal” command must be a single line containing three strings.

    The first string is “newanimal”.
    The second string is an arbitrary string which will be the name of the new animal.
    The third string is the type of the new animal, either “cow”, “bird”, or “snake”.

    Your program should process each newanimal command by creating the new animal
    and printing “Created it!” on the screen.

    Each “query” command must be a single line containing 3 strings.

    The first string is “query”.
    The second string is the name of the animal.
    The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.

    Your program should process each query command by printing out the requested data.

    Define an interface type called Animal which describes the methods of an animal.

    Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
    which take no arguments and return no values.

    The Eat() method should print the animal’s food,
    the Move() method should print the animal’s locomotion, and
    the Speak() method should print the animal’s spoken sound.

    Define three types Cow, Bird, and Snake.
    For each of these three types, define methods Eat(), Move(), and Speak()
    so that the types Cow, Bird, and Snake all satisfy the Animal interface.

    When the user creates an animal, create an object of the appropriate type.

    Your program should call the appropriate method when the user issues a query command.
*/

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct {
    name string
    food string
    locomotion string
    noise string
}

type Bird struct {
    name string
    food string
    locomotion string
    noise string
}

type Snake struct {
    name string
    food string
    locomotion string
    noise string
}

func (c *Cow) Eat() {
    fmt.Println(c.food)
}

func (c *Cow) Move() {
    fmt.Println(c.locomotion)
}

func (c *Cow) Speak() {
    fmt.Println(c.noise)
}

func (b *Bird) Eat() {
    fmt.Println(b.food)
}

func (b *Bird) Move() {
    fmt.Println(b.locomotion)
}

func (b *Bird) Speak() {
    fmt.Println(b.noise)
}

func (s *Snake) Eat() {
    fmt.Println(s.food)
}

func (s *Snake) Move() {
    fmt.Println(s.locomotion)
}

func (s *Snake) Speak() {
    fmt.Println(s.noise)
}

func createNewAnimal(name, animalType string) Animal {
    var animal Animal

    switch animalType {
    case "cow":
        animal = &Cow{name, "grass", "walk", "moo"}
    case "bird":
        animal = &Bird{name, "worms", "fly", "peep"}
    case "snake":
        animal = &Snake{name, "mice", "slither", "hsss"}
    }

    return animal
}

func queryAnimal(animals []Animal, name, action string) {
    for _, animal := range animals {
        switch t := animal.(type) {
        case *Cow:
            if t.name == name {
                ExecuteAnimalAction(t, action)
            }
        case *Bird:
            if t.name == name {
                ExecuteAnimalAction(t, action)
            }
        case *Snake:
            if t.name == name {
                ExecuteAnimalAction(t, action)
            }
        }
    }
}

func ExecuteAnimalAction(a Animal, action string) {
    switch action {
    case "eat":
        a.Eat()
    case "move":
        a.Move()
    case "speak":
        a.Speak()
    }
}

func main() {
    var command, name, args string
    animals := make([]Animal, 0, 10)

    for {
        fmt.Print("> ")
        fmt.Scanln(&command, &name, &args)

        switch command {
        case "newanimal":
            animals = append(animals, createNewAnimal(name, args))
        case "query":
            queryAnimal(animals, name, args)
        }

        fmt.Println()
    }
}
