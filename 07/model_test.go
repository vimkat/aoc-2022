package main

import (
	"testing"
)

func TestFile(t *testing.T) {
	file := NewFile("f", 100)

	if file.Size() != 100 {
		t.Fatalf("file size is %d, expected 100", file.Size())
	}
}

func TestDirectory(t *testing.T) {
	oneLevel := NewDirectory(
		"d1",
		NewFile("f1", 10),
		NewFile("f2", 20),
		NewFile("f3", 30),
	)

	twoLevel := NewDirectory(
		"d2",
		NewFile("f4", 40),
		oneLevel,
	)

	sizeOne := oneLevel.Size()
	if sizeOne != 60 {
		t.Fatalf("size of one-level directory is %d, expected 60", sizeOne)
	}

	sizeTwo := twoLevel.Size()
	if sizeTwo != 100 {
		t.Fatalf("size of two-level directory is %d, expected 60", sizeTwo)
	}
}

func TestDirectoryAdd(t *testing.T) {
	dir := NewDirectory("/")
	dir.AddAt("", NewFile("test.txt", 10))
	dir.AddAt("etc", NewFile("foo.txt", 20))

	tmp := NewDirectory("tmp")

	tmp.AddAt("", NewFile("tmp1", 10))
	dir.Add(tmp)
	dir.AddAt("tmp", NewFile("tmp2", 10))

	dir.PrettyPrint()
}
