package ptt

import (
    // "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
    "time"
    "strconv"
)

type Fish struct {
    Board Board
    BoardName string
    Status int
    Max int
    CreatedTime time.Time
}

type Pool struct {
    Fishes []*Fish
}

func NewPool() Pool {
    p := Pool{}
    return p
}

func (p *Pool) Count() map[string]int{

    result := make(map[string]int)
    for index, fish := range p.Fishes {
        result[strconv.Itoa(index) + "-" + fish.BoardName] = len(fish.Board.Urls)
    }
    return result
}

func (p *Pool) AddFish(board_name string, max int) {
    f := &Fish{Status: 0, BoardName: board_name, Max: max, CreatedTime: time.Now()}
    p.Fishes = append(p.Fishes, f)
    color.Blue("%s\n", p)
}

func (f *Fish) Swim() {
    if (f.Status == 0) {
        board := GetBoard(f.BoardName, f.Max)
        f.Board = board
    } else if (f.Status == 1) {
        GetArticles(f)
    }
    f.Status += 1
}
