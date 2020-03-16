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

package v1alpha1

import (
	knerrors "knative.dev/client/pkg/errors"
	v1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1alpha1"
	clientv1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/client/clientset/versioned/typed/sources/v1alpha1"
)

// KafkaSourcesClient interface for working with Kafka sources
type KafkaSourcesClient interface {
	// Create an KafkaSource by object
	CreateKafkaSource(kafkaSource *v1alpha1.KafkaSource) error

	// Get namespace for this client
	Namespace() string
}

type kafkaSourcesClient struct {
	client    clientv1alpha1.KafkaSourceInterface
	namespace string
}

// newKnAPIServerSourcesClient is to invoke Eventing Sources Client API to create object
func newKafkaSourcesClient(client clientv1alpha1.KafkaSourceInterface, namespace string) KafkaSourcesClient {
	return &kafkaSourcesClient{
		client:    client,
		namespace: namespace,
	}
}

//CreateKafkaSource is used to create an instance of ApiServerSource
func (c *kafkaSourcesClient) CreateKafkaSource(kafkaSource *v1alpha1.KafkaSource) error {
	_, err := c.client.Create(kafkaSource)
	if err != nil {
		return knerrors.GetError(err)
	}

	return nil
}

// Return the client's namespace
func (c *kafkaSourcesClient) Namespace() string {
	return c.namespace
}
