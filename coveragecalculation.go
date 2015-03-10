package main

import (
	"math"
)

func AggregateLineHits(lineCoverages []LineCoverage) (int, int) {
	numberOfLinesCovered := 0
	numberOfRelevantLines := 0
	for i := range lineCoverages {
		currentLine := lineCoverages[i]
		if !currentLine.Ignore {
			numberOfRelevantLines++
			if currentLine.Hits > 0 {
				numberOfLinesCovered++
			}
		}
	}
	return numberOfRelevantLines, numberOfLinesCovered
}

func CalculateCoveragePercentage(linesCovered int, relevantLines int) float64 {
	percentage := (float64(linesCovered) / float64(relevantLines))
	if math.IsNaN(percentage) {
		return float64(0)
	}
	return percentage
}

func CalculateAndSetAllFileCoverages(report *CoverageReport) {
	totalNumberOfRelevantLines := 0
	totalNumberOfLinesCovered := 0
	fileReports := &report.FileReports
	for i := range *fileReports {
		fileReport := &(*fileReports)[i]
		numberOfRelevantLines, numberOfLinesCovered := AggregateLineHits((*fileReport).LineCoverages)
		fileReport.CoveragePercentage = CalculateCoveragePercentage(numberOfLinesCovered, numberOfRelevantLines)
		totalNumberOfLinesCovered += numberOfLinesCovered
		totalNumberOfRelevantLines += numberOfRelevantLines
	}
	report.CoveragePercentage = CalculateCoveragePercentage(totalNumberOfLinesCovered, totalNumberOfRelevantLines)
}
