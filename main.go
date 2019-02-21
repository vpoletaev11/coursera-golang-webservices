package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type dir struct {
	Name   string
	Path   string
	IsDir  bool
	Tabs   int
	Size   int64
	IsLast bool
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

//combainer
func dirTree(out io.Writer, path string, printFiles bool) error {
	// if printfiles true
	pathFI, err := os.Stat(path)
	if err != nil {
		return err
	}
	pathDir, err := FIToDir(pathFI, true, 0)
	if err != nil {
		return err
	}
	dirRecursiveFinder(out, pathDir)
	return nil
}

//return massive with collection of dirs
func dirContent(path dir, tabs int) ([]dir, error) {
	files, err := ioutil.ReadDir(path.Path) // warning !!!
	if err != nil {
		return nil, err
	}
	var dirs []dir
	for i, file := range files {
		var isLast bool
		if i == len(files)-1 {
			isLast = true
		}
		a := dir{
			IsDir:  file.IsDir(),
			Name:   file.Name(),
			Path:   path.Name + "/" + file.Name(),
			Size:   file.Size(),
			Tabs:   tabs,
			IsLast: isLast,
		}
		dirs = append(dirs, a)
	}
	return dirs, nil
}

//output formated string of dir in output
func dirPrinter(out io.Writer, Dir dir) {
	if Dir.IsDir == true {
		fmt.Fprintln(out, Dir.Name)
	} else {
		if Dir.Size == 0 {
			str := fmt.Sprintf("%v (empty)", Dir.Name)
			fmt.Fprintln(out, str)
		} else {
			str := fmt.Sprintf("%v (%vb)", Dir.Name, Dir.Size)
			fmt.Fprintln(out, str)
		}
	}
}

//return in output formated tree
func dirRecursiveFinder(out io.Writer, current dir) error {
	dirs, err := dirContent(current, 0) //change!!
	if err != nil {
		return err
	}
	for _, subDir := range dirs {
		dirPrinter(out, subDir)
		dirRecursiveFinder(out, subDir)
	}
	return nil
}

//input FileInfo output dir
func FIToDir(file os.FileInfo, isLast bool, tabs int) (dir, error) {
	var a dir
	a.IsDir, a.Name, a.Size = file.IsDir(), file.Name(), file.Size()
	a.Path = file.Name()
	a.Tabs, a.IsLast = tabs, isLast
	return a, nil
}
