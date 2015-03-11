# coverageanalysis

Analyzing coverage reports with go. This project resembles the way
[coveralls](https://coveralls.io) might be analysing coverage reports generated
by their client libraries.

## Usage

```go
import "github.com/flower-pot/coverageanalysis"
```

You can either directly analyze a json string .

```go
coverageanalysis.AnalyzeJson(`{
  "source_files": [{
    "name": "test.rb",
    "source": "puts 'test'",
    "source_digest": "*md5*",
    "coverage": [null, 1, null]
  }]
}`)
```

Or a `CoverageReport` struct.

> Beware: when analyzing a struct directly the values of the struct itself
> will be changed

```go
report := &coverageanalysis.CoverageReport{FileReports: []FileReport{{LineCoverages: []LineCoverage{{Hits: 1, Ignore: false}, {Hits: 0, Ignore: false}, {Hits: 0, Ignore: false}, {Hits: 0, Ignore: true}}}}}
coverageanalysis.AnalyzeReport(report)
```

## Contributing

1. Fork it ( https://github.com/flower-pot/coverageanalysis/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
