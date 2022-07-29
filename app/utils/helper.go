package utils

import (
	"os"
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)

	return regex.MatchString(email)
}

func DateFormat(layout string, d float64) string {
	intTime := int64(d)
	t := time.Unix(intTime, 0)
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return t.Format(layout)
}

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
