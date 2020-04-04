package treehole_test

import (
	"github.com/iochen/jandanTreeholeRSS/jandan/treehole"
	"testing"
)

func TestGetFromID(t *testing.T) {
	th, err := treehole.GetFromID(treehole.ID(4516747))
	if err != nil {
		t.Error(err)
	}
	t.Log(th)
	t.Log(th.Comments)
}
