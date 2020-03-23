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
	"github.com/maximilien/kn-source-pkg/pkg/types"
)

type DefautParamsFactory struct {
	knSourceParams *types.KnSourceParams
}

func NewDefaultParamsFactory() types.ParamsFactory {
	return &DefautParamsFactory{}
}

func (f *DefautParamsFactory) KnSourceParams() *types.KnSourceParams {
	return f.knSourceParams
}

func (f *DefautParamsFactory) CreateKnSourceParams() *types.KnSourceParams {
	f.knSourceParams = &types.KnSourceParams{}
	f.knSourceParams.Initialize()
	return f.knSourceParams
}
