package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	// "github.com/spf13/cobra"
	"journal-cli/uinp"
)

func main() {
	now := time.Now()
	fmt.Printf("%v %v %v %v %v:%v:%v\n", now.Weekday().String(), now.Day(), now.Month().String(), now.Year(), now.Hour(), now.Minute(), now.Second())
	p := tea.NewProgram(uinp.Model{})
	_, err := p.Run()
	if err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}

	file_name := now.Format(time.RFC3339)
	_, err = os.Create(file_name)
	if err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}
