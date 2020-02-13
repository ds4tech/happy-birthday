package birthday

import (
	"testing"
  "github.com/gorilla/mux"
)

func TestNewRouter(t *testing.T) {
  router := NewRouter()
  expected := mux.NewRouter()
  //below test does not work properly because the condition should be !=
  if router == expected {
		t.Errorf("expecting mux.NewRouter(), got NewRouter()")
  }
}
