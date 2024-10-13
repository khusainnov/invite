package helpers

import "time"

func BuildDatePtr(date string) (*time.Time, error) {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return &time.Time{}, err
	}

	return &parsedDate, err
}
