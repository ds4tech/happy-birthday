package birthday

type HelloMan struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"dateOfBirth"`
}

type HelloPeople []HelloMan
