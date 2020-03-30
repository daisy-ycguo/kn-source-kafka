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
	"github.com/daisy-ycguo/kn-source-kafka/pkg/client"
	"github.com/daisy-ycguo/kn-source-kafka/pkg/types"
	sourcefactories "github.com/maximilien/kn-source-pkg/pkg/factories"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type kafkaClientFactory struct {
	kafkaSourceParams *types.KafkaSourceParams
	kafkaSourceClient types.KafkaSourceClient

	knSourceFactory sourcetypes.KnSourceFactory
}

func NewKafkaSourceFactory() types.KafkaSourceFactory {
	return &kafkaClientFactory{
		kafkaSourceParams: nil,
		kafkaSourceClient: nil,
		knSourceFactory:   sourcefactories.NewDefaultKnSourceFactory(),
	}
}

func (f *kafkaClientFactory) CreateKafkaSourceClient(namespace string) types.KafkaSourceClient {
	if f.kafkaSourceClient == nil {
		f.initKafkaSourceClient(namespace)
	}
	return f.kafkaSourceClient
}

func (f *kafkaClientFactory) KafkaSourceParams() *types.KafkaSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}

	return f.kafkaSourceParams
}

func (f *kafkaClientFactory) KafkaSourceClient() types.KafkaSourceClient {
	return f.kafkaSourceClient
}

func (f *kafkaClientFactory) CreateKafkaSourceParams() *types.KafkaSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}

	return f.kafkaSourceParams
}

//KnSources
func (f *kafkaClientFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams.KnSourceParams
}

func (f *kafkaClientFactory) CreateKnSourceParams() *sourcetypes.KnSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams.KnSourceParams
}

func (f *kafkaClientFactory) CreateKnSourceClient(namespace string) sourcetypes.KnSourceClient {
	return f.CreateKafkaSourceClient(namespace)
}

// Private

func (f *kafkaClientFactory) initKafkaSourceClient(namespace string) {
	if f.kafkaSourceClient == nil {
		f.kafkaSourceClient = client.NewKafkaSourceClient(f.KafkaSourceParams(), namespace)
	}
}

// Private

func (f *kafkaClientFactory) initKafkaSourceParams() {
	f.kafkaSourceParams = &types.KafkaSourceParams{
		KnSourceParams: f.knSourceFactory.CreateKnSourceParams(),
	}
	f.kafkaSourceParams.KnSourceParams.Initialize()
}
