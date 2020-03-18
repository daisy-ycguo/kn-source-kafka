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
	"testing"

	"gotest.tools/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	"github.com/daisy-ycguo/kn-source-kafka/pkg/kafka/v1alpha1"
	dynamicfake "knative.dev/client/pkg/dynamic/fake"

	"knative.dev/client/pkg/util"
)

func TestSimpleCreateKafkaSource(t *testing.T) {
	mysvc := &servingv1.Service{
		TypeMeta:   v1.TypeMeta{Kind: "Service", APIVersion: "serving.knative.dev/v1"},
		ObjectMeta: v1.ObjectMeta{Name: "mysvc", Namespace: "default"},
	}
	dynamicClient := dynamicfake.CreateFakeKnDynamicClient("default", mysvc)

	kafkaClient := v1alpha1.NewMockKafkaSourcesClient(t)

	kafkaRecorder := kafkaClient.Recorder()
	kafkaRecorder.CreateKafkaSource(createKafkaSource("testsource", "myhost.com", "mytopic", "mygroup", "mysvc"), nil)

	out, err := executeKafkaSourceCommand(kafkaClient, dynamicClient, "create", "--servers", "myhost.com", "--topics", "mytopic", "--consumergroup", "mygroup", "--sink", "mysvc", "testsource")
	assert.NilError(t, err, "Source should have been created")
	util.ContainsAll(out, "created", "default", "testsource")

	kafkaRecorder.Validate()
}
