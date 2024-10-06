# openapi

Developer-friendly & type-safe Go SDK specifically catered to leverage *openapi* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/dveracity/dveracity). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Storage Service: Storage service which handles the metadata ingestion in the OSDU data platform
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get openapi
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Records.CreateOrUpdateRecords(ctx, "<id>", []components.Record{
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
	}, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateUpdateRecordsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Health](docs/sdks/health/README.md)

* [HealthMessage](docs/sdks/health/README.md#healthmessage) - Health Message

### [Info](docs/sdks/info/README.md)

* [Info](docs/sdks/info/README.md#info) - Version info

### [Query](docs/sdks/query/README.md)

* [GetAllRecords](docs/sdks/query/README.md#getallrecords) - Get all record from kind
* [GetRecords](docs/sdks/query/README.md#getrecords) - Fetch records
* [FetchRecords](docs/sdks/query/README.md#fetchrecords) - Fetch multiple records

### [Records](docs/sdks/records/README.md)

* [CreateOrUpdateRecords](docs/sdks/records/README.md#createorupdaterecords) - Create or Update Records
* [UpdateRecordsMetadata1](docs/sdks/records/README.md#updaterecordsmetadata1) - Modify record data and/or metadata attributes using patch operations
* [DeleteRecord](docs/sdks/records/README.md#deleterecord) - Delete Record
* [BulkDeleteRecords](docs/sdks/records/README.md#bulkdeleterecords) - Soft delete of multiple records
* [GetLatestRecordVersion](docs/sdks/records/README.md#getlatestrecordversion) - Get latest record version data
* [PurgeRecord](docs/sdks/records/README.md#purgerecord) - Purge Record
* [GetSpecificRecordVersion](docs/sdks/records/README.md#getspecificrecordversion) - Get Specific record
* [GetRecordVersions](docs/sdks/records/README.md#getrecordversions) - Get record versions


</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"log"
	"models/operations"
	"openapi"
	"openapi/models/components"
	"openapi/retry"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Records.CreateOrUpdateRecords(ctx, "<id>", []components.Record{
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
	}, nil, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateUpdateRecordsResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
	"openapi/retry"
)

func main() {
	s := openapi.New(
		openapi.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Records.CreateOrUpdateRecords(ctx, "<id>", []components.Record{
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
	}, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateUpdateRecordsResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateOrUpdateRecords` function may return the following errors:

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

### Example

```go
package main

import (
	"context"
	"errors"
	"log"
	"openapi"
	"openapi/models/components"
	"openapi/models/sdkerrors"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Records.CreateOrUpdateRecords(ctx, "<id>", []components.Record{
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
	}, nil, nil)
	if err != nil {

		var e *sdkerrors.AppError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https:///api/storage/v2/` | None |

#### Example

```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	s := openapi.New(
		openapi.WithServerIndex(0),
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Records.CreateOrUpdateRecords(ctx, "<id>", []components.Record{
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
	}, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateUpdateRecordsResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	s := openapi.New(
		openapi.WithServerURL("https:///api/storage/v2/"),
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Records.CreateOrUpdateRecords(ctx, "<id>", []components.Record{
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
	}, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateUpdateRecordsResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name            | Type            | Scheme          |
| --------------- | --------------- | --------------- |
| `Authorization` | http            | HTTP Bearer     |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Records.CreateOrUpdateRecords(ctx, "<id>", []components.Record{
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
		components.Record{
			ACL: components.ACL{},
			Data: map[string]components.Data{
				"key": components.Data{},
			},
		},
	}, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateUpdateRecordsResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go)
