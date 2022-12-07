package solutions

import (
	"strings"

	"github.com/wazeemwoz/advent2022/file"
	"github.com/wazeemwoz/advent2022/utils"
)

type myFile struct {
	label  string
	parent *directory
	size   int
}

type directory struct {
	label          string
	dictionary     map[string]int
	subdirectories []*directory
	files          []myFile
	parent         *directory
	size           int
}

func FilteredSize(d directory) int {
	size := 0

	for _, sub := range d.subdirectories {
		size += FilteredSize(*sub)
	}

	if d.size <= 100000 {
		size += d.size
	}

	return size
}

func FindUnder(d directory) int {
	var collect func(directory, int) int
	collect = func(d directory, target int) int {
		size := d.size

		for _, sub := range d.subdirectories {
			bestSize := collect(*sub, target)
			if bestSize >= target {
				size = utils.Min(bestSize, size)
			}
		}
		return size
	}
	return collect(d, 30000000-(70000000-d.size))
}

func Solution7(fn func(directory) int) func(string) int {
	return func(filepath string) int {
		fileStream := file.NewStream(filepath)
		rootNode := newDir("/")
		currentNode := rootNode
		fileStream.ForEach(func(line string) {
			tokens := strings.Split(line, " ")
			switch tokens[0] {
			case "$":
				if tokens[1] == "cd" {
					switch tokens[2] {
					case "/":
						currentNode = rootNode
					case "..":
						currentNode = currentNode.parent
					default:
						currentNode = currentNode.subdirectories[currentNode.dictionary[tokens[2]]]
					}
				}
			case "dir":
				currentNode.addDir(tokens[1])
			default:
				currentNode.add(tokens[1], utils.ToInt(tokens[0]))
			}
		})
		return fn(*rootNode)
	}
}

func newDir(label string) *directory {
	return &directory{label, make(map[string]int), make([]*directory, 0), make([]myFile, 0), nil, 0}
}

func (dir *directory) addDir(label string) {
	dir.subdirectories = append(dir.subdirectories, newDir(label))
	dir.dictionary[label] = len(dir.subdirectories) - 1
	dir.subdirectories[len(dir.subdirectories)-1].parent = dir
}

func (dir *directory) add(label string, size int) {
	dir.files = append(dir.files, myFile{label, dir, size})
	dir.dictionary[label] = len(dir.files) - 1
	dir.size += size
	parent := dir.parent
	for parent != nil {
		parent.size += size
		parent = parent.parent
	}
}
