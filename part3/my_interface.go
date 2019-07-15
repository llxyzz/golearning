package part3

import (
	"fmt"
	"time"
)

type DateTime struct {
	month                           time.Month
	year, day, hour, minute, second int
}

func getStr() string {
	s := "123"
	return fmt.Sprintf("this is %v\n", s)
}

type DataArray [3]string

type DateFormatter interface {
	DateFormat(t time.Time) string
}

func (d DataArray) String() string {
	return fmt.Sprintf("%v-%v-%v", d[0], d[1], d[2])
}

func (d DateTime) DateFormat(t time.Time) string {
	return fmt.Sprintf("%v-%02d-%02v %02v:%02v:%02v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
