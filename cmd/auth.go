package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/abhishekkushwahaa/secure-cloud-cli/db"
	"github.com/abhishekkushwahaa/secure-cloud-cli/internal/auth"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var register = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("👤 Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("🔑 Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		_, err := db.DB.Exec("INSERT INTO users (username, hashedPassword) VALUES ($1, $2)", username, string(hashedPassword))

		if err != nil {
			fmt.Println("👍 User already exists")
		} else {
			fmt.Println("✅ User registered successfully")
		}
	},
}

var login = &cobra.Command{
	Use:   "login",
	Short: "Login a user",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("👤 Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("🔑 Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		var hashedPassword string
		err := db.DB.QueryRow("SELECT hashedPassword FROM users WHERE username=$1", username).Scan(&hashedPassword)
		if err == sql.ErrNoRows {
			fmt.Println("👎 User does not exist.")
			return
		} else if err != nil {
			log.Fatal("Database error:", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			fmt.Println("❌ Invalid credentials")
			return
		}

		err = auth.SaveSession(username)
		if err != nil {
			log.Fatalf("❌ Failed to save session: %v", err)
		}

		fmt.Println("✅ Login successful")
	},
}

var logout = &cobra.Command{
	Use:   "logout",
	Short: "Logout a user",
	Run: func(cmd *cobra.Command, args []string) {
		err := auth.ClearSession()
		if err != nil {
			log.Fatal("❌ Logout failed:", err)
		}
		fmt.Println("✅ Logged out successfully!")
	},
}
