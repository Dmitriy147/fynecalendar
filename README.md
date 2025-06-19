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


```go
import "github.com/Dmitriy147/fynecalendar"
...
// NewCalendar создаёт виджет календаря и возвращает выбранную дату
// (активная дата, начальная дата активного интервала, конечная дата активного интервала)
calendar := fynecalendar.NewMyCalendar(current_date, range_start, range_end, func(t time.Time) {
    label1.Text = t.Format("Выбрана дата 02.01.2006")
})
```

**ВНИМАНИЕ!!!**! При первой сборке бинарника возможно ожидание запуска до 5 мин (особенности Fyne) - не спешите с Ctrl+C.

Так как цвета виджета основаны на системной теме, то ,возможно, потребуется создание пользовательской темы для ваших задач.

Проверено: Go 1.21.3, 1.23.3; Fyne 2.4.1, 2.6.1
