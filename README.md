# M3 Chat Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/m3-chat/go-sdk.svg)](https://pkg.go.dev/github.com/m3-chat/go-sdk)   [![Go Report Card](https://goreportcard.com/badge/github.com/m3-chat/go-sdk)](https://goreportcard.com/report/github.com/m3-chat/go-sdk)   [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

The **M3 Chat Go SDK** is a Go client for interacting with the [M3 Chat](https://github.com/m3-chat) AI chat platform. It offers an easy-to-use, idiomatic Go interface compatible with the official TypeScript SDK, allowing you to send chat messages, stream responses, and batch multiple requests.

## Table of Contents
* [Introduction](https://github.com/m3-chat/go-sdk?tab=readme-ov-file)
* [Table of Contents](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#table-of-contents)
* [Features](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#features)
* [Installation](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#installation)
* [Quick Start](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#quick-start)
* [Usage](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#usage)
    * [Client Creation](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#client-creation)
    * [GetResponse](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#getresponse)
    * [BatchRequests](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#batchrequests)
* [Available Models](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#available-models)
* [Development](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#development)
* [Contribution](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#contribution)
* [License](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#license)
* [Contact](https://github.com/m3-chat/go-sdk?tab=readme-ov-file#contact)


## Features

- Send chat requests with configurable models and content
- Support for streaming and non-streaming responses
- Batch multiple chat requests sequentially
- Validate available models
- Easy to use and integrate into Go projects
- Compatible with Go 1.21+

## Installation

```bash
go get github.com/m3-chat/go-sdk
````

## Quick start

```go
package main

import (
	"log"

	"github.com/m3-chat/go-sdk/client"
	"github.com/m3-chat/go-sdk/types"
)

func main() {
	c := client.NewClient(&types.ClientOptions{
		Stream: true,
	})

	err := c.GetResponse(types.RequestParams{
		Model:   "mistral",
		Content: "Hello, how are you?",
	})

	if err != nil {
		log.Fatal(err)
	}
}
```

## Usage

### Client Creation

```go
c := client.NewClient(&types.ClientOptions{
	Stream: false, // Set to true to stream responses
})
```

### GetResponse

```go
err := c.GetResponse(types.RequestParams{
	Model:   "mistral",
	Content: "Explain quantum computing in simple terms.",
})
```

* Returns an error if the model is invalid or the request fails.
* Streams output to stdout if `Stream` is enabled.

### BatchRequests

```go
messages := []string{
	"Who won the 2022 World Cup?",
	"What is the capital of France?",
	"Tell me a joke.",
}

results, err := c.BatchRequests(messages, types.BatchOptions{
	Model: "dolphin3",
})
```

> Returns an array of responses, one for each message.

## Available Models

The M3 Chat Go SDK internally validates models against this list:

```ts
[
  "llama3:8b",
  "llama2-uncensored",
  "gemma3",
  "gemma",
  "phi3:mini",
  "mistral",
  "gemma:2b",
  "gemma:7b",
  "qwen:7b",
  "qwen2.5-coder",
  "qwen3",
  "deepseek-coder:6.7b",
  "deepseek-v2:16b",
  "dolphin-mistral:7b",
  "dolphin3",
  "starcoder2:7b",
  "magistral",
  "devstral",
];
```

## Development

To build or contribute:

```bash
git clone https://github.com/m3-chat/go-sdk.git
cd go-sdk
go build ./...
```

## Contribution

Contributions, issues, and feature requests are welcome! Please open issues or pull requests on the [GitHub repo](https://github.com/m3-chat/go-sdk).

## License

This project is licensed under the Apache License 2.0 — see the [LICENSE](LICENSE) file for details.

## Contact

Join the [M3 Chat Discussions](https://github.com/orgs/m3-chat/discussions) or open an issue on GitHub for support and questions.

---

Thank you for using **M3 Chat Go SDK** — build powerful AI chat applications with ease!
