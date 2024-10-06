# Records
(*Records*)

## Overview

Records management operations

### Available Operations

* [CreateOrUpdateRecords](#createorupdaterecords) - Create or Update Records
* [UpdateRecordsMetadata1](#updaterecordsmetadata1) - Modify record data and/or metadata attributes using patch operations
* [DeleteRecord](#deleterecord) - Delete Record
* [BulkDeleteRecords](#bulkdeleterecords) - Soft delete of multiple records
* [GetLatestRecordVersion](#getlatestrecordversion) - Get latest record version data
* [PurgeRecord](#purgerecord) - Purge Record
* [GetSpecificRecordVersion](#getspecificrecordversion) - Get Specific record
* [GetRecordVersions](#getrecordversions) - Get record versions

## CreateOrUpdateRecords

The API represents the main injection mechanism into the Data Ecosystem. 
It allows records creation and/or update.When no record id is provided or when the provided id is not already present in the Data Ecosystem then a new record is created. 
 If the id is related to an existing record in the Data Ecosystem then an update operation takes place and a new version of the record is created.

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

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `dataPartitionID`                                          | *string*                                                   | :heavy_check_mark:                                         | Tenant Id                                                  |
| `requestBody`                                              | [][components.Record](../../models/components/record.md)   | :heavy_check_mark:                                         | N/A                                                        |
| `xCollaboration`                                           | **string*                                                  | :heavy_minus_sign:                                         | x-collaboration                                            |
| `skipdupes`                                                | **bool*                                                    | :heavy_minus_sign:                                         | Skip duplicates when updating records with the same value. |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.CreateOrUpdateRecordsResponse](../../models/operations/createorupdaterecordsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## UpdateRecordsMetadata1

The API represents the patch update mechanism for records. It allows updating multiple records in one request. The API supports metadata update only (Legal Tags, ACLs and Tags) if the request body media type is `application/json`. The API supports metadata and data update (Legal Tags, ACLs, Tags, Ancestry, Kind, Meta and Data) if the request body media type is `application/json-patch+json`. Please choose the appropriate media type from the Request body dropdown. The currently supported operations are replace, add, and remove. 
Required roles: `users.datalake.editors` or `users.datalake.admins`.

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
    res, err := s.Records.UpdateRecordsMetadata1(ctx, "<id>", components.RecordBulkUpdateParam{
        Query: components.RecordQuery{
            Ids: []string{
                "<value>",
            },
        },
        Ops: []components.PatchOperation{
            components.PatchOperation{
                Value: []string{
                    "<value>",
                    "<value>",
                    "<value>",
                },
            },
            components.PatchOperation{
                Value: []string{
                    "<value>",
                    "<value>",
                    "<value>",
                },
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `dataPartitionID`                                                                    | *string*                                                                             | :heavy_check_mark:                                                                   | Tenant Id                                                                            |
| `recordBulkUpdateParam`                                                              | [components.RecordBulkUpdateParam](../../models/components/recordbulkupdateparam.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `xCollaboration`                                                                     | **string*                                                                            | :heavy_minus_sign:                                                                   | x-collaboration                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateRecordsMetadata1Response](../../models/operations/updaterecordsmetadata1response.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## DeleteRecord

The API performs a logical deletion of the record using recordId. This operation can be reverted later. 
Allowed roles: `service.storage.creator` and `service.storage.admin` who is the OWNER of the record.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Records.DeleteRecord(ctx, "tenant1:well:123456789", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Record id                                                | tenant1:well:123456789                                   |
| `dataPartitionID`                                        | *string*                                                 | :heavy_check_mark:                                       | Tenant Id                                                |                                                          |
| `xCollaboration`                                         | **string*                                                | :heavy_minus_sign:                                       | x-collaboration                                          |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteRecordResponse](../../models/operations/deleterecordresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## BulkDeleteRecords

The API performs a soft deletion of the given list of records. 
Required roles: `users.datalake.editors` or `users.datalake.admins` who is the OWNER of the record.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Records.BulkDeleteRecords(ctx, "<id>", []string{
        "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `dataPartitionID`                                        | *string*                                                 | :heavy_check_mark:                                       | Tenant Id                                                |
| `requestBody`                                            | []*string*                                               | :heavy_check_mark:                                       | N/A                                                      |
| `xCollaboration`                                         | **string*                                                | :heavy_minus_sign:                                       | x-collaboration                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.BulkDeleteRecordsResponse](../../models/operations/bulkdeleterecordsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLatestRecordVersion

This API returns the latest version of the given record.
Allowed roles: `service.storage.viewer`, `service.storage.creator` and `service.storage.admin`.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Records.GetLatestRecordVersion(ctx, "tenant1:well:123456789", "<id>", nil, []string{
        "data.wellName",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `id`                                                                                                    | *string*                                                                                                | :heavy_check_mark:                                                                                      | Record id                                                                                               | tenant1:well:123456789                                                                                  |
| `dataPartitionID`                                                                                       | *string*                                                                                                | :heavy_check_mark:                                                                                      | Tenant Id                                                                                               |                                                                                                         |
| `xCollaboration`                                                                                        | **string*                                                                                               | :heavy_minus_sign:                                                                                      | x-collaboration                                                                                         |                                                                                                         |
| `attribute`                                                                                             | []*string*                                                                                              | :heavy_minus_sign:                                                                                      | Filter attributes to restrict the returned fields of the record.  Usage: data.{record-data-field-name}. | data.wellName                                                                                           |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.GetLatestRecordVersionResponse](../../models/operations/getlatestrecordversionresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## PurgeRecord

The API performs the physical deletion of the given record and all of its versions.
 This operation cannot be undone. 
Allowed roles: `service.storage.admin` who is the OWNER of the record.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Records.PurgeRecord(ctx, "tenant1:well:123456789", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Record id                                                | tenant1:well:123456789                                   |
| `dataPartitionID`                                        | *string*                                                 | :heavy_check_mark:                                       | Tenant Id                                                |                                                          |
| `xCollaboration`                                         | **string*                                                | :heavy_minus_sign:                                       | x-collaboration                                          |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PurgeRecordResponse](../../models/operations/purgerecordresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.AppError           | 401, 403, 404, 500, 502, 503 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetSpecificRecordVersion

The API retrieves the specific version of the given record. 
Allowed roles: `service.storage.viewer`, `service.storage.creator` and `service.storage.admin`.

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
    res, err := s.Records.GetSpecificRecordVersion(ctx, operations.GetSpecificRecordVersionRequest{
        ID: "tenant1:well:123456789",
        Version: 123456789,
        Attribute: []string{
            "data.wellName",
        },
        DataPartitionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetSpecificRecordVersionRequest](../../models/operations/getspecificrecordversionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetSpecificRecordVersionResponse](../../models/operations/getspecificrecordversionresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## GetRecordVersions

The API returns a list containing all versions for the given record id. 
Allowed roles: `service.storage.viewer`, `service.storage.creator` and `service.storage.admin`.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Records.GetRecordVersions(ctx, "tenant1:well:123456789", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RecordVersions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Record id                                                | tenant1:well:123456789                                   |
| `dataPartitionID`                                        | *string*                                                 | :heavy_check_mark:                                       | Tenant Id                                                |                                                          |
| `xCollaboration`                                         | **string*                                                | :heavy_minus_sign:                                       | x-collaboration                                          |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetRecordVersionsResponse](../../models/operations/getrecordversionsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.AppError                | 400, 401, 403, 404, 500, 502, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |