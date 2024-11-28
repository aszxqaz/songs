package date

import "time"

func TimeToDMY(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02.01.2006")
}

func DmyToTime(dmy string) time.Time {
	if dmy == "" {
		return time.Time{}
	}
	t, _ := time.Parse("02.01.2006", dmy)
	return t
}

func DmyNow() string {
	return TimeToDMY(time.Now())
}

func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}
