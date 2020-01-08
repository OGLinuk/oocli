package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSH executes given cmd in an SSH session with the Onion Omega2
// NOTE: There are many ways to store credentials, I suggest
// environment variables ~ ie export OOHOST="192.168.3.1"
func SSH(cmd string) {
	host := os.Getenv("OOHOST")              // args[0] //"192.168.3.1"
	port := "22"                             // SSH port
	addr := fmt.Sprintf("%s:%s", host, port) // "192.168.3.1:22"
	user := os.Getenv("OOUSER")              // args[1] //"root"
	pass := os.Getenv("OOPASS")              // args[2] //"onioneer"

	// get Onion Omega2 public key
	ooHostKey := getOOPK(host)

	// ssh client config
	conf := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		// allow any host key to be used (non-prod)
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		// verify host public key
		HostKeyCallback: ssh.FixedHostKey(ooHostKey),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", addr, conf)
	if err != nil {
		log.Printf("ssh.go::SSH::ssh.Dial(\"tcp\", %s, %v)::ERROR: %s", addr, conf, err.Error())
	}
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		log.Printf("ssh.go::SSH::client.NewSession()::ERROR: %s", err.Error())
	}
	defer sess.Close()

	// Refactor below to output to log file
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	log.Printf("Executing command %s ...", cmd)

	err = sess.Run(cmd)
	if err != nil {
		log.Printf("ssh.go::SSH::sess.Run(%s)::ERROR: %s", cmd, err.Error())
	}
}

// getOOPK gets and returns the public key for the Onion Omega2
func getOOPK(host string) ssh.PublicKey {
	knownHostsPath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	f, err := os.Open(knownHostsPath)
	if err != nil {
		log.Printf("ssh.go::getOOPK::os.Open(%s)::ERROR: %s", knownHostsPath, err.Error())
	}
	defer f.Close()

	var pk ssh.PublicKey
	bscanner := bufio.NewScanner(f)

	for bscanner.Scan() {
		scanned := bscanner.Text()
		fields := strings.Split(scanned, " ")
		if len(fields) != 3 {
			continue
		}

		if strings.Contains(fields[0], host) {
			var err error
			b := bscanner.Bytes()
			pk, _, _, _, err = ssh.ParseAuthorizedKey(b)
			if err != nil {
				log.Printf("ssh.go::getOOPK::ssh.ParseAuthorizedKey(%v)::ERROR: %s | %s", b, fields[2], err.Error())
			}
			break
		}
	}

	if pk == nil {
		log.Printf("ssh.go::getOOPK::pk::NIL: %s", pk)
	}

	return pk
}
