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
    url_list := ptt.GetAllDocUrl("Gossiping", 0, "", 10)
    for index, url := range url_list.Urls {
        fmt.Printf("%d: %s\n", index, url)
    }
}

func printhotboard() {

    s := ptt.GetHotBoardList()

    for name, url := range s {
        fmt.Printf("%s: %s\n", name, url)
    }

}
