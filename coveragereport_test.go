package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestLineCoverageWithInteger(t *testing.T) {
	coverage := BuildLineCoverage(float64(1))
	assert.Equal(t, float64(1), coverage.Hits)
	assert.False(t, coverage.Ignore)
}

func TestLineCoverageWithNil(t *testing.T) {
	coverage := BuildLineCoverage(nil)
	assert.Equal(t, float64(0), coverage.Hits)
	assert.True(t, coverage.Ignore)
}

func TestLineCoverageWithDefault(t *testing.T) {
	coverage := BuildLineCoverage("")
	assert.Equal(t, float64(0), coverage.Hits)
	assert.False(t, coverage.Ignore)
}

func TestLineCoverageWithArray(t *testing.T) {
	json_string := `[null, 1]`
	var jsonArray []interface{}
	err := json.Unmarshal([]byte(json_string), &jsonArray)
	if err != nil {
		t.Error(err)
	}
	report := BuildAllLineCoverages(jsonArray)
	expected_report := []LineCoverage{{Hits: 0, Ignore: true}, {Hits: 1, Ignore: false}}
	if !reflect.DeepEqual(report, expected_report) {
		t.Errorf("")
	}
}

func TestCoverageReportWithExpectedJson(t *testing.T) {
	json_string := `{
		"source_files": [{
			"name": "test.rb",
			"source": "puts 'test'",
			"source_digest": "*md5*",
			"coverage": [null, 1, null]
		}]
	}`
	var jsonObject map[string]interface{}
	err := json.Unmarshal([]byte(json_string), &jsonObject)
	if err != nil {
		t.Error(err)
	}
	report := BuildCoverageReport(jsonObject)
	expected_report := &CoverageReport{
		FileReports: []FileReport{{
			Name:       "test.rb",
			FileDigest: "*md5*",
			FileSource: "puts 'test'",
			Coverage:   []LineCoverage{{Hits: 0, Ignore: true}, {Hits: 1, Ignore: false}, {Hits: 0, Ignore: true}},
		}},
	}
	if !reflect.DeepEqual(report, expected_report) {
		t.Errorf("")
	}
}

func TestCoverageReportWithWrongJson(t *testing.T) {
	json_string := `{
		"something_else": [{
			"test": "some string"
		}]
	}`
	var jsonObject map[string]interface{}
	err := json.Unmarshal([]byte(json_string), &jsonObject)
	if err != nil {
		t.Error(err)
	}
	assert.Panics(t, func() {
		BuildCoverageReport(jsonObject)
	}, "Building coverage report with malformed json should panic")
}
