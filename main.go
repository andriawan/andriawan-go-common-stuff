package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/golang-module/carbon/v2"
)

func printBanner() {
	color.Set(color.FgHiYellow, color.Bold).Println(printHeader("Andriawan Stuff"))
	color.Set(color.FgCyan).Println("Please choose the options below:")
	color.Set(color.FgCyan).Println("1) Age Calculator")
	color.Set(color.FgCyan).Println("2) Date Diff Calculator")
	color.Set(color.FgCyan).Println("3) URL GET Fetcher")
	fmt.Println()
}

func printHeader(message string) string {
	var text bytes.Buffer
	text.WriteString("\n==========================\n")
	text.WriteString(message)
	text.WriteString("\n==========================\n")
	return text.String()
}

func doAgeCalculation(scanner *bufio.Scanner) {
	color.Set(color.FgHiYellow, color.Bold)
	fmt.Println(printHeader("Age Calculator"))
	color.Unset()
	color.Set(color.FgHiGreen)
	fmt.Print("Please input year: ")
	scanner.Scan()
	input := scanner.Text()
	color.Unset()
	date := carbon.Parse(input)

	if date.Error != nil {
		color.HiRed(date.Error.Error())
		os.Exit(1)
	}

	yearNow := carbon.Now().Year()
	yearInput := date.Year()

	if yearInput > yearNow {
		log.Fatal("You are a liar")
	}

	age := yearNow - yearInput

	color.Set(color.FgHiGreen, color.Underline)
	fmt.Printf("You are %d years old\n", age)
	color.Unset()
	fmt.Print("\nPlease enter to continue...")
	scanner.Scan()
	renderMainMenu()
}

func doDateDiffCalculation(scanner *bufio.Scanner) {
	color.New(color.FgHiYellow, color.Bold).Println(printHeader("Date Diff Calculator"))
	color.HiGreen("Please input start date: ")
	scanner.Scan()
	startDate := scanner.Text()
	color.HiGreen("Please input end date: ")
	scanner.Scan()
	endDate := scanner.Text()
	carbonStartDate := carbon.Parse(startDate)
	if carbonStartDate.Error != nil {
		color.HiRed(carbonStartDate.Error.Error())
		os.Exit(1)
	}
	carbonEndDate := carbon.Parse(endDate)
	if carbonEndDate.Error != nil {
		color.HiRed(carbonEndDate.Error.Error())
		os.Exit(1)
	}
	diff := carbonStartDate.DiffAbsInDays(carbonEndDate)
	color.New(color.FgHiYellow, color.Bold).Printf("diff %d days\n", diff)
	fmt.Print("\nPlease enter to continue...")
	scanner.Scan()
	renderMainMenu()
}

func renderMainMenu() {
	printBanner()
	scanner := bufio.NewScanner(os.Stdin)
	color.HiGreen("Please input option: ")
	scanner.Scan()
	option := scanner.Text()
	optionInt, err := strconv.Atoi(option)
	if err != nil {
		color.HiRed(err.Error())
	}

	switch optionInt {
	case 1:
		doAgeCalculation(scanner)
	case 2:
		doDateDiffCalculation(scanner)
	case 3:
		doUrlFetch(scanner)
	default:
		color.HiRed("incorrect option")
		renderMainMenu()
	}
}

func doUrlFetch(scanner *bufio.Scanner) {
	color.New(color.FgHiYellow, color.Bold).Println(printHeader("URL GET Fetcher"))
	color.HiGreen("Please input URL: ")
	scanner.Scan()
	url := scanner.Text()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		color.HiRed(err.Error())
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.HiRed(err.Error())
		os.Exit(1)
	}
	elapsed := time.Since(start)
	color.HiGreen(string(body))
	color.HiGreen("execution time : " + elapsed.String())
	fmt.Print("\nPlease enter to continue...")
	scanner.Scan()
	renderMainMenu()
}

func main() {
	renderMainMenu()
}
