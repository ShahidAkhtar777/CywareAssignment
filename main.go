package main

import (
	"CywareAssignment/repositories"
	"CywareAssignment/services"
	"CywareAssignment/utils"
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Cron CLI")
	fmt.Println("Type 'exit' to quit the application")

	// TODO: Move configuration's to some utils file main is cluttered.
	// Initialize Viper
	viper.SetConfigFile("local.env") // Name of your local.env file
	viper.AddConfigPath(".")         // Search for the local.env file in the current directory
	viper.AutomaticEnv()             // Automatically read in environment variables

	// Load the local.env file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading local.env file:", err)
	}

	// Initialize the database
	utils.InitializeDB()
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(utils.DB)

	// Create a repository with the Gorm DB
	cronRepository := repositories.NewCronRepository(utils.DB)
	// Create and run the service
	cronService := services.NewCronService(cronRepository)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> Enter Expression (e.g., '* * * *'), 'fetch', or 'exit': ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		} else if strings.ToLower(input) == "fetch" {
			displayRecentCronRuns(cronService)
		} else {
			input = strings.TrimSpace(input)

			if utils.IsValidCronExpression(input) {
				_, err := cronService.CalculateNextRun(input)
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Invalid input. Please enter a valid cron expression or 'fetch'.")
			}
		}
	}

	fmt.Println("Goodbye! ~ Interview Candidate: Shahid Akhtar")
}

func displayRecentCronRuns(cronService *services.CronService) {
	recentCronRuns := cronService.GetRecentCronRuns()
	fmt.Println("\nRecent Cron Job Runs:")
	for _, run := range recentCronRuns {
		fmt.Printf("ID: %d, Next Run At: %s, Cron Expression: %s\n", run.ID, run.RunAt.Format("02-01-2006 15:04:05"), run.Expression)
	}
}

//TODO: General extensions --
// 1. Add more tables to perform more complex operations as well.
// 2. Add indexing in the tables and caching as well if possible.
