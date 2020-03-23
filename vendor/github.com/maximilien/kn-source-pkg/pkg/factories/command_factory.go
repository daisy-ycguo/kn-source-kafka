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
	"github.com/maximilien/kn-source-pkg/pkg/commands/source"
	"github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
)

type DefautCommandFactory struct {
	knSourceParams *types.KnSourceParams
}

func NewDefaultCommandFactory(knSourceParams *types.KnSourceParams) types.CommandFactory {
	return &DefautCommandFactory{
		knSourceParams: knSourceParams,
	}
}

func (f *DefautCommandFactory) KnSourceParams() *types.KnSourceParams {
	return f.knSourceParams
}

func (f *DefautCommandFactory) SourceCommand() *cobra.Command {
	return source.NewSourceCommand(f.KnSourceParams())
}

func (f *DefautCommandFactory) CreateCommand() *cobra.Command {
	return source.NewCreateCommand(f.KnSourceParams())
}

func (f *DefautCommandFactory) DeleteCommand() *cobra.Command {
	return source.NewDeleteCommand(f.KnSourceParams())
}

func (f *DefautCommandFactory) UpdateCommand() *cobra.Command {
	return source.NewUpdateCommand(f.KnSourceParams())
}

func (f *DefautCommandFactory) DescribeCommand() *cobra.Command {
	return source.NewDescribeCommand(f.KnSourceParams())
}
