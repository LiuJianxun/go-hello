package main

import (
    "fmt"
    "os"
    "strings"
)
/*func main() {
    fmt.Println(os.Args[1:])
}*/

func main() {
    fmt.Println(strings.Join(os.Args[1:],"-+-"))
}
