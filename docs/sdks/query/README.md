# Query
(*Query*)

## Overview

Querying Records operations

### Available Operations

* [GetAllRecords](#getallrecords) - Get all record from kind
* [GetRecords](#getrecords) - Fetch records
* [FetchRecords](#fetchrecords) - Fetch multiple records

## GetAllRecords

The API returns a list of all record ids which belong to the specified kind.
Allowed roles: `service.storage.admin`.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Query.GetAllRecords(ctx, operations.GetAllRecordsRequest{
        Limit: openapi.Int(10),
        Kind: "tenant1:public:well:1.0.2",
        DataPartitionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreQueryResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAllRecordsRequest](../../models/operations/getallrecordsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetAllRecordsResponse](../../models/operations/getallrecordsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## GetRecords

The API fetches multiple records at once.
Allowed roles: `service.storage.viewer`,`service.storage.creator` and `service.storage.admin`.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/components"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Query.GetRecords(ctx, "<id>", components.MultiRecordIds{
        Records: []string{
            "<value>",
            "<value>",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MultiRecordInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `dataPartitionID`                                                      | *string*                                                               | :heavy_check_mark:                                                     | Tenant Id                                                              |
| `multiRecordIds`                                                       | [components.MultiRecordIds](../../models/components/multirecordids.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `xCollaboration`                                                       | **string*                                                              | :heavy_minus_sign:                                                     | x-collaboration                                                        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.GetRecordsResponse](../../models/operations/getrecordsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## FetchRecords

The API fetches multiple records at once in the specific {data-partition-id}.The value of {frame-of-reference} indicates whether normalization is applied.
Required roles: `users.datalake.viewers` or `users.datalake.editors` or `users.datalake.admins`.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/components"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Query.FetchRecords(ctx, "<id>", "units=SI;crs=wgs84;elevation=msl;azimuth=true north;dates=utc;", components.MultiRecordRequest{
        Records: []string{
            "<value>",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MultiRecordResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                       | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     | Example                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                           | :heavy_check_mark:                                                                                                                              | The context to use for the request.                                                                                                             |                                                                                                                                                 |
| `dataPartitionID`                                                                                                                               | *string*                                                                                                                                        | :heavy_check_mark:                                                                                                                              | Tenant Id                                                                                                                                       |                                                                                                                                                 |
| `frameOfReference`                                                                                                                              | *string*                                                                                                                                        | :heavy_check_mark:                                                                                                                              | This value indicates whether normalization applies, should be either `none` or `units=SI;crs=wgs84;elevation=msl;azimuth=true north;dates=utc;` | units=SI;crs=wgs84;elevation=msl;azimuth=true north;dates=utc;                                                                                  |
| `multiRecordRequest`                                                                                                                            | [components.MultiRecordRequest](../../models/components/multirecordrequest.md)                                                                  | :heavy_check_mark:                                                                                                                              | N/A                                                                                                                                             |                                                                                                                                                 |
| `xCollaboration`                                                                                                                                | **string*                                                                                                                                       | :heavy_minus_sign:                                                                                                                              | x-collaboration                                                                                                                                 |                                                                                                                                                 |
| `opts`                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                        | :heavy_minus_sign:                                                                                                                              | The options for this request.                                                                                                                   |                                                                                                                                                 |

### Response

**[*operations.FetchRecordsResponse](../../models/operations/fetchrecordsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |