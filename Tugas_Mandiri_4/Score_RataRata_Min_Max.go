package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

func (s Student) AverageScore(students []Student) float64 {
	total := 0
	for _, student := range students {
		total += student.Score
	}
	average := float64(total) / float64(len(students))
	return average
}

func (s Student) MinScoreStudent(students []Student) Student {
	minScore := math.MaxInt32
	var minStudent Student
	for _, student := range students {
		if student.Score < minScore {
			minScore = student.Score
			minStudent = student
		}
	}
	return minStudent
}

func (s Student) MaxScoreStudent(students []Student) Student {
	maxScore := math.MinInt32
	var maxStudent Student
	for _, student := range students {
		if student.Score > maxScore {
			maxScore = student.Score
			maxStudent = student
		}
	}
	return maxStudent
}

func main() {
	var students []Student

	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scanln(&name)

		fmt.Printf("Input %d Student's Score: ", i+1)
		fmt.Scanln(&score)

		student := Student{Name: name, Score: score}
		students = append(students, student)
	}

	average := students[0].AverageScore(students)
	minStudent := students[0].MinScoreStudent(students)
	maxStudent := students[0].MaxScoreStudent(students)

	fmt.Printf("Average Score: %.2f\n", average)
	fmt.Printf("Min Score of Students: %s (%d)\n", minStudent.Name, minStudent.Score)
	fmt.Printf("Max Score of Students: %s (%d)\n", maxStudent.Name, maxStudent.Score)
}
