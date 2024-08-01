package quiz

import "time"

var mockIDTokenUserIDMap = map[string]string{
	"user":  "user",
	"admin": "admin",
}

var mockUserIDRoleMap = map[string]Role{
	"user":  User,
	"admin": Admin,
}

var mockGameSessions = []GameSession{
	{
		ID: "test-session",
		Config: GameSessionConfig{
			BrandID:       "Test",
			MaxPlayers:    40,
			CurrentStage:  0,
			StagePeriod:   10,
			OpenQueueTime: time.Now(),
			Stages: []Stage{
				{
					Question:    "What animal can sleep for most of the winter?",
					Explanation: "Bears hibernate during winter, while the other options are typically active year-round.",
					Options:     []string{"Lion", "Bear", "Monkey", "Dolphin"},
					AnswerIndex: 1,
					Points:      5,
				},
				{
					Question:    "What is the largest country in the world by land area?",
					Explanation: "Russia is the largest country in the world, spanning across two continents.",
					Options:     []string{"Russia", "Canada", "China", "United States"},
					AnswerIndex: 0,
					Points:      5,
				},
				{
					Question:    "Who painted the Mona Lisa?",
					Explanation: "Leonardo da Vinci is the renowned artist behind the iconic Mona Lisa.",
					Options:     []string{"Michelangelo", "Raphael", "Leonardo da Vinci", "Vincent van Gogh"},
					AnswerIndex: 2,
					Points:      7,
				},
				{
					Question:    "What is the capital of Australia?",
					Explanation: "Canberra is the capital city of Australia.",
					Options:     []string{"Sydney", "Melbourne", "Brisbane", "Canberra"},
					AnswerIndex: 3,
					Points:      6,
				},
				{
					Question:    "Which planet is closest to the Sun?",
					Explanation: "Mercury is the closest planet to the Sun.",
					Options:     []string{"Venus", "Earth", "Mars", "Mercury"},
					AnswerIndex: 3,
					Points:      8,
				},
				{
					Question:    "What is the highest mountain in the world?",
					Explanation: "Mount Everest is the tallest mountain above sea level.",
					Options:     []string{"K2", "Mount Kilimanjaro", "Mount McKinley", "Mount Everest"},
					AnswerIndex: 3,
					Points:      10,
				},
			},
		},
		Status: Playing,
	},
}
