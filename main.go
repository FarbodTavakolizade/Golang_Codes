/*
cd [optional path]       
mkdir Zombie_Attack_Game
cd Zombie_Attack_Game
go mod init Game
go get github.com/gdamore/tcell
code .
create main.go  file in vscode
*/


package main
import (
    "fmt"
    "math/rand"
    "os"
    "time"
    "github.com/gdamore/tcell"
)

const (
    Width      = 80
    Height     = 25
    MaxZombies = 3
)

type Point struct {
    row, col int
    ch       rune
}

var screen tcell.Screen
var player Point
var zombies []Point
var bullets []Point
var isGameOver, isPaused bool
var score int

func main() {
    rand.Seed(time.Now().UnixNano())

    var err error
    screen, err = tcell.NewScreen()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    if err = screen.Init(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    defer screen.Fini()

    resetGame()

    inputChan := make(chan *tcell.EventKey)
    go func() {
        for {
            ev := screen.PollEvent()
            if key, ok := ev.(*tcell.EventKey); ok {
                inputChan <- key
            }
        }
    }()

    ticker := time.NewTicker(150 * time.Millisecond) 
    defer ticker.Stop()


loop:
    for {
        select {
        case key := <-inputChan:
            if isGameOver {
                if key.Rune() == 'r' {
                    resetGame()
                } else if key.Rune() == 'q' {
                    break loop
                }
                continue
            }

            switch key.Rune() {
            case 'q':
                break loop
            case 'p':
                isPaused = !isPaused
            case 'w':
                if player.row > 1 {
                    player.row--
                }
            case 's':
                if player.row < Height-2 {
                    player.row++
                }
            case 'a':
                if player.col > 1 {
                    player.col--
                }
            case 'd':
                if player.col < Width-2 {
                    player.col++
                }
            case ' ':
                bullets = append(bullets, Point{player.row, player.col + 1, '➤'})
            }
        case <-ticker.C:
            if !isPaused && !isGameOver {
                updateGame()
                drawGame()
            }
        }
    }

    screen.Clear()
    drawText(Height/2, Width/2-7, "Thanks for playing!")
    screen.Show()
    time.Sleep(2 * time.Second)
}

func resetGame() {
    screen.Clear()
    player = Point{Height / 2, 2, '👾'}
    zombies = []Point{}
    bullets = []Point{}
    score = 0
    isGameOver = false
    isPaused = false
    drawGame()
    drawText(Height/2, Width/2-10, "Press q to quit, r to restart")
    screen.Show()
}


func updateGame() {

    newBullets := []Point{}
    for _, b := range bullets {
        if b.col < Width-2 {
            b.col++
            newBullets = append(newBullets, b)
        }
    }
    bullets = newBullets

    newZombies := []Point{}
    for _, z := range zombies {
        if z.col > 1 {
            z.col--
            newZombies = append(newZombies, z)
            if z.col == 1 {
                isGameOver = true
            }
        }
    }
    zombies = newZombies

    if len(zombies) < MaxZombies && rand.Intn(10) < 1 {
        zombies = append(zombies, Point{rand.Intn(Height-2) + 1, Width - 2, '🧟'})
    }


    bulletsToRemove := make(map[int]bool)
    zombiesToRemove := make(map[int]bool)


    for zi, z := range zombies {
        for bi, b := range bullets {
            if b.row == z.row && b.col >= z.col {
                score++
                bulletsToRemove[bi] = true
                zombiesToRemove[zi] = true
                break
            }
        }
    }


    filteredBullets := []Point{}
    for i, b := range bullets {
        if !bulletsToRemove[i] {
            filteredBullets = append(filteredBullets, b)
        }
    }
    bullets = filteredBullets

    filteredZombies := []Point{}
    for i, z := range zombies {
        if !zombiesToRemove[i] {
            filteredZombies = append(filteredZombies, z)
        }
    }
    zombies = filteredZombies

    for _, z := range zombies {
        if z.row == player.row && z.col == player.col {
            isGameOver = true
        }
    }
}


func drawGame() {
    screen.Clear()

    for c := 0; c < Width; c++ {
        screen.SetContent(c, 0, '-', nil, tcell.StyleDefault)
        screen.SetContent(c, Height-1, '-', nil, tcell.StyleDefault)
    }
    for r := 0; r < Height; r++ {
        screen.SetContent(0, r, '|', nil, tcell.StyleDefault)
        screen.SetContent(Width-1, r, '|', nil, tcell.StyleDefault)
    }

    screen.SetContent(player.col, player.row, player.ch, nil, tcell.StyleDefault)

    for _, z := range zombies {
        screen.SetContent(z.col, z.row, z.ch, nil, tcell.StyleDefault)
    }

    for _, b := range bullets {
        screen.SetContent(b.col, b.row, b.ch, nil, tcell.StyleDefault)
    }

    if isGameOver {
        drawText(Height/2, Width/2-6, "Game Over! Press 'r' to restart or 'q' to quit")
        drawText(Height/2+1, Width/2-7, fmt.Sprintf("Your score: %d", score))
    } else {
        drawText(Height, 2, fmt.Sprintf("Score: %d  (q=quit, p=pause, w/a/s/d=move, space=shoot)", score))
    }


    screen.Show()
}


func drawText(row, col int, text string) {
    for i, ch := range text {
        screen.SetContent(col+i, row, ch, nil, tcell.StyleDefault)
    }
}
