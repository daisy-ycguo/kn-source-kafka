// Copyright © 2019 The Knative Authors
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

package client

import (
	"github.com/daisy-ycguo/kn-source-kafka/pkg/types"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
	knerrors "knative.dev/client/pkg/errors"
	v1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1alpha1"
	clientv1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/client/clientset/versioned/typed/sources/v1alpha1"
)

type kafkaSourceClient struct {
	kafkaSourceParams *types.KafkaSourceParams
	client            clientv1alpha1.KafkaSourceInterface
	namespace         string
}

func NewKafkaSourceClient(kafkaParams *types.KafkaSourceParams, ns string) types.KafkaSourceClient {
	return &kafkaSourceClient{
		kafkaSourceParams: kafkaParams,
		namespace:         ns,
	}
}

func (client *kafkaSourceClient) KnSourceParams() *sourcetypes.KnSourceParams {
	return client.kafkaSourceParams.KnSourceParams
}

func (client *kafkaSourceClient) KafkaSourceParams() *types.KafkaSourceParams {
	return client.kafkaSourceParams
}

//CreateKafkaSource is used to create an instance of ApiServerSource
func (c *kafkaSourceClient) CreateKafkaSource(kafkaSource *v1alpha1.KafkaSource) error {
	_, err := c.client.Create(kafkaSource)
	if err != nil {
		return knerrors.GetError(err)
	}

	return nil
}

// Return the client's namespace
func (c *kafkaSourceClient) Namespace() string {
	return c.namespace
}
