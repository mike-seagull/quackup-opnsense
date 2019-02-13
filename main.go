package main

import (
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func scpFile(localFile string) {
	SERVER_USER := os.Getenv("SERVER_USER")
	SERVER_HOST := os.Getenv("SERVER_HOST")

	// Use SSH key authentication from the auth package
	// we ignore the host key in this example, please change this if you use this library
	clientConfig, _ := auth.PrivateKey(SERVER_USER, "/path/to/rsa/key", ssh.InsecureIgnoreHostKey())

	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod

	// Create a new SCP client
	client := scp.NewClient(SERVER_HOST+":22", &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}

	// Open a file
	f, _ := os.Open(localFile)

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	err = client.CopyFile(f, "/home/server/test.txt", "0655")

	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}
func copyFile(sourceFile string, destFile string) {
	/*
	 * copies a file
	 * https://opensource.com/article/18/6/copying-files-go
	 */
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destFile)
		fmt.Println(err)
		return
	}
}

func main() {
	current_time := strconv.FormatInt((time.Now().Local().Unix()), 10) // current unix time
	destFile := "/usr/local/etc/backup-opnsense/config." + current_time + ".xml"
	copyFile("/conf/config.xml", destFile)
}
