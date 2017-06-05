package school

import "sort"

const testVersion = 1

// School represents a school
type School struct {
	grades []Grade
}

// Grade represents a grade at a school
type Grade struct {
	grade    int
	students []string
}

// New creates a new school
func New() *School {
	return &School{[]Grade{}}
}

// Add adds a student to a grade
func (s *School) Add(student string, grade int) {
	for g := range s.grades {
		if s.grades[g].grade == grade {
			s.grades[g].add(student)
			return
		}
	}

	s.grades = append(s.grades, Grade{grade: grade, students: []string{student}})
}

func (g *Grade) add(student string) {
	s := sort.SearchStrings(g.students, student)
	if s < len(g.students) {
		g.students = append(g.students[0:s], append([]string{student}, g.students[s:]...)...)
	} else {
		g.students = append(g.students, student)
	}
}

// Grade returns the members of a grade
func (s *School) Grade(grade int) []string {
	for _, g := range s.grades {
		if g.grade == grade {
			return g.students
		}
	}

	return []string{}
}

// Enrollment returns the grades for the school
func (s *School) Enrollment() []Grade {
	sort.Slice(s.grades, func(i, j int) bool { return s.grades[i].grade < s.grades[j].grade })
	return s.grades
}
