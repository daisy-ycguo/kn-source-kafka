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
	"errors"
	"fmt"

	"github.com/daisy-ycguo/kn-source-kafka/pkg/client"
	"github.com/daisy-ycguo/kn-source-kafka/pkg/types"
	"knative.dev/client/pkg/kn/commands/flags"

	basicfactories "github.com/maximilien/kn-source-pkg/pkg/factories"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
)

type kafkaSourceRunEFactory struct {
	basicfactories.DefautRunEFactory
	kafkaSourceParams  *types.KafkaSourceParams
	kafkaSourceClient  types.KafkaSourceClient
	kafkaClientFactory types.KafkaSourceFactory
}

func NewKafkaSourceRunEFactory(kafkaSrcParams *types.KafkaSourceParams,
	kafkaSourceClientFactory types.KafkaSourceFactory) types.KafkaSourceRunEFactory {
	return &kafkaSourceRunEFactory{
		kafkaSourceParams:  kafkaSrcParams,
		kafkaClientFactory: kafkaSourceClientFactory,
		kafkaSourceClient:  nil,
	}
}

func (f *kafkaSourceRunEFactory) KafkaSourceParams() *types.KafkaSourceParams {
	return f.kafkaSourceParams
}

func (f *kafkaSourceRunEFactory) KafkaSourceFactory() types.KafkaSourceFactory {
	return f.kafkaClientFactory
}

func (f *kafkaSourceRunEFactory) KafkaSourceClient(cmd *cobra.Command) error {
	if f.kafkaSourceClient == nil {
		p := f.KnSourceParams().KnParams
		namespace, err := p.GetNamespace(cmd)
		if err != nil {
			return err
		}
		f.kafkaSourceClient = f.kafkaClientFactory.CreateKafkaSourceClient(namespace)
	}
	return nil
}

func (f *kafkaSourceRunEFactory) CreateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		f.KafkaSourceClient(cmd)
		fmt.Printf("%s RunE function called for Kafka source: args: %#v, client: %#v\n", cmd.Name(), args, f.kafkaSourceClient)

		if len(args) != 1 {
			return errors.New("requires the name of the source to create as single argument")
		}
		name := args[0]

		dynamicClient, err := f.KnSourceParams().KnParams.NewDynamicClient(f.kafkaSourceClient.Namespace())
		if err != nil {
			return err
		}
		objectRef, err := f.KnSourceParams().SinkFlag.ResolveSink(dynamicClient, f.kafkaSourceClient.Namespace())
		if err != nil {
			return fmt.Errorf(
				"cannot create Kafka '%s' in namespace '%s' "+
					"because: %s", name, f.kafkaSourceClient.Namespace(), err)
		}

		b := client.NewKafkaSourceBuilder(name).
			BootstrapServers(f.kafkaSourceParams.BootstrapServers).
			Topics(f.kafkaSourceParams.Topics).
			ConsumerGroup(f.kafkaSourceParams.ConsumerGroup).
			Sink(flags.SinkToDuckV1Beta1(objectRef))

		err = f.kafkaSourceClient.CreateKafkaSource(b.Build())

		if err != nil {
			return fmt.Errorf(
				"cannot create KafkaSource '%s' in namespace '%s' "+
					"because: %s", name, f.kafkaSourceClient.Namespace(), err)
		}

		if err == nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Kafka source '%s' created in namespace '%s'.\n", args[0], f.kafkaSourceClient.Namespace())
		}

		return err
	}
}

func (f *kafkaSourceRunEFactory) DeleteRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s RunE function called for Kafka source: args: %#v, client: %#v\n", cmd.Name(), args, f.kafkaSourceClient)
		return nil
	}
}

func (f *kafkaSourceRunEFactory) UpdateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s RunE function called for Kafka source: args: %#v, client: %#v\n", cmd.Name(), args, f.kafkaSourceClient)
		return nil
	}
}

func (f *kafkaSourceRunEFactory) DescribeRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s RunE function called for Kafka source: args: %#v, client: %#v\n", cmd.Name(), args, f.kafkaSourceClient)
		return nil
	}
}
