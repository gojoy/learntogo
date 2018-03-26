package inotifyTree

import (
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"

	"fmt"
)

func AddWatchAll(dir string, w *fsnotify.Watcher) error {

	var (
		err error
	)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
			fmt.Printf("add dir:%v\n", path)
			if err = w.Add(path); err != nil {
				fmt.Println(err)
				return err
			}
		}
		return nil
	})

	return err
}

func UpdateDirTree(w *fsnotify.Watcher) error {
	for {
		select {
		case events := <-w.Events:
			if events.Op&fsnotify.Create == fsnotify.Create {
				if i, err := os.Stat(events.Name); err == nil && i.IsDir() {
					fmt.Printf("update add create dir:%v\n", events.Name)
					if err = w.Add(events.Name); err != nil {
						panic(err)
						return err
					}
				}
			}
			if events.Op&fsnotify.Remove == fsnotify.Remove ||
				events.Op&fsnotify.Rename == fsnotify.Rename {
				if i, err := os.Stat(events.Name); err == nil && i.IsDir() {
					fmt.Printf("update remove create dir:%v\n", events.Name)
					if err = w.Remove(events.Name); err != nil {
						panic(err)
						return err
					}
				}
			}
		case err := <-w.Errors:
			return err
		}
	}
}
