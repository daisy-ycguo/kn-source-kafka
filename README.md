# kn-source_kafka

`kn-source_kafka` is a plugin of Knative Client, which allows you to management
of Kafka event source interactively from the command line.

## Description

`kn-source_kafka` is a plugin of Knative Client. You could create, update,
describe and delete Kafka event sources in Knative Eventing. Go to
[Knative Eventing document](https://knative.dev/docs/eventing/samples/kafka/source/)
to understand more about Kafka event sources.

## Build and Install

You must
[set up your development environment](https://github.com/knative/client/blob/master/docs/DEVELOPMENT.md#prerequisites)
before you build `kn-source_kafka`.

**Building:**

Once you've set up your development environment, let's build `kn-source_kafka`.
Run below command under the directory of `client-contrib/plugins/source-kafka`.

```sh
$ hack/build.sh
```

**Installing:**

You will get an excuatable file `kn-source_kafka` under the directory of
`client-contrib/plugins/source-kafka` after you run the build command. Then
let's install it to become a Knative Client `kn` plugin.

Install a plugin by simply copying the excuatable file `kn-source_kafka` to the
folder of the `kn` plugins directory. You will be able to invoke it by
`kn source_kafka`.

## Usage
### kafka

Knative eventing Kafka source plugin

#### Synopsis

Manage your Knative Kafka eventing sources

#### Options

```
  -h, --help   help for kafka
```

#### SEE ALSO

* [kafka create](#kafka-create)	 - create NAME
* [kafka delete](#kafka-delete)	 - delete NAME
* [kafka describe](#kafka-describe)	 - describe NAME
* [kafka update](#kafka-update)	 - update NAME

### kafka create

create NAME

#### Synopsis

create NAME

```
kafka create NAME [flags]
```

#### Examples

```
#Creates a new Kafka source with NAME
kn source kafka create kafka-name
```

#### Options

```
  -A, --all-namespaces         If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
      --consumergroup string   the consumer group ID
  -h, --help                   help for create
  -n, --namespace string       Specify the namespace to operate in.
      --servers string         Kafka bootstrap servers that the consumer will connect to, consist of a hostname plus a port pair, e.g. my-kafka-bootstrap.kafka:9092
  -s, --sink string            Addressable sink for events
      --topics string          Topics to consume messages from
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing Kafka source plugin

### kafka delete

delete NAME

#### Synopsis

delete a Kafka source

```
kafka delete NAME [flags]
```

#### Examples

```
#Deletes a Kafka source with NAME
kn source kafka delete kafka-name
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for delete
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing Kafka source plugin

### kafka describe

describe NAME

#### Synopsis

update a Kafka source

```
kafka describe NAME [flags]
```

#### Examples

```
#Describes a Kafka source with NAME
kn source kafka describe kafka-name
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for describe
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing Kafka source plugin

### kafka update

update NAME

#### Synopsis

update a Kafka source

```
kafka update NAME [flags]
```

#### Examples

```
#Updates a Kafka source with NAME
kn source kafka update kafka-name
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for update
  -n, --namespace string   Specify the namespace to operate in.
  -s, --sink string        Addressable sink for events
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing Kafka source plugin

