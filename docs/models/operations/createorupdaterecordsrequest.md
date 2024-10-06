# CreateOrUpdateRecordsRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `XCollaboration`                                           | **string*                                                  | :heavy_minus_sign:                                         | x-collaboration                                            |
| `Skipdupes`                                                | **bool*                                                    | :heavy_minus_sign:                                         | Skip duplicates when updating records with the same value. |
| `DataPartitionID`                                          | *string*                                                   | :heavy_check_mark:                                         | Tenant Id                                                  |
| `RequestBody`                                              | [][components.Record](../../models/components/record.md)   | :heavy_check_mark:                                         | N/A                                                        |