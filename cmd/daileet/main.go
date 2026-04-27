package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/0b10headedcalf/daileet/internal/storage"
	"github.com/0b10headedcalf/daileet/internal/tui"
)

func main() {
	db, err := storage.OpenDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	app := tui.NewApp(db)
	p := tea.NewProgram(app)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running program: %v\n", err)
		os.Exit(1)
	}
}
