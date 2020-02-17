package main

import ( 
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
	"time"
	"log"
)

type ServerInfo struct {
	Cpu string
	ProcessList string 
	Users string
	NameOS string
	VersionOS string 
}

func saveInfo(w http.ResponseWriter, r *http.Request){
	var serverInfo ServerInfo
	currentTime := time.Now()
	filename := fmt.Sprintf("%s_%s.json", r.RemoteAddr, currentTime.Format("2006-01-02"))
	_ = json.NewDecoder(r.Body).Decode(&serverInfo)
	jsonServer, errBody := json.Marshal(serverInfo)
	if errBody != nil {
		log.Fatal(errBody)
		http.Error(w, "Error parsing body", http.StatusBadRequest)
	}

	err := ioutil.WriteFile(filename, jsonServer, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error saving data", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonServer)
	}
	
}