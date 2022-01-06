package main

import (
    "encoding/json"
    "fmt"
)

/*
    Write a program which prompts the user to first enter a name, and then enter an address.
    Your program should create a map and add the name and address to the map using the keys “name” and “address”,
    respectively. Your program should use Marshal() to create a JSON object from the map,
    and then your program should print the JSON object.
*/

type Person struct {
    Name string
    Address string
}

func main() {
    var p1 = new(Person)

    fmt.Print("Enter a name: ")
    fmt.Scan(&p1.Name)

    fmt.Print("Enter an address: ")
    fmt.Scan(&p1.Address)

    bArr, _ := json.Marshal(p1)

    fmt.Println(string(bArr))
}
