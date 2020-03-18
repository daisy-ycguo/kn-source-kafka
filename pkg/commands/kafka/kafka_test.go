// Copyright © 2020 The Knative Authors
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
	"bytes"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"

	clientv1alpha1 "github.com/daisy-ycguo/kn-source-kafka/pkg/kafka/v1alpha1"
	kndynamic "knative.dev/client/pkg/dynamic"
	"knative.dev/client/pkg/kn/commands"
	v1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1alpha1"
)

// Helper methods
var blankConfig clientcmd.ClientConfig

// TODO: Remove that blankConfig hack for tests in favor of overwriting GetConfig()
// Remove also in service_test.go
func init() {
	var err error
	blankConfig, err = clientcmd.NewClientConfigFromBytes([]byte(`kind: Config
version: v1
users:
- name: u
clusters:
- name: c
  cluster:
    server: example.com
contexts:
- name: x
  context:
    user: u
    cluster: c
current-context: x
`))
	if err != nil {
		panic(err)
	}
}

func executeKafkaSourceCommand(kafkaClient clientv1alpha1.KafkaSourcesClient, dynamicClient kndynamic.KnDynamicClient, args ...string) (string, error) {
	knParams := &commands.KnParams{}
	knParams.ClientConfig = blankConfig

	output := new(bytes.Buffer)
	knParams.Output = output
	knParams.NewDynamicClient = func(namespace string) (kndynamic.KnDynamicClient, error) {
		return dynamicClient, nil
	}

	cmd := NewKafkaCommand(knParams)
	cmd.SetArgs(args)
	cmd.SetOutput(output)

	kafkaSourceClientFactory = func(config clientcmd.ClientConfig, namespace string) (clientv1alpha1.KafkaSourcesClient, error) {
		return kafkaClient, nil
	}
	defer cleanupKafkaMockClient()

	err := cmd.Execute()

	return output.String(), err
}

func cleanupKafkaMockClient() {
	kafkaSourceClientFactory = nil
}

func createKafkaSource(name, server, topics, group, service string) *v1alpha1.KafkaSource {
	sink := &duckv1beta1.Destination{
		Ref: &corev1.ObjectReference{
			Kind:       "Service",
			Name:       service,
			APIVersion: "serving.knative.dev/v1",
			Namespace:  "default",
		}}
	return clientv1alpha1.NewKafkaSourceBuilder(name).BootstrapServers(server).Topics(topics).ConsumerGroup(group).
		Sink(sink).
		Build()
}
