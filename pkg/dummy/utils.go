package dummy

import (
	"time"
)

type DummyData struct {
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
	Message   string    `json:"message"`
}

func createDummyArrays() []*DummyData {
	dummyDummies := make([]*DummyData, 0)

	dummyDummies = append(dummyDummies, &DummyData{
		Timestamp: time.Now(),
		ID:        "1",
		Message:   "This is the first dummy",
	})

	dummyDummies = append(dummyDummies, &DummyData{
		Timestamp: time.Now(),
		ID:        "2",
		Message:   "This is the second dummy",
	})

	dummyDummies = append(dummyDummies, &DummyData{
		Timestamp: time.Now(),
		ID:        "3",
		Message:   "This is the third dummy",
	})

	return dummyDummies
}
