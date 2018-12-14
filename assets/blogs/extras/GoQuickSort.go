package main

import (
    "fmt"
    "math/rand"
)

func quickSort(a []int) []int{
    if len(a)<2{
        fmt.Println("Ret Array and Len:", a, len(a))
        return a
    }

    left, right := 0, len(a)-1

    pivot := rand.Int() % len(a)
    fmt.Println("Pivot:            ", pivot)

    fmt.Println("Unsorted Array:   ", a)

    a[pivot], a[right] = a[right], a[pivot]
    fmt.Println("[1]Feed Array:    ", a)
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
    fmt.Println("[2]Output Array:  ", a)

    a[left], a[right] = a[right], a[left]
    fmt.Println("[3]Pivoted Array: ", a)

    quickSort(a[:left])
    quickSort(a[left+1:])

    return a
}

func main(){
    a := []int{10, 14, 5, 2, 1, 19, 15, 14, 11, 21, 12}
    quickSort(a)
    fmt.Println("Sorted Array:     ", a)
}
