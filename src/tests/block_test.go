package tests

import (
	"strings"
	"testing"
	"time"
)

func TestMineBlock(t *testing.T) {
	mined := bc.MineBlock("", false)
	time.Sleep(time.Second * 4)
	mined = bc.MineBlock("", false)
	if !strings.HasPrefix(mined.GetHash(), strings.Repeat("0", int(4))) {
		t.Errorf("Difficulty  was %s", mined.GetHash())
	}
}
