# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [trans/v1/errors.proto](#trans_v1_errors-proto)
    - [Error](#trans-v1-Error)
  
- [trans/v1/model.proto](#trans_v1_model-proto)
    - [Identity](#trans-v1-Identity)
    - [TransRecord](#trans-v1-TransRecord)
  
    - [TransStatus](#trans-v1-TransStatus)
  
- [trans/v1/trans.proto](#trans_v1_trans-proto)
    - [AddProcessRequest](#trans-v1-AddProcessRequest)
    - [AddProcessResponse](#trans-v1-AddProcessResponse)
    - [ApproveRequest](#trans-v1-ApproveRequest)
    - [ApproveResponse](#trans-v1-ApproveResponse)
    - [BatchBurnRequest](#trans-v1-BatchBurnRequest)
    - [BatchBurnResponse](#trans-v1-BatchBurnResponse)
    - [BatchMintRequest](#trans-v1-BatchMintRequest)
    - [BatchMintResponse](#trans-v1-BatchMintResponse)
    - [BatchProcessRequest](#trans-v1-BatchProcessRequest)
    - [BatchProcessResponse](#trans-v1-BatchProcessResponse)
    - [BatchTransformRequest](#trans-v1-BatchTransformRequest)
    - [BatchTransformResponse](#trans-v1-BatchTransformResponse)
    - [BurnRequest](#trans-v1-BurnRequest)
    - [BurnResponse](#trans-v1-BurnResponse)
    - [GetTransRequest](#trans-v1-GetTransRequest)
    - [GetTransResponse](#trans-v1-GetTransResponse)
    - [GrantRoleRequest](#trans-v1-GrantRoleRequest)
    - [GrantRoleResponse](#trans-v1-GrantRoleResponse)
    - [ListTransByObjRequest](#trans-v1-ListTransByObjRequest)
    - [ListTransByObjResponse](#trans-v1-ListTransByObjResponse)
    - [MintRequest](#trans-v1-MintRequest)
    - [MintResponse](#trans-v1-MintResponse)
    - [NameRequest](#trans-v1-NameRequest)
    - [NameResponse](#trans-v1-NameResponse)
    - [SetApproveForAllRequest](#trans-v1-SetApproveForAllRequest)
    - [SetApproveForAllResponse](#trans-v1-SetApproveForAllResponse)
    - [SymbolRequest](#trans-v1-SymbolRequest)
    - [SymbolResponse](#trans-v1-SymbolResponse)
    - [TransferRequest](#trans-v1-TransferRequest)
    - [TransferResponse](#trans-v1-TransferResponse)
    - [UpdateTransRequest](#trans-v1-UpdateTransRequest)
    - [UpdateTransResponse](#trans-v1-UpdateTransResponse)
  
    - [OperationResult](#trans-v1-OperationResult)
  
    - [TransService](#trans-v1-TransService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="trans_v1_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## trans/v1/errors.proto


 


<a name="trans-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| RECORD_NOT_FOUND | 0 | 未找到交易记录 |
| DUPLICATE_ERR | 1 | 记录重复 |
| TRANS_ALREADY_EXIST | 2 | 已存在的交易 |


 

 

 



<a name="trans_v1_model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## trans/v1/model.proto



<a name="trans-v1-Identity"></a>

### Identity
用于标识上链操作者身份


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cert | [bytes](#bytes) |  | 用户证书sign |
| key | [bytes](#bytes) |  | 用户私钥sign |
| username | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 用户名 |
| tls_cert | [bytes](#bytes) |  | tls证书 |
| tls_key | [bytes](#bytes) |  | tls私钥 |






<a name="trans-v1-TransRecord"></a>

### TransRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| trans_id | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | transID |
| sender | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 交易发送者 |
| contract | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 调用合约名 |
| method | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 调用方法名 |
| params | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 参数 |
| status | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 状态 |
| tx_hash | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 交易hash |
| tx_params_hash | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 交易参数hash |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更改时间 |





 


<a name="trans-v1-TransStatus"></a>

### TransStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| waiting | 0 |  |
| failed | 1 |  |
| success | 2 |  |


 

 

 



<a name="trans_v1_trans-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## trans/v1/trans.proto



<a name="trans-v1-AddProcessRequest"></a>

### AddProcessRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  |  |
| tokenId | [int64](#int64) |  |  |
| digest | [string](#string) |  |  |
| trans_id | [string](#string) |  |  |






<a name="trans-v1-AddProcessResponse"></a>

### AddProcessResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  |  |






<a name="trans-v1-ApproveRequest"></a>

### ApproveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  | 操作者 |
| to | [Identity](#trans-v1-Identity) |  | to |
| token_id | [int64](#int64) |  | tokenId |
| callback_url | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | callback url which is a golang url.URL type |






<a name="trans-v1-ApproveResponse"></a>

### ApproveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  |  |






<a name="trans-v1-BatchBurnRequest"></a>

### BatchBurnRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  |  |
| token_ids | [int64](#int64) | repeated |  |






<a name="trans-v1-BatchBurnResponse"></a>

### BatchBurnResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  |  |






<a name="trans-v1-BatchMintRequest"></a>

### BatchMintRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  |  |
| to | [Identity](#trans-v1-Identity) |  |  |
| token_ids | [int64](#int64) | repeated |  |
| transIds | [string](#string) | repeated |  |
| digest | [string](#string) |  |  |






<a name="trans-v1-BatchMintResponse"></a>

### BatchMintResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  |  |






<a name="trans-v1-BatchProcessRequest"></a>

### BatchProcessRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  |  |
| token_ids | [int64](#int64) | repeated |  |
| transIds | [string](#string) | repeated |  |
| digest | [string](#string) |  |  |






<a name="trans-v1-BatchProcessResponse"></a>

### BatchProcessResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  |  |






<a name="trans-v1-BatchTransformRequest"></a>

### BatchTransformRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  |  |
| from | [Identity](#trans-v1-Identity) |  |  |
| to | [Identity](#trans-v1-Identity) |  |  |
| token_ids | [int64](#int64) | repeated |  |
| transIds | [string](#string) | repeated |  |
| digest | [string](#string) |  |  |






<a name="trans-v1-BatchTransformResponse"></a>

### BatchTransformResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  |  |






<a name="trans-v1-BurnRequest"></a>

### BurnRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  | 身份标识 |
| token_id | [int64](#int64) |  | 要销毁的tokenId |
| trans_id | [string](#string) |  | transId |
| callback_url | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | callback url which is a golang url.URL type |






<a name="trans-v1-BurnResponse"></a>

### BurnResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  | 操作结果 |






<a name="trans-v1-GetTransRequest"></a>

### GetTransRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans_id | [string](#string) |  |  |






<a name="trans-v1-GetTransResponse"></a>

### GetTransResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans | [TransRecord](#trans-v1-TransRecord) |  |  |






<a name="trans-v1-GrantRoleRequest"></a>

### GrantRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| granter | [Identity](#trans-v1-Identity) |  | 授予角色者 |
| account | [Identity](#trans-v1-Identity) |  | 被授予角色者 |
| trans_id | [string](#string) |  | transId |
| callback_url | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | callback url which is a golang url.URL type |






<a name="trans-v1-GrantRoleResponse"></a>

### GrantRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  | 操作结果 |






<a name="trans-v1-ListTransByObjRequest"></a>

### ListTransByObjRequest







<a name="trans-v1-ListTransByObjResponse"></a>

### ListTransByObjResponse







<a name="trans-v1-MintRequest"></a>

### MintRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  | 身份标识 |
| to | [Identity](#trans-v1-Identity) |  | to |
| token_id | [int64](#int64) |  | tokenId, 需保证唯一性 |
| trans_id | [string](#string) |  | transId |
| digest | [string](#string) |  | digest |






<a name="trans-v1-MintResponse"></a>

### MintResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  | 操作结果 |






<a name="trans-v1-NameRequest"></a>

### NameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  | 操作者身份 |
| trans_id | [string](#string) |  | transId |






<a name="trans-v1-NameResponse"></a>

### NameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  | 操作结果 |






<a name="trans-v1-SetApproveForAllRequest"></a>

### SetApproveForAllRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  | 授权人 |
| operator | [Identity](#trans-v1-Identity) |  | 被授权人operator |
| approved | [bool](#bool) |  | 是否授权 |
| trans_id | [string](#string) |  | transId |
| callback_url | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | callback url which is a golang url.URL type |






<a name="trans-v1-SetApproveForAllResponse"></a>

### SetApproveForAllResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  | 操作结果 |






<a name="trans-v1-SymbolRequest"></a>

### SymbolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  | 身份标识 |
| trans_id | [string](#string) |  | transId |






<a name="trans-v1-SymbolResponse"></a>

### SymbolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  | 操作结果 |






<a name="trans-v1-TransferRequest"></a>

### TransferRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| caller | [Identity](#trans-v1-Identity) |  | 操作者 |
| from | [Identity](#trans-v1-Identity) |  | 转账者 |
| to | [Identity](#trans-v1-Identity) |  | 接收转账者 |
| token_id | [int64](#int64) |  | tokenId |
| trans_id | [string](#string) |  | transId |
| digest | [string](#string) |  | digest |






<a name="trans-v1-TransferResponse"></a>

### TransferResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  | 操作结果 |






<a name="trans-v1-UpdateTransRequest"></a>

### UpdateTransRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans_id | [string](#string) |  |  |
| tx_hash | [string](#string) |  |  |
| success | [bool](#bool) |  |  |






<a name="trans-v1-UpdateTransResponse"></a>

### UpdateTransResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#trans-v1-OperationResult) |  |  |





 


<a name="trans-v1-OperationResult"></a>

### OperationResult


| Name | Number | Description |
| ---- | ------ | ----------- |
| SUCCESS | 0 |  |
| FAIL | 1 |  |


 

 


<a name="trans-v1-TransService"></a>

### TransService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Name | [NameRequest](#trans-v1-NameRequest) | [NameResponse](#trans-v1-NameResponse) |  |
| Symbol | [SymbolRequest](#trans-v1-SymbolRequest) | [SymbolResponse](#trans-v1-SymbolResponse) |  |
| Mint | [MintRequest](#trans-v1-MintRequest) | [MintResponse](#trans-v1-MintResponse) |  |
| Burn | [BurnRequest](#trans-v1-BurnRequest) | [BurnResponse](#trans-v1-BurnResponse) |  |
| Transfer | [TransferRequest](#trans-v1-TransferRequest) | [TransferResponse](#trans-v1-TransferResponse) |  |
| Approve | [ApproveRequest](#trans-v1-ApproveRequest) | [ApproveResponse](#trans-v1-ApproveResponse) |  |
| SetApproveForAll | [SetApproveForAllRequest](#trans-v1-SetApproveForAllRequest) | [SetApproveForAllResponse](#trans-v1-SetApproveForAllResponse) |  |
| GrantRole | [GrantRoleRequest](#trans-v1-GrantRoleRequest) | [GrantRoleResponse](#trans-v1-GrantRoleResponse) |  |
| GetTrans | [GetTransRequest](#trans-v1-GetTransRequest) | [GetTransResponse](#trans-v1-GetTransResponse) |  |
| UpdateTrans | [UpdateTransRequest](#trans-v1-UpdateTransRequest) | [UpdateTransResponse](#trans-v1-UpdateTransResponse) |  |
| AddProcess | [AddProcessRequest](#trans-v1-AddProcessRequest) | [AddProcessResponse](#trans-v1-AddProcessResponse) |  |
| ListTransByObj | [ListTransByObjRequest](#trans-v1-ListTransByObjRequest) | [ListTransByObjResponse](#trans-v1-ListTransByObjResponse) |  |
| BatchMint | [BatchMintRequest](#trans-v1-BatchMintRequest) | [BatchMintResponse](#trans-v1-BatchMintResponse) |  |
| BatchBurn | [BatchBurnRequest](#trans-v1-BatchBurnRequest) | [BatchBurnResponse](#trans-v1-BatchBurnResponse) |  |
| BatchProcess | [BatchProcessRequest](#trans-v1-BatchProcessRequest) | [BatchProcessResponse](#trans-v1-BatchProcessResponse) |  |
| BatchTransform | [BatchTransformRequest](#trans-v1-BatchTransformRequest) | [BatchTransformResponse](#trans-v1-BatchTransformResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

