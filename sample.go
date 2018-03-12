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
    board := ptt.GetBoard("Gossiping", 0, "", 10)
    for index, url := range board.Urls {
        fmt.Printf("%d: %s\n", index, url)
    }
}

func printhotboard() {

    s := ptt.GetHotBoardList()

    for name, url := range s {
        fmt.Printf("%s: %s\n", name, url)
    }

}
