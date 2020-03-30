module github.com/daisy-ycguo/kn-source-kafka

require (
	github.com/maximilien/kn-source-pkg v0.1.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.17.0
	knative.dev/client v0.13.0
	knative.dev/eventing-contrib v0.13.2
	knative.dev/pkg v0.0.0-20200323231609-0840da9555a3
	knative.dev/test-infra v0.0.0-20200324183109-d81c66eea4e6
)

// Temporary pinning certain libraries. Please check periodically, whether these are still needed
// ----------------------------------------------------------------------------------------------

replace github.com/spf13/cobra => github.com/chmouel/cobra v0.0.0-20191021105835-a78788917390

replace github.com/maximilien/kn-source-pkg => /Users/Daisy/go/src/github.com/maximilien/kn-source-pkg

replace knative.dev/client => /Users/Daisy/go/src/knative.dev/client

go 1.13
