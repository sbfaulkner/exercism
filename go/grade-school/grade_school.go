package school

import "sort"

const testVersion = 1

// School represents a school
// type School struct {
// 	grades map[int]*Grade
// }
type School map[int]*Grade

// Grade represents a grade at a school
type Grade struct {
	grade    int
	students []string
}

// New creates a new school
func New() *School {
	return &School{}
}

// Add adds a student to a grade
func (s *School) Add(student string, grade int) {
	if _, ok := (*s)[grade]; !ok {
		(*s)[grade] = &Grade{grade: grade, students: []string{}}
	}

	(*s)[grade].add(student)
}

// Grade returns the members of a grade
func (s *School) Grade(grade int) []string {
	if _, ok := (*s)[grade]; !ok {
		return []string{}
	}

	return (*s)[grade].students
}

// Enrollment returns the grades for the school
func (s *School) Enrollment() (enrollment []Grade) {
	var grades []int

	for g := range *s {
		grades = append(grades, g)
	}

	sort.Ints(grades)

	for _, g := range grades {
		enrollment = append(enrollment, *(*s)[g])
	}

	return
}

func (g *Grade) add(student string) {
	g.students = append(g.students, student)
	sort.Strings(g.students)
}
