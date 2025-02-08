package cmd

import (
	"fmt"
	"log"

	"github.com/abhishekkushwahaa/secure-cloud-cli/internal/auth"
	"github.com/abhishekkushwahaa/secure-cloud-cli/internal/cloud"
	"github.com/spf13/cobra"
)

var download = &cobra.Command{
	Use:   "download [filename]",
	Short: "Download and Decrypt a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := auth.LoadSession()
		if err != nil {
			log.Fatal("❌ You must login first. Use `go run . login`")
		}

		filename := args[0]
		fmt.Println("🔄 Downloading file:", filename)

		err = cloud.DownloadFromS3(filename)
		if err != nil {
			log.Fatalf("❌ Download failed: %v", err)
		}

		fmt.Println("✅ File downloaded and decrypted successfully")
	},
}
