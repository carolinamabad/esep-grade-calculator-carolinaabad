package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestComputeAverageEmpty(t *testing.T) {
	if got := computeAverage([]Grade{}); got != 0 {
		t.Fatalf("avg(empty) = %d, want 0", got)
	}
}

func TestComputeAverageSingle(t *testing.T) {
	got := computeAverage([]Grade{{Name: "x", Grade: 80, Type: Assignment}})
	if got != 80 {
		t.Fatalf("avg(single) = %d, want 80", got)
	}
}

func TestAddGradeRoutes(t *testing.T) {
	calc := NewGradeCalculator()
	calc.AddGrade("a1", 90, Assignment)
	calc.AddGrade("e1", 85, Exam)
	calc.AddGrade("s1", 88, Essay)

	if computeAverage(calc.assignments) != 90 { t.Fatalf("assignments avg wrong") }
	if computeAverage(calc.exams) != 85 { t.Fatalf("exams avg wrong") }
	if computeAverage(calc.essays) != 88 { t.Fatalf("essays avg wrong") }
}

func TestLetterBoundaries(t *testing.T) {
	type in struct{ a, e, s int }
	cases := []struct{
		name string
		g in
		want string
	}{
		{"A_90", in{90,90,90}, "A"},
		{"B_80", in{80,80,80}, "B"},
		{"C_70", in{70,70,70}, "C"},
		{"D_60", in{60,60,60}, "D"},
		{"F_59", in{59,59,59}, "F"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calc := NewGradeCalculator()
			calc.AddGrade("a", tc.g.a, Assignment)
			calc.AddGrade("e", tc.g.e, Exam)
			calc.AddGrade("s", tc.g.s, Essay)
			if got := calc.GetGrade(); got != tc.want {
				t.Fatalf("GetGrade() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestGradeTypeString(t *testing.T) {
    if Assignment.String() != "assignment" {
        t.Errorf("Expected Assignment.String() to be 'assignment'")
    }
    if Exam.String() != "exam" {
        t.Errorf("Expected Exam.String() to be 'exam'")
    }
    if Essay.String() != "essay" {
        t.Errorf("Expected Essay.String() to be 'essay'")
    }
}