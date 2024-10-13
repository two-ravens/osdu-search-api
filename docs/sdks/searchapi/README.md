# SearchAPI
(*SearchAPI*)

## Overview

Service endpoints to search data in datalake

### Available Operations

* [QueryWithCursor](#querywithcursor) - Queries the index using cursor for the input request criteria.
* [QueryRecords](#queryrecords) - Queries the index for the input request criteria.

## QueryWithCursor

The API supports full text search on string fields, range queries on date, numeric or string fields, along with geo-spatial search. 
Required roles: `users.datalake.viewers` or `users.datalake.editors` or `users.datalake.admins` or `users.datalake.ops`. In addition, users must be a member of data groups to access the data. 
It can be used to retrieve large numbers of results (or even all results) from a single search request, in much the same way as you would use a cursor on a traditional database.

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
    res, err := s.SearchAPI.QueryWithCursor(ctx, "<id>", components.CursorQueryRequest{
        Kind: components.Kind{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `dataPartitionID`                                                              | *string*                                                                       | :heavy_check_mark:                                                             | Tenant Id                                                                      |
| `cursorQueryRequest`                                                           | [components.CursorQueryRequest](../../models/components/cursorqueryrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.QueryWithCursorResponse](../../models/operations/querywithcursorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## QueryRecords

The API supports full text search on string fields, range queries on date, numeric or string fields, along with geo-spatial search. 
 Required roles: `users.datalake.viewers` or  `users.datalake.editors` or `users.datalake.admins` or `users.datalake.ops`. In addition, users must be a member of data groups to access the data.

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
    res, err := s.SearchAPI.QueryRecords(ctx, "<id>", components.QueryRequest{
        Kind: components.QueryRequestKind{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `dataPartitionID`                                                  | *string*                                                           | :heavy_check_mark:                                                 | Tenant Id                                                          |
| `queryRequest`                                                     | [components.QueryRequest](../../models/components/queryrequest.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.QueryRecordsResponse](../../models/operations/queryrecordsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |