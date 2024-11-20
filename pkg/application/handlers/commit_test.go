package handlers

import (
	"errors"
	"repo-metrics/pkg/application/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCommitHandler_ActivityScoreCLI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tcs := []struct {
		description string
		mockCommit  func() *MockCommit
		mockWriter  func() *MockWriter
		expectedErr string
	}{
		{
			description: "when ActivityScore and WriteScores succeed",
			mockCommit: func() *MockCommit {
				m := NewMockCommit(ctrl)
				m.EXPECT().ActivityScore().Return([]models.ScoreRepository{
					{Repository: "repo1", Score: 10.5},
					{Repository: "repo2", Score: 20.3},
				}, nil)
				return m
			},
			mockWriter: func() *MockWriter {
				m := NewMockWriter(ctrl)
				m.EXPECT().WriteScores([]models.ScoreRepository{
					{Repository: "repo1", Score: 10.5},
					{Repository: "repo2", Score: 20.3},
				}).Return(nil)
				return m
			},
			expectedErr: "",
		},
		{
			description: "when ActivityScore fails",
			mockCommit: func() *MockCommit {
				m := NewMockCommit(ctrl)
				m.EXPECT().ActivityScore().Return(nil, errors.New("failed to calculate scores"))
				return m
			},
			mockWriter: func() *MockWriter {
				return NewMockWriter(ctrl)
			},
			expectedErr: "failed to calculate activity scores: failed to calculate scores",
		},
		{
			description: "when WriteScores fails",
			mockCommit: func() *MockCommit {
				m := NewMockCommit(ctrl)
				m.EXPECT().ActivityScore().Return([]models.ScoreRepository{
					{Repository: "repo1", Score: 10.5},
				}, nil)
				return m
			},
			mockWriter: func() *MockWriter {
				m := NewMockWriter(ctrl)
				m.EXPECT().WriteScores([]models.ScoreRepository{
					{Repository: "repo1", Score: 10.5},
				}).Return(errors.New("failed to write scores"))
				return m
			},
			expectedErr: "failed to write scores",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			mockCommit := tc.mockCommit()
			mockWriter := tc.mockWriter()

			handler := NewCommitHandler(mockCommit, mockWriter)

			err := handler.ActivityScoreCLI()
			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
