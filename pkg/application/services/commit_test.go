package services

import (
	"errors"
	"repo-metrics/pkg/application/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCommitService_ActivityScore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tcs := []struct {
		description    string
		mockCommit     func() *MockCommit
		expectedScores []models.ScoreRepository
		expectedErr    string
	}{
		{
			description: "when commits are successfully read and scores are calculated",
			mockCommit: func() *MockCommit {
				m := NewMockCommit(ctrl)

				m.EXPECT().ReadAll().Return([]models.Commit{
					{Repo: "repo1", Add: 10, Deletions: 5, Files: 2},
					{Repo: "repo2", Add: 20, Deletions: 10, Files: 3},
					{Repo: "repo1", Add: 5, Deletions: 5, Files: 1},
				}, nil)

				return m
			},
			expectedScores: []models.ScoreRepository{
				{Repository: "repo2", Score: 16.6},
				{Repository: "repo1", Score: 15.099999},
			},
			expectedErr: "",
		},
		{
			description: "when ReadAll fails",
			mockCommit: func() *MockCommit {

				m := NewMockCommit(ctrl)

				m.EXPECT().ReadAll().Return([]models.Commit{}, errors.New("failed to read commits"))

				return m
			},
			expectedScores: []models.ScoreRepository{},
			expectedErr:    "failed to read commits",
		},
		{
			description: "when no commits are present",
			mockCommit: func() *MockCommit {
				m := NewMockCommit(ctrl)

				m.EXPECT().ReadAll().Return([]models.Commit{}, nil)

				return m
			},
			expectedScores: nil,
			expectedErr:    "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			mockCommit := tc.mockCommit()
			service := NewCommitService(mockCommit)

			scores, err := service.ActivityScore()

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				assert.Nil(t, scores)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedScores, scores)
			}
		})
	}
}
