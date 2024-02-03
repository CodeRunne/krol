package main

import (
	"log"

	"github.com/D3Ext/maldev/process"
)

type Process struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	PID  int    `json:"pid"`
	PPID int    `json:"ppid"`
}

var Processes []Process

func init() {

	// Get all processes
	processes, err := process.GetProcesses()
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, p := range processes {
		var proc Process
		proc.ID = count
		proc.Name, _ = process.FindNameByPid(p.Pid())
		proc.PID = p.Pid()
		proc.PPID = p.PPid()

		// increment count
		count += 1

		// append to struct
		Processes = append(Processes, proc)
	}

}
