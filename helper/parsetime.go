package helper

import "time"

const layout = "2006-02-01"

func ParseTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(layout, timeStr)
	return t, err
}

func ToString(t time.Time) string {
	str := t.Format(layout)
	return str
}
