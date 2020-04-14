#!/usr/bin/env bash

# Copyright 2020 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

export PATH=$PWD:$PATH

dir=$(dirname "${BASH_SOURCE[0]}")
base=$(cd "$dir/.." && pwd)

# Strimzi installation config template used for starting up Kafka clusters.
readonly STRIMZI_INSTALLATION_CONFIG_TEMPLATE="test/config/100-strimzi-cluster-operator-0.17.0.yaml"
# Strimzi installation config.
readonly STRIMZI_INSTALLATION_CONFIG="$(mktemp)"
# Kafka cluster CR config file.
readonly KAFKA_INSTALLATION_CONFIG="test/config/100-kafka-ephemeral-triple-2.4.0.yaml"
readonly KAFKA_TOPIC_INSTALLATION_CONFIG="test/config/100-kafka-topic.yaml"
# Kafka cluster URL for our installation
readonly KAFKA_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka:9092"
# Kafka channel CRD config template directory.
readonly KAFKA_CRD_CONFIG_TEMPLATE_DIR="kafka/channel/config"
# Kafka channel CRD config template file. It needs to be modified to be the real config file.
readonly KAFKA_CRD_CONFIG_TEMPLATE="400-kafka-config.yaml"
# Real Kafka channel CRD config , generated from the template directory and modified template file.
readonly KAFKA_CRD_CONFIG_DIR="$(mktemp -d)"
# Kafka channel CRD config template directory.
readonly KAFKA_SOURCE_CRD_CONFIG_DIR="kafka/source/config"

function kafka_setup() {
  echo "Installing Kafka cluster"
  kubectl create namespace kafka || return 1
  sed 's/namespace: .*/namespace: kafka/' ${STRIMZI_INSTALLATION_CONFIG_TEMPLATE} > ${STRIMZI_INSTALLATION_CONFIG}
  kubectl apply -f "${STRIMZI_INSTALLATION_CONFIG}" -n kafka
  kubectl apply -f ${KAFKA_INSTALLATION_CONFIG} -n kafka
  wait_until_pods_running kafka || fail_test "Failed to start up a Kafka cluster"
}

function kafka_teardown() {
  echo "Uninstalling Kafka cluster"
  kubectl delete -f ${KAFKA_TOPIC_INSTALLATION_CONFIG} -n kafka
  kubectl delete -f ${KAFKA_INSTALLATION_CONFIG} -n kafka
  kubectl delete -f "${STRIMZI_INSTALLATION_CONFIG}" -n kafka
  kubectl delete namespace kafka
}

function plugin_test_setup() {
  kafka_setup || return 1
}

function plugin_test_teardown() {
  kafka_teardown
}

# TODO: Check and Install Kn if it's not present

# Will create and delete this namespace (used for all tests, modify if you want a different one used)
export KN_E2E_NAMESPACE=kne2etests

echo "🧪  Setup"
plugin_test_setup
echo "🧪  Testing"
go test ${base}/test/e2e/ -test.v -tags "e2e ${E2E_TAGS}" "$@"
echo "🧪  Teardown"
plugin_test_teardown