package services

import (
	"bytes"
	"errors"
	"repo-metrics/pkg/application/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriterService_WriteScores(t *testing.T) {
	t.Run("writes scores successfully", func(t *testing.T) {
		var buffer bytes.Buffer
		service := NewWriterService(&buffer)

		scores := []models.ScoreRepository{
			{Repository: "repo1", Score: 10.5},
			{Repository: "repo2", Score: 20.3},
		}

		err := service.WriteScores(scores)
		assert.NoError(t, err)

		expectedOutput := "Repository: repo1, Score: 10.50\nRepository: repo2, Score: 20.30\n"
		assert.Equal(t, expectedOutput, buffer.String())
	})

	t.Run("returns error when writer fails", func(t *testing.T) {
		errorWriter := &failingWriter{}
		service := NewWriterService(errorWriter)

		scores := []models.ScoreRepository{
			{Repository: "repo1", Score: 10.5},
		}

		err := service.WriteScores(scores)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to write scores")
	})
}

// Simulates a writer that always fails
type failingWriter struct{}

func (f *failingWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("simulated writer failure")
}
