package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

func ScanSSH(ip string, port int, timeout time.Duration, service, username, password string) (result bool, err error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout: timeout,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", ip, port), config)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		errRet := session.Run("echo xsec")
		if err == nil && errRet == nil {
			defer session.Close()
			result = true
		}
	}

	return result, err
}

func main() {
	ip := "10.99.60.15"
	port := 22
	timeout := 3 * time.Second
	service := "ssh"
	username := "yunying"
	password := "Yunzxa@188"
	result, err := ScanSSH(ip, port, timeout, service, username, password)
	fmt.Printf("check %v service, %v:%v, result: %v, err: %v\n", service, ip, port, result, err)
}
