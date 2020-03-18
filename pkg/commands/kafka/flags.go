// Copyright © 2019 The Knative Authors
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

package kafka

import (
	"github.com/spf13/cobra"
)

type kafkaUpdateFlags struct {
	bootstrapServers string
	topics           string
	consumerGroup    string
}

func (c *kafkaUpdateFlags) addKafkaFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&c.bootstrapServers, "servers", "", "Kafka bootstrap servers that the consumer will connect to, consist of a hostname plus a port pair, e.g. my-kafka-bootstrap.kafka:9092")
	cmd.Flags().StringVar(&c.topics, "topics", "", "Topics to consume messages from")
	cmd.Flags().StringVar(&c.consumerGroup, "consumergroup", "", "the consumer group ID")
}
