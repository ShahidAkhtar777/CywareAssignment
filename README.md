
# Cron CLI Application

This is a simple Command Line Interface (CLI) application that allows users to input cron expressions, calculate the next run time, and retrieve recent cron job runs.

## Getting Started

To run the application, you need a working Go environment.

1. Clone the repository:
    ```
    git clone https://github.com/ShahidAkhtar777/CywareAssignment.git
    cd CywareAssignment
   ```

2. Build and run the application:
   ```
    go build
    ./CywareAssignment
   ```

## Usage

The application provides the following features:

### Input Cron Expressions

You can input cron expressions in the format `"* * * *"` where each asterisk represents a field (minute, hour, day of the month, month).

For example, to schedule a cron job to run every minute, you can input `"* * * *"`.

### Calculate Next Run

The application calculates and displays the next run time based on the provided cron expression. The result is shown in the format `"dd-mm-yyyy hh:mm:ss"`.

### Fetch Recent Cron Job Runs

You can enter "fetch" to retrieve and display all the recent cron job runs. This option provides information about past executions, including their IDs, next run times, and cron expressions.

### Exit the Application

To exit the application, simply enter "exit."

## Code Structure

The application consists of the following components:

- `main.go`: Contains the main entry point for the CLI application. It handles user input, including cron expressions, "fetch" requests, and exiting the application.

- `models`: This package defines the data models used in the application, such as `CronRun` to represent cron job runs.

- `repositories`: The repository package manages data storage for cron job runs.

- `services`: The service package contains the `CronService`, which handles the core logic of calculating next run times and fetching recent cron job runs.

## Customization

You can customize the code to add additional features or modify the behavior as needed. The application currently provides a basic framework for working with cron expressions and storing execution history.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests. We welcome improvements, bug fixes, and new features.
