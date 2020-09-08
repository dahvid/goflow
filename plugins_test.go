package goflow

import (
	"testing"
	"fmt"
)


func TestBadParams(t *testing.T) {
	paths := []string{"boogie"}
	_,err := LoadComponents(paths, NewFactory())
	if err != nil {
		t.Error("LoadComponenets succeeded with bad parameters")
	}
	paths = []string{"/usr/lib","/lib"}
	_,err = LoadComponents(paths, NewFactory())
	if err != nil {
		t.Error("LoadComponenets succeeded with bad parameters")
	}
}

func TestOpening(t *testing.T) {
	paths := []string{"/home/david/go/src/github.com/dahvid/goflow/test_plugins"}
	plugs,err := LoadComponents(paths, NewFactory())
	if err != nil {
		t.Error("Failed loading compononents in ./test_plugins",err)
	}
	fmt.Println("Opened plugs",plugs)
}