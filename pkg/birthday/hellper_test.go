package birthday

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// testing WhenIsBirthday
func TestWhenIsBirthday(t *testing.T) {
	birthdayDate := time.Now() //Date(1992, 2, 14, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
	result := WhenIsBirthday(birthday, "John")
	expected := "Hello, John! Happy birthday!"
	//if result != expected {
	if strings.Compare(result, expected) != 0 {
		t.Errorf("Expected: \n%v, got \n%v", expected, result)
	}
}

func TestBirthdayIsIn5Days(t *testing.T) {
	birthdayDate := time.Date(1992, 2, 14, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
	name := "John"
	result := WhenIsBirthday(birthday, name)
	expected := fmt.Sprintf("Hello, %s! Your birthday is in 5 days!", name)
	if result != expected {
		t.Errorf("Expected: \n%v\n, got \n%v", expected, result)
	}
}

func TestBirthdayIsNotIn5Days(t *testing.T) {
	birthdayDate := time.Date(1989, 6, 14, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
	daysDifference := calculateDaysDifference(birthday)
	name := "John"
	result := WhenIsBirthday(birthday, name)
	expected := fmt.Sprintf("Hello, %s! Your birthday is in %d days.", name, daysDifference)
	if result != expected {
		t.Errorf("Expected: \n%v\n, got \n%v", expected, result)
	}
}

// testing calculateDaysDifference
func TestTodayIsYourBirthday(t *testing.T) {
	birthdayDate := time.Now()
	birthday := birthdayDate.Format(timeISO)
	result := calculateDaysDifference(birthday)
	expecting := 0
	if result != expecting {
		t.Errorf("birthday is: %s", birthday)
		t.Errorf("expecting %v, got %v", expecting, result)
	}
}
