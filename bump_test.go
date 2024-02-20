package bump

import (
	"fmt"
	"testing"
)

func TestBumps(t *testing.T) {
	s := "1.2.3"
	s2, err := BumpString(s, &Options{Part: "patch"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s2)
	if s2 != "1.2.4" {
		t.Fatal("Expected 1.2.4")
	}
	s = "1.2.3"
	s2, err = BumpString(s, &Options{Part: "minor"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s2)
	if s2 != "1.3.0" {
		t.Fatal("Expected 1.3.0, got", s2)
	}

	s = "1.2.3-alpha"
	s2, err = BumpString(s, &Options{Part: "patch"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s2)
	if s2 != "1.2.4" {
		t.Fatal("Expected 1.2.4, got", s2)
	}

	s = "1.2.3-alpha"
	s2, err = BumpString(s, &Options{Part: "minor", PreRelease: "alpha"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s2)
	if s2 != "1.3.0-alpha" {
		t.Fatal("Expected 1.3.0-alpha, got", s2)
	}

	s = "1.2.3-alpha"
	s2, err = BumpString(s, &Options{Part: "minor", PreRelease: "alpha", Metadata: "build123"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s2)
	if s2 != "1.3.0-alpha+build123" {
		t.Fatal("Expected 1.3.0-alpha+build123, got", s2)
	}

}
