package main

import (
	"github.com/spf13/viper"
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
		log.Printf("Scenario %s created\n", name)
	} else {
		_, err := io.WriteString(w, "EXISTS")
		must(err)
		log.Printf("Scenario %s already exists\n", name)
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
		log.Printf("Scenario %s loaded\n", name)
	} else {
		_, err := io.WriteString(w, "NOT FOUND")
		must(err)
		log.Printf("Scenario %s not found\n", name)
	}
}

func clearHandler(w http.ResponseWriter, req *http.Request) {
	must(os.RemoveAll(viper.GetString("datapath")))
	must(os.Mkdir(viper.GetString("datapath"), os.ModePerm))

	var _, err = io.WriteString(w, "OK\n")
	must(err)

	log.Println("Scenarios cleared")
}

func server() {
	http.HandleFunc("/api/v1/clear", clearHandler)
	http.HandleFunc("/api/v1/backup", backupHandler)
	http.HandleFunc("/api/v1/restore", restoreHandler)

	log.Println("Listening on 8077")
	must(http.ListenAndServe(":8077", nil))
}
