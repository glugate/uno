package uno

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	kh "golang.org/x/crypto/ssh/knownhosts"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys the website on Digital Ocean",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deploy...")

		user, _ := os.LookupEnv("DO_USER")
		address, _ := os.LookupEnv("DO_ADDRESS")
		keyPath, _ := os.LookupEnv("DO_KEY_PATH")
		knownHostsPath, _ := os.LookupEnv("KNOWN_HOSTS_PATH")
		command, _ := os.LookupEnv("DO_DEPLOY_CMD")
		port := "22"

		key, err := os.ReadFile(keyPath)
		if err != nil {
			log.Fatalf("unable to read private key: %v", err)
		}

		// Create the Signer for this private key.
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("unable to parse private key: %v", err)
		}

		hostKeyCallback, err := kh.New(knownHostsPath)
		if err != nil {
			log.Fatal("could not create hostkeycallback function: ", err)
		}

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				// Add in password check here for moar security.
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: hostKeyCallback,
		}
		// Connect to the remote server and perform the SSH handshake.
		client, err := ssh.Dial("tcp", address+":"+port, config)
		if err != nil {
			log.Fatalf("unable to connect: %v", err)
		}
		defer client.Close()
		ss, err := client.NewSession()
		if err != nil {
			log.Fatal("unable to create SSH session: ", err)
		}
		defer ss.Close()
		// Creating the buffer which will hold the remotly executed command's output.
		var stdoutBuf bytes.Buffer
		ss.Stdout = &stdoutBuf
		ss.Run(command)
		// Let's print out the result of command.
		fmt.Println(stdoutBuf.String())

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
