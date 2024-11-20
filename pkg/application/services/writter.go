package services

import (
	"fmt"
	"io"
	"repo-metrics/pkg/application/models"
)

type WriterService struct {
	Output io.Writer
}

func NewWriterService(output io.Writer) WriterService {
	return WriterService{
		Output: output,
	}
}

func (w WriterService) WriteScores(scores []models.ScoreRepository) error {
	for _, score := range scores {
		_, err := fmt.Fprintf(w.Output, "Repository: %s, Score: %.2f\n", score.Repository, score.Score)
		if err != nil {
			return fmt.Errorf("failed to write scores: %v", err)
		}
	}
	return nil
}
