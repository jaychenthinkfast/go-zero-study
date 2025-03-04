package logic

import "testing"

// go test -run=TestStockKey -v
func TestStockKey(t *testing.T) {
	tests := []struct {
		name   string
		input  uint64
		output string
	}{
		{"test one", 1, "stock:1"},
		{"test two", 2, "stock:2"},
		{"test three", 3, "stock:3"},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			ret := stockKey(ts.input)
			if ret != ts.output {
				t.Errorf("input: %d expectd: %s got: %s", ts.input, ts.output, ret)
			}
		})
	}
}

func TestStockKeyParallel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  uint64
		output string
	}{
		{"test one", 1, "stock:1"},
		{"test two", 2, "stock:2"},
		{"test three", 3, "stock:3"},
	}

	for _, ts := range tests {
		ts := ts
		t.Run(ts.name, func(t *testing.T) {
			t.Parallel()
			ret := stockKey(ts.input)
			if ret != ts.output {
				t.Errorf("input: %d expectd: %s got: %s", ts.input, ts.output, ret)
			}
		})
	}
}
