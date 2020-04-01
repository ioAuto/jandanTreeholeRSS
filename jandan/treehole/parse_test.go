package treehole_test

import (
	"github.com/iochen/jandanTreeholeRSS/jandan/treehole"
	"testing"
	"time"
)

func TestGetFromID(t *testing.T) {
	th, err := treehole.GetFromID(treehole.ID(4516747))
	if err != nil {
		t.Error(err)
	}
	t.Log(th)
	t.Log(th.Comments)
	t.Log(time.Unix(50*4516747+1.359e9, 0))
}
