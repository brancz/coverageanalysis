package coverageanalysis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineCoverageAggregation(t *testing.T) {
	coverages := []LineCoverage{{Hits: 1, Ignore: false}, {Hits: 0, Ignore: false}, {Hits: 0, Ignore: true}}
	relevantLines, linesCovered := AggregateLineHits(coverages)
	assert.Equal(t, 2, relevantLines)
	assert.Equal(t, 1, linesCovered)
}

func TestAllFileCoverageCalculation(t *testing.T) {
	report := &CoverageReport{FileReports: []FileReport{{LineCoverages: []LineCoverage{{Hits: 1, Ignore: false}, {Hits: 0, Ignore: false}, {Hits: 0, Ignore: false}, {Hits: 0, Ignore: true}}}, {LineCoverages: []LineCoverage{{Hits: 1, Ignore: false}, {Hits: 0, Ignore: false}, {Hits: 0, Ignore: true}}}}}
	CalculateAndSetAllFileCoverages(report)
	assert.Equal(t, 0.3333333333333333, report.FileReports[0].CoveragePercentage)
	assert.Equal(t, 0.5, report.FileReports[1].CoveragePercentage)
	assert.Equal(t, 0.4, report.CoveragePercentage)
}

func TestCoverageCalculationZeroDivision(t *testing.T) {
	report := &CoverageReport{FileReports: []FileReport{{LineCoverages: []LineCoverage{}}}}
	CalculateAndSetAllFileCoverages(report)
	assert.Equal(t, 0, report.FileReports[0].CoveragePercentage)
	assert.Equal(t, 0, report.CoveragePercentage)
}
