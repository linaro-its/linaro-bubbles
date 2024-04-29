go test $(go list ./... | grep -v /examples/) -coverprofile coverage.out
go tool cover -func coverage.out
