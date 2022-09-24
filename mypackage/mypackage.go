package mypackage

import (
	"errors"
	"time"
)

type MyDate struct {
	day   int
	month int
	year  int
}

func (d *MyDate) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("Incorrect day value")
	}
	d.day = day
	return nil
}
func (d *MyDate) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New("Incorrect month value")
	}
	d.month = month
	return nil
}
func (d *MyDate) SetYear(year int) error {
	if (year < 1875) || (year > time.Now().Year()) {
		return errors.New("Incorrect Year value")
	}
	d.year = year
	return nil
}

func (d *MyDate) Day() int {
	return d.day
}
func (d *MyDate) Month() int {
	return d.month
}
func (d *MyDate) Year() int {
	return d.year
}
