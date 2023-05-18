test:
	go test -v ./... --cover
count-line:
	find . -name "*.go" ! -name "*_test.go" | xargs wc -l
count-test-line:
	find . -name "*_test.go" | xargs wc -l