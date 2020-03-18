// Copyright © 2020 The Knative Authors
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
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/daisy-ycguo/kn-source-kafka/pkg/kafka/v1alpha1"
	"knative.dev/client/pkg/kn/commands"
	"knative.dev/client/pkg/kn/commands/flags"
)

// NewKafkaCreateCommand for creating source
func NewKafkaCreateCommand(p *commands.KnParams) *cobra.Command {
	var updateFlags kafkaUpdateFlags
	var sinkFlags flags.SinkFlags

	cmd := &cobra.Command{
		Use:   "create NAME --servers BootstrapServers --topics TOPICS --consumergroup CONSUMERGROUP --sink SINK",
		Short: "Create a Kafka source.",
		Example: `
  # Create a Kafka 'kafkasrc' which consumes Kafka events and sends message to service 'mysvc' as a cloudevent
  kn source kafka create kafkasrc --servers my-kafka-bootstrap.kafka:9092 --topics demo-topic --consumergroup knative-group --sink svc:mysvc`,

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) != 1 {
				return errors.New("requires the name of the source to create as single argument")
			}
			name := args[0]

			// get client
			kafkaClient, err := newKafkaSourceClient(p, cmd)
			if err != nil {
				return err
			}

			namespace := kafkaClient.Namespace()

			dynamicClient, err := p.NewDynamicClient(namespace)
			if err != nil {
				return err
			}
			objectRef, err := sinkFlags.ResolveSink(dynamicClient, namespace)
			if err != nil {
				return fmt.Errorf(
					"cannot create ApiServerSource '%s' in namespace '%s' "+
						"because: %s", name, namespace, err)
			}

			b := v1alpha1.NewKafkaSourceBuilder(name).
				BootstrapServers(updateFlags.bootstrapServers).
				Topics(updateFlags.topics).
				ConsumerGroup(updateFlags.consumerGroup).
				Sink(toDuckV1Beta1(objectRef))

			err = kafkaClient.CreateKafkaSource(b.Build())

			if err != nil {
				return fmt.Errorf(
					"cannot create KafkaSource '%s' in namespace '%s' "+
						"because: %s", name, namespace, err)
			}

			if err == nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Kafka source '%s' created in namespace '%s'.\n", args[0], namespace)
			}

			return err
		},
	}
	commands.AddNamespaceFlags(cmd.Flags(), false)
	updateFlags.addKafkaFlags(cmd)
	sinkFlags.Add(cmd)
	cmd.MarkFlagRequired("servers")
	cmd.MarkFlagRequired("topics")
	cmd.MarkFlagRequired("consumergroup")
	return cmd
}
