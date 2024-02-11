package markdown

import (
	d "JimD/models"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func IsExerciseFormat(input string) bool {
	// Define the regular expression for the generic pattern
	pattern := regexp.MustCompile(`^[a-zA-Z]+ - \d+ x \d+ of \d+ kgs\.$`)

	// Check if the input string matches the pattern
	return pattern.MatchString(input)
}

func ParseExerciseString(input string) (*d.Exercise, error) {
	fmt.Println("one")
	// Define the regular expression pattern to extract information
	pattern := `^(?i)([a-z]+) - (\d+) x (\d+) of (\d+(\.\d+)?) kgs\.$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Find matches in the input string
	matches := regex.FindStringSubmatch(input)

	// Check if the input string matches the expected pattern
	if len(matches) != 6 {
		return nil, fmt.Errorf("input string does not match the expected pattern")
	}

	// Extract information from the matches
	name := matches[1]
	reps, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	sets, err := strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}

	weight, err := strconv.ParseFloat(matches[4], 64)
	if err != nil {
		panic(err)
	}

	// Create and return the Exercise struct
	exercise := &d.Exercise{
		Name:   name,
		Reps:   reps,
		Sets:   sets,
		Weight: weight,
	}
	fmt.Println("name - ", exercise.Name)
	return exercise, nil
}

func HasDateInFirstLine(filename string) (bool, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the first line
	if scanner.Scan() {
		// Extract the date from the first line using a regular expression
		datePattern := `\d{2}/\d{2}/\d{4} \d{2}:\d{2}`
		firstLine := scanner.Text()
		matched, err := regexp.MatchString(datePattern, firstLine)
		if err != nil {
			return false, err
		}

		// If a date is found, check if it's a valid date
		if matched {
			_, err := time.Parse("02/01/2006 15:04", firstLine)
			if err == nil {
				// Valid date format
				return true, nil
			}
		}
	}

	// No date found in the first line
	return false, nil
}
