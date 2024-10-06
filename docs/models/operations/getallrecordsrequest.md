# GetAllRecordsRequest


## Fields

| Field                     | Type                      | Required                  | Description               | Example                   |
| ------------------------- | ------------------------- | ------------------------- | ------------------------- | ------------------------- |
| `XCollaboration`          | **string*                 | :heavy_minus_sign:        | x-collaboration           |                           |
| `Cursor`                  | **string*                 | :heavy_minus_sign:        | Cursor                    |                           |
| `Limit`                   | **int*                    | :heavy_minus_sign:        | Page Size                 | 10                        |
| `Kind`                    | *string*                  | :heavy_check_mark:        | Filter Kind               | tenant1:public:well:1.0.2 |
| `DataPartitionID`         | *string*                  | :heavy_check_mark:        | Tenant Id                 |                           |