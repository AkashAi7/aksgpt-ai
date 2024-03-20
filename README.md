<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./images/banner-white.png" width="600px;">
  <img alt="Text changing depending on mode. Light: 'So light!' Dark: 'So dark!'" src="./images/banner-black.png" width="600px;">
</picture>
<br/>

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/aksgpt-ai/aksgpt)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/aksgpt-ai/aksgpt/release.yaml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/aksgpt-ai/aksgpt)
[![OpenSSF Best Practices](https://bestpractices.coreinfrastructure.org/projects/7272/badge)](https://bestpractices.coreinfrastructure.org/projects/7272)
[![Link to documentation](https://img.shields.io/static/v1?label=%F0%9F%93%96&message=Documentation&color=blue)](https://docs.aksgpt.ai/)  
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Faksgpt-ai%2Faksgpt.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Faksgpt-ai%2Faksgpt?ref=badge_shield)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go version](https://img.shields.io/github/go-mod/go-version/aksgpt-ai/aksgpt.svg)](https://github.com/aksgpt-ai/aksgpt)
[![codecov](https://codecov.io/github/aksgpt-ai/aksgpt/graph/badge.svg?token=ZLR7NG8URE)](https://codecov.io/github/aksgpt-ai/aksgpt)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/aksgpt-ai/aksgpt/main)  

`aksgpt` is a tool for scanning your Kubernetes clusters, diagnosing, and triaging issues in simple English.

It has SRE experience codified into its analyzers and helps to pull out the most relevant information to enrich it with AI.

_Out of the box integration with OpenAI, Azure, Cohere, Amazon Bedrock and local models._

<a href="https://www.producthunt.com/posts/aksgpt?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-aksgpt" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=389489&theme=light" alt="aksgpt - aksgpt&#0032;gives&#0032;Kubernetes&#0032;Superpowers&#0032;to&#0032;everyone | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

<img src="images/demo4.gif" width=650px; />

# CLI Installation


### Linux/Mac via brew

```
brew tap aksgpt-ai/aksgpt
brew install aksgpt
```

<details>
  <summary>RPM-based installation (RedHat/CentOS/Fedora)</summary>

  **32 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/aksgpt-ai/aksgpt/releases/download/v0.3.27/aksgpt_386.rpm
  sudo rpm -ivh aksgpt_386.rpm
  ```
  <!---x-release-please-end-->

  **64 bit:**

  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/aksgpt-ai/aksgpt/releases/download/v0.3.27/aksgpt_amd64.rpm
  sudo rpm -ivh -i aksgpt_amd64.rpm
  ```
  <!---x-release-please-end-->
</details>

<details>
  <summary>DEB-based installation (Ubuntu/Debian)</summary>

  **32 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/aksgpt-ai/aksgpt/releases/download/v0.3.27/aksgpt_386.deb
  sudo dpkg -i aksgpt_386.deb
  ```
  <!---x-release-please-end-->
  **64 bit:**

  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/aksgpt-ai/aksgpt/releases/download/v0.3.27/aksgpt_amd64.deb
  sudo dpkg -i aksgpt_amd64.deb
  ```
  <!---x-release-please-end-->
</details>

<details>

  <summary>APK-based installation (Alpine)</summary>

  **32 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/aksgpt-ai/aksgpt/releases/download/v0.3.27/aksgpt_386.apk
  apk add aksgpt_386.apk
  ```
  <!---x-release-please-end-->
  **64 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/aksgpt-ai/aksgpt/releases/download/v0.3.27/aksgpt_amd64.apk
  apk add aksgpt_amd64.apk
  ```
  <!---x-release-please-end-->x
</details>

<details>
  <summary>Failing Installation on WSL or Linux (missing gcc)</summary>
  When installing Homebrew on WSL or Linux, you may encounter the following error:

  ```
  ==> Installing aksgpt from aksgpt-ai/aksgpt Error: The following formula cannot be installed from a bottle and must be
  built from the source. aksgpt Install Clang or run brew install gcc.
  ```

If you install gcc as suggested, the problem will persist. Therefore, you need to install the build-essential package.
  ```
     sudo apt-get update
     sudo apt-get install build-essential
  ```
</details>


### Windows

* Download the latest Windows binaries of **aksgpt** from the [Release](https://github.com/aksgpt-ai/aksgpt/releases)
  tab based on your system architecture.
* Extract the downloaded package to your desired location. Configure the system *path* variable with the binary location

## Operator Installation

To install within a Kubernetes cluster please use our `aksgpt-operator` with installation instructions available [here](https://github.com/aksgpt-ai/aksgpt-operator)

_This mode of operation is ideal for continuous monitoring of your cluster and can integrate with your existing monitoring such as Prometheus and Alertmanager._


## Quick Start

* Currently, the default AI provider is OpenAI, you will need to generate an API key from [OpenAI](https://openai.com)
  * You can do this by running `aksgpt generate` to open a browser link to generate it
* Run `aksgpt auth add` to set it in aksgpt.
  * You can provide the password directly using the `--password` flag.
* Run `aksgpt filters` to manage the active filters used by the analyzer. By default, all filters are executed during analysis.
* Run `aksgpt analyze` to run a scan.
* And use `aksgpt analyze --explain` to get a more detailed explanation of the issues.
* You also run `aksgpt analyze --with-doc` (with or without the explain flag) to get the official documentation from Kubernetes.

## Analyzers

aksgpt uses analyzers to triage and diagnose issues in your cluster. It has a set of analyzers that are built in, but
you will be able to write your own analyzers.

### Built in analyzers

#### Enabled by default

- [x] podAnalyzer
- [x] pvcAnalyzer
- [x] rsAnalyzer
- [x] serviceAnalyzer
- [x] eventAnalyzer
- [x] ingressAnalyzer
- [x] statefulSetAnalyzer
- [x] deploymentAnalyzer
- [x] cronJobAnalyzer
- [x] nodeAnalyzer
- [x] mutatingWebhookAnalyzer
- [x] validatingWebhookAnalyzer

#### Optional

- [x] hpaAnalyzer
- [x] pdbAnalyzer
- [x] networkPolicyAnalyzer
- [x] gatewayClass
- [x] gateway
- [x] httproute

## Examples

_Run a scan with the default analyzers_

```
aksgpt generate
aksgpt auth add
aksgpt analyze --explain
aksgpt analyze --explain --with-doc
```

_Filter on resource_

```
aksgpt analyze --explain --filter=Service
```

_Filter by namespace_
```
aksgpt analyze --explain --filter=Pod --namespace=default
```

_Output to JSON_

```
aksgpt analyze --explain --filter=Service --output=json
```

_Anonymize during explain_

```
aksgpt analyze --explain --filter=Service --output=json --anonymize
```

<details>
<summary> Using filters </summary>

_List filters_

```
aksgpt filters list
```

_Add default filters_

```
aksgpt filters add [filter(s)]
```

### Examples :

- Simple filter : `aksgpt filters add Service`
- Multiple filters : `aksgpt filters add Ingress,Pod`

_Remove default filters_

```
aksgpt filters remove [filter(s)]
```

### Examples :

- Simple filter : `aksgpt filters remove Service`
- Multiple filters : `aksgpt filters remove Ingress,Pod`

</details>

<details>

<summary> Additional commands </summary>

_List configured backends_

```
aksgpt auth list
```

_Update configured backends_

```
aksgpt auth update $MY_BACKEND1,$MY_BACKEND2..
```

_Remove configured backends_

```
aksgpt auth remove -b $MY_BACKEND1,$MY_BACKEND2..
```

_List integrations_

```
aksgpt integrations list
```

_Activate integrations_

```
aksgpt integrations activate [integration(s)]
```

_Use integration_

```
aksgpt analyze --filter=[integration(s)]
```

_Deactivate integrations_

```
aksgpt integrations deactivate [integration(s)]
```

_Serve mode_

```
aksgpt serve
```

_Analysis with serve mode_

```
grpcurl -plaintext -d '{"namespace": "aksgpt", "explain": false}' localhost:8080 schema.v1.ServerService/Analyze
```
</details>

## LLM AI Backends

aksgpt uses the chosen LLM, generative AI provider when you want to explain the analysis results using --explain flag e.g. `aksgpt analyze --explain`. You can use `--backend` flag to specify a configured provider (it's `openai` by default).

You can list available providers using `aksgpt auth list`:

```
Default: 
> openai
Active: 
Unused: 
> openai
> localai
> azureopenai
> cohere
> amazonbedrock
> amazonsagemaker
> google
> huggingface
> noopai
```

For detailed documentation on how to configure and use each provider see [here](https://docs.aksgpt.ai/reference/providers/backend/).

_To set a new default provider_

```
aksgpt auth default -p azureopenai
Default provider set to azureopenai
```

## Key Features

<details>

With this option, the data is anonymized before being sent to the AI Backend. During the analysis execution, `aksgpt` retrieves sensitive data (Kubernetes object names, labels, etc.). This data is masked when sent to the AI backend and replaced by a key that can be used to de-anonymize the data when the solution is returned to the user.

<summary> Anonymization </summary>

1. Error reported during analysis:
```bash
Error: HorizontalPodAutoscaler uses StatefulSet/fake-deployment as ScaleTargetRef which does not exist.
```

2. Payload sent to the AI backend:
```bash
Error: HorizontalPodAutoscaler uses StatefulSet/tGLcCRcHa1Ce5Rs as ScaleTargetRef which does not exist.
```

3. Payload returned by the AI:
```bash
The Kubernetes system is trying to scale a StatefulSet named tGLcCRcHa1Ce5Rs using the HorizontalPodAutoscaler, but it cannot find the StatefulSet. The solution is to verify that the StatefulSet name is spelled correctly and exists in the same namespace as the HorizontalPodAutoscaler.
```

4. Payload returned to the user:
```bash
The Kubernetes system is trying to scale a StatefulSet named fake-deployment using the HorizontalPodAutoscaler, but it cannot find the StatefulSet. The solution is to verify that the StatefulSet name is spelled correctly and exists in the same namespace as the HorizontalPodAutoscaler.
```

Note: **Anonymization does not currently apply to events.**

### Further Details

**Anonymization does not currently apply to events.**

*In a few analysers like Pod, we feed to the AI backend the event messages which are not known beforehand thus we are not masking them for the **time being**.*

- The following is the list of analysers in which data is **being masked**:-

  - Statefulset
  - Service
  - PodDisruptionBudget
  - Node
  - NetworkPolicy
  - Ingress
  - HPA
  - Deployment
  - Cronjob

- The following is the list of analysers in which data is **not being masked**:-

  - RepicaSet
  - PersistentVolumeClaim
  - Pod
  - **_*Events_**

***Note**:
  - k8gpt will not mask the above analysers because they do not send any identifying information except **Events** analyser.
  - Masking for **Events** analyzer is scheduled in the near future as seen in this [issue](https://github.com/aksgpt-ai/aksgpt/issues/560). _Further research has to be made to understand the patterns and be able to mask the sensitive parts of an event like pod name, namespace etc._

- The following is the list of fields which are not **being masked**:-

  - Describe
  - ObjectStatus
  - Replicas
  - ContainerStatus
  - **_*Event Message_**
  - ReplicaStatus
  - Count (Pod)

***Note**:
  - It is quite possible the payload of the event message might have something like "super-secret-project-pod-X crashed" which we don't currently redact _(scheduled in the near future as seen in this [issue](https://github.com/aksgpt-ai/aksgpt/issues/560))_.

### Proceed with care

  - The K8gpt team recommends using an entirely different backend **(a local model) in critical production environments**. By using a local model, you can rest assured that everything stays within your DMZ, and nothing is leaked.
  - If there is any uncertainty about the possibility of sending data to a public LLM (open AI, Azure AI) and it poses a risk to business-critical operations, then, in such cases, the use of public LLM should be avoided based on personal assessment and the jurisdiction of risks involved.


</details>

<details>
<summary> Configuration management</summary>

`aksgpt` stores config data in the `$XDG_CONFIG_HOME/aksgpt/aksgpt.yaml` file. The data is stored in plain text, including your OpenAI key.

Config file locations:
| OS      | Path                                             |
| ------- | ------------------------------------------------ |
| MacOS   | ~/Library/Application Support/aksgpt/aksgpt.yaml |
| Linux   | ~/.config/aksgpt/aksgpt.yaml                     |
| Windows | %LOCALAPPDATA%/aksgpt/aksgpt.yaml                |
</details>

<details>
There may be scenarios where caching remotely is preferred.
In these scenarios aksgpt supports AWS S3 or Azure Blob storage Integration.

<summary> Remote caching </summary>  
<em>Note: You can only configure and use only one remote cache at a time</em>

_Adding a remote cache_

 * AWS S3
   *  _As a prerequisite `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` are required as environmental variables._
   * Configuration, ``` aksgpt cache add s3 --region <aws region> --bucket <name> ```
     * aksgpt will create the bucket if it does not exist
 * Azure Storage
   * We support a number of [techniques](https://learn.microsoft.com/en-us/azure/developer/go/azure-sdk-authentication?tabs=bash#2-authenticate-with-azure) to authenticate against Azure
   * Configuration, ``` aksgpt cache add azure --storageacc <storage account name> --container <container name> ```
     * aksgpt assumes that the storage account already exist and it will create the container if it does not exist
     * It is the **user** responsibility have to grant specific permissions to their identity in order to be able to upload blob files and create SA containers (e.g Storage Blob Data Contributor) 
  * Google Cloud Storage
    * _As a prerequisite `GOOGLE_APPLICATION_CREDENTIALS` are required as environmental variables._
    * Configuration, ``` aksgpt cache add gcs --region <gcp region> --bucket <name> --projectid <project id>```
      * aksgpt will create the bucket if it does not exist   

_Listing cache items_
```
aksgpt cache list
```

_Purging an object from the cache_
Note: purging an object using this command will delete upstream files, so it requires appropriate permissions.
```
aksgpt cache purge $OBJECT_NAME
```

_Removing the remote cache_
Note: this will not delete the upstream S3 bucket or Azure storage container
```
aksgpt cache remove
```
</details>

<details>
<summary> Custom Analyzers</summary>

There may be scenarios where you wish to write your own analyzer in a language of your choice.
aksgpt now supports the ability to do so by abiding by the [schema](https://github.com/aksgpt-ai/schemas/blob/main/protobuf/schema/v1/analyzer.proto) and serving the analyzer for consumption.
To do so, define the analyzer within the aksgpt configuration and it will add it into the scanning process.
In addition to this you will need to enable the following flag on analysis:
```
aksgpt analyze --custom-analysis
```

Here is an example local host analyzer in [Rust](https://github.com/aksgpt-ai/host-analyzer)
When this is run on `localhost:8080` the aksgpt config can pick it up with the following additions:

```
custom_analyzers:
  - name: host-analyzer
    connection:
      url: localhost
      port: 8080
```

This now gives the ability to pass through hostOS information ( from this analyzer example ) to aksgpt to use as context with normal analysis.

_See the docs on how to write a custom analyzer_

</details>

## Documentation

Find our official documentation available [here](https://docs.aksgpt.ai)


## Contributing

Please read our [contributing guide](./CONTRIBUTING.md).
## Community
Find us on [Slack](https://join.slack.com/t/aksgpt/shared_invite/zt-276pa9uyq-pxAUr4TCVHubFxEvLZuT1Q)

<a href="https://github.com/aksgpt-ai/aksgpt/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aksgpt-ai/aksgpt" />
</a>


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Faksgpt-ai%2Faksgpt.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Faksgpt-ai%2Faksgpt?ref=badge_large)
