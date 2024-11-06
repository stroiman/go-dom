module github.com/stroiman/go-dom

go 1.23.2

require (
	github.com/onsi/ginkgo/v2 v2.21.0
	github.com/onsi/gomega v1.35.1
	github.com/tommie/v8go v0.22.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20241029153458-d1b30febd7db // indirect
	github.com/tommie/v8go/deps/android_amd64 v0.0.0-20241023013435-d8e1c56d9e6a // indirect
	github.com/tommie/v8go/deps/android_arm64 v0.0.0-20241023013435-d8e1c56d9e6a // indirect
	github.com/tommie/v8go/deps/darwin_amd64 v0.0.0-20241023013435-d8e1c56d9e6a // indirect
	github.com/tommie/v8go/deps/darwin_arm64 v0.0.0-20241023013435-d8e1c56d9e6a // indirect
	github.com/tommie/v8go/deps/linux_amd64 v0.0.0-20241023013435-d8e1c56d9e6a // indirect
	github.com/tommie/v8go/deps/linux_arm64 v0.0.0-20241023013435-d8e1c56d9e6a // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Use a WIP version of the code
// replace github.com/tommie/v8go v0.22.0 => /Users/peter/go/src/github/stroiman/v8go

// Use the version from github
// Replace the actual line with this line, and run `go mod tidy`
// replace github.com/tommie/v8go v0.22.0 => github.com/stroiman/v8go external-support
replace github.com/tommie/v8go v0.22.0 => github.com/stroiman/v8go v0.0.0-20241106214940-8fca5578e930
