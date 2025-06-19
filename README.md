# fynecalendar
Моя версия Go Fyne Calendar widget.
Основа кодовой базы взята из репозитория
https://github.com/fyne-io/fyne-x

<div align="center">
  <img src="https://github.com/user-attachments/assets/9ce5fa21-3350-4c87-a8dc-90a36e56c462" alt="Описание изображения">
</div>

Основные изменения:
- локализация - russian
- указание интервала доступных дат для выбора (недоступные дни неактивны)
- выделение цветом выбранной даты

**HELLO WORLD!**!
```go
package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Dmitriy147/fynecalendar"
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

	// NewCalendar создаёт виджет календаря и возвращает выбранную дату
        // (активная дата, начальная дата активного интервала, конечная дата активного интервала)
	calendar_start := fynecalendar.NewMyCalendar(range_end, range_start, range_end, func(t time.Time) {
		filt_label.Text = t.Format("Выбрана дата начала 02.01.2006")
		filt_label.Refresh()
	})

	calendar_end := fynecalendar.NewMyCalendar(range_end, range_start, range_end, func(t time.Time) {
		filt_label.Text = t.Format("Выбрана дата окончания 02.01.2006")
		filt_label.Refresh()
	})

	cal := container.NewHBox(calendar_start, calendar_end)

	c := container.NewVBox(cal, filt_label)
	c.Refresh()
	w.SetContent(c)
	w.ShowAndRun()
}
```

**ВНИМАНИЕ!!!**! При первой сборке бинарника возможно ожидание запуска до 5 мин (особенности Fyne) - не спешите с Ctrl+C.

Так как цвета виджета основаны на системной теме, то ,возможно, потребуется создание пользовательской темы для ваших задач.

Проверено: Go 1.21.3, 1.23.3; Fyne 2.4.1, 2.6.1
