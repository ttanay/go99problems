package main

import "fmt"

// P04: Find the number of elements in a list
func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    fmt.Println(len(arr))
}
