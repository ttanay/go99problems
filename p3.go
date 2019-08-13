package main

import "fmt"

// P03: Find the Kth element of a list
func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    var k int
    fmt.Scan(&k)

    fmt.Println(getElement(k, arr))
}

func getElement(k int, arr []int) int {
    return arr[k]
}
