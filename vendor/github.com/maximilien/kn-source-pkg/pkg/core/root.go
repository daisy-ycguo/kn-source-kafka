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
	"github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewKnSourceCommand(knSourceParams *types.KnSourceParams,
	commandFactory types.CommandFactory,
	flagsFactory types.FlagsFactory,
	runEFactory types.RunEFactory) *cobra.Command {
	params := &types.KnSourceParams{}

	rootCmd := commandFactory.SourceCommand()

	// Disable docs header
	rootCmd.DisableAutoGenTag = true

	// Affects children as well
	rootCmd.SilenceUsage = true

	// Prevents Cobra from dealing with errors as we deal with them in main.go
	rootCmd.SilenceErrors = true

	if params.Output != nil {
		rootCmd.SetOutput(params.Output)
	}

	//TODO: add common source commands flags here

	createCmd := commandFactory.CreateCommand()
	createCmd.Flags().AddFlagSet(flagsFactory.CreateFlags())
	createCmd.RunE = runEFactory.CreateRunE()
	rootCmd.AddCommand(createCmd)

	deleteCmd := commandFactory.DeleteCommand()
	deleteCmd.Flags().AddFlagSet(flagsFactory.DeleteFlags())
	deleteCmd.RunE = runEFactory.DeleteRunE()
	rootCmd.AddCommand(deleteCmd)

	updateCmd := commandFactory.UpdateCommand()
	updateCmd.Flags().AddFlagSet(flagsFactory.UpdateFlags())
	updateCmd.RunE = runEFactory.UpdateRunE()
	rootCmd.AddCommand(updateCmd)

	describeCmd := commandFactory.DescribeCommand()
	describeCmd.Flags().AddFlagSet(flagsFactory.DescribeFlags())
	describeCmd.RunE = runEFactory.DescribeRunE()
	rootCmd.AddCommand(describeCmd)

	// Initialize default `help` cmd early to prevent unknown command errors
	rootCmd.InitDefaultHelpCmd()

	return rootCmd
}

// Private

func addFlags(cmd *cobra.Command, flags []*pflag.Flag) {
	for _, flag := range flags {
		cmd.Flags().AddFlag(flag)
	}
}
