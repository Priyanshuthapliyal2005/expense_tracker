package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//import the required packages

// for reading and writing csv files
// for io operation like printing to the console
// for file operations

// for time operations

// gloabl score that is outside the main function
// transaction struct to hold each transaction
type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string // income or expense
	Message  string // optional message
}

// budget tracker struct to mange transaction
type BudgetTracker struct {
	transactions []Transaction
	nextId       int
}

// INTERFACE FOR COMMON BEHAVIOUR
// interface is a key to achieve polymorphism in golang
type FinancialTracker interface {
	GetAmount() float64
	GetType() string
}

// implement interface methods for transation structure
func (t Transaction) GetAmount() float64 {
	return t.Amount
}
func (t Transaction) GetType() string {
	return t.Type
}

// add a new transaction
func (bt *BudgetTracker) AddTransaction(amount float64, category, tType, message string) {
	newTransaction := Transaction{
		ID:       bt.nextId,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
		Message:  message,
	}
	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextId++
}

// creating displayTransaction method
func (bt BudgetTracker) DisplayTransaction() {
	fmt.Println("ID \t Amount\t Category\t Date\t Type\t Message")
	//range
	for _, transaction := range bt.transactions {
		fmt.Printf("%d \t %.2f \t %s \t %s \t %s \t %s\n", transaction.ID, transaction.Amount, transaction.Category, transaction.Date.Format("02-01-2006"), transaction.Type, transaction.Message)
	}
}

// get total incode or expense
func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

// save the transaction to a csv file
func (bt BudgetTracker) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file) //creating a new csv file
	defer writer.Flush()          //flush is very important to make sure that
	//data is written before the file is closed

	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type", "Message"})

	//Write Data
	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("02-01-2006"),
			t.Type,
			t.Message,
		}
		writer.Write(record)
	}
	fmt.Println("Transaction Saved to filename : ", filename)
	return nil
}

func main() {
	// initialisation of budgetTracker struct
	bt := BudgetTracker{}
	reader := bufio.NewReader(os.Stdin) // Added bufio.Reader

	for {
		fmt.Println("Budget Tracker")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Display Transaction")
		fmt.Println("3. Calculate Total Income")
		fmt.Println("4. Calculate Total Expense")
		fmt.Println("5. Save Transaction to CSV")
		fmt.Println("6. Exit")
		fmt.Println("Enter your choice : ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var amount float64
			var category, tType, message string
			fmt.Println("Enter Amount : ")
			fmt.Scanln(&amount)
			fmt.Println("Enter Category : ")
			categoryInput, _ := reader.ReadString('\n') // Changed input method
			category = strings.TrimSpace(categoryInput)
			fmt.Println("Enter Type : ")
			typeInput, _ := reader.ReadString('\n') // Changed input method
			tType = strings.TrimSpace(typeInput)
			fmt.Println("Enter Message (optional) : ")
			messageInput, _ := reader.ReadString('\n')
			message = strings.TrimSpace(messageInput)
			bt.AddTransaction(amount, category, tType, message)
		case 2:
			bt.DisplayTransaction()
		case 3:
			fmt.Println("Total Income : ", bt.CalculateTotal("income"))
		case 4:
			fmt.Println("Total Expense : ", bt.CalculateTotal("expense"))
		case 5:
			fmt.Println("Enter the filename : ")
			var filename string
			fmt.Scanln(&filename)
			bt.SaveToCSV(filename)
		case 6:
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
