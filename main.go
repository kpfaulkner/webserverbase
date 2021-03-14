package main

import (
	"flag"
	"fmt"
	"github.com/kpfaulkner/webserverbase/pkg/server"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func InitializeLogging(logFile string) {

	var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	fmt.Printf("so it begins...\n")

	useTLS := flag.Bool("usetls", false, "Listen for TLS traffic")
  certCRTPath := flag.String("crtpath", "", "Path to Certificate CRT file")
	certKeyPath := flag.String("keypath", "", "Path to Certificate Key file")
	jwtSecret := flag.String("jwtSecret", server.DefaultJWTSecret, "JWT Secret")
	port := flag.String("port", "8080", "Port for server to listen on")

	InitializeLogging("server.log")

	svr := server.NewServer(*jwtSecret)
	svr.Use(server.WithLogging())

	var err error
	portStr := fmt.Sprintf(":%s", *port)
	if useTLS != nil && *useTLS == true  &&
		certCRTPath != nil && *certCRTPath != "" &&
		certKeyPath != nil && *certKeyPath != "" {
		err = http.ListenAndServeTLS(portStr, *certCRTPath, *certKeyPath, svr.GetServerWithMiddleware())
	} else {
		err = http.ListenAndServe(portStr, svr.GetServerWithMiddleware())
	}

	fmt.Printf("err is %s\n", err.Error())
}
