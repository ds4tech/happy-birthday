package birthday

import (
	"testing"
	"time"
)

func TestBirthdayIsIn5Days(t *testing.T) {
	birthdayDate := time.Date(1992, 1, 14, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
  result := calculateDaysDifference(birthday)
  if result > 5 {
		t.Errorf("expecting anything lower than 5, got %v", result)
  }
}

func TestBirthdayIsNotIn5Days(t *testing.T) {
	birthdayDate := time.Date(1989, 6, 14, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
  result := calculateDaysDifference(birthday)
  if result <= 5 {
		t.Errorf("expecting anything higher than 5, got %v", result)
  }
}

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

func TestWhenIsBirthday(t *testing.T) {
	birthdayDate := time.Now()//Date(1992, 2, 14, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
	result :=  WhenIsBirthday(birthday, "John")
	expected := "Hello, John! Happy birthday!"
	//expected := "Hello, John! Your birthday is in 5 days!"
	if result != expected {
		t.Errorf("Expected: \n%v, got \n%v", expected, result)
  }
}
