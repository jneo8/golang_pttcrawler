package ptt

import (
    // "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
    "time"
    "strconv"
)

type Fish struct {
    BoardName string
    Url string
    Status int
    Article *Article
    CreatedTime time.Time
}

type BabyFish struct {
    BoardName string
}


type Pool struct {
    Fishes []*Fish
    BabyFishes []*BabyFish
}

func NewPool() Pool {
    p := Pool{}
    return p
}

func (p *Pool) Status() {
    color.Magenta("\nPool info: \n\n")
    for index, bady_fish := range p.BabyFishes {
        color.Magenta("BabyFish %s :", strconv.Itoa(index))
        color.Magenta("\t BoardName: %s", bady_fish.BoardName)
    }
}

func (p *Pool) AddBabyFish(board_name string, num int) {
    for i := 1; i <= num; i++ {
        babyfish := &BabyFish{BoardName: board_name}
        p.BabyFishes = append(p.BabyFishes, babyfish)
    }
}

// func (p *Pool) AddFish(board_name string, max int) {
//     f := &Fish{Status: 0, BoardName: board_name, Max: max, CreatedTime: time.Now()}
//     p.Fishes = append(p.Fishes, f)
//     color.Blue("%s\n", p)
// }
// 
// func (f *Fish) Swim() {
//     // In Step 1, Get all url in Give BoardName
//     if (f.Status == 0) {
//         board := GetBoard(f.BoardName, f.Max)
//         f.Board = board
//     }
//     // In Step 2, Get all article.
//     if (f.Status == 1) {
//         articles := GetArticles(f)
//         f.Articles = articles
//     }
//     f.Status += 1
// }
