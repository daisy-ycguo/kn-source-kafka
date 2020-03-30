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
	"github.com/maximilien/kn-source-pkg/pkg/client"
	"github.com/maximilien/kn-source-pkg/pkg/types"
)

type DefautKnSourceFactory struct {
	knSourceParams *types.KnSourceParams
}

func NewDefaultKnSourceFactory() types.KnSourceFactory {
	return &DefautKnSourceFactory{}
}

func (f *DefautKnSourceFactory) KnSourceParams() *types.KnSourceParams {
	if f.knSourceParams == nil {
		f.knSourceParams = f.CreateKnSourceParams()
	}

	return f.knSourceParams
}

func (f *DefautKnSourceFactory) CreateKnSourceParams() *types.KnSourceParams {
	f.knSourceParams = &types.KnSourceParams{}
	f.knSourceParams.Initialize()
	return f.knSourceParams
}

func (f *DefautKnSourceFactory) CreateKnSourceClient(namespace string) types.KnSourceClient {
	return client.NewKnSourceClient(f.knSourceParams, namespace)
}
