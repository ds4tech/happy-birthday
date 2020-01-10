package birthday

import (
	"testing"
	 "time"
)


func TestBirthdayIsIn5Days(t *testing.T) {
	birthdayDate := time.Date(1992, 1, 15, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
  result := IsBirthdayIn5Days(birthday)
	expecting := true
  if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
  }
}

func TestBirthdayIsNotIn5Days(t *testing.T) {
	birthdayDate := time.Date(1989, 6, 30, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
  result := IsBirthdayIn5Days(birthday)
	expecting := false
  if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
  }
}

func TestTodayIsYourBirthday(t *testing.T) {
	birthdayDate := time.Date(1999, 1, 10, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
  result := IsBirthdayIn5Days(birthday)
	expecting := true
  if result != expecting {
		t.Errorf("birthday is: %s", birthday)
		t.Errorf("expecting %v, got %v", expecting, result)
  }
}
