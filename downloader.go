// necdet.uygur@gmail.com
//# GET EXECUTABLE BINARY FILES
//```
//https://yt-dl.org/downloads/latest/youtube-dl
//https://www.ffmpeg.org/download.html
//```
//
//# MAKE EXECUTABLE BINARY FILES
//```
//sudo chmod a+rx /usr/local/bin/youtube-dl
//sudo chmod a+rx /usr/local/bin/ffmpeg
//```
//
//# BUILD THIS FILE
//```
//go build downloader.go
//sudo cp -r downloader /usr/local/bin/
//sudo chmod a+rx /usr/local/bin/downloader
//```
//or
//```
//go build downloader.go
//mv downloader /usr/bin/
//chmod a+rx /usr/bin/downloader
//```

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	thisFileName := GetThisFileName()
	isHelp := IsHelp()
	if isHelp {
		help := thisFileName + " -mp3 <url>\n" + thisFileName + " -480 <url>\n" + thisFileName + " <url>"
		fmt.Println(help)
	} else {
		_url := ""
		if len(Contains("http", os.Args)) > 0 {
			_url = Contains("http", os.Args)
		}
		s := "youtube-dl"
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" || runtime.GOOS == "android" {
			s += " -o %(title)s.%(ext)s"
		} else {
			s += " -o '%(title)s.%(ext)s'"
		}
		if len(Contains("480", os.Args)) > 0 {
			if runtime.GOOS == "windows" || runtime.GOOS == "darwin" || runtime.GOOS == "android" {
				s += " -f bestvideo[height<=480]+bestaudio/best[height<=480]"
			} else {
				s += " -f 'bestvideo[height<=480]+bestaudio/best[height<=480]'"
			}
		} else {
			if runtime.GOOS == "windows" || runtime.GOOS == "darwin" || runtime.GOOS == "android" {
				s += " -f bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best"
			} else {
				s += " -f 'bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best'"
			}
		}
		if len(Contains("mp3", os.Args)) > 0 {
			s += " -x --audio-format mp3"
		}
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" || runtime.GOOS == "android" {
			s += " " + _url
		} else {
			s += " '" + _url + "'"
		}
		fmt.Println(s)
		Execute(s)
	}
}

func Contains(a string, list []string) string {
	for _, b := range list {
		if strings.Contains(b, a) {
			return b
		}
	}
	return ""
}

func GetThisFileName() string {
	return os.Args[0]
}

func IsHelp() bool {
	return len(os.Args) < 2 ||
		len(Contains("-h", os.Args)) > 0 ||
		len(Contains("-H", os.Args)) > 0 ||
		len(Contains("--help", os.Args)) > 0 ||
		len(Contains("--HELP", os.Args)) > 0
}

func Execute(s string) {
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdBuffer.String())
}
