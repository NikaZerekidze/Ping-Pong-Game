package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"os"
)

//// This is just starter code copied from tcell package demo folder. For rendering simple screen

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func displayWelcomeSceen(screen tcell.Screen) {
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

func main() {

	screen := InintScreen()

	displayWelcomeSceen(screen)

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
			if ev.Key() == tcell.KeyEnter {
				screen.Beep()

			}

		}
	}
}

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
