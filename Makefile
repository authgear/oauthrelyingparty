.PHONY: vendor
vendor:
	go mod download
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: go-mod-outdated
go-mod-outdated:
	# https://stackoverflow.com/questions/55866604/whats-the-go-mod-equivalent-of-npm-outdated
	go list -u -m -f '{{if .Update}}{{if not .Indirect}}{{.}}{{end}}{{end}}' all

.PHONY: test
test:
	go test ./pkg/...

.PHONY: vet
vet:
	go vet ./pkg/...

.PHONY:	fmt
fmt:
	# Ignore generated files, such as wire_gen.go and *_mock_test.go
	find ./pkg -name '*.go' -not -name 'wire_gen.go' -not -name '*_mock_test.go' | sort | xargs goimports -w -format-only -local github.com/authgear/oauthrelyingparty

.PHONY: govulncheck
govulncheck:
	govulncheck -show traces,version,verbose ./...

.PHONY: check-tidy
check-tidy:
	$(MAKE) fmt
	go mod tidy
	git status --porcelain | grep '.*'; test $$? -eq 1
