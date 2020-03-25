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
	kafkaSourceParams *types.KafkaSourceParams
}

func NewKafkaSourceFlagsFactory(kafkaParams *types.KafkaSourceParams) sourcetypes.FlagsFactory {
	return &kafkaSourceFlagsFactory{
		kafkaSourceParams: kafkaParams,
	}
}

func (f *kafkaSourceFlagsFactory) KafkaSourceParams() *types.KafkaSourceParams {
	return f.kafkaSourceParams
}

func (f *kafkaSourceFlagsFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.kafkaSourceParams.KnSourceParams
}

func (f *kafkaSourceFlagsFactory) CreateFlags() *pflag.FlagSet {
	flagSet := pflag.NewFlagSet("create", pflag.ExitOnError)
	flagSet.StringVar(&f.kafkaSourceParams.BootstrapServers, "servers", "", "Kafka bootstrap servers that the consumer will connect to, consist of a hostname plus a port pair, e.g. my-kafka-bootstrap.kafka:9092")
	flagSet.StringVar(&f.kafkaSourceParams.Topics, "topics", "", "Topics to consume messages from")
	flagSet.StringVar(&f.kafkaSourceParams.ConsumerGroup, "consumergroup", "", "the consumer group ID")
	commands.AddNamespaceFlags(flagSet, false)
	f.kafkaSourceParams.KnSourceParams.SinkFlag.Add(flagSet)
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
