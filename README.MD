# ESEP Grade Calculator

This library is a simple utility to calculate your grades based on averaging assignments, exams, and essays. The distribution for these is 50%, 35%, and 15% respectively. The average of the above is then converted into the final letter grade for the course.

Runtime requirements:
Golang 1.24+
Must be installed within the path in order to work properly

Run unit tests:
```
All tests
cd grade_calculator/
go test ./...

Verbose Output
go test -v ./...

Code Coverage 
go test -cover ./...

Coverage Report
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
go tool cover -html=coverage.txt
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```