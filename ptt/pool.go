package ptt

import (
    // "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
)

type Fish struct {
    Board Board
    BoardName string
    Status int
    Max int
}

type Pool struct {
    Fishes []Fish
}

func NewPool() Pool {
    p := Pool{}
    return p
}

func (p *Pool) AddFish(board_name string, max int) {
    f := Fish{Status: 0, BoardName: board_name, Max: max}
    p.Fishes = append(p.Fishes, f)
    color.Blue("%s\n", p)
}

func (f *Fish) Swim() {
    if (f.Status == 0) {
        board := GetBoard(f.BoardName, f.Max)
        f.Board = board
    }
    f.Status += 1
}

