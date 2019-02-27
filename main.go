package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type dir struct {
	Name         string
	Path         string
	IsDir        bool
	Tabs         int
	Size         int64
	IsLast       bool
	PrevDirsLast []bool // map of IsLast for dirs from path struct dir
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
	pathDir, err := strToDir(path)
	if err != nil {
		return err
	}
	//logic for -f
	if printFiles == true {
		dirRecursiveFinder(out, pathDir)
	}
	return nil
}

//return massive with collection of dirs
func dirContent(path dir) ([]dir, error) {
	files, err := ioutil.ReadDir(path.Path) // warning !!!
	if err != nil {
		return nil, err
	}
	var dirs []dir
	for i, file := range files {
		isLast := i == len(files)-1
		a := dir{
			IsDir:  file.IsDir(),
			Name:   file.Name(),
			Path:   path.Path + "/" + file.Name(),
			Size:   file.Size(),
			Tabs:   path.Tabs + 1,
			IsLast: isLast,
		}
		a.PrevDirsLast = append(path.PrevDirsLast, path.IsLast)
		dirs = append(dirs, a)
		//prevDirIsLast for dir
	}
	return dirs, nil
}

//output formated string of dir in output
func dirPrinter(out io.Writer, path dir) {
	if path.IsDir == true {
		if path.Tabs == 0 {
			subPrinter(out, "", path)
		} else if path.Tabs == 1 {
			subPrinter(out, "│⎸  ", path)
		} else if path.Tabs == 2 {
			subPrinter(out, "│⎸  │⎸  ", path)
		}
	}
}

//return in output formated tree
func dirRecursiveFinder(out io.Writer, current dir) error {
	dirs, err := dirContent(current)
	if err != nil {
		return err
	}
	for _, subDir := range dirs {
		// fmt.Fprintf(out, "%v | %v\n", subDir.Name, subDir.PrevDirsLast)
		dirPrinter(out, subDir)
		dirRecursiveFinder(out, subDir)
	}
	return nil
}

//input FileInfo output dir
func strToDir(path string) (dir, error) {
	pathFI, err := os.Stat(path)
	var emptyDir dir
	if err != nil {
		return emptyDir, err
	}
	a := dir{
		IsDir:  pathFI.IsDir(),
		Name:   pathFI.Name(),
		Size:   pathFI.Size(),
		Path:   pathFI.Name(),
		Tabs:   -1,
		IsLast: true,
	}
	return a, nil
}

func subPrinter(out io.Writer, tabs string, path dir) {
	fmt.Fprintf(out, "%v%v\n", tabGen(path), path.Name)
}

func tabGen(file dir) string {
	prevDirs := append(file.PrevDirsLast[1:], file.IsLast)
	pipeAndTab := "│⎸  "
	tabs := ""
	for i, currentDirLast := range prevDirs {
		// create last symbol
		if i == len(prevDirs)-1 {
			if currentDirLast == false {
				tabs = "├───" + tabs
			} else {
				tabs += "└───"
			}
			// creating tabs
		} else if currentDirLast == false {
			tabs += pipeAndTab
		} else {
			tabs += "	"
		}
	}
	return tabs
}
