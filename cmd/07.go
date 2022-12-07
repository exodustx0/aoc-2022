package cmd

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

type fileinfo struct {
	name string
	size uint
}

type dirinfo struct {
	name    string
	parent  *dirinfo
	subdirs []dirinfo
	files   []fileinfo
}

func (d *dirinfo) allDirs(yield func(*dirinfo) bool) {
	if !yield(d) {
		return
	}
	for i := range d.subdirs {
		d.subdirs[i].allDirs(yield)
	}
}

func (d dirinfo) size() uint {
	var size uint
	for _, sd := range d.subdirs {
		size += sd.size()
	}
	for _, f := range d.files {
		size += f.size
	}
	return size
}

func day07(input *bufio.Reader) error {
	var root dirinfo
	cur := &root
	s := bufio.NewScanner(input)
	for s.Scan() {
		str := s.Text()
		if !strings.HasPrefix(str, "$ ") {
			panic("not a command")
		}

	newCmd:
		cmd, arg, _ := strings.Cut(str[2:], " ")
		switch cmd {
		case "cd":
			switch arg {
			case "/":
				cur = &root
			case "..":
				if cur.parent == nil {
					panic("nil parent")
				}
				cur = cur.parent
			default:
				for i := range cur.subdirs {
					if cur.subdirs[i].name == arg {
						cur = &cur.subdirs[i]
						break
					}
				}
			}

		case "ls":
			for s.Scan() {
				str = s.Text()
				if strings.HasPrefix(str, "$ ") {
					goto newCmd
				}

				typ, name, ok := strings.Cut(str, " ")
				if !ok {
					panic("ls line without space")
				}

				if typ == "dir" {
					cur.subdirs = append(cur.subdirs, dirinfo{name: name, parent: cur})
				} else {
					size, err := strconv.ParseUint(typ, 10, 64)
					if err != nil {
						return err
					}
					cur.files = append(cur.files, fileinfo{name, uint(size)})
				}
			}
		}
	}

	var smallSizesSum uint
	root.allDirs(func(d *dirinfo) bool {
		if size := d.size(); size <= 100_000 {
			smallSizesSum += size
		}
		return true
	})

	partOne(smallSizesSum)

	spaceNeeded := root.size() - 40_000_000
	if spaceNeeded > 70_000_000 {
		panic("no space needed")
	}

	smallestDelDirSize := uint(math.MaxUint)
	root.allDirs(func(d *dirinfo) bool {
		if size := d.size(); size < spaceNeeded {
			return false
		} else if size < smallestDelDirSize {
			smallestDelDirSize = size
		}
		return true
	})

	partTwo(smallestDelDirSize)

	return nil
}
