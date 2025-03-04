package logic

import (
	"testing"
	"time"
)

// go test -run=TestGenOrderID
// go test -v
func TestGenOrderID(t *testing.T) {
	oid := genOrderID(time.Now())
	if len(oid) != 24 {
		t.Errorf("oid len expected 24, got: %d", len(oid))
	}
}
