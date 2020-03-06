package birthday

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

const (
	timeISO = "2006-01-02"
	timeUS  = "January 2, 2006"
)

func WhenIsBirthday(birthday string, name string) string {
	daysDifference := calculateDaysDifference(birthday)
	message := ""
	if daysDifference == 0 {
		message = fmt.Sprintf("Hello, %s! Happy birthday!", name)
	} else if daysDifference <= 5 {
		message = fmt.Sprintf("Hello, %s! Your birthday is in 5 days!", name)
	} else {
		message = fmt.Sprintf("Hello, %s! Your birthday is in %d days.", name, daysDifference)
	}
	return message
}

func calculateDaysDifference(birthday string) int {
	birthdayDate, _ := time.Parse(timeISO, birthday)
	todayDate := time.Now()
	_, bmonth, bday := birthdayDate.Date()
	tyear, tmonth, tday := todayDate.Date()

	dayOfBirth := time.Date(tyear, bmonth, bday, 0, 0, 0, 0, time.UTC)
	dayOfToday := time.Date(tyear, tmonth, tday, 0, 0, 0, 0, time.UTC)
	//log.Printf("\ntodayDate: %v,\n\tday: %d\n\tmonth: %s\n", todayDate, tday, tmonth)
	//log.Printf("\nbirthdayDate: %v,\n\tday: %d\n\tmonth: %v\n", birthdayDate, bday, bmonth)

	daysDifference := math.Floor(dayOfBirth.Sub(dayOfToday).Hours() / 24)
	return int(daysDifference)
}

func CreateFile(p string) *os.File {
	fmt.Println("Logging to a file.")
	f, err := os.Create(p)
	// file, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// panic(err)
		log.Panic(err)
	}
	return f
}
func CloseFile(f *os.File) {
	fmt.Println("Closing file.")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
