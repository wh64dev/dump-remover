package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/devproje/plog/log"
)

var (
	path     string
	interval int
)

func init() {
	flag.StringVar(&path, "path", "/mnt", "dump backup directory")
	flag.IntVar(&interval, "interval", 2, "remove dump backup n days before")

	flag.Parse()
}

func main() {
	cur := time.Now()
	fs, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Stat("/etc/pve")
	if os.IsNotExist(err) {
		log.Fatalln("current system is not proxmox!")
		return
	}

	for _, f := range fs {
		if f.IsDir() {
			continue
		}

		if !strings.Contains(f.Name(), ".zst") && !strings.Contains(f.Name(), ".gz") && !strings.Contains(f.Name(), ".lzo") {
			continue
		}

		info, ferr := f.Info()
		if ferr != nil {
			log.Errorln(err)
			continue
		}

		mt := info.ModTime()
		if cur.Day()-mt.Day() <= int(interval) {
			os.Remove(filepath.Join(path, f.Name()))
			log.Infof("removed backup file for: %s\n", f.Name())
			continue
		}
	}
}
