package cmd

import (
	"fmt"
	"log"

	"github.com/abhishekkushwahaa/secure-cloud-cli/internal/auth"
	"github.com/abhishekkushwahaa/secure-cloud-cli/internal/cloud"
	"github.com/spf13/cobra"
)

var upload = &cobra.Command{
	Use:   "upload [path]",
	Short: "Upload and Encrypt a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		_, err := auth.LoadSession()
		if err != nil {
			log.Fatal("❌ You must login first. Use `go run . login`")
		}

		filePath := args[0]
		fmt.Println("🔄 Uploading file:", filePath)

		err = cloud.UploadToS3(filePath)
		if err != nil {
			log.Fatalf("❌ Upload failed: %v", err)
		}

		fmt.Println("✅ File uploaded successfully")
	},
}
