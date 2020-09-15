package goflow

import (
	"testing"
	"path"
	"fmt"
	"os"
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
	test_plugs := path.Join(os.Getenv("GOPATH"),"bin")
	paths := []string{test_plugs}
	factory := NewFactory()
	plugs,err := LoadComponents(paths, factory)
	if err != nil {
		t.Error("Failed loading compononents in ./test_plugins",err)
	}
	fmt.Println("Opened plugs",plugs)
	any,err := factory.Create("Plug1")
	if err != nil {
		t.Error("Failed to create object Plug1",err)
	}
	plug1 := any.(PlugIn)
	fmt.Println(plug1.Info())


}