package birthday

import "fmt"

var currentId int
var helloPeople HelloPeople


func RepoCreateMan(man HelloMan, name string) HelloMan {
	currentId += 1
	man.Id = currentId
	man.Name = name
	helloPeople = append(helloPeople, man)
	return man
}

func RepoUpdateMan(birthday string, name string) {
	for i, it := range helloPeople {
		if it.Name == name {
			helloPeople[i].DateOfBirth = birthday
		}
	}
}

func RepoFindMan(name string) HelloMan {
	for _, man := range helloPeople {
		if man.Name == name {
			return man
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
