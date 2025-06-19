package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var range_start time.Time = time.Now().Add(-240 * time.Hour)
	var range_end time.Time = time.Now().Add(-24 * time.Hour)
	filt_label := widget.NewLabel("===")

	// New app
	a := app.New()
	w := a.NewWindow("Календарь")
	w.Resize(fyne.NewSize(500, 400))
	w.CenterOnScreen()

	///////////// CALENDAR ////////////////////////////////////
	calendar_start := NewMyCalendar(range_end, range_start, range_end, func(t time.Time) {
		filt_label.Text = t.Format("Выбрана дата начала 02.01.2006")
		filt_label.Refresh()
	})

	calendar_end := NewMyCalendar(range_end, range_start, range_end, func(t time.Time) {
		filt_label.Text = t.Format("Выбрана дата окончания 02.01.2006")
		filt_label.Refresh()
	})

	cal := container.NewHBox(calendar_start, calendar_end)

	c := container.NewVBox(cal, filt_label)
	c.Refresh()
	w.SetContent(c)
	w.ShowAndRun()
}
