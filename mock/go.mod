module github.com/valkey-io/valkey-go/mock

go 1.22.0

replace github.com/valkey-io/valkey-go => ../

require (
	github.com/valkey-io/valkey-go v1.0.54
	go.uber.org/mock v0.5.0
)

require golang.org/x/sys v0.29.0 // indirect
