module github.com/gost-dom/browser

go 1.23.4

toolchain go1.23.6

require (
	github.com/dop251/goja v0.0.0-20241024094426-79f3a7efcdbd
	github.com/ericchiang/css v1.4.0
	github.com/onsi/ginkgo/v2 v2.22.2
	github.com/onsi/gomega v1.36.2
	github.com/tommie/v8go v0.28.0
	golang.org/x/net v0.34.0
)

require (
	github.com/dlclark/regexp2 v1.11.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20241210010833-40e02aabc2ad // indirect
	github.com/tommie/v8go/deps/android_amd64 v0.0.0-20250116090852-b2a925bc7092 // indirect
	github.com/tommie/v8go/deps/android_arm64 v0.0.0-20250116090852-b2a925bc7092 // indirect
	github.com/tommie/v8go/deps/darwin_amd64 v0.0.0-20250116090852-b2a925bc7092 // indirect
	github.com/tommie/v8go/deps/darwin_arm64 v0.0.0-20250116090852-b2a925bc7092 // indirect
	github.com/tommie/v8go/deps/linux_amd64 v0.0.0-20250116090852-b2a925bc7092 // indirect
	github.com/tommie/v8go/deps/linux_arm64 v0.0.0-20250116090852-b2a925bc7092 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/tommie/v8go => github.com/stroiman/v8go v0.0.0-20250205095925-1fbfd47bcf87

replace github.com/ericchiang/css => github.com/gost-dom/css v0.0.0-20250131142618-e71a4110d83f
