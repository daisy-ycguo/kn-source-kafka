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

package kafka

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/daisy-ycguo/kn-source-kafka/pkg/kafka/v1alpha1"
	"knative.dev/client/pkg/kn/commands"
	clientv1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/client/clientset/versioned/typed/sources/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	v1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/apis/duck/v1beta1"
)

// NewKafkaCommand for managing ApiServer source
func NewKafkaCommand(p *commands.KnParams) *cobra.Command {
	kafkaCmd := &cobra.Command{
		Use:   "kafka",
		Short: "Kafka Source command group",
	}
	return kafkaCmd
}

var kafkaSourceClientFactory func(config clientcmd.ClientConfig, namespace string) (v1alpha1.KafkaSourcesClient, error)

func newKafkaSourceClient(p *commands.KnParams, cmd *cobra.Command) (v1alpha1.KafkaSourcesClient, error) {
	namespace, err := p.GetNamespace(cmd)
	if err != nil {
		return nil, err
	}

	if kafkaSourceClientFactory != nil {
		config, err := p.GetClientConfig()
		if err != nil {
			return nil, err
		}
		return kafkaSourceClientFactory(config, namespace)
	}

	clientConfig, err := p.RestConfig()
	if err != nil {
		return nil, err
	}

	client, err := clientv1alpha1.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	return v1alpha1.NewSourcesClient(client, namespace).KafkaSourcesClient(), nil
}

func toDuckV1Beta1(destination *v1.Destination) *v1beta1.Destination {
	r := destination.Ref
	return &v1beta1.Destination{
		Ref: &corev1.ObjectReference{
			Kind:       r.Kind,
			Namespace:  r.Namespace,
			Name:       r.Name,
			APIVersion: r.APIVersion,
		},
		URI: destination.URI,
	}
}
