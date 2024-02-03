package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProcessesController(c *gin.Context) {
	c.JSON(http.StatusOK, Processes)
}

func GetProcessByPIDController(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, process := range Processes {
		if process.ID == id {
			c.JSON(http.StatusOK, process)
			return
		}
	}

	// Returns error if PID is invalid
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Process ID Not Found!",
	})
}

func KillProcessByPIDController(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for key, process := range Processes {
		if process.ID == id {
			proc, err := os.FindProcess(process.PID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			Processes = append(Processes[:key], Processes[key+1:]...)
			proc.Kill() // kill process
			c.JSON(http.StatusOK, gin.H{
				"message": "Process has been successfully terminated!",
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Process ID Not Found!",
	})
}

func KillAllProcessesController(c *gin.Context) {
	for _, process := range Processes {
		proc, err := os.FindProcess(process.PID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Kill all running processes
		proc.Kill()
	}
}
