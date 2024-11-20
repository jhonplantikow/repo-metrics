package file

import (
	"io"
	"os"
	"repo-metrics/pkg/application/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileReader_ReadAll(t *testing.T) {
	input := `timestamp,username,repository,files,additions,deletions
1610969774,user123,repo123,3,10,5
1610969774,user123,,3,10,5
1610969774,user123,repo123,0,10,5
1610969774,user123,repo123,2,invalid,5
1610969774,user123,repo123,2,10,invalid`

	inputFile, err := os.CreateTemp("", "input.csv")
	assert.NoError(t, err)
	defer os.Remove(inputFile.Name())

	outputFile, err := os.CreateTemp("", "output.csv")
	assert.NoError(t, err)
	defer os.Remove(outputFile.Name())

	_, err = inputFile.WriteString(input)
	assert.NoError(t, err)

	_, err = inputFile.Seek(0, 0)
	assert.NoError(t, err)

	reader := NewFileReader(inputFile, outputFile)

	t.Run("successfully reads valid rows and skips invalid ones", func(t *testing.T) {
		commits, err := reader.ReadAll()

		assert.NoError(t, err)
		assert.Len(t, commits, 1)

		expectedCommit := models.Commit{
			TS:        time.Unix(1610969774, 0),
			User:      "user123",
			Repo:      "repo123",
			Files:     3,
			Add:       10,
			Deletions: 5,
		}
		assert.Equal(t, expectedCommit, commits[0])

		_, err = outputFile.Seek(0, 0)
		assert.NoError(t, err)

		outputData, err := io.ReadAll(outputFile)
		assert.NoError(t, err)

		expectedOutput := `1610969774,user123,,3,10,5,error:invalid data: repository must not be empty
1610969774,user123,repo123,0,10,5,error:invalid data: files must be greater than 0
1610969774,user123,repo123,2,invalid,5,error:failed to sanitize additions: strconv.Atoi: parsing "invalid": invalid syntax
1610969774,user123,repo123,2,10,invalid,error:failed to sanitize deletions: strconv.Atoi: parsing "invalid": invalid syntax
`
		assert.Equal(t, expectedOutput, string(outputData))
	})
}
