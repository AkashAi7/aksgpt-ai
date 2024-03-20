# serve

The serve commands allow you to run aksgpt in a grpc server mode.
This would be enabled typically through `aksgpt serve` and is how the in-cluster aksgpt deployment functions when managed by the [aksgpt-operator](https://github.com/aksgpt-ai/aksgpt-operator)

The grpc interface that is served is hosted on [buf](https://buf.build/aksgpt-ai/schemas) and the repository for this is [here](https://github.com/aksgpt-ai/schemas)

## grpcurl

A fantastic tool for local debugging and development is `grpcurl` 
It allows you to form curl like requests that are http2
e.g. 

```
grpcurl -plaintext -d '{"namespace": "aksgpt", "explain" : "true"}' localhost:8080 schema.v1.ServerService/Analyze
```

```
grpcurl -plaintext  localhost:8080 schema.v1.ServerService/ListIntegrations 
{
  "integrations": [
    "trivy"
  ]
}

```

```
grpcurl -plaintext -d '{"integrations":{"trivy":{"enabled":"true","namespace":"default","skipInstall":"false"}}}' localhost:8080 schema.v1.ServerService/AddConfig
```
