package ini

import (
	"fmt"
	"testing"
)

func TestIniParser(t *testing.T) {
	ini :=
		`
Version = 1.0
[global]
Author=Kevin
`
	cfg := ParseString(ini)
	if cfg["global"]["Author"] != "Kevin" {
		t.Fatal("Expected Author=Kevin, got", cfg["global"]["Author"])
	}
	if cfg[""]["Version"] != "1.0" {
		t.Fatal("Expected Version=1.0, got", cfg[""]["Version"] != "1.0")
	}
}

func TestIniString(t *testing.T) {
	ini :=
		`
Version = 1.0
[global]
Author=Kevin
`
	cfg := ParseString(ini)
	fmt.Print(cfg)
}
