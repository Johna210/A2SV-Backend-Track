package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main(){
	reader := bufio.NewReader(os.Stdin)
	student := Student{}
	
	// Read Student Name
	for {
		fmt.Print("Enter Your Full Name: ")
		studentName, _ := reader.ReadString('\n')
		studentName = strings.TrimSpace(studentName)

		if studentName == "" {
			fmt.Println("Invalid Input. Please enter a valid name.")
		}else {
			student.Name = studentName
			break
		}
	}
	
	// Read number of subjects
	var subjectCount int

	for {
		fmt.Print("Enter number of subjects: ")
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parsedInt, err := strconv.Atoi(input)
		if err != nil || parsedInt <= 0 || parsedInt > 10 {
			fmt.Println("Invalid input. Please enter a number between 1 and 10.")
		} else {
			subjectCount = parsedInt
			break
		}
	}
	
	fmt.Println()

	// Read Subject Name and Grade
	for i:=0; i < subjectCount; i++ {
		subject := Subject{}
		for {
			fmt.Print("Enter Subject Name: ")
			subjectName, _ := reader.ReadString('\n')
			subjectName = strings.TrimSpace(subjectName)

			if subjectName == "" {
				fmt.Println("Invalid Input. Please enter a valid name.")
			}else {
				subject.Name = subjectName
				break
			}
		}

		for {
			fmt.Print("Enter Grade: ")
			input,_ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			parsedInt, err := strconv.Atoi(input)
			if err != nil || parsedInt < 0 || parsedInt > 100 {
				fmt.Println("Invalid input. Please enter a number between 0 and 100.")
			} else {
				subject.Grade = parsedInt
				break
			}
		}
		student.AddSubject(subject)
		fmt.Println("-----------------------")
	}

	// Display Grade
	fmt.Println()
	student.display()

}