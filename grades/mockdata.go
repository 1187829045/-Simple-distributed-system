package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Roberto",
			LastName:  "Baagio",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 100,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 81,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Emma",
			LastName:  "Stone",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 67,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 0,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 75,
				},
			},
		},
	}
}
