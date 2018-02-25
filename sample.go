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
    pages := ptt.GetAllDocUrl("Gossiping", 0, "", 10)
    for index, v := range pages {
        fmt.Printf("%d: %s\n", index, v.Url)
    }
}

func printhotboard() {

    s := ptt.GetHotBoardList()

    for name, url := range s {
        fmt.Printf("%s: %s\n", name, url)
    }

}
