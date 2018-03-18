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
    Articles []*Article
    CreatedTime time.Time
}

type Pool struct {
    Fishes []*Fish
}

func NewPool() Pool {
    p := Pool{}
    return p
}

func (p *Pool) Status() {
    color.Magenta("\nPool info: \n\n")
    for index, fish := range p.Fishes {
        color.Magenta("Fish %s :", strconv.Itoa(index))
        color.Magenta("\tBoardName: %s", fish.BoardName)
        color.Magenta("\tStatus: %d", fish.Status)
        color.Magenta("\tMax: %d", fish.Max)
        color.Magenta("\tCount: %d", len(fish.Board.Urls))
        color.Magenta("\tArticle Count: %d", len(fish.Articles))
        color.Magenta("\tCreate_time: %v", fish.CreatedTime)
    }
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
        articles := GetArticles(f)
        f.Articles = articles
    }
    f.Status += 1
}
