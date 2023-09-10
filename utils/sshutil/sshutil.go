package sshutil

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func decodeBase64(s string) string {

	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "GRONDO B64 DECODING ERROR"
	}

	return string(decodedBytes)
}

func RunCommandOverSSH(sshHost string, sshPort uint, sshUser string, sshPasswordOrKey string, command string) (status string, output string) {

	sshPasswordOrKeyDecoded := decodeBase64(sshPasswordOrKey)
	if sshPasswordOrKeyDecoded == "GRONDO B64 DECODING ERROR" {
		return "ERROR", "Error decoding Base64 string"
	}
	auth := ssh.Password(strings.ReplaceAll(sshPasswordOrKeyDecoded, "\n", ""))

	if strings.Contains(sshPasswordOrKeyDecoded, "-----BEGIN OPENSSH PRIVATE KEY-----") {
		sshPasswordOrKeyDecoded = strings.ReplaceAll(sshPasswordOrKeyDecoded, "\t", "")
		privateKey, err := ssh.ParsePrivateKey([]byte(sshPasswordOrKeyDecoded))
		if err != nil {
			return "ERROR", "Error parsing Private Key: " + err.Error()
		}

		auth = ssh.PublicKeys(privateKey)
	}
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			auth,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err := ssh.Dial("tcp", sshHost+":"+strconv.FormatUint(uint64(sshPort), 10), sshConfig)
	if err != nil {
		return "ERROR", "Error connection to remote host: " + err.Error()
	}
	defer sshClient.Close()

	session, err := sshClient.NewSession()
	if err != nil {
		return "ERROR", "Creating SSH session error: " + err.Error()
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	err = session.Run(command)
	if err != nil {
		return "ERROR", "Error running remote command:" + err.Error()
	}
	defer session.Close()

	output = stderr.String() + stdout.String()

	return "SUCCESS", output
}
