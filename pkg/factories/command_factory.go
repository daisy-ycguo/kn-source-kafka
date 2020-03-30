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
	"github.com/daisy-ycguo/kn-source-kafka/pkg/types"

	sourcefactories "github.com/maximilien/kn-source-pkg/pkg/factories"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
)

type kafkaSourceCommandFactory struct {
	kafkaSourceFactory    types.KafkaSourceFactory
	defaultCommandFactory sourcetypes.CommandFactory
}

func NewKafkaSourceCommandFactory(kafkaFactory types.KafkaSourceFactory) types.KafkaSourceCommandFactory {
	return &kafkaSourceCommandFactory{
		kafkaSourceFactory:    kafkaFactory,
		defaultCommandFactory: sourcefactories.NewDefaultCommandFactory(kafkaFactory),
	}
}

func (f *kafkaSourceCommandFactory) KnSourceFactory() sourcetypes.KnSourceFactory {
	return f.kafkaSourceFactory
}

func (f *kafkaSourceCommandFactory) KafkaSourceFactory() types.KafkaSourceFactory {
	return f.kafkaSourceFactory
}

func (f *kafkaSourceCommandFactory) KafkaSourceParams() *types.KafkaSourceParams {
	return f.kafkaSourceFactory.KafkaSourceParams()
}

func (f *kafkaSourceCommandFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.kafkaSourceFactory.KnSourceParams()
}

func (f *kafkaSourceCommandFactory) SourceCommand() *cobra.Command {
	sourceCmd := f.defaultCommandFactory.SourceCommand()
	sourceCmd.Use = "kafka"
	sourceCmd.Short = "Knative eventing Kafka source plugin"
	sourceCmd.Long = "Manage your Knative Kafka eventing sources"
	return sourceCmd
}

func (f *kafkaSourceCommandFactory) CreateCommand() *cobra.Command {
	createCmd := f.defaultCommandFactory.CreateCommand()
	createCmd.Short = "create NAME"
	createCmd.Example = `#Creates a new Kafka source with NAME
kn source kafka create kafka-name`
	return createCmd
}

func (f *kafkaSourceCommandFactory) DeleteCommand() *cobra.Command {
	deleteCmd := f.defaultCommandFactory.DeleteCommand()
	deleteCmd.Short = "delete NAME"
	deleteCmd.Long = "delete a Kafka source"
	deleteCmd.Example = `#Deletes a Kafka source with NAME
kn source kafka delete kafka-name`
	return deleteCmd
}

func (f *kafkaSourceCommandFactory) UpdateCommand() *cobra.Command {
	updateCmd := f.defaultCommandFactory.UpdateCommand()
	updateCmd.Short = "update NAME"
	updateCmd.Long = "update a Kafka source"
	updateCmd.Example = `#Updates a Kafka source with NAME
kn source kafka update kafka-name`
	return updateCmd
}

func (f *kafkaSourceCommandFactory) DescribeCommand() *cobra.Command {
	describeCmd := f.defaultCommandFactory.DescribeCommand()
	describeCmd.Short = "describe NAME"
	describeCmd.Long = "update a Kafka source"
	describeCmd.Example = `#Describes a Kafka source with NAME
kn source kafka describe kafka-name`
	return describeCmd
}
