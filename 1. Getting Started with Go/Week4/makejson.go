package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    var m map[string]string
    m = make(map[string]string)
    var name,adress string
    fmt.Println("Please input your name:")
    fmt.Scanln(&name)
    m["name"]=name
    fmt.Println("Then input your adress:")
    fmt.Scanln(&adress)
    m["adress"]=adress
    b, err := json.Marshal(m)
    if err != nil {
        fmt.Println("Encoding faild")
    } else {
        fmt.Println("Encoded data : ")
        fmt.Println(b)
        fmt.Println("Decoded data : ")
        fmt.Println(string(b))
    }
}