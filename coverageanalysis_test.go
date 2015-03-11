package coverageanalysis

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonAnalysis(t *testing.T) {
	report, err := AnalyzeJson([]byte(`{
		"source_files": [{
			"name": "test.rb",
			"source": "puts 'test'",
			"source_digest": "*md5*",
			"coverage": [null, 1, null]
		}]
	}`))
	if err != nil {
		t.Error(err)
	}
	expectedReport := &CoverageReport{
		CoveragePercentage: float64(1),
		FileReports: []FileReport{
			{
				Name:               "test.rb",
				FileSource:         "puts 'test'",
				FileDigest:         "*md5*",
				CoveragePercentage: float64(1),
				LineCoverages: []LineCoverage{
					{Hits: 0, Ignore: true},
					{Hits: 1, Ignore: false},
					{Hits: 0, Ignore: true}},
			},
		},
	}
	assert.Equal(t, report, expectedReport)
}

func TestJsonAnalysisWithMalformedJson(t *testing.T) {
	_, err := AnalyzeJson([]byte("{"))
	assert.NotNil(t, err)
}

func TestJsonAnalysisWithWrongStructureOfJson(t *testing.T) {
	_, err := AnalyzeJson([]byte("{}"))
	if assert.Error(t, err, "bad JSON") {
		assert.Equal(t, err, errors.New("bad JSON"))
	}
}
