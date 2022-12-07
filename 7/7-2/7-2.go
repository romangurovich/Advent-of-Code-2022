package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := common.GetData(7)
	root := Directory{Name: "/", Resources: make([]IResource, 0)}
	os := Os{CurrentDirectory: &root, Root: &root}
	terminal := Terminal{}
	for _, entry := range data {
		terminal.parseEntry(entry, &os)
	}
	spaceFree := 70_000_000 - root.GetSize()
	spaceNeeded := 30_000_000 - spaceFree
	total := getSmallestDirectorySizeToDelete(&root, root.GetSize(), spaceNeeded)
	fmt.Println(total)
}

type Os struct {
	Root             *Directory
	CurrentDirectory *Directory
}

type Terminal struct{}

func (t *Terminal) parseEntry(entry string, os *Os) {
	if string(entry[0]) == "$" {
		t.parseInput(entry[2:], os)
	} else {
		t.parseOutput(entry, os)
	}
}

func (t *Terminal) parseInput(entry string, os *Os) {
	if entry[0:2] == "cd" {
		if string(entry[3]) == "/" {
			os.CurrentDirectory = os.Root
		} else if entry[3:5] == ".." {
			os.CurrentDirectory = os.CurrentDirectory.Parent
		} else {
			dirName := entry[3:]
			directory := createDirectory(dirName, os)
			os.CurrentDirectory = directory
		}
	}
}

func (t *Terminal) parseOutput(entry string, os *Os) {
	info := strings.Split(entry, " ")
	if info[0] == "dir" {
		createDirectory(info[1], os)
	} else {
		size, _ := strconv.Atoi(info[0])
		createFile(info[1], size, os)
	}
}

func createDirectory(dirName string, os *Os) *Directory {
	selectedDirectory := os.CurrentDirectory.FindDirectory(dirName)

	if selectedDirectory == nil {
		selectedDirectory = &Directory{Name: dirName, Parent: os.CurrentDirectory}
		os.CurrentDirectory.Resources = append(os.CurrentDirectory.Resources, selectedDirectory)
		os.CurrentDirectory.Subdirectories = append(os.CurrentDirectory.Subdirectories, selectedDirectory)
	}

	return selectedDirectory
}

func createFile(fileName string, size int, os *Os) {
	selectedFile := os.CurrentDirectory.FindFile(fileName)
	if selectedFile == nil {
		selectedFile = &File{Name: fileName, Size: size}
		os.CurrentDirectory.Resources = append(os.CurrentDirectory.Resources, selectedFile)
		os.CurrentDirectory.Files = append(os.CurrentDirectory.Files, selectedFile)
	}
}

type IResource interface {
	GetSize() int64
	GetName() string
}

type Directory struct {
	Name           string
	Resources      []IResource
	Subdirectories []*Directory
	Files          []*File
	Parent         *Directory
}

func (d *Directory) GetSize() int64 {
	var result int64
	for _, resource := range d.Resources {
		result += resource.GetSize()
	}

	return result
}

func (d *Directory) GetName() string {
	return d.Name
}

func (d *Directory) FindFile(name string) *File {
	for _, file := range d.Files {
		if file.GetName() == name {
			return file
		}
	}
	return nil
}

func (d *Directory) FindDirectory(name string) *Directory {
	for _, subdirectory := range d.Subdirectories {
		if subdirectory.GetName() == name {
			return subdirectory
		}
	}
	return nil
}

type File struct {
	Name string
	Size int
}

func (f *File) GetSize() int64 {
	return int64(f.Size)
}

func (f *File) GetName() string {
	return f.Name
}

func getFilteredSizeSum(dir *Directory, limit int64) int64 {
	var total int64
	size := dir.GetSize()

	if size < limit {
		total += size
	}

	for _, subdirectory := range dir.Subdirectories {
		total += getFilteredSizeSum(subdirectory, limit)
	}

	return total
}

func getSmallestDirectorySizeToDelete(dir *Directory, smallestCapacitySoFar int64, limit int64) int64 {
	size := dir.GetSize()
	if size > limit && size < smallestCapacitySoFar {
		smallestCapacitySoFar = size
	}

	for _, subdirectory := range dir.Subdirectories {
		smallestCapacitySoFar = getSmallestDirectorySizeToDelete(subdirectory, smallestCapacitySoFar, limit)
	}

	return smallestCapacitySoFar
}
