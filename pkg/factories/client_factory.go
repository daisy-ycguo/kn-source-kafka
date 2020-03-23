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
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type kafkaClientFactory struct {
	kafkaSourceClient  types.KafkaSourceClient
	kafkaParamsFactory types.KafkaSourceParamsFactory
}

func NewKafkaSourceClientFactory(kafkaPFactory types.KafkaSourceParamsFactory) types.KafkaSourceClientFactory {
	return &kafkaClientFactory{
		kafkaParamsFactory: kafkaPFactory,
	}
}

func (f *kafkaClientFactory) CreateKnSourceClient(namespace string) sourcetypes.KnSourceClient {
	return f.CreateKafkaSourceClient(namespace)
}

func (f *kafkaClientFactory) CreateKafkaSourceClient(namespace string) types.KafkaSourceClient {
	if f.kafkaSourceClient == nil {
		f.initKafkaSourceClient(namespace)
	}
	return f.kafkaSourceClient
}

func (f *kafkaClientFactory) KafkaSourceParams() *types.KafkaSourceParams {
	return f.kafkaParamsFactory.KafkaSourceParams()
}

func (f *kafkaClientFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.kafkaParamsFactory.KnSourceParams()
}

// Private

func (f *kafkaClientFactory) initKafkaSourceClient(namespace string) {
	if f.kafkaSourceClient == nil {
		f.kafkaSourceClient = client.NewKafkaSourceClient(f.kafkaParamsFactory, namespace)
	}
}
