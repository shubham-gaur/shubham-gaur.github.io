package main

import "fmt"


func main() {
    // expected output: 18
    
    str := "aykjjjklbbatyyytalyu"
    size := len(str)
    fmt.Println(size)
    var index int
    for i := (size-1); i >= 0; i-- {
        if str[i] == byte('y') {
            index = i
            break
        }
    }
    fmt.Println(index)
}
