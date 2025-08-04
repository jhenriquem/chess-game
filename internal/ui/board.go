package ui

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func inverterBoard(arr [8][8]string) [8][8]string {
	var board [8][8]string
	start := 0
	end := 7

	for start < end {
		board[start], board[end] = arr[end], arr[start]
		start++
		end--
	}

	for i, line := range board {

		start = 0
		end = 7

		for range line {
			if start < end {
				board[i][start], board[i][end] = board[i][end], board[i][start]
				start++
				end--
			}
		}
	}

	return board
}

var re = lipgloss.NewRenderer(os.Stdout)

var icons = map[string]string{
	"r": "󰡛", "n": "󰡘", "b": "󰡜",
	"q": "󰡚", "k": "󰡗", "p": "󰡙",
}

var labelStyle = re.NewStyle().Foreground(lipgloss.Color("241"))

var (
	whiteSquare = lipgloss.NewStyle().Background(lipgloss.Color("#edd6b0"))
	blackSquare = lipgloss.NewStyle().Background(lipgloss.Color("#b88762"))
)

func styledPiece(piece string, row, col int) string {
	icon := icons[strings.ToLower(piece)]
	style := lipgloss.NewStyle().Width(3).Height(1).Align(lipgloss.Center)

	bgStyle := whiteSquare
	if (row+col)%2 == 0 {
		bgStyle = blackSquare
	}

	if unicode.IsUpper(rune(piece[0])) {
		style = style.Foreground(lipgloss.Color("#FAFAFA"))
	} else {
		style = style.Foreground(lipgloss.Color("#000000"))
	}

	return style.Background(bgStyle.GetBackground()).Render(icon)
}

func buildStyledBoard(board [8][8]string) [][]string {
	styled := make([][]string, 8)
	for i := range board {
		styled[i] = make([]string, 8)
		for j := range board[i] {
			p := board[i][j]
			if p == " " {
				p = " "
			}
			styled[i][j] = styledPiece(p, i, j)
		}
	}
	return styled
}

func invert(arr []string) []string {
	start := 0
	end := 7

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}

func Board(gameBoard [8][8]string, color string) {
	board := gameBoard

	columnsFields := []string{" A", "B", "C", "D", "E", "F", "G", "H "}
	rowsFields := []string{"  8", "  7", "  6", "  5", "  4", "  3", "  2", "  1"}

	if color == "B" {
		columnsFields = invert(columnsFields)
		rowsFields = invert(rowsFields)
		board = inverterBoard(board)
	}

	styledBoard := buildStyledBoard(board)

	t := table.New().StyleFunc(func(_, _ int) lipgloss.Style {
		return lipgloss.NewStyle()
	}).Rows(styledBoard...).
		BorderColumn(false).
		BorderRow(false)

	columns := labelStyle.Render(strings.Join(columnsFields, "  "))
	rows := labelStyle.Render(strings.Join(rowsFields, "\n"))

	fmt.Println(
		lipgloss.JoinVertical(
			lipgloss.Right,
			lipgloss.JoinHorizontal(lipgloss.Center, rows, t.Render()),
			columns,
		) + "\n")
}
