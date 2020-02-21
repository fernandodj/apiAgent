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

type CpuInfo struct {
	Usage, Model, Type, Uptime string
}

type ProcessInfo struct {
	Mem_Perc 		string `json: "MEM_PERC"`
	Cmd 		string `json: "CMD"` 
	User 		string `json: "USER"`
	Pid 		string `json: "PID"`
	Ppid 		string `json: "PPID"`
	Time 		string `json: "TIME"`
	Cpu_Perc 		string `json: "CPU_PERC"`
}

type UserInfo struct {
	Date, Line, Name, Hour string
}

type ServerInfo struct {
	Cpu 		CpuInfo
	ProcessData []ProcessInfo 
	Users 		[]UserInfo
	OsName 		string
	OsVersion 	string 
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
