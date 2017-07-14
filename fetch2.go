package main

import (
    "fmt"
    "strings"
    //"io/ioutil"
    //"net/http"
    //"os"
)

func main() {
   str := "Hello World"
   head := "http://"
   fmt.Println(str)
   fmt.Println(strings.HasPrefix(str, "He"))
   tmp := strings.HasPrefix(str, "He")
   if tmp == true {
       fmt.Println("Ok,come in")
       head += str
       fmt.Println(head)
   }
}



