package models

import "time"

type Commit struct {
	TS        time.Time `csv:"timestamp"`
	User      string    `csv:"user"`
	Repo      string    `csv:"repository"`
	Files     int       `csv:"files"`
	Add       int       `csv:"additions"`
	Deletions int       `csv:"deletions"`
}

type ScoreRepository struct {
	Score      float32
	Repository string
}
