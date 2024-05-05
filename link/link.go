package link

import (
	"errors"
	"fmt"
	"os"
	"path"
)

var (
	ErrFileExists = errors.New("file exists")
)

func Link(homeDir, dotDir string, replace bool) error {
	files, err := os.ReadDir(dotDir)
	if err != nil {
		return err
	}

	dots, err := os.ReadFile(path.Join(dotDir, ".dots"))
	if os.IsNotExist(err) {
		// fallthrough
	} else if err != nil {
		return err
	}

	if dots != nil {
		return symlink(dotDir, homeDir, replace)
	}

	err = os.MkdirAll(homeDir, 0o777)
	if err != nil {
		return err
	}

	for _, f := range files {
		homeFile := path.Join(homeDir, f.Name())
		dotFile := path.Join(dotDir, f.Name())
		if f.IsDir() {
			if f.Name() == ".git" {
				continue
			}

			err = Link(homeFile, dotFile, replace)
			if err != nil {
				return err
			}
		} else {
			err = symlink(dotFile, homeFile, replace)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func symlink(dotFile, homeFile string, replace bool) error {
	fileExists := true
	fi, err := os.Lstat(homeFile)
	if os.IsNotExist(err) {
		fileExists = false
	} else if err != nil {
		return err
	} else if fi.Mode()&os.ModeSymlink != 0 {
		l, err := os.Readlink(homeFile)
		if err != nil {
			return err
		}
		if l == dotFile {
			return nil
		}
	}

	if fileExists {
		if replace {
			err = os.Rename(homeFile, homeFile+".backup")
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("%w: %s", ErrFileExists, homeFile)
		}

	}

	err = os.Symlink(dotFile, homeFile)
	if err != nil {
		return err
	}
	return nil
}
