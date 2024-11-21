package main

import "fmt"

func main() {
    // Three-component loop
    for i := 1; i < 5; i ++ {
        fmt.Println(i)
    }

    // While loop
    n := 1
    for n < 5 {
        n *= 2
        fmt.Println(n)
    }
}
