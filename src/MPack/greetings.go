package MPack

import "time"

func MDay(name string) string {
	return "good day " + name
}

func MNight(name string) string {
	return "good night " + name
}

func IsAm() bool {
	now := time.Now()
	return now.Hour() <= 12
}

func IsPm() bool {
	now := time.Now()
	return now.Hour() >= 12
}
