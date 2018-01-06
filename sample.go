package main

import (
    "fmt"
    "./ptt"
)

func main() {
    ptt.GetHotBoardList()
    fmt.Println(ptt.BASE_URL)
    ptt.GetTitleList("Gossiping")
}
