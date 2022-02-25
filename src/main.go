package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 80
)

func main() {
	var hour = time.Now().Hour()
	var minute = time.Now().Minute()
	var minutesInADay = 1440.00

	var time = float64(((hour * 60) + minute))
	var per = float64(time / minutesInADay)

	prog(int(time), per)
}

type model struct {
	curtime  int
	percent  float64
	progress progress.Model
}

func prog(curtime int, completion float64) {
	const (
		gradient1 = "#FF7CCB"
		gradient2 = "#AE7CFE"
	)

	prog := progress.New(progress.WithScaledGradient(gradient1, gradient2))

	if err := tea.NewProgram(model{curtime: curtime, percent: completion, progress: prog}).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

func (e model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" + pad + e.progress.ViewAs(e.percent) + "\n\n"
}

func (_ model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		return m, tea.Quit

	default:
		return m, nil
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
