package workspaces

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"

	systemconfigs "github.com/torabian/fireback/modules/workspaces/systemconfigs"
	"github.com/urfave/cli"
)

// ~/Library/LaunchAgents Per-user agents provided by the user.
var SCOPE_USER_FOR_USER = "~/Library/LaunchAgents"

// /Library/LaunchAgents Per-user agents provided by the administrator.
var SCOPE_SUDO_FOR_USER = "/Library/LaunchAgents"

// /Library/LaunchDaemons System-wide daemons provided by the administrator.
// /System/Library/LaunchAgents Per-user agents provided by Mac OS X.
// /System/Library/LaunchDaemons System-wide daemons provided by Mac OS X.

type SystemServiceInfo struct {
	Label       string
	Program     string
	StdOut      string
	CONFIG_PATH string
	StdErr      string
}

func SystemServiceHandler(action string, c *cli.Context) {
	switch os := runtime.GOOS; os {
	case "darwin":
		if action == "load" {
			ServiceLoadMac(c)
		} else if action == "unload" {
			ServiceUnloadMac(c)
		} else if action == "reload" {
			ServiceUnloadMac(c)
			ServiceLoadMac(c)
		}
	case "linux":
		if action == "load" {
			ServiceLoadDebian(c)
		} else if action == "unload" {
			ServiceUnloadDebian(c)
		} else if action == "reload" {
			ServiceUnloadDebian(c)
			ServiceLoadDebian(c)
		}
	default:
		fmt.Printf("Other %v", os)
	}
}

func ServiceUnloadDebian(c *cli.Context) error {

	serviceName := config.DebianIdentifier
	serviceFileName := serviceName + ".service"

	{
		app := "service"
		cmd := exec.Command(app, serviceName, "stop")

		fmt.Println(app, serviceName, "stop")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		app := "systemctl"
		cmd := exec.Command(app, "disable", serviceFileName)
		fmt.Println(app, "disable", serviceFileName)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func ServiceLoadDebian(c *cli.Context) error {
	binaryPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	uri, err := ResolveConfigurationUri()
	td := SystemServiceInfo{
		Label:       config.DebianIdentifier,
		Program:     binaryPath,
		CONFIG_PATH: uri,
	}

	// I think this is no longer needed
	td.StdErr = config.StdErr
	td.StdOut = config.StdOut

	if c.IsSet("stderr") {
		td.StdErr = c.String("stderr")
	}

	if c.IsSet("stdout") {
		td.StdOut = c.String("stdout")
	}

	t, err := template.ParseFS(systemconfigs.SystemConfigs, "debian.service.tpl")
	if err != nil {
		panic(err)
	}

	serviceName := config.DebianIdentifier
	serviceFileName := serviceName + ".service"

	daemonPath := filepath.Join("/lib/systemd/system/", serviceFileName)

	fmt.Println("Daemon", daemonPath)
	file, err := os.OpenFile(daemonPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, td)
	if err != nil {
		panic(err)
	}

	{
		app := "service"
		cmd := exec.Command(app, serviceName, "start")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	{
		app := "systemctl"
		cmd := exec.Command(app, "enable", serviceFileName)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func GetMacDaemon() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	daemonPath := filepath.Join(dirname, "Library/LaunchAgents", config.MacIdentifier+".plist")

	return daemonPath
}

func ServiceLoadMac(c *cli.Context) error {
	binaryPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	td := SystemServiceInfo{
		Label:   config.MacIdentifier,
		Program: binaryPath,
	}

	uri, err := ResolveConfigurationUri()
	possibleConfigPath := filepath.Join(workingDirectory, uri)

	if Exists(possibleConfigPath) {
		td.CONFIG_PATH = possibleConfigPath
	}

	td.StdErr = config.StdErr
	td.StdOut = config.StdOut

	if c.IsSet("stderr") {
		td.StdErr = c.String("stderr")
	}

	if c.IsSet("stdout") {
		td.StdOut = c.String("stdout")
	}

	t, err := template.ParseFS(systemconfigs.SystemConfigs, "macos-daemon.plist.tpl")
	if err != nil {
		panic(err)
	}

	daemonPath := GetMacDaemon()
	file, err := os.OpenFile(daemonPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, td)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	app := "launchctl"

	cmd := exec.Command(app, "load", "-w", daemonPath)

	err = cmd.Run()

	if err != nil {
		fmt.Println("Error was here")
		log.Fatal(err)
	}

	return nil
}

func ServiceUnloadMac(c *cli.Context) {

	daemonPath := GetMacDaemon()
	cmd := exec.Command("launchctl", "unload", "-w", daemonPath)

	fmt.Println("Path:", daemonPath)
	fmt.Println("Cmd:", cmd)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error unloading the system service on:", daemonPath)
		log.Fatal(err)
	}
}
