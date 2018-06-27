package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/goji/httpauth"
	"github.com/gorilla/handlers"
)

func getParam(arg, env string) string {
	if s := os.Getenv(env); s != "" {
		return s
	}

	if path := os.Getenv(env + "_FILE"); path != "" {
		if b, err := ioutil.ReadFile(path); err != nil {
			log.Fatalf("Unable to read file \"%s\": %v\n", path, err)
		} else {
			return strings.TrimSpace(string(b))
		}
	}

	if arg != "" {
		return arg
	}

	return ""
}

func main() {
	addrFlag := flag.String("addr", ":80", "address to listen on, e.g. [<ip>]:<port>")
	dirFlag := flag.String("dir", "/http", "folder to serve")
	userFlag := flag.String("user", "", "set username for basic auth")
	passwdFlag := flag.String("passwd", "", "set password for basic auth")

	flag.Parse()

	addr := getParam(*addrFlag, "HTTP_ADDRESS")
	dir := getParam(*dirFlag, "HTTP_DIRECTORY")
	user := getParam(*userFlag, "HTTP_USERNAME")
	passwd := getParam(*passwdFlag, "HTTP_PASSWORD")

	log.Println("Serving files from", dir)
	fileHandler := http.FileServer(http.Dir(dir))

	var server http.Handler

	if user != "" && passwd != "" {
		log.Println("Using Basic Auth")
		server = handlers.CombinedLoggingHandler(os.Stdout, httpauth.SimpleBasicAuth(user, passwd)(fileHandler))
	} else {
		server = handlers.CombinedLoggingHandler(os.Stdout, fileHandler)
	}

	log.Println("Listening on", addr)
	log.Fatalln(http.ListenAndServe(addr, server))
}
