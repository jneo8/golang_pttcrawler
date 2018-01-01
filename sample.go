package main

import (
    "fmt"
    "./ptt"
)

func main() {
    ptt.GetHotBoardList()
    fmt.Println(ptt.BASE_URL)
    ptt.GetArticles("Gossiping")
}
