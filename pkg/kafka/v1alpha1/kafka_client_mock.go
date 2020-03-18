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
	"testing"

	v1alpha1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1alpha1"

	"knative.dev/client/pkg/util/mock"
)

// MockKafkaSourcesClient is a combine of test object and recorder
type MockKafkaSourcesClient struct {
	t         *testing.T
	recorder  *KafkaSourcesRecorder
	namespace string
}

// NewMockKafkaSourcesClient returns a new mock instance which you need to record for
func NewMockKafkaSourcesClient(t *testing.T, ns ...string) *MockKafkaSourcesClient {
	namespace := "default"
	if len(ns) > 0 {
		namespace = ns[0]
	}
	return &MockKafkaSourcesClient{
		t:        t,
		recorder: &KafkaSourcesRecorder{mock.NewRecorder(t, namespace)},
	}
}

// Ensure that the interface is implemented
var _ KafkaSourcesClient = &MockKafkaSourcesClient{}

// KafkaSourcesRecorder is recorder for eventing objects
type KafkaSourcesRecorder struct {
	r *mock.Recorder
}

// Recorder returns the recorder for registering API calls
func (c *MockKafkaSourcesClient) Recorder() *KafkaSourcesRecorder {
	return c.recorder
}

// Namespace of this client
func (c *MockKafkaSourcesClient) Namespace() string {
	return c.recorder.r.Namespace()
}

// CreateKafkaSource records a call for CreateKafkaSource with the expected error
func (sr *KafkaSourcesRecorder) CreateKafkaSource(kafkaSource interface{}, err error) {
	sr.r.Add("CreateKafkaSource", []interface{}{kafkaSource}, []interface{}{err})
}

// CreateKafkaSource performs a previously recorded action, failing if non has been registered
func (c *MockKafkaSourcesClient) CreateKafkaSource(kafkaSource *v1alpha1.KafkaSource) error {
	call := c.recorder.r.VerifyCall("CreateKafkaSource", kafkaSource)
	return mock.ErrorOrNil(call.Result[0])
}

// Validate validates whether every recorded action has been called
func (sr *KafkaSourcesRecorder) Validate() {
	sr.r.CheckThatAllRecordedMethodsHaveBeenCalled()
}
