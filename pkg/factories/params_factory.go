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
)

type kafkaSourceParamsFactory struct {
	defaultParamsFactory sourcetypes.ParamsFactory
	kafkaSourceParams    *types.KafkaSourceParams
}

func NewKafkaSourceParamsFactory() types.KafkaSourceParamsFactory {
	return &kafkaSourceParamsFactory{
		defaultParamsFactory: sourcefactories.NewDefaultParamsFactory(),
	}
}

func (f *kafkaSourceParamsFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams.KnSourceParams
}

func (f *kafkaSourceParamsFactory) KafkaSourceParams() *types.KafkaSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams
}

func (f *kafkaSourceParamsFactory) CreateKnSourceParams() *sourcetypes.KnSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams.KnSourceParams
}

func (f *kafkaSourceParamsFactory) CreateKafkaSourceParams() *types.KafkaSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams
}

// Private

func (f *kafkaSourceParamsFactory) initKafkaSourceParams() {
	f.kafkaSourceParams = &types.KafkaSourceParams{
		KnSourceParams: f.defaultParamsFactory.CreateKnSourceParams(),
	}
	f.kafkaSourceParams.KnSourceParams.Initialize()
}
