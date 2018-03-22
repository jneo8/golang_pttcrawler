package main

import (
    "fmt"
    "./ptt"
    "github.com/fatih/color"
)

func main() {
    // ptt.GetTitleList("Gossiping")
    // printhotboard()
    // getalldocurl()
    pool()
}


func getalldocurl() {
    board := ptt.GetBoard("Gossiping", 10)
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

func pool() {
    color.Red("Sample: Pool\n")
    p := ptt.NewPool()
    p.AddBabyFish("Gossiping", 3)
    p.AddBabyFish("NBA", 3)

    // for i := range p.Fishes {
    //     p.Fishes[i].Swim()
    //     p.Fishes[i].Swim()
    // }

    // Status
    p.Status()
}
