package commands

import (
	"github.com/mtardy/kdigger/pkg/bucket"
	"github.com/spf13/cobra"
)

// VERSION indicates which version of the binary is running
var VERSION string

// GITCOMMIT indicates which git hash the binary was built off of
var GITCOMMIT string

// GOVERSION indicates the golang version the binary was built with
var GOVERSION string

// ARCH indicates the arch the binary was built on
var ARCH string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version information",
	Long:    `Print the version tag and git commit hash.`,
	Run: func(cmd *cobra.Command, args []string) {
		// leveraging bucket results to print even if it's not a plugin
		res := bucket.NewResults("Version")
		res.SetHeaders([]string{"Tag", "GITCommit", "GoVersion", "Architecture"})
		res.AddContent([]string{VERSION, GITCOMMIT, GOVERSION, ARCH})

		showName := false
		showComment := false
		printResults(*res, bucket.ResultsOpts{ShowName: &showName, ShowComment: &showComment})
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
