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

func RepoUpdateMan(man HelloMan, name string) {
	fmt.Println("update action on the way")
	for _, it := range helloPeople {
		if it.Name == name {
			fmt.Println("found", name)
			it.DateOfBirth = man.DateOfBirth
			//todo
			//update object in the table
		}
	}
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
