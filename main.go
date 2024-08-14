package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/devproje/plog/log"
	"github.com/joho/godotenv"
)

func init() {
	exec, _ := os.Executable()
	current := filepath.Dir(exec)
	f, err := os.Stat(filepath.Join(current, ".env"))
	if err != nil {
		ex, _ := os.OpenFile(filepath.Join(current, ".env.example"), 0644, 'r')
		var buf []byte
		_, _ = ex.Read(buf)
		_ = os.WriteFile(filepath.Join(current, ".env"), buf, 0644)

		log.Fatalf(".env config not founded! Please write file first!\n")
		return
	}

	godotenv.Load(f.Name())
}

func main() {
	cur := time.Now()
	fs, err := os.ReadDir(os.Getenv("TARGET_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	interval, err := strconv.ParseInt(os.Getenv("INTERVAL_DAY"), 10, 32)
	if err != nil {
		log.Fatalln("value must be to integer")
		return
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
			os.Remove(f.Name())
			log.Infof("removed backup file for: %s\n", f.Name())
			continue
		}
	}
}
