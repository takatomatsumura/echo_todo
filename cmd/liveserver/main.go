package main

import (
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const ignoreFileName = ".liveserverignore"

var (
	ignore *Ignore
	files  = map[string]int64{}
)

func init() {
	var err error
	ignore, err = ParseIgnoreFile(ignoreFileName)
	if err != nil {
		ignore = &Ignore{}
	}

	// cache mtimes
	_ = shouldRebuild()
}

func shouldRebuild() bool {
	var cond bool
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// watch only golang files
		if d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}
		if ignore.Match(path) {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}
		mtime := info.ModTime().Unix()
		if old_mtime, ok := files[path]; !ok || mtime > old_mtime {
			files[path] = mtime
			cond = true
		}
		return nil
	})
	return cond
}

func main() {
	for {
		exec.Command("make", "build").Run()
		cmd := exec.Command("make", "run")
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

		for {
			t := time.NewTimer(1 * time.Second)
			s := make(chan os.Signal, 1)
			signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
			select {
			case <-s:
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
				os.Exit(0)
			case <-t.C:
				break
			}
			if shouldRebuild() {
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
				break
			}
		}
	}
}
