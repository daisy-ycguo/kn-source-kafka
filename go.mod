module github.com/daisy-ycguo/kn-source-kafka

require (
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/spf13/cobra v0.0.5
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	google.golang.org/appengine v1.6.2 // indirect
	k8s.io/client-go v0.17.0
	knative.dev/test-infra v0.0.0-20200229011351-4dac123b9a3d
)

// Temporary pinning certain libraries. Please check periodically, whether these are still needed
// ----------------------------------------------------------------------------------------------

// Fix for `[` in help messages and shell completion code
// See https://github.com/spf13/cobra/pull/899
replace github.com/spf13/cobra => github.com/chmouel/cobra v0.0.0-20191021105835-a78788917390

go 1.13
