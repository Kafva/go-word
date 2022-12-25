package main

import (
    "bufio"
    "flag"
    "html/template"
    "log"
    "math/rand"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

const WEBROOT = "./public"
const WORD_LIST = "./words.txt"
const DEFAULT_PORT = 2112
const DEFAULT_ADDR = "127.0.0.1"

var VERBOSE = false
var WORDS = []string{}

func LoadWordList() {
    f, err := os.Open(WORD_LIST)
    if err == nil {
        defer f.Close()
        scanner := bufio.NewScanner(f)

        for scanner.Scan() {
            WORDS = append(WORDS, strings.TrimSpace(scanner.Text()))
        }
    } else {
        log.Fatal("Can not open '" + WORD_LIST + "'")
    }
}

func Hook(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        w.Header().Add("Access-Control-Allow-Origin", "*")

        if VERBOSE {
            log.Printf("\033[97m%-5s %-20s\033[0m %-20s", r.Method,
                r.URL.RequestURI(), r.UserAgent())
        }

        if filepath.Base(r.URL.Path) == "index.html" || r.URL.Path == "/" {
            tmpl := template.Must(template.ParseFiles(WEBROOT + "/index.html"))
            word := WORDS[rand.Intn(len(WORDS))]
            tmpl.Execute(w, word)

        } else {
            next.ServeHTTP(w, r)
        }
    })
}

func main() {
    vflag := flag.Bool("verbose", false, "Log all requests")
    port := flag.Int("port", DEFAULT_PORT, "Port to listen on")
    addr := flag.String("addr", DEFAULT_ADDR, "Bind address")
    flag.Parse()
    VERBOSE = *vflag

    rand.Seed(time.Now().UTC().UnixNano())

    // Read the wordlist into memory once during startup
    LoadWordList()

    fs := http.FileServer(http.Dir(WEBROOT + "/"))
    http.Handle("/", Hook(fs))

    listener := *addr + ":" + strconv.Itoa(*port)

    if VERBOSE {
        log.Println("Listening on " + listener + "...")
    }
    if err := http.ListenAndServe(listener, nil); err != nil {
        log.Fatal(err)
    }
}
