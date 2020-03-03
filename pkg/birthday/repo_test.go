package birthday

import "testing"

func TestRepoDestroyMan(t *testing.T) {
	// RepoCreateMan(name, birthday)
	result := RepoDestroyMan(1)
	if result == nil {
		t.Errorf("Expecting nil object, got: %v", result)
	}
}

// func TestRepoUpdateMan(t *testing.T) {
// 	name := "John"
// 	birthday := `{ "dateOfBirth": "2001-01-10" }`
// 	result := RepoUpdateMan(name, birthday)

// 	if result != true {
// 		t.Errorf("expecting true, got %v", result)
// 	}
// }
