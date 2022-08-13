package uno

import (
	"fmt"
	"os"

	"github.com/glugate/uno/pkg/uno/log"
	"github.com/spf13/cobra"
)

var RootVerbose bool
var CMDLogger = log.DefaultLogFactory().NewLogger()

var rootCmd = &cobra.Command{
	Use:   "uno",
	Short: "Uno is a simple web framework written in Go.",
	Long: `A Fast and with minimal dependencies web framework 
				  SQL database migrations
				  Complete documentation is available at http://uno.ba`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Uno! To run the server run: `go run main.go serve`, or set CMD_DEFAULT=serve in .env file")
	},
}

func Execute() {
	// Try to attach the default specified
	// command to main command args
	cDft, _ := os.LookupEnv("CMD_DEFAULT")
	if cDft != "" {
		rootCmd.SetArgs([]string{cDft})
	}

	// Execute main
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
