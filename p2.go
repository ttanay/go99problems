package main

import "fmt"

// P02: Find last but one element
func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    fmt.Println(arr[len(arr)-2])
}
