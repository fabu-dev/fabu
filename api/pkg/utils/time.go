package utils

import "time"

func GetCurrentTimeStamp() time.Time {
	location, _ := time.LoadLocation("Timezone")
	return time.Now().In(location)
}
