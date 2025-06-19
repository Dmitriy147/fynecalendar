package fynecalendar

import (
	"math"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Layout = (*calendarLayout)(nil)

const daysPerWeek int = 7

type calendarLayout struct {
	cellSize float32
}

func newCalendarLayout() fyne.Layout {
	return &calendarLayout{}
}

// Позицируем (top or left)
func (g *calendarLayout) getLeading(offset int) float32 {
	ret := (g.cellSize) * float32(offset)

	return float32(math.Round(float64(ret)))
}

// Позицируем (bottom or right)
func (g *calendarLayout) getTrailing(offset int) float32 {
	return g.getLeading(offset + 1)
}

// Layout is called to pack all child objects into a specified size.
func (g *calendarLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	g.cellSize = size.Width / float32(daysPerWeek)
	row, col := 0, 0
	i := 0
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		x1 := g.getLeading(col)
		y1 := g.getLeading(row)
		x2 := g.getTrailing(col)
		y2 := g.getTrailing(row)

		child.Move(fyne.NewPos(x1, y1))
		child.Resize(fyne.NewSize(x2-x1, y2-y1))

		if (i+1)%daysPerWeek == 0 {
			row++
			col = 0
		} else {
			col++
		}
		i++
	}
}

// Минимальный размер виджета
func (g *calendarLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(250, 250)
}

// Calendar creates a new date time picker which returns a time object
type Calendar struct {
	widget.BaseWidget
	currentTime time.Time
	startRange  time.Time
	endRange    time.Time

	monthPrevious *widget.Button
	monthNext     *widget.Button
	monthLabel    *widget.Label

	dates *fyne.Container

	onSelected   func(time.Time)
	selectedDate time.Time
}

func (c *Calendar) daysOfMonth() []fyne.CanvasObject {
	start := time.Date(c.currentTime.Year(), c.currentTime.Month(), 1, 0, 0, 0, 0, c.currentTime.Location())
	buttons := []fyne.CanvasObject{}

	//account for Go time pkg starting on sunday at index 0
	dayIndex := int(start.Weekday())
	if dayIndex == 0 {
		dayIndex += daysPerWeek
	}

	//add spacers if week doesn't start on Monday
	for i := 0; i < dayIndex-1; i++ {
		buttons = append(buttons, layout.NewSpacer())
	}

	for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {

		dayNum := d.Day()
		s := strconv.Itoa(dayNum)
		b := widget.NewButton(s, func() {
			// даём выбрать дату если она в допущенном интервале
			if !c.dateForButton(dayNum).Before(c.startRange) && !c.dateForButton(dayNum).After(c.endRange) {
				c.selectedDate = c.dateForButton(dayNum)
				c.onSelected(c.selectedDate)
				c.Refresh()
			}
		})

		// подсвечиваем даты выбранная и из интервала допущенных дат
		if d.Equal(time.Date(c.selectedDate.Year(), c.selectedDate.Month(), c.selectedDate.Day(), 0, 0, 0, 0, c.selectedDate.Location())) {
			b.Importance = widget.HighImportance
		} else {
			b.Importance = widget.MediumImportance
		}

		// серым цветом даты не в допущенном интервале дат
		if d.Before(c.startRange) || d.After(c.endRange) {
			b.Disable() //widget.LowImportance
		}

		buttons = append(buttons, b)
	}

	return buttons
}

func (c *Calendar) dateForButton(dayNum int) time.Time {
	oldName, off := c.currentTime.Zone()
	return time.Date(c.currentTime.Year(), c.currentTime.Month(), dayNum, c.currentTime.Hour(), c.currentTime.Minute(), 0, 0, time.FixedZone(oldName, off)).In(c.currentTime.Location())
}

// название месяцев
func (c *Calendar) monthYear() string {

	mon := c.currentTime.Month().String()[:3]
	switch mon {
	case "Jan":
		mon = "Январь"
	case "Feb":
		mon = "Февраль"
	case "Mar":
		mon = "Март"
	case "Apr":
		mon = "Апрель"
	case "May":
		mon = "Май"
	case "Jun":
		mon = "Июнь"
	case "Jul":
		mon = "Июль"
	case "Aug":
		mon = "Август"
	case "Sep":
		mon = "Сентябрь"
	case "Oct":
		mon = "Октябрь"
	case "Nov":
		mon = "Ноябрь"
	case "Dec":
		mon = "Декабрь"
	}
	// заголовок месяца и года
	// return c.currentTime.Format("January 2006")
	return mon + "  " + strconv.Itoa(c.currentTime.Year())
}

// дни недели
func (c *Calendar) calendarObjects() []fyne.CanvasObject {
	columnHeadings := []fyne.CanvasObject{}
	for i := 0; i < daysPerWeek; i++ {
		j := i + 1
		if j == daysPerWeek {
			j = 0
		}
		t := widget.NewLabel("-")
		wek := time.Weekday(j).String()[:3]
		switch wek {
		case "Mon":
			t = widget.NewLabel("Пон")
		case "Tue":
			t = widget.NewLabel("Втр")
		case "Wed":
			t = widget.NewLabel("Срд")
		case "Thu":
			t = widget.NewLabel("Чтв")
		case "Fri":
			t = widget.NewLabel("Птн")
		case "Sat":
			t = widget.NewLabel("Суб")
		case "Sun":
			t = widget.NewLabel("Вск")
		}

		t.Alignment = fyne.TextAlignCenter
		columnHeadings = append(columnHeadings, t)
	}
	columnHeadings = append(columnHeadings, c.daysOfMonth()...)

	return columnHeadings
}

// CreateRenderer returns a new WidgetRenderer for this widget.
func (c *Calendar) CreateRenderer() fyne.WidgetRenderer {
	c.monthPrevious = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		c.currentTime = c.currentTime.AddDate(0, -1, 0)
		// Dates are 'normalised', forcing date to start from the start of the month ensures move from March to February
		c.currentTime = time.Date(c.currentTime.Year(), c.currentTime.Month(), 1, 0, 0, 0, 0, c.currentTime.Location())
		c.monthLabel.SetText(c.monthYear())
		c.dates.Objects = c.calendarObjects()
	})
	c.monthPrevious.Importance = widget.LowImportance

	c.monthNext = widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		c.currentTime = c.currentTime.AddDate(0, 1, 0)
		c.monthLabel.SetText(c.monthYear())
		c.dates.Objects = c.calendarObjects()
	})
	c.monthNext.Importance = widget.LowImportance

	c.monthLabel = widget.NewLabel(c.monthYear())

	nav := container.New(layout.NewBorderLayout(nil, nil, c.monthPrevious, c.monthNext),
		c.monthPrevious, c.monthNext, container.NewCenter(c.monthLabel))

	c.dates = container.New(newCalendarLayout(), c.calendarObjects()...)

	dateContainer := container.NewVBox(nav, c.dates)

	return widget.NewSimpleRenderer(dateContainer)
}

// NewCalendar создаёт виджет календаря (активная дата, начальная дата активного интервала, конечная дата активного интервала)
func NewMyCalendar(currentTm time.Time, startTm time.Time, endTm time.Time, onSelected func(time.Time)) *Calendar {
	c := &Calendar{
		currentTime: currentTm,  // стартовая дата
		onSelected:  onSelected, // функция выбора даты
		startRange:  startTm,    // начало интервала доступных дат
		endRange:    endTm,      // конец интервала доступных дат
	}
	c.selectedDate = c.currentTime // назначаем выбранную дату при создании календаря чтобы подсветить кнопку
	c.ExtendBaseWidget(c)

	return c
}

// Перерисовка кнопок с днями при выборе
func (c *Calendar) Refresh() {
	c.dates.Objects = c.calendarObjects()
}
