package ui

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/corentings/chess/v2"
	"github.com/gdamore/tcell/v2"
)

var (
	lightSquare = tcell.NewRGBColor(240, 217, 181)
	darkSquare  = tcell.NewRGBColor(181, 136, 99)
)

func InvesionBoard(ranks []string) []string {
	var board []string = []string{"", "", "", "", "", "", "", ""}
	start, end := 0, 7

	for rankI, rank := range ranks {
		line := ""
		for charI := range rank {
			line += string(rank[len(rank)-1-charI])
		}
		ranks[rankI] = line
	}

	for start < end {
		board[start], board[end] = ranks[end], ranks[start]

		start++
		end--
	}

	return board
}

func RenderBoard(fen string, color chess.Color) {
	light := tcell.StyleDefault.Background(lightSquare)
	dark := tcell.StyleDefault.Background(darkSquare)
	base := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)

	coordsFiles := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	ranks := strings.Split(fen, "/")

	if color == chess.Black {
		light, dark = dark, light

		ranks = InvesionBoard(strings.Split(fen, "/"))
		coordsFiles = []rune{'H', 'G', 'F', 'E', 'D', 'C', 'B', 'A'}
	}

	if len(ranks) != 8 {
		return
	}

	startX, startY := 4, 5

	for y, rank := range ranks {
		x := 1

		num := strconv.Itoa(8 - y)
		if color == chess.Black {
			num = strconv.Itoa(y + 1)
		}
		screen.SetContent(startX-2, startY+y, rune(num[0]), nil, base)

		for _, c := range rank {
			if c >= '1' && c <= '8' {
				empty := int(c - '0')
				for i := 0; i < empty; i++ {
					drawSquare(startX+x*3, startY+y, ' ', x, y, light, dark)
					x++
				}
			} else {
				drawSquare(startX+x*3, startY+y, c, x, y, light, dark)
				x++
			}
		}
	}

	for i, col := range coordsFiles {
		colX := startX + i*3 + 3
		screen.SetContent(colX, startY+9, col, nil, base)
	}

	screen.Show()
}

func ReturnIcon(piece rune) rune {
	icons := map[string]rune{
		"r": '󰡛',
		"b": '󰡜',
		"n": '󰡘',
		"k": '󰡗',
		"q": '󰡚',
		"p": '󰡙',
	}
	if val, ok := icons[strings.ToLower(string(piece))]; ok {
		return val
	}
	return piece
}

func drawSquare(baseX, baseY int, piece rune, file, rank int, light, dark tcell.Style) {
	style := light.Foreground(tcell.ColorWhite)

	if (file+rank)%2 == 1 {
		style = dark.Foreground(tcell.ColorWhite)
	}

	if unicode.IsLower(piece) {
		style = style.Foreground(tcell.ColorBlack)
	}

	char := ' '
	if piece != ' ' {
		char = ReturnIcon(piece)
	}

	screen.SetContent(baseX-1, baseY, ' ', nil, style)
	screen.SetContent(baseX, baseY, char, nil, style)
	screen.SetContent(baseX+1, baseY, ' ', nil, style)
}
