package main

import (
	"os"

	"github.com/kr/pretty"
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/grip"
	"github.com/sparkymat/grip/event"
	"github.com/sparkymat/grip/size"
	"github.com/sparkymat/sparkymud/config"
)

// OnEvent is the event handler for the app
func OnEvent(app *grip.App, e event.Event) {
	switch e.Type {
	case event.EventKeyPress:
		termboxEvent := e.Data.(termbox.Event)
		if termboxEvent.Type == termbox.EventKey {
			switch termboxEvent.Key {
			case termbox.KeyEsc:
				app.Confirm(
					"Are you sure you want to quit?",
					func(app *grip.App) {
						termbox.Close()
						os.Exit(0)
					},
					func(app *grip.App) {},
				)
			}
		}
	}
}

func main() {
	appConfig := config.Load()

	app := grip.App{}

	mainGrid := grip.Grid{
		ColumnSizes:     []size.Size{size.Auto},
		RowSizes:        []size.Size{size.Auto, size.WithPercent(30)},
		HasBackground:   true,
		BackgroundColor: termbox.ColorDefault,
	}

	mainGrid.AddView("title-text", &grip.TextView{
		Text:            "SparkyMUD",
		ForegroundColor: termbox.ColorRed,
		BackgroundColor: termbox.ColorDefault,
		TextAlignment:   grip.TextAlignmentCenter,
	}, grip.Area{0, 0, 0, 0})

	app.View = &mainGrid
	app.RegisterEventListener(event.EventKeyPress, OnEvent)

	pretty.Log(appConfig)

	app.Run()
}
