package birthday

import "fmt"

var currentId int
var helloPeople HelloPeople


func RepoCreateMan(t HelloMan, name string) HelloMan {
	currentId += 1
	t.Id = currentId
	t.Name = name
	helloPeople = append(helloPeople, t)
	return t
}

func RepoFindMan(name string) HelloMan {
	for _, t := range helloPeople {
		if t.Name == name {
			return t
		}
	}
	return HelloMan{}
}

func RepoDestroyMan(id int) error {
	for i, t := range helloPeople {
		if t.Id == id {
			helloPeople = append(helloPeople[:i], helloPeople[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find Man with provided name to delete")
}
