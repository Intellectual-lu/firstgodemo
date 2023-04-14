package mystruct

import (
	"errors"
	"unicode/utf8"
)

type MyStruct struct {
	Name   string
	Number float64
	Flag   bool
}

type Date struct {
	year  int
	month int
	day   int
}

type Event struct {
	title string
	Date
}

// SetYear set方法
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 9 {
		return errors.New("title must less five")
	}
	e.title = title
	return nil
}

// Year get方法
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (e *Event) Title() string {
	return e.title
}
