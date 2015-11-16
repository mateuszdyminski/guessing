package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var configPath string

// Config holds configuration of feeder.
type Config struct {
	Host    string
	Statics string
	Results string
}

func init() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.StringVar(&configPath, "config", "config/conf.toml", "config path")
}

var conf Config

func main() {
	// load config
	flag.Parse()

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Can't open config file!")
	}

	if err := toml.Unmarshal(bytes, &conf); err != nil {
		log.Fatalf("Can't decode config file!")
	}

	log.Printf("Config: %+v", conf)

	launchServer(&conf)
}

func launchServer(conf *Config) {
	r := mux.NewRouter()

	// Handle routes
	var statics StaticRoutes
	r.HandleFunc("/restapi/results", saveResults)
	r.Handle("/{path:.*}", http.FileServer(append(statics, http.Dir(conf.Statics)))).Name("static")

	http.Handle("/", loggingHandler{r})

	// Listen on hostname:port
	log.Printf("Listening on %s...", conf.Host)
	if err := http.ListenAndServe(conf.Host, nil); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

type StaticRoutes []http.FileSystem

func (sr StaticRoutes) Open(name string) (f http.File, err error) {
	for _, s := range sr {
		if f, err = s.Open(name); err == nil {
			f = disabledDirListing{f}
			return
		}
	}
	return
}

type disabledDirListing struct {
	http.File
}

func (f disabledDirListing) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

type loggingHandler struct {
	http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// path := req.URL.Path
	t := time.Now()
	h.Handler.ServeHTTP(w, req)

	elapsed := time.Since(t)
	log.Printf("%s [%s] \"%s %s %s\" \"%s\" \"%s\" \"Took: %s\"", req.RemoteAddr,
		t.Format("02/Jan/2006:15:04:05 -0700"), req.Method, req.RequestURI, req.Proto, req.Referer(), req.UserAgent(), elapsed)
}

func saveResults(w http.ResponseWriter, req *http.Request) {
	var res result
	json.NewDecoder(req.Body).Decode(&res)

	f, err := os.OpenFile(conf.Results, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		w.Write([]byte("Can't save results! Err: " + err.Error()))
	}

	fmt.Fprintf(f, "%s %+v \n", time.Now(), res)

	f.Close()
}

type result struct {
	Answer   int    `json:"answer,omitempty"`
	Round    int    `json:"round,omitempty"`
	Steps    []int  `json:"steps,omitempty"`
	Username string `json:"username,omitempty"`
}
