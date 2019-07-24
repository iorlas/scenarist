package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func backupHandler(w http.ResponseWriter, req *http.Request) {
	var start = time.Now()

	name := req.URL.Query().Get("name")
	created, err := backup(name)
	must(err)

	if created {
		var took = time.Now().Sub(start).String()
		_, err := io.WriteString(w, "CREATED, TOOK "+took)
		must(err)
	} else {
		_, err := io.WriteString(w, "EXISTS")
		must(err)
	}
}

func restoreHandler(w http.ResponseWriter, req *http.Request) {
	var start = time.Now()

	name := req.URL.Query().Get("name")
	restored, err := restore(name)
	must(err)

	if restored {
		var took = time.Now().Sub(start).String()
		_, err := io.WriteString(w, "RESTORED, TOOK "+took)
		must(err)
	} else {
		_, err := io.WriteString(w, "NOT FOUND")
		must(err)
	}
}

func clearHandler(w http.ResponseWriter, req *http.Request) {
	must(os.RemoveAll("scenarios"))
	must(os.Mkdir("scenarios", os.ModePerm))

	var _, err = io.WriteString(w, "OK\n")
	must(err)
}

func server() {
	http.HandleFunc("/api/v1/clear", clearHandler)
	http.HandleFunc("/api/v1/backup", backupHandler)
	http.HandleFunc("/api/v1/restore", restoreHandler)

	log.Println("Listening on 8077")
	must(http.ListenAndServe(":8077", nil))
}
