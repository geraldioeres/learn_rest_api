package encrypt

import "testing"

func TestAddition(t *testing.T) {
	if Addition(4, 5) != 9 {
		t.Error("4 + 5 harusnya 9")
	}

	if Addition(4, 1) != 5 {
		t.Error("4 + 1 harusnya 5")
	}

	if Addition(-4, 1) != 0 {
		t.Error("-4 + 1 harusnya 0")
	}
}
