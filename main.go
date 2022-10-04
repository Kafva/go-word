package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"log"
)

// Note: the executable must be in the same directory as the public folder
var CWD string 		    = GetCwd()
var SHUF_CMD string 	= GetShufPath()
var VERBOSE = false
const PORT = "2112"

func GetWord(w http.ResponseWriter, r *http.Request) {
  log_request(r)
	// Wordlist from: https://github.com/dwyl/english-words/blob/master/words.txt
	cmd 	  := exec.Command(SHUF_CMD, "-n", "1", CWD + "/words.txt")
	out,_ 	:= cmd.Output()

	//w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w, strings.Title(string(out)))
}

func main(){
  vflag := flag.Bool("verbose", false, "Print request information.")
  flag.Parse()
  VERBOSE = *vflag

	fs := http.FileServer(http.Dir(CWD + "/public/"))
	http.Handle("/", LogRequest(fs))
	http.HandleFunc("/word", GetWord)

  if VERBOSE { log.Println("Listening on "+PORT+"...") }
	http.ListenAndServe("127.0.0.1:"+PORT, nil)
}

//============================================================================//

func GetCwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetShufPath() string {
	path, err := exec.LookPath("shuf")
	if err != nil {
    panic("Could not find 'shuf' executable")
	}
	return path
}

func log_request(r *http.Request) {
	if VERBOSE {
		log.Printf("\033[97m%-5s %-20s\033[0m %-20s", r.Method,
							 r.URL.RequestURI(), r.UserAgent())
	}
}

func LogRequest(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log_request(r)
    next.ServeHTTP(w, r)
  })
}

