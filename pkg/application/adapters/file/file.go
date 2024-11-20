package file

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"repo-metrics/pkg/application/models"
	"strconv"
	"strings"
	"time"
)

const (
	TS = iota
	User
	Repo
	Files
	Add
	Deletions

	RecNumber = 6
)

type FileReader struct {
	in  *os.File
	out *os.File
}

func NewFileReader(in, out *os.File) *FileReader {
	return &FileReader{
		in:  in,
		out: out,
	}
}

func (f *FileReader) ReadAll() ([]models.Commit, error) {
	if _, err := f.out.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek to the beginning of the file: %w", err)
	}

	writer := bufio.NewWriter(f.out)

	csvReader := csv.NewReader(f.in)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var commits []models.Commit
	for i, record := range records {
		// Skip header
		if i == 0 || len(record) != RecNumber {
			continue
		}

		commit, err := sanitizeRow(record)
		if err != nil {
			writer.WriteString(fmt.Sprintf("%s,error:%s\n", strings.Join(record, ","), err.Error()))
			continue
		}

		commits = append(commits, commit)
	}

	if err := writer.Flush(); err != nil {
		return nil, fmt.Errorf("failed to flush writer: %w", err)
	}

	return commits, nil
}

func sanitizeRow(row []string) (models.Commit, error) {
	var empty models.Commit

	utimestamp, err := strconv.ParseInt(row[TS], 10, 64)
	if err != nil {
		return empty, fmt.Errorf("failed to sanitize timestamp: %w", err)
	}

	if row[Repo] == "" {
		return empty, errors.New("invalid data: repository must not be empty")
	}

	f, err := strconv.Atoi(row[Files])
	if err != nil {
		return empty, fmt.Errorf("failed to sanitize files: %w", err)
	}

	if f == 0 {
		return empty, errors.New("invalid data: files must be greater than 0")
	}

	add, err := strconv.Atoi(row[Add])
	if err != nil {
		return empty, fmt.Errorf("failed to sanitize additions: %w", err)
	}

	del, err := strconv.Atoi(row[Deletions])
	if err != nil {
		return empty, fmt.Errorf("failed to sanitize deletions: %w", err)
	}

	return models.Commit{
		TS:        time.Unix(utimestamp, 0),
		User:      row[User],
		Repo:      row[Repo],
		Files:     f,
		Add:       add,
		Deletions: del,
	}, nil
}
