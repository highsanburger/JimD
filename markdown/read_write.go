package markdown

import (
	d "JimD/models"
	"bufio"
	"fmt"
	"os"
	"time"
)

func EnterDate(filename string) {
	time := time.Now().Format("02/01/2006 15:04")
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err2 := file.WriteString(time + "\n")
	if err2 != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func EnterEx(filename string, ex d.Exercise) {
	exs := ex.Name + " - " + fmt.Sprint(ex.Sets) + " x " + fmt.Sprint(ex.Reps) + " of " + fmt.Sprint(ex.Weight) + " kgs."
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err2 := file.WriteString(exs + "\n")
	if err2 != nil {
		fmt.Println("Error writing to file:", err)
	}

}

func ReadLinesFromFile(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize an empty slice to store lines
	var lines []string

	// Read lines and append to the slice
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	lines = lines[1:]

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
