package coverageanalysis

import (
	"encoding/json"
	"errors"
)

func AnalyzeJson(jsonString []byte) (report *CoverageReport, err error) {
	var jsonObject map[string]interface{}
	err = json.Unmarshal(jsonString, &jsonObject)
	if err != nil {
		return &CoverageReport{}, errors.New("malformed JSON")
	}
	defer func() { //catch or finally
		if innerErr := recover(); innerErr != nil { //catch
			err = errors.New("bad JSON")
		}
	}()
	report = BuildCoverageReport(jsonObject)
	AnalyzeReport(report)
	return report, err
}

func AnalyzeReport(report *CoverageReport) {
	CalculateAndSetAllFileCoverages(report)
}
