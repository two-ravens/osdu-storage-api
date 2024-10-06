<!-- Start SDK Example Usage [usage] -->
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