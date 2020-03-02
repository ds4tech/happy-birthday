package birthday

import (
	"testing"
	"time"
)

func TestRepoCreateMan(t *testing.T) {
  var helloMan HelloMan
  man := RepoCreateMan(helloMan, "John")

  if man.Id != 1 {
		t.Errorf("expecting 1, got %d", man.Id)
  }
}

func TestRepoUpdateMan(t *testing.T) {
	birthdayDate := time.Now()//Date(1992, 2, 14, 0, 0, 0, 0, time.UTC)
	birthday := birthdayDate.Format(timeISO)
	result :=  RepoUpdateMan(birthday, "John")

  if result != true {
		t.Errorf("expecting true, got %v", result)
  }
}
