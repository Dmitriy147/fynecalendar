# fynecalendar 

Моя версия Go Fyne Calendar widget. 
Основа кодовой базы взята из репозитория
https://github.com/fyne-io/fyne-x

<div align="center">
  <img src="https://github.com/user-attachments/assets/be2b99fd-5f55-4d7b-8df4-b8897641f321" alt="Описание изображения">
</div>

**Основные изменения:**
- локализация - russian|english
- указание интервала доступных дат для выбора (недоступные дни неактивны)
- выделение цветом выбранной даты


**HELLO WORLD**
```go
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/Dmitriy147/fynecalendar"
)

func main() {
	var range_start time.Time = time.Now().Add(-240 * time.Hour)
	var range_end time.Time = time.Now().Add(-24 * time.Hour)

	a := app.New()
	w := a.NewWindow("Календарь")
	w.CenterOnScreen()

	lbl := widget.NewLabel("")

	// create new calendar widget (rus/eng language -true/false, selected date, start active date interval, end active date interval)
	calendar_start := fynecalendar.NewMyCalendar(true, range_end, range_start, range_end, func(t time.Time) {
		lbl.Text = t.Format("selected: 02.01.2006")
		lbl.Refresh()
	})

	// create new calendar widget (rus/eng language -true/false, selected date, start active date interval, end active date interval)
	calendar_end := fynecalendar.NewMyCalendar(false, range_end, range_start, range_end, func(t time.Time) {
		lbl.Text = t.Format("selected: 02.01.2006")
		lbl.Refresh()
	})

	cal := container.NewHBox(calendar_start, calendar_end)

	c := container.NewVBox(cal, lbl)
	c.Refresh()
	w.SetContent(c)
	w.ShowAndRun()
}
```
Проверено:
- Go 1.21.3, 1.23.3
- Fyne 2.4.1, 2.6.1
