package services

import (
	"repo-metrics/pkg/application/models"
	"sort"
)

//go:generate mockgen -source=commit.go -destination=mock_commit_generated.go -package=services

const (
	CommitScore = 1
)

type Commit interface {
	ReadAll() ([]models.Commit, error)
}

type CommitService struct {
	Commit Commit
}

func NewCommitService(c Commit) CommitService {
	return CommitService{
		Commit: c,
	}
}

func (s CommitService) ActivityScore() ([]models.ScoreRepository, error) {
	var (
		shortedCommits []models.ScoreRepository
		repoScores     = make(map[string]float32)
	)

	commits, err := s.Commit.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, c := range commits {
		repoScores[c.Repo] = repoScores[c.Repo] + (CommitScore + 0.5*float32(c.Add) + 0.5*float32(c.Deletions) + 0.2*float32(c.Files))
	}

	for r, sc := range repoScores {
		shortedCommits = append(shortedCommits, models.ScoreRepository{Repository: r, Score: sc})
	}

	sort.Slice(shortedCommits, func(i, j int) bool {
		return shortedCommits[i].Score > shortedCommits[j].Score
	})

	return shortedCommits, nil
}
