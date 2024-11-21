package main

import "fmt"

func main() {
    // Three-component loop
    // for i := 1; i < 5; i ++ {
    //     fmt.Println(i)
    // }

    // While loop
    // n := 1
    // for n < 5 {
    //     n *= 2
    //     fmt.Println(n)
    // }

    // Infinite loop
    // sum := 0
    // for {
    //     sum ++
    //     break
    // }

    // For-each range loop
    // strings := []string{"hello", "world"}
    // for i, s := range strings {
    //     fmt.Println(i, s)
    // }

    // Exit a loop
    sum := 0
    for i := 1; i < 5; i ++ {
        if i % 2 != 0 {
            continue
        }
        sum += i
    }
    fmt.Println(sum)
}
