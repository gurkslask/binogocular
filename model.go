package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// common "github.com/gurkslask/common-go-libs/common"
)
const (
	hotPink    = lipgloss.Color("#FF06B7")
	darkGray   = lipgloss.Color("#767676")
	ColorR1    = lipgloss.Color("#DC243C")
	ColorR2    = lipgloss.Color("#F75270")
	ColorR3    = lipgloss.Color("#F7CAC9")
	ColorR4    = lipgloss.Color("#FDEBD0")
	ColorBG1   = lipgloss.Color("#BADA55")
	ColorBG2   = lipgloss.Color("#FFFA55")
	ColorGreen = lipgloss.Color("#00FF11")
	ColorWhite = lipgloss.Color("#FFFFFF")
	ColorBlack = lipgloss.Color("#000000")
)
type (
	errMsg error
	)
type model struct {
	variables string
	program string
	inputs []textinput.Model
	focused int
	dynvars []dynvarbool
}
func InitialModel() model {
	var dynvars []dynvarbool = make([]dynvarbool, 3)
	dynvars = []dynvarbool {
			Newdynvarbool("kontaktor1", &kontaktor1),
			Newdynvarbool("startknapp1", &startknapp1),
			Newdynvarbool("stoppknapp1", &stoppknapp1),
	}
	var inputs []textinput.Model = make([]textinput.Model, len(dynvars))
	inputs[0] = textinput.New()
	vars, progs := readFile("./program.go")
	
	return model{
		variables: vars,
		program: progs,
		dynvars:dynvars,
		focused: 0,
	}

}
func (m model) Init() tea.Cmd {
	return textinput.Blink
}
func (m model) Update(msg tea.Msg)  ( tea.Model, tea.Cmd) {

	var cmds []tea.Cmd = make([]tea.Cmd , len(m.inputs))
		switch msg := msg.(type) {
	case tea.KeyMsg: 
	switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		case tea.KeySpace:
			m.dynvars[m.focused].Toggle()
		}
	}
	for k, _ := range m.dynvars {
		if k == m.focused {
			m.dynvars[k].active = true
		} else {
			m.dynvars[k].active = false
		}
	}
	m.dynvars[m.focused].active = true
	program()
	return m, tea.Batch(cmds...)
}
func (m model) View()  string{
	return fmt.Sprintf(`
--- VARIABLER ---
%s

--- PROGRAMMET ---
%s

%s
%s
%s
%d
		`, 
		m.variables,
		m.program,
		m.dynvars[0].Print(),
		m.dynvars[1].Print(),
		m.dynvars[2].Print(),
		m.focused,
	)
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.dynvars)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.dynvars) - 1
	}
}
