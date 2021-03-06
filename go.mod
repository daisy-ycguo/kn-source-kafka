module github.com/daisy-ycguo/kn-source-kafka

require (
	github.com/spf13/cobra v0.0.5
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	knative.dev/client v0.13.0
	knative.dev/eventing v0.13.3 // indirect
	knative.dev/eventing-contrib v0.13.1
	knative.dev/pkg v0.0.0-20200314000429-7cbd0bcc1a30
	knative.dev/serving v0.13.0
	knative.dev/test-infra v0.0.0-20200229011351-4dac123b9a3d
)

// Temporary pinning certain libraries. Please check periodically, whether these are still needed
// ----------------------------------------------------------------------------------------------

// Fix for `[` in help messages and shell completion code
// See https://github.com/spf13/cobra/pull/899
replace github.com/spf13/cobra => github.com/chmouel/cobra v0.0.0-20191021105835-a78788917390

go 1.13
