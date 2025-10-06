package main

import (
	"github.com/charmbracelet/lipgloss"
	"fmt"
)
type line struct {
	text string
	variables []dynvarbool
}
func NewLine(l string)

type dynvarbool struct {
	name string
	truestyle lipgloss.Style
	falsestyle lipgloss.Style
	activestyle lipgloss.Style
	pp *bool
	active bool
}
func (d dynvarbool) Print() string {
	data := fmt.Sprintf("%s: %s", d.name, boolToString(*d.pp))
	if d.active {
		return d.activestyle.Render(data)
	} else if *d.pp {
		return d.truestyle.Render(data)
	} else if !*d.pp {
		return d.falsestyle.Render(data)
	} 
	return d.activestyle.Render("undefined")
}
func (d dynvarbool) Toggle() {
	*d.pp = !*d.pp
}

func Newdynvarbool(n string, varpointer *bool) dynvarbool {
	return dynvarbool {
		truestyle: lipgloss.NewStyle().Foreground(lipgloss.Color("BADA55")).Background(ColorGreen),
		falsestyle: lipgloss.NewStyle().Foreground(lipgloss.Color("00ff55")).Background(ColorR1),
		activestyle: lipgloss.NewStyle().Foreground(lipgloss.Color("ff0055")).Background(hotPink),
		name: n,
		pp: varpointer,
	}
}

func boolToString(b bool)string{
	if b {
		return "True"
	}else {
		return "False"
	}
}
