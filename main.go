package main

import (
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
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
	SERVER_USER := os.Getenv("SERVER_USER")
	SERVER_IP := os.Getenv("SERVER_IP")

	clientConfig, auth_err := auth.PrivateKey(SERVER_USER, "/root/.ssh/id_rsa", ssh.InsecureIgnoreHostKey())
	// clientConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	if auth_err != nil {
		fmt.Println("Error in auth", auth_err)
		return
	}
	client := scp.NewClient(SERVER_IP+":22", &clientConfig)
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}
	f, _ := os.Open(localFile)
	defer client.Close()
	defer f.Close()
	err = client.CopyFile(f, remoteFile, "0644")
	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}

func main() {
	current_time := strconv.FormatInt((time.Now().Local().Unix()), 10) // current unix time
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Need a remote_back_up_dir")
		os.Exit(1)
	}
	remote_back_up_dir := args[0]
	remoteFile := path.Join(remote_back_up_dir, "config." + current_time + ".xml")
	scpFile("/conf/config.xml", remoteFile)
}
