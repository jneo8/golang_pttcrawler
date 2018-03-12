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
    urls := ptt.GetAllDocUrl("Gossiping", 0, "", 10)
    for index, url := range urls {
        fmt.Printf("%d: %s\n", index, url)
    }
}

func printhotboard() {

    s := ptt.GetHotBoardList()

    for name, url := range s {
        fmt.Printf("%s: %s\n", name, url)
    }

}
