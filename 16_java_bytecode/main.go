package main

import (
    "fmt"
    "jvm"
)

func main() {
    // Read the Java bytecode file.
    bytecode, err := jvm.ReadFile("hello.class")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Parse the bytecode.
    classFile, err := jvm.ParseClassFile(bytecode)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the class file.
    fmt.Println(classFile)
}

