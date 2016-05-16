package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

const (
	dfltAddr = ":3000"
	prefix   = "/serve/"
)

var (
	addr    = dfltAddr
	pattern = regexp.MustCompile(fmt.Sprintf(`^%s([[:alpha:]]+)/(\d+)$`, prefix))
)

func handler(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Println("HTTP is being SERVED! Path:", path)
	matches := pattern.FindStringSubmatch(path)
	if len(matches) != 3 {
		log.Println("bad path:", path)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	var (
		id, err   = strconv.Atoi(matches[2])
		typeParam = matches[1]
	)
	if err != nil {
		log.Println("bad ID:", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	log.Printf("id: %d, type: %s", id, typeParam)
	rw.Header().Set("content-type", "application/json")
	rw.Write([]byte(fmt.Sprintf(`{"id":%d,"type":%q}`, id, typeParam)))
}

func main() {
	http.HandleFunc("/serve/", handler)
	log.Println("serving on", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func init() {
	if envaddr := os.Getenv("ADDR"); envaddr != "" {
		addr = envaddr
		return
	}
	if envport := os.Getenv("PORT"); envport != "" {
		addr = fmt.Sprintf(":%s", envport)
	}
}
