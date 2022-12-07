package nospaceleft

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func PartOne(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var err error
	sh := newShell()

	for scanner.Scan() {
		cmd := scanner.Text()
		sh, err = sh.ParseCommand(cmd)
		if err != nil {
			return "", err
		}
	}

	sh, _ = sh.ParseCommand("$ cd /")

	filter := func(d dir) bool { return d.Size() < 100000 }
	size := sh.Size(filter)

	return strconv.FormatInt(int64(size), 10), nil
}

func newShell() Shell {
	return &shell{
		currentDirectory: newDir("/", nil),
		executingCommand: nil,
	}
}

type Shell interface {
	ParseCommand(string) (Shell, error)
	Size(func(dir) bool) int
}

type shell struct {
	currentDirectory *dir
	executingCommand func(string)
}

func (sh *shell) ParseCommand(cmd string) (Shell, error) {
	if strings.HasPrefix(cmd, "$ ") {
		sh.executingCommand = nil
		sh.exec(cmd[2:])
		return sh, nil
	} else if sh.executingCommand != nil {
		sh.executingCommand(cmd)
		return sh, nil
	}

	panic("unknown state")
}

func (sh *shell) Size(filter func(dir) bool) int {
	matchingDirectories := sh.currentDirectory.Find(filter)
	sizeOfMatchingDirs := 0

	for _, d := range matchingDirectories {
		sizeOfMatchingDirs += d.Size()
	}

	return sizeOfMatchingDirs
}

func (sh *shell) exec(cmd string) {
	if strings.HasPrefix(cmd, "cd ") {
		target := cmd[3:]
		if target == "/" {
			sh.currentDirectory = sh.currentDirectory.Root()
		} else if target == ".." {
			sh.currentDirectory = sh.currentDirectory.Parent()
		} else {
			sh.currentDirectory = sh.currentDirectory.SubDirectory(target)
		}
	} else if cmd == "ls" {
		sh.executingCommand = sh.handleListDirectoryOutput
	}
}

func (sh *shell) handleListDirectoryOutput(output string) {
	if strings.HasPrefix(output, "dir ") {
		dirName := output[4:]
		dir := newDir(dirName, sh.currentDirectory)
		sh.currentDirectory.subdirs = append(sh.currentDirectory.subdirs, *dir)
		return
	}

	var fileSize int
	var fileName string

	fmt.Sscanf(output, "%d %s", &fileSize, &fileName)
	sh.currentDirectory.sizeOfFiles += fileSize
}

type dir struct {
	name        string
	sizeOfFiles int

	parent  *dir
	subdirs []dir
}

func newDir(name string, parent *dir) *dir {
	return &dir{name: name, parent: parent}
}

func (d *dir) Find(filter func(dir) bool) []dir {
	result := []dir{}

	if filter(*d) {
		result = append(result, *d)
	}

	for _, subdir := range d.subdirs {
		result = append(result, subdir.Find(filter)...)
	}

	return result
}

func (d *dir) Name() string {
	return d.name
}

func (d *dir) Parent() *dir {
	return d.parent
}

func (d *dir) Root() *dir {
	root := d

	for root.parent != nil {
		root = root.parent
	}

	return root
}

func (d *dir) Size() int {
	var sizeOfSubDirs int
	for _, subdir := range d.subdirs {
		sizeOfSubDirs += subdir.Size()
	}
	return d.sizeOfFiles + sizeOfSubDirs
}

func (d *dir) SubDirectory(name string) *dir {
	for idx, subdir := range d.subdirs {
		if subdir.Name() == name {
			return &d.subdirs[idx]
		}
	}

	panic("no such sub directory")
}
