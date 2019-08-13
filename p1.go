package main

import "fmt"


//P01: Find the last element of a list

func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    fmt.Println(arr[len(arr)-1])
}
