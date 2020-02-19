package main

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Scoring implements Viewport
type Scoring struct {
	title              string
	card               *widgets.Paragraph
	selectionCursorPos int
}

func (self Scoring) Handler(e <-chan ui.Event) (Viewport, error) {
	event := <-e
	switch event.ID {
	case "<C-c>":
		return self, Quit{}
	case "<Enter>":
		return createSelection(self.selectionCursorPos), nil
	}
	return self, nil
}

func (self Scoring) Render() {
	ui.Render(self.card)
}

func Cpm(correctCharacters int, duration time.Duration) float64 {
	return 60.0 * float64(correctCharacters) / float64(duration.Seconds())
}

func Accuracy(correctCharacters int, typedCharacters int) float64 {
	return float64(correctCharacters) / float64(typedCharacters)
}

func CreateScoring(correctCharacters int, totalCharacters int, duration time.Duration, selectionCursorPos int) Scoring {
	cpm := Cpm(correctCharacters, duration)
	accuracy := Accuracy(correctCharacters, totalCharacters)
	card := widgets.NewParagraph()
	card.Title = "Scoring Card"
	card.Text = fmt.Sprintf("CPM: %.0f\nAccuracy: %.2f%%", cpm, 100.0*accuracy)
	card.SetRect(MainMinX, MainMinY, MainMaxX, MainMaxY)
	return Scoring{
		title:              "Scoring",
		card:               card,
		selectionCursorPos: selectionCursorPos,
	}
}