package main

import (
    "fmt"
    "os"
    "runtime"
)

func main() {
    fmt.Printf("Running: %s\n", os.Args[0])

    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())
}
