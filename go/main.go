package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "-h", "--help":
		showHelp()

	case "-l", "--history":
		showHistory()

	case "-s", "--stats":
		if len(os.Args) < 3 {
			fmt.Println("Error: Invalid format. Use 'dice_tools -s <dice_roll>' for stats.")
			os.Exit(1)
			fmt.Println("ops")
		}
		mathStats(os.Args[2])

	default:
		dicePattern := regexp.MustCompile(`^(\d+)d(\d+)$`)
		if dicePattern.MatchString(arg) {
			parts := strings.Split(arg, "d")
			diceRoll(parts)
		} else {
			fmt.Println("Error: Invalid format. Use 'dice_tools -h' for help.")
			os.Exit(1)
			fmt.Println("ops")
		}
	}
}

func diceRoll(parts []string) int {
	result := []int{}
	total := 0

	dice_num, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Error: ", err)
		return 0
	}
	dice_face, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return 0
	}

	for i := 0; i < dice_num; i++ {
		result = append(result, rand.IntN(dice_face)+1)
		total += result[i]
	}

	fmt.Println("Result:", result)
	fmt.Println("Total:", total)

	writeLog(total, result)
	return 1
}

func showHelp() {
	helpText := `Dice Tools - A simple dice rolling tool

Usage:
  <N>d<M>        Roll N dice with M sides
  --history      Show roll history (last 100 rolls)
  --stats <M>    Roll 1d<M> 10,000 times and show statistics
  --help         Show this help message

Examples:
  2d6            Roll two 6-sided dice
  1d20           Roll one 20-sided dice
  --stats 6      Get statistics for 10,000 rolls of 1d6

Options:
  -l, --history            Show roll history
  -s, --stats <M>          Show statistics
  -h, --help               Show help`
	fmt.Println(helpText)
}

func showHistory() {
	content, err := os.ReadFile("dice_log.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("=== Roll History (Last 100 rolls) ===")
	fmt.Println()
	fmt.Println(string(content))
}

func writeLog(total int, result []int) {
	timestamp := time.Now().Format("2006-01-02 15:04")
	newLog := fmt.Sprintf("%s: %d %v", timestamp, total, result)

	lines := []string{}
	content, err := os.ReadFile("dice_log.txt")
	if err == nil && len(content) > 0 {
		text := strings.TrimSpace(string(content))
		if text != "" {
			lines = strings.Split(text, "\n")
		}
	}

	lines = append(lines, newLog)

	if len(lines) > 100 {
		lines = lines[len(lines)-100:]
	}

	file, err := os.OpenFile("dice_log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		file.WriteString(line + "\n")
	}
}

func mathStats(arg string) {

}
