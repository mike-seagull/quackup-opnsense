package main

import (
	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/apsdehal/go-logger"
	"golang.org/x/crypto/ssh"
	"os"
	"strconv"
	"time"
	"path"
)
/*
 * pushes the OPNsense config file to a remote server using scp
 */
func scpFile(localFile string, remoteFile string) {
	log, _ := logger.New("scpFile", 0, os.Stdout)
	log.Info("about to scp backup")
	SERVER_USER := os.Getenv("SERVER_USER")
	SERVER_IP := os.Getenv("SERVER_IP")

	clientConfig, err := auth.PrivateKey(SERVER_USER, "/root/.ssh/id_rsa", ssh.InsecureIgnoreHostKey())
	if err != nil {
		log.Fatalf("Error in auth: "+ err.Error())
	}
	client := scp.NewClient(SERVER_IP+":22", &clientConfig)
	err = client.Connect()
	if err != nil {
		log.Fatalf("Couldn't establish a connection to the remote server: "+ err.Error())
	}
	f, _ := os.Open(localFile)
	defer client.Close()
	defer f.Close()
	err = client.CopyFile(f, remoteFile, "0644")
	if err != nil {
		log.Fatalf("Error while copying file: "+ err.Error())
	}
	log.Info("done scp'ing backup")
}

func main() {
	log, _ := logger.New("main", 0, os.Stdout)
	log.Info("started")
	current_time := strconv.FormatInt((time.Now().Local().Unix()), 10) // current unix time
	args := os.Args[1:]
	if len(args) < 1 {
		log.Error("Need a remote_back_up_dir")
		os.Exit(1)
	}
	remote_back_up_dir := args[0]
	log.Info("remote backup directory is: "+ remote_back_up_dir)
	remoteFile := path.Join(remote_back_up_dir, "config." + current_time + ".xml")
	log.Info("remoteFile = "+remoteFile)
	scpFile("/conf/config.xml", remoteFile)
	log.Info("done.")
}
