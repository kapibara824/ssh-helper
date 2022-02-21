package sshhelper

import (
	"bytes"

	"golang.org/x/crypto/ssh"
)

type Config struct {
	Host     string
	User     string
	Password string
}

func Run(config Config, cmd string) (string, error) {
	c := &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", config.Host, c)
	if err != nil {
		return "", err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var buf bytes.Buffer

	if err := session.Run(cmd); err != nil {
		return buf.String(), err
	}
	return buf.String(), nil
}
