package main

import "fmt"


func main() {
    // expected output: aykklattalyu

    str := "aykjjjklbbatyyytalyu"

    size := len(str)
    var final string
    // fmt.Println(size)
   
    if str[0] != str[1] {
        final = final + string(str[0])
    }
   
    for i:= 1; i<size-1; i++ {
        // fmt.Println(string(str[i]))
        if str[i] != str[i-1] && str[i] != str[i+1] {
            final = final + string(str[i])
        }
    }
   
    if str[size-1] != str[size-2] {
        final = final + string(str[size-1])
    }
    fmt.Println(final)
}
