package main

type LineCoverage struct {
	Hits   float64
	Ignore bool
}

type FileReport struct {
	Name       string
	FileDigest string
	FileSource string
	Coverage   []LineCoverage
}

type CoverageReport struct {
	FileReports []FileReport
}

func BuildCoverageReport(jsonObject map[string]interface{}) *CoverageReport {
	report := &CoverageReport{}
	source_files := jsonObject["source_files"].([]interface{})
	for i := range source_files {
		fileReport := BuildFileReport(source_files[i].(map[string]interface{}))
		report.FileReports = append(report.FileReports, *fileReport)
	}
	return report
}

func BuildFileReport(jsonObject map[string]interface{}) *FileReport {
	report := &FileReport{}
	report.Name = jsonObject["name"].(string)
	report.FileDigest = jsonObject["source_digest"].(string)
	report.FileSource = jsonObject["source"].(string)
	report.Coverage = BuildAllLineCoverages(jsonObject["coverage"].([]interface{}))
	return report
}

func BuildAllLineCoverages(coverages []interface{}) []LineCoverage {
	report := []LineCoverage{}
	for i := range coverages {
		c := BuildLineCoverage(coverages[i])
		report = append(report, *c)
	}
	return report
}

func BuildLineCoverage(c interface{}) *LineCoverage {
	switch c.(type) {
	case float64:
		return &LineCoverage{Ignore: false, Hits: c.(float64)}
	case nil:
		return &LineCoverage{Ignore: true, Hits: 0}
	default:
		return &LineCoverage{Ignore: false, Hits: 0}
	}
}
