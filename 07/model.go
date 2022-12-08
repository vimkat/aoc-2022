package main

import (
	"fmt"
	"strings"
)

type FSNode interface {
	Name() string
	Size() int
	Parent() *Directory
	setParent(*Directory)
}

func IsFile(node FSNode) bool {
	_, isFile := node.(*File)
	return isFile
}

func IsDirectory(node FSNode) bool {
	_, isDir := node.(*Directory)
	return isDir
}

type Directory struct {
	name     string
	children map[string]FSNode
	parent   *Directory
}

func NewDirectory(name string, nodes ...FSNode) *Directory {
	dir := Directory{
		name:     name,
		children: make(map[string]FSNode),
		parent:   nil,
	}

	for _, node := range nodes {
		dir.children[node.Name()] = node
	}

	return &dir
}

func (d Directory) String() string {
	return fmt.Sprintf("%s (dir, size=%d)", d.name, d.Size())
}

func (d Directory) PrettyPrint() {
	fmt.Printf("- %v\n", d)
	d.prettyPrint("")
}

func (d Directory) prettyPrint(prefix string) {
	for _, node := range d.children {
		fmt.Printf("%v  - %v\n", prefix, node)
		if dir, isDir := node.(*Directory); isDir {
			dir.prettyPrint(prefix + "  ")
		}
	}
}

func (d Directory) Name() string       { return d.name }
func (d Directory) Parent() *Directory { return d.parent }

func (d Directory) Size() int {
	size := 0
	for _, node := range d.children {
		size += node.Size()
	}
	return size
}

func (d *Directory) setParent(parent *Directory) {
	d.parent = parent
}

func (d Directory) Walk(walkFunc func(node FSNode)) {
	for _, node := range d.children {
		walkFunc(node)
		if dir, isDir := node.(*Directory); isDir {
			dir.Walk(walkFunc)
		}
	}
}

func (d Directory) WalkFiltered(filterFunc func(node FSNode) bool, walkFunc func(node FSNode)) {
	for _, node := range d.children {
		if filterFunc(node) {
			walkFunc(node)
		}
		if dir, isDir := node.(*Directory); isDir {
			dir.WalkFiltered(filterFunc, walkFunc)
		}
	}
}

func (d Directory) Get(path string) (FSNode, error) {
	components := strings.SplitN(path, "/", 2)
	current := d.children[components[0]]

	if len(components) == 1 {
		return current, nil
	} else {
		dir, isDir := current.(*Directory)
		if !isDir {
			return nil, fmt.Errorf("not foundL: %v", path)
		}
		return dir.Get(components[1])
	}
}

func (d *Directory) AddAt(path string, node FSNode) error {
	if path == "" {
		d.Add(node)
		return nil
	}

	dirname, subpath, _ := strings.Cut(path, "/")
	if _, hasNode := d.children[dirname]; !hasNode {
		newDir := NewDirectory(dirname)
		newDir.setParent(d)
		d.children[dirname] = newDir
	}
	dir, isDir := d.children[dirname].(*Directory)
	if !isDir {
		return fmt.Errorf("couldn't add %s: %s isn't a directory", node.Name(), dirname)
	}
	dir.AddAt(subpath, node)

	return nil
}

func (d *Directory) Add(node FSNode) {
	node.setParent(d)
	d.children[node.Name()] = node
}

type File struct {
	name   string
	size   int
	parent *Directory
}

func NewFile(name string, size int) *File {
	return &File{
		name:   name,
		size:   size,
		parent: nil,
	}
}

func (f File) String() string {
	return fmt.Sprintf("%s (file, size=%d)", f.name, f.size)
}

func (f *File) setParent(parent *Directory) {
	f.parent = parent
}

func (f File) Name() string       { return f.name }
func (f File) Size() int          { return f.size }
func (f File) Parent() *Directory { return f.parent }
