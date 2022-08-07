package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	path := flag.String("path", "DataFolder/logs.log", "The path to the log file to be analyzed")
	level := flag.String("level", "Error", "Log level to look for - Error, Info or Warn ")
	flag.Parse()

//logfile
	logFileName := "DataFolder/" + time.Now().Format("2006-01-02") + "log.log"
 
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)	

// end logfile

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	log.Print("File opened successfully")

	r := bufio.NewReader(f)

	newString := ""

	for {
		s,err := r.ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		s = strings.Replace(s,"Information", "Info", -1)	

		newString += s

		if strings.Contains(s, *level){			
			println(s)
		}
	}
	openOrCreateFile("DataFolder/Logs1.log", newString)
}

func openOrCreateFile(path string, toSave string) {
    // check if file exists
    _, err := os.Stat(path)

    // create file if not exists
    if os.IsNotExist(err) {
        file, err := os.Create(path)
        if err!= nil {
            log.Println("Error - " + err.Error())
			return 
        }
        defer file.Close()
		log.Println("File Created Successfully", path)
		file.WriteString(toSave)
    } else {
		file, err := os.OpenFile(path, os.O_RDWR, 0644)
		if err!=nil {
			log.Println("Erroe = " + err.Error())
			return
		}
		defer file.Close()
		log.Println("File opened Successfully", path)	
		file.WriteString(toSave)	
	}	
}
