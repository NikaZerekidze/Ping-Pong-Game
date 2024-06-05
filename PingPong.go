package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"os"
	"sync"
)

const (
	PaddleSymbol = 0x2588
	PaddleHeight = 4
)

type Paddle struct {
	x, y, width, height int
}

var (
	paddle1 *Paddle
	paddle2 *Paddle
)

var mutex sync.M
utex

// / this is not written by me this comes with tcell library
func emitStr(screen tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		screen.SetContent(x, y, c, comb, style)
		x += w
	}
}

// / here paddles are paddle object is created
func PrintGameObjectPaddles(screen tcell.Screen, x, y, width, height int, character rune) {

	for col := 0; col < width; col++ {
		for row := 0; row < height; row++ {
			screen.SetContent(x+col, y+row, character, nil, tcell.StyleDefault)

		}
	}

}

// / this is main manu screen where welcome text is located
func drawWelcomeSceen(screen tcell.Screen) {
	w, h := screen.Size()
	screen.Clear()
	welcomeText := "Welcome to my ping pong game"
	exitText := "Press Enter to exit"
	playSomeoneText := "Press Enter to play with someone"

	welcomeTextPosition := (w - len(welcomeText)) / 2
	playSomeoneTextPosition := (w - len(playSomeoneText)) / 2
	exitTextPosition := (w - len(exitText)) / 2

	style := tcell.StyleDefault.Foreground(tcell.Color30).Background(tcell.ColorBlack)
	emitStr(screen, welcomeTextPosition, h/2, style, welcomeText)
	emitStr(screen, exitTextPosition, h/2+1, tcell.StyleDefault, exitText)
	emitStr(screen, playSomeoneTextPosition, h/2+2, tcell.StyleDefault, playSomeoneText)

	screen.Show()
}

// / this is function to display game objects (game screen) paddles, ball
func drawGameScreen(screen tcell.Screen) {
	screen.Clear()

	PrintGameObjectPaddles(screen, paddle1.x, paddle1.y, paddle1.width, paddle1.height, PaddleSymbol)
	PrintGameObjectPaddles(screen, paddle2.x, paddle2.y, paddle2.width, paddle2.height, PaddleSymbol)

	screen.Show()
}

func main() {

	screen := InintScreen()

	drawWelcomeSceen(screen)
	InitGameState(screen)

	isWelcomeScreen := true

	for {

		if isWelcomeScreen {
			drawWelcomeSceen(screen)
		} else {
			drawGameScreen(screen)
		}

		switch ev := screen.PollEvent().(type) {

		case *tcell.EventKey:

			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
			if ev.Key() == tcell.KeyEnter {

				if isWelcomeScreen {
					drawGameScreen(screen)
					isWelcomeScreen = false

				} else {
					screen.Beep()
					drawWelcomeSceen(screen)
					isWelcomeScreen = true
					positionPaddles(screen)
				}
			}
			if ev.Key() == tcell.KeyUp {
				paddle2.y--
			}
			if ev.Key() == tcell.KeyDown {
				paddle2.y++

			}
			if ev.Rune() == 'w' {
				paddle1.y--

			}
			if ev.Rune() == 's' {
				paddle1.y++

			}

		case *tcell.EventResize:
			if isWelcomeScreen {
				drawWelcomeSceen(screen)
			} else {
				positionPaddles(screen) /// when sceen size is changed I need to call screen again to update widht and height
				drawGameScreen(screen)
			}
		}

	}
}

func InitUserInput(screen tcell.Screen) {
	for {
		switch ev := screen.PollEvent().(type) {

		case *tcell.EventKey:

			if ev.Key() == tcell.KeyUp {
				paddle2.y--
			}
			if ev.Key() == tcell.KeyDown {
				paddle2.y++

			}
			if ev.Rune() == 'w' {
				paddle1.y--

			}
			if ev.Rune() == 's' {
				paddle1.y++

			}
		}
	}
}

// / this is for creating screen in general
func InintScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	return screen
}

// / created this separate function because I did not want to call InitGameState every time I wanted to re-render object (this is mainly for resizing)
func positionPaddles(screen tcell.Screen) {
	width, height := screen.Size()
	paddleStart := height/2 - PaddleHeight/2

	paddle1.x = 0
	paddle1.y = paddleStart

	paddle2.x = width - 1
	paddle2.y = paddleStart
}

// / in here I am creating initializing game objects for displaying game object function
func InitGameState(screen tcell.Screen) {
	width, height := screen.Size()
	paddleStart := height/2 - PaddleHeight/2

	paddle1 = &Paddle{
		x: 0, y: paddleStart, width: 1, height: PaddleHeight,
	}

	paddle2 = &Paddle{
		x: width - 1, y: paddleStart, width: 1, height: PaddleHeight,
	}

}
