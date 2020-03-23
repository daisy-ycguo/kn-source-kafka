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
	"fmt"

	"github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
)

type DefautRunEFactory struct {
	knSourceParams        *types.KnSourceParams
	knSourceClientFactory types.ClientFactory
	knSourceClient        types.KnSourceClient
}

func NewDefaultRunEFactory(knSourceParams *types.KnSourceParams, clientFactory types.ClientFactory) types.RunEFactory {
	return &DefautRunEFactory{
		knSourceParams:        knSourceParams,
		knSourceClientFactory: clientFactory,
		knSourceClient:        nil,
	}
}

func (f *DefautRunEFactory) KnSourceParams() *types.KnSourceParams {
	return f.knSourceParams
}

func (f *DefautRunEFactory) KnSourceClientFactory() types.ClientFactory {
	return f.knSourceClientFactory
}

func (f *DefautRunEFactory) KnSourceClient(cmd *cobra.Command) error {
	if f.knSourceClient == nil {
		p := f.KnSourceParams().KnParams
		namespace, err := p.GetNamespace(cmd)
		if err != nil {
			return err
		}
		f.knSourceClient = f.knSourceClientFactory.CreateKnSourceClient(namespace)
	}
	return nil
}

func (f *DefautRunEFactory) CreateRunE() types.RunE {
	return func(cmd *cobra.Command, args []string) error {
		f.KnSourceClient(cmd)
		fmt.Printf("%s RunE called: args: %#v, client: %#v\n", cmd.Name(), args, f.knSourceClient)
		return nil
	}
}

func (f *DefautRunEFactory) DeleteRunE() types.RunE {
	return func(cmd *cobra.Command, args []string) error {
		f.KnSourceClient(cmd)
		fmt.Printf("%s RunE called: args: %#v, client: %#v\n", cmd.Name(), args, f.knSourceClient)
		return nil
	}
}

func (f *DefautRunEFactory) UpdateRunE() types.RunE {
	return func(cmd *cobra.Command, args []string) error {
		f.KnSourceClient(cmd)
		fmt.Printf("%s RunE called: args: %#v, client: %#v\n", cmd.Name(), args, f.knSourceClient)
		return nil
	}
}

func (f *DefautRunEFactory) DescribeRunE() types.RunE {
	return func(cmd *cobra.Command, args []string) error {
		f.KnSourceClient(cmd)
		fmt.Printf("%s RunE called: args: %#v, client: %#v\n", cmd.Name(), args, f.knSourceClient)
		return nil
	}
}
