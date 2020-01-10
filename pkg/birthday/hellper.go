package birthday

import (
	"time"
	"math"
	 "log"
)

const (
    timeISO = "2006-01-02"
    timeUS  = "January 2, 2006"
)

func IsBirthdayIn5Days(birthday string) bool {
	birthdayDate, _ := time.Parse(timeISO, birthday)
	todayDate := time.Now()
	_ , bmonth, bday := birthdayDate.Date()
	tyear , tmonth, tday := todayDate.Date()

	dayOfBirth := time.Date(tyear, bmonth, bday, 0, 0, 0, 0, time.UTC)
	dayOfToday := time.Date(tyear, tmonth, tday, 0, 0, 0, 0, time.UTC)
	//log.Printf("\ntodayDate: %v,\n\tday: %d\n\tmonth: %s\n", todayDate, tday, tmonth)
	//log.Printf("\nbirthdayDate: %v,\n\tday: %d\n\tmonth: %v\n", birthdayDate, bday, bmonth)

	daysDifference := math.Floor(dayOfBirth.Sub(dayOfToday).Hours() / 24)

	if( daysDifference <=5 ) {
		return true
	} else {
		log.Printf("dayOfBirth is in %v days\n", daysDifference)
		return false
	}
}
