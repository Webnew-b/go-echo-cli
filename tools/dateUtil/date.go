package dateUtil

import "time"

func FormatTicketDate(date time.Time) string {
	newFormat := "January 2, 2006"
	formattedDate := date.Format(newFormat)
	return formattedDate
}

func ParseTicketDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
