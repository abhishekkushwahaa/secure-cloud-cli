package cmd

import (
	"fmt"
	"log"

	"github.com/abhishekkushwahaa/secure-cloud-cli/internal/cloud"
	"github.com/spf13/cobra"
)

var download = &cobra.Command{
	Use:   "download [filename]",
	Short: "Download and Decrypt a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		fmt.Println("🔄 Downloading file:", filename)

		err := cloud.DownloadFromS3(filename)
		if err != nil {
			log.Fatalf("❌ Download failed: %v", err)
		}

		fmt.Println("✅ File downloaded successfully")
	},
}
