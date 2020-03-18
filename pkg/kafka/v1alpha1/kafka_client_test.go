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

package v1alpha1

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clienttesting "k8s.io/client-go/testing"
	fake "knative.dev/eventing-contrib/kafka/source/pkg/client/clientset/versioned/typed/sources/v1alpha1/fake"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

var testKafkaSourceNamespace = "test-ns"

func setupKafkaSourcesClient(t *testing.T) (fakeSources fake.FakeSourcesV1alpha1, client KafkaSourcesClient) {
	fakeSources = fake.FakeSourcesV1alpha1{Fake: &clienttesting.Fake{}}
	client = NewSourcesClient(&fakeSources, testKafkaSourceNamespace).KafkaSourcesClient()
	assert.Equal(t, client.Namespace(), testKafkaSourceNamespace)
	return
}

func TestCreateKafkaSource(t *testing.T) {
	sourcesServer, client := setupKafkaSourcesClient(t)

	sourcesServer.AddReactor("create", "kafka",
		func(a clienttesting.Action) (bool, runtime.Object, error) {
			newSource := a.(clienttesting.CreateAction).GetObject()
			name := newSource.(metav1.Object).GetName()
			if name == "errorSource" {
				return true, nil, fmt.Errorf("error while creating KafkaSource %s", name)
			}
			return true, newSource, nil
		})
	err := client.CreateKafkaSource(newKafkaSource("foo", "Event"))
	assert.NilError(t, err)

	err = client.CreateKafkaSource(newKafkaSource("errorSource", "Event"))
	assert.ErrorContains(t, err, "errorSource")

}

func newKafkaSource(name, topics string) *v1alpha1.KafkaSource {
	b := NewKafkaSourceBuilder(name).BootstrapServers("mykafka.com").Topics(topics).ConsumerGroup("mygroup")
	b.Sink(&duckv1beta1.Destination{
		Ref: &v1.ObjectReference{
			Kind:      "Service",
			Name:      "foosvc",
			Namespace: "default",
		}})
	return b.Build()
}
