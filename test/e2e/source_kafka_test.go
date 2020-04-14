// Copyright 2020 The Knative Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build e2e
// +build !eventing

package e2e

import (
	"os"
	"path/filepath"
	"testing"

	"gotest.tools/assert"
	"knative.dev/client/lib/test"
	"knative.dev/client/pkg/util"
)

func SourceKafkaTest(t *testing.T) {
	t.Parallel()

	currentDir, err := os.Getwd()
	assert.NilError(t, err)

	it, err := NewE2ETest("kn-source_kafka", filepath.Join(currentDir, "../.."), false)
	assert.NilError(t, err)
	defer func() {
		assert.NilError(t, it.KnTest().Teardown())
	}()

	r := test.NewKnRunResultCollector(t)
	defer r.DumpIfFailed()

	err = it.KnPlugin().Install()
	assert.NilError(t, err)

	serviceCreate(r, "sinksvc")

	t.Log("kn-source_kafka create 'source-name'")
	knSourceKafkaCreate(r, "mykafka", "sinksvc")

	err = it.KnPlugin().Uninstall()
	assert.NilError(t, err)
}

// Private

func knSourceKafkaCreate(r *test.KnRunResultCollector, sourceName, server, topic, string) {
	out := r.KnTest().KnPlugin().Run("create", sourceName, "--sink", sinkName)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "create", sourceName))
}

func serviceCreate(r *test.KnRunResultCollector, serviceName string) {
	out := r.KnTest().Kn().Run("service", "create", serviceName, "--image", test.KnDefaultTestImage)
	r.AssertNoError(out)
	assert.Check(r.T(), util.ContainsAllIgnoreCase(out.Stdout, "service", serviceName, "creating", "namespace", r.KnTest().Kn().Namespace(), "ready"))
}