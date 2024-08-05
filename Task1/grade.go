package main

import "fmt"

type Subject struct {
	Name string
	Grade float64
}

type Student struct {
	Name string
	Subjects []Subject
}

func (s *Student) AddSubject(subject Subject) {
	s.Subjects = append(s.Subjects, subject)
}

func (s *Student) GetTotal() float64 {
	var total float64
	for _, subject := range s.Subjects {
		total += subject.Grade
	}
	return total
}

func (s *Student) GetAverage() float64 {
	total := s.GetTotal()
	return float64(total) / float64(len(s.Subjects))
}

func (s *Student) display() {
	fmt.Println("Student Name:", s.Name)
	fmt.Printf("%-20s : %s\n", "Subjects", "Grades")
	for _, subject := range s.Subjects {
		fmt.Printf("%-20s : %f\n", subject.Name, subject.Grade)
	}
	fmt.Printf("%-20s : %f\n", "Total", s.GetTotal())
	fmt.Printf("%-20s : %.2f\n", "Average", s.GetAverage())
}