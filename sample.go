package main

import (
    "fmt"
    "./ptt"
)

func main() {
    // ptt.GetTitleList("Gossiping")
    // printhotboard()
    getalldocurl()
}


func getalldocurl() {
    ptt.GetAllDocUrl("Gossiping", 0, "", 10)
}

func printhotboard() {

    s := ptt.GetHotBoardList()

    for name, url := range s {
        fmt.Printf("%s: %s\n", name, url)
    }

}
