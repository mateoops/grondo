package sshutil

import "testing"

func TestDecodeBase64(t *testing.T) {
	validBase64 := "dGVzdA==" // "test"
	invalidBase64 := "NotB64"

	validResult := decodeBase64(validBase64)
	invalidResult := decodeBase64(invalidBase64)

	if validResult != "test" {
		t.Errorf("decodeBase64(%s) = %s; want test", validBase64, validResult)
	}
	if invalidResult != "GRONDO B64 DECODING ERROR" {
		t.Errorf("decodeBase64(%s) = %s; want GRONDO B64 DECODING ERROR", invalidBase64, invalidResult)
	}
}

func TestRunCommandOverSSH(t *testing.T) {
	sshHost := "test"
	sshPort := uint(22)
	sshUser := "test"
	sshPasswordOrKey := "LS0tLS1CRUdJTiBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS1OT1R2YWxpZEI2NAo=" // not valid B64 key, but with valid prefix
	command := "test"

	result, status := RunCommandOverSSH(sshHost, sshPort, sshUser, sshPasswordOrKey, command)

	if status != "Error parsing Private Key: ssh: no key found" {
		t.Errorf("Incorrect status: %s; want 'Error parsing Private Key: ssh: no key found'", status)
	}

	if result != "ERROR" {
		t.Errorf("Incorrect result: %s; want 'ERROR'", result)
	}

}
