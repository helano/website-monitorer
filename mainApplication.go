package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const waitBetweenEachMonitoringSession = 5
const monitorCounter = 5

func main() {

	exibirIntroducao()
	for {

		exibirMenu()
		command := readCommand()

		switch command {
		case 1:
			monitorWebSite()
		case 2:
			fmt.Println("Show logs")
			showLogs()
		case 0:
			fmt.Println("Exiting program")
			os.Exit(0)
		default:
			fmt.Println("Command not found - exiting with error")
			os.Exit(-1)

		}
	}

}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error when reading file:", err)

	}
	fmt.Println("Printing log information...")
	fmt.Println(string(file))
}

func monitorWebSite() {
	fmt.Println("which web-site do you wanna check-out?")
	listOfWebSites := readWebSitesFromFile()

	for i := 0; i < monitorCounter; i++ {
		for i, webSite := range listOfWebSites {
			fmt.Println("Testing item:", i, webSite)
			testWebSite(webSite)
		}
		time.Sleep(waitBetweenEachMonitoringSession)
	}
}

func readWebSitesFromFile() []string {
	file, err := os.Open("web-sites.txt")

	var webSites []string

	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			println("Finish reading file", err)
			break
		}
		webSites = append(webSites, line)
	}
	file.Close()
	return webSites
}

func testWebSite(webSite string) {
	response, _ := http.Get(webSite)

	if response.StatusCode == 200 {
		fmt.Println("Website is doing fine, status-code:", response.StatusCode)
		persistLog(webSite, true)
	} else {
		fmt.Println("Website return a different status code:", response.Status)
		persistLog(webSite, false)
	}
}

func persistLog(webSite string, flag bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error when reading log.txt file:", err)
	}
	file.WriteString(time.Now().Format("02-01-2006 15:04:05") + "-" + webSite + " - online :" + strconv.FormatBool(flag) + "\n")

	file.Close()
}

func readCommand() int16 {
	var command int16 = -1
	fmt.Scan(&command)
	return command
}

func exibirMenu() {
	fmt.Println(" 1  - Start monitoring")
	fmt.Println(" 2 - Show logs")
	fmt.Println(" 0 - Quit program")
}

func exibirIntroducao() {
	name := "helano"
	version := 1.1
	fmt.Println(" Hello, Sr.", name)
	fmt.Println("Welcome to the log monitorer - version:", version)

}
