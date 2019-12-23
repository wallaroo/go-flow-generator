package flow_test

import (
	"testing"
	"os"
	
	"github.com/wallaroo/flowgenerator/flow"
)

func TestCreate(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log("wd", wd)
	err := flow.Create("test.flow.yaml", "test.flow.out", 10)
	if err != nil {
		t.Error(err)
	}
}
