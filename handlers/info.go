package handlers

import (
	"fmt"
	"github.com/takama/router"
	"net/http"
	"os"
	"runtime"
)

type ServiceInfo struct {
	Host    string       `json:"host"`
	Runtime *RuntimeInfo `json:"runtime"`
	Version string       `json:"version"`
	Repo    string       `json:"repo"`
	Commit  string       `json:"commit"`
}

type RuntimeInfo struct {
	Compiler   string `json:"compiler"`
	CPU        int    `json:"cpu"`
	Memory     string `json:"memory"`
	Goroutines int    `json:"goroutines"`
}

func Info(c *router.Control, version, repo, commit string) {
	host, _ := os.Hostname()
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)

	rt := &RuntimeInfo{
		CPU:        runtime.NumCPU(),
		Memory:     fmt.Sprintf("%.2fMB", float64(m.Alloc)/(1<<(10*2))),
		Goroutines: runtime.NumGoroutine(),
	}

	info := ServiceInfo{
		Host:    host,
		Runtime: rt,
		Version: version,
		Repo:    repo,
		Commit:  commit,
	}

	c.Code(http.StatusOK).Body(info)
}
