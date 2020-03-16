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

package core

import (
	"github.com/daisy-ycguo/kn-source-kafka/pkg/commands"
	"github.com/daisy-ycguo/kn-source-kafka/pkg/commands/kafka"

	"github.com/spf13/cobra"
)

// NewKnSourceCommand creates the rootCmd which is the base command when called without any subcommands
func NewKnSourceCommand(params ...commands.KnSourceParams) *cobra.Command {
	var p *commands.KnSourceParams
	p = &commands.KnSourceParams{}

	rootCmd := &cobra.Command{
		Use:   "source",
		Short: "Knative Source plugin",
		Long:  `Manage your Knative eventing sources`,
	}
	if p.Output != nil {
		rootCmd.SetOutput(p.Output)
	}

	rootCmd.AddCommand(kafka.NewKafkaCommand(p))

	return rootCmd
}
