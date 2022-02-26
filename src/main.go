package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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
	var mode = "day" // Mode can be 'day' or 'shift'

	var hour = time.Now().Hour()
	var minute = time.Now().Minute()
	var length = 1440.

	var shiftL = 8.0

	// Check if --shift is in args
	if len(os.Args) >= 2 {
		if "--shift" == os.Args[1] {
			if len(os.Args) >= 3 {
				shiftL, _ = strconv.ParseFloat(os.Args[2], 64)
			}
			mode = "shift"
		} else {
			fmt.Println("Usage: timeleft [--shift [Shift length]]")
			os.Exit(1)
		}
	}

	if mode == "shift" {
		c, b := exec.Command("uptime"), new(strings.Builder)
		c.Stdout = b
		c.Run()

		var time = strings.Split(b.String(), " ")[1]

		hour, _ = strconv.Atoi(strings.Split(time, ":")[0])
		minute, _ = strconv.Atoi(strings.Split(time, ":")[1])

		length = (shiftL * 60)
	}

	var time = float64(((hour * 60) + minute))
	var per = float64(time / length)

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
