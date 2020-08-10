package main

import "testing"

func TestFind(t *testing.T) {
	menu := make([]string, 3)
	menu = append(menu, "Go")
	menu = append(menu, "Python")
	menu = append(menu, "Ruby")

	findGo := find("Go", menu)
	findPHP := find("PHP", menu)

	if findGo != true {
		t.Errorf("find function was incorrect, got: %t, want: %t.",
			findGo, true)
	}

	if findPHP != false {
		t.Errorf("find function was incorrect, got: %t, want: %t.",
			findGo, false)
	}
}
