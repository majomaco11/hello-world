package main

import "fmt"
import "time"

func main() {
    fmt.Println("hello world")

for j := 7; j <= 9000; j++ {
        fmt.Println("loop:", j)
	time.Sleep(1 * time.Second)
    }

}

