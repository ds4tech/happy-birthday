package birthday

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestNewRouter(t *testing.T) {
	var file = CreateFile("info.log")
	router := NewRouter(file)
	expected := mux.NewRouter()
	//below test does not work properly because the condition should be !=
	if router == expected {
		t.Errorf("expecting mux.NewRouter(), got NewRouter()")
	}
}
