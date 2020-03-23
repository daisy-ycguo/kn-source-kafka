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

package types

import (
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
	flag "github.com/spf13/pflag"
	"knative.dev/client/pkg/kn/commands/flags"
)

type KafkaSourceParams struct {
	KnSourceParams   *sourcetypes.KnSourceParams
	SinkFlag         flags.SinkFlags
	BootstrapServers string
	Topics           string
	ConsumerGroup    string
}

func (p *KafkaSourceParams) AddFlags(flagset *flag.FlagSet) {
	p.SinkFlag.Add(flagset)
	flagset.StringVar(&p.BootstrapServers, "servers", "", "Kafka bootstrap servers that the consumer will connect to, consist of a hostname plus a port pair, e.g. my-kafka-bootstrap.kafka:9092")
	flagset.StringVar(&p.Topics, "topics", "", "Topics to consume messages from")
	flagset.StringVar(&p.ConsumerGroup, "consumergroup", "", "the consumer group ID")
}
