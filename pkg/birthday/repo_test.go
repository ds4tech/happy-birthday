package birthday

import "testing"

func TestRepoDestroyMan(t *testing.T) {
	name := "John"
	man := RepoCreateMan(name)
	result := RepoDestroyMan(1)
	//expected := "Could not find Man with provided name to delete."
	if result != nil {
		t.Errorf("Only %s name is in the database. Therefore expecting id: 1, got id: %d.", man.Name, man.Id)
	}
}

// func TestRepoUpdateMan(t *testing.T) {
// 	birthdayDate := time.Now() //Date(1992, 2, 14, 0, 0, 0, 0, time.UTC)
// 	birthday := birthdayDate.Format(timeISO)
// 	result := RepoUpdateMan(birthday, "John")

// 	if result != true {
// 		t.Errorf("expecting true, got %v", result)
// 	}
// }
