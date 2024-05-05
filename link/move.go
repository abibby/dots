package link

import (
	"os"
	"path"

	"github.com/davecgh/go-spew/spew"
)

func Move(homeFile, dotFile string, replace bool) error {
	// _, err := os.Stat(dotFile)
	// if !os.IsNotExist(err) {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return fmt.Errorf("%w: %s", ErrFileExists, dotFile)
	// }

	s, err := os.Stat(homeFile)
	if err != nil {
		return err
	}

	// TODO: check if the file is already in a committed folder

	if s.IsDir() {
		err = os.WriteFile(path.Join(homeFile, ".dots"), []byte("*\n"), 0o644)
		if err != nil {
			return err
		}
	}

	err = move(homeFile, dotFile)
	if err != nil {
		return err
	}

	return symlink(dotFile, homeFile, replace)
}

func move(from, to string) error {
	s, err := os.Stat(from)
	if err != nil {
		spew.Dump("a")
		return err
	}

	if !s.IsDir() {
		return os.Rename(from, to)
	}

	files, err := os.ReadDir(from)
	if err != nil {
		spew.Dump("b")
		return err
	}

	err = os.Mkdir(to, 0o777)
	if os.IsExist(err) {
		// fallthrough
	} else if err != nil {
		spew.Dump("c")
		return err
	}
	for _, f := range files {
		err = move(path.Join(from, f.Name()), path.Join(to, f.Name()))
		if err != nil {
			spew.Dump("d")
			return err
		}
	}

	return os.Remove(from)
}
