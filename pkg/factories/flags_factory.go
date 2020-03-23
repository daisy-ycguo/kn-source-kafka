// Copyright © 2018 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package factories

import (
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
	"knative.dev/client/pkg/kn/commands"

	"github.com/daisy-ycguo/kn-source-kafka/pkg/types"

	"github.com/spf13/pflag"
)

type kafkaSourceFlagsFactory struct {
	kafkaSourceParamsFactory types.KafkaSourceParamsFactory
}

func NewKafkaSourceFlagsFactory(kafkaParamsFactory types.KafkaSourceParamsFactory) sourcetypes.FlagsFactory {
	return &kafkaSourceFlagsFactory{
		kafkaSourceParamsFactory: kafkaParamsFactory,
	}
}

func (f *kafkaSourceFlagsFactory) KafkaSourceParams() *types.KafkaSourceParams {
	return f.kafkaSourceParamsFactory.KafkaSourceParams()
}

func (f *kafkaSourceFlagsFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.kafkaSourceParamsFactory.KnSourceParams()
}

func (f *kafkaSourceFlagsFactory) CreateFlags() *pflag.FlagSet {
	flagSet := pflag.NewFlagSet("create", pflag.ExitOnError)
	commands.AddNamespaceFlags(flagSet, false)
	f.KafkaSourceParams().AddFlags(flagSet)
	return flagSet
}

func (f *kafkaSourceFlagsFactory) DeleteFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("delete", pflag.ExitOnError)
}

func (f *kafkaSourceFlagsFactory) UpdateFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("create", pflag.ExitOnError)
}

func (f *kafkaSourceFlagsFactory) DescribeFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("create", pflag.ExitOnError)
}
