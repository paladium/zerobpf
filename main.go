package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

/*
CPUs: sockets, cores, hardware threads (virtual CPUs)
Memory: capacity
Network interfaces
Storage devices: I/O, capacity
Controllers: storage, network cards
*/

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			updateCPUMetrics(textView)
		}
	}()
	flex := tview.NewFlex().
		AddItem(textView, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Memory"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Storage"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Network"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Misc"), 20, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func updateCPUMetrics(textView *tview.TextView) {
	fmt.Fprintf(textView, "hello world")
	// Get
}
