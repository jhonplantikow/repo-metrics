package handlers

import (
	"fmt"
	"repo-metrics/pkg/application/models"
)

//go:generate mockgen -source=commit.go -destination=mock_commit_generated.go -package=handlers

type Commit interface {
	ActivityScore() ([]models.ScoreRepository, error)
}

type Writer interface {
	WriteScores(scores []models.ScoreRepository) error
}

type CommitHandler struct {
	Commit Commit
	Writer Writer
}

func NewCommitHandler(c Commit, w Writer) CommitHandler {
	return CommitHandler{
		Commit: c,
		Writer: w,
	}
}

func (h CommitHandler) ActivityScoreCLI() error {
	scores, err := h.Commit.ActivityScore()
	if err != nil {
		return fmt.Errorf("failed to calculate activity scores: %v", err)
	}

	if err := h.Writer.WriteScores(scores); err != nil {
		return err
	}

	fmt.Println("Activity Scores:")
	for _, score := range scores {
		fmt.Printf("Repository: %s, Score: %.2f\n", score.Repository, score.Score)
	}

	return nil
}
