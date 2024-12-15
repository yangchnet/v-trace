# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [circ/v1/model.proto](#circ_v1_model-proto)
    - [CircRecord](#circ-v1-CircRecord)
  
    - [CircType](#circ-v1-CircType)
    - [RecordStatus](#circ-v1-RecordStatus)
  
- [circ/v1/circ.proto](#circ_v1_circ-proto)
    - [BatchCircRequest](#circ-v1-BatchCircRequest)
    - [BatchCircResponse](#circ-v1-BatchCircResponse)
    - [BatchTransIdRequest](#circ-v1-BatchTransIdRequest)
    - [BatchTransIdResponse](#circ-v1-BatchTransIdResponse)
    - [BatchTransIdResponse.TransIdsEntry](#circ-v1-BatchTransIdResponse-TransIdsEntry)
    - [CreateCircRequest](#circ-v1-CreateCircRequest)
    - [CreateCircResponse](#circ-v1-CreateCircResponse)
    - [GetCircByGoodsIdRequest](#circ-v1-GetCircByGoodsIdRequest)
    - [GetCircByGoodsIdResponse](#circ-v1-GetCircByGoodsIdResponse)
    - [GetCircByTransIdRequest](#circ-v1-GetCircByTransIdRequest)
    - [GetCircRequest](#circ-v1-GetCircRequest)
    - [TransIdRequest](#circ-v1-TransIdRequest)
    - [TransIdResponse](#circ-v1-TransIdResponse)
    - [UpdateCircStatusRequest](#circ-v1-UpdateCircStatusRequest)
    - [UpdateCircStatusResponse](#circ-v1-UpdateCircStatusResponse)
  
    - [CircService](#circ-v1-CircService)
  
- [circ/v1/errors.proto](#circ_v1_errors-proto)
    - [Error](#circ-v1-Error)
  
- [Scalar Value Types](#scalar-value-types)



<a name="circ_v1_model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## circ/v1/model.proto



<a name="circ-v1-CircRecord"></a>

### CircRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| object_id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 流转主体id（一般为商品id） |
| circ_type | [CircType](#circ-v1-CircType) |  | 流转类型 |
| Operator | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 流转操作者 |
| from | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | from |
| to | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | to |
| from_value | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | formValue |
| tx_hash | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | txHash |
| trans_id | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | transId |
| times | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | times |
| status | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | status |





 


<a name="circ-v1-CircType"></a>

### CircType


| Name | Number | Description |
| ---- | ------ | ----------- |
| produce | 0 | 生产 |
| process | 1 | 一般流程 |
| transfer | 3 | 转交 |
| exam | 4 | 核验 |



<a name="circ-v1-RecordStatus"></a>

### RecordStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| pending | 0 | 进行中 |
| failed | 1 | 已失败 |
| success | 2 | 已成功 |


 

 

 



<a name="circ_v1_circ-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## circ/v1/circ.proto



<a name="circ-v1-BatchCircRequest"></a>

### BatchCircRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans_ids | [string](#string) | repeated | 流转号 |
| circ_type | [CircType](#circ-v1-CircType) |  | 流转类型 |
| operator | [string](#string) |  | 流转参与者 |
| from | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | from, 当流转类型为&#34;Produce&#34;时为空 |
| to | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | to, 当流转类型为Transfer时不为空 |
| form_info | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 表单json数据 |






<a name="circ-v1-BatchCircResponse"></a>

### BatchCircResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| circ_ids | [int32](#int32) | repeated |  |






<a name="circ-v1-BatchTransIdRequest"></a>

### BatchTransIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_ids | [int32](#int32) | repeated |  |






<a name="circ-v1-BatchTransIdResponse"></a>

### BatchTransIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transIds | [BatchTransIdResponse.TransIdsEntry](#circ-v1-BatchTransIdResponse-TransIdsEntry) | repeated |  |






<a name="circ-v1-BatchTransIdResponse-TransIdsEntry"></a>

### BatchTransIdResponse.TransIdsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [string](#string) |  |  |






<a name="circ-v1-CreateCircRequest"></a>

### CreateCircRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans_id | [string](#string) |  | transId 用于标记此次流转 其格式为：[流转次数]-[产品id]-[流转时间] |
| circ_type | [CircType](#circ-v1-CircType) |  | 流转类型 |
| operator | [string](#string) |  | 流转执行者 |
| from | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 实体来源用户 流转类型为&#34;produce&#34;时，此字段为空 流转类型为&#34;process&#34;、&#34;exam&#34;或&#34;transfer&#34;时，此字段为操作者本身用户名 |
| to | [string](#string) |  | 实体目标用户 流转类型为&#34;produce&#34;、&#34;process&#34;, &#34;exam&#34;时，此字段为操作者 流转类型为&#34;transfer&#34;时，此字段为要转交的用户的用户名 |
| form_info | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 表单json数据 |






<a name="circ-v1-CreateCircResponse"></a>

### CreateCircResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| circ_id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 流转历史记录编号 |






<a name="circ-v1-GetCircByGoodsIdRequest"></a>

### GetCircByGoodsIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int64](#int64) |  | 商品id |






<a name="circ-v1-GetCircByGoodsIdResponse"></a>

### GetCircByGoodsIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| records | [CircRecord](#circ-v1-CircRecord) | repeated | 流转历史记录列表 |






<a name="circ-v1-GetCircByTransIdRequest"></a>

### GetCircByTransIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans_id | [string](#string) |  | transId |






<a name="circ-v1-GetCircRequest"></a>

### GetCircRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| circ_id | [int64](#int64) |  | 流转历史记录编号 |






<a name="circ-v1-TransIdRequest"></a>

### TransIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int32](#int32) |  | 商品id |






<a name="circ-v1-TransIdResponse"></a>

### TransIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans_id | [string](#string) |  |  |






<a name="circ-v1-UpdateCircStatusRequest"></a>

### UpdateCircStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trans_id | [string](#string) |  |  |
| status | [RecordStatus](#circ-v1-RecordStatus) |  |  |






<a name="circ-v1-UpdateCircStatusResponse"></a>

### UpdateCircStatusResponse






 

 

 


<a name="circ-v1-CircService"></a>

### CircService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCirc | [CreateCircRequest](#circ-v1-CreateCircRequest) | [CreateCircResponse](#circ-v1-CreateCircResponse) | CreateCirc 新增流转记录 |
| BatchCirc | [BatchCircRequest](#circ-v1-BatchCircRequest) | [BatchCircResponse](#circ-v1-BatchCircResponse) | 批量流转 |
| GetCirc | [GetCircRequest](#circ-v1-GetCircRequest) | [CircRecord](#circ-v1-CircRecord) | GetCirc 根据流转id获取流转记录 |
| GetCircByTransId | [GetCircByTransIdRequest](#circ-v1-GetCircByTransIdRequest) | [CircRecord](#circ-v1-CircRecord) | 根据transId获取流转记录 |
| GetCircByGoodsId | [GetCircByGoodsIdRequest](#circ-v1-GetCircByGoodsIdRequest) | [GetCircByGoodsIdResponse](#circ-v1-GetCircByGoodsIdResponse) | 根据商品id获取流转记录 |
| TransId | [TransIdRequest](#circ-v1-TransIdRequest) | [TransIdResponse](#circ-v1-TransIdResponse) | 为商品生成transId |
| BatchTransId | [BatchTransIdRequest](#circ-v1-BatchTransIdRequest) | [BatchTransIdResponse](#circ-v1-BatchTransIdResponse) | 批量为商品生成transId |
| UpdateCircStatus | [UpdateCircStatusRequest](#circ-v1-UpdateCircStatusRequest) | [UpdateCircStatusResponse](#circ-v1-UpdateCircStatusResponse) | 更新流转记录状态 |

 



<a name="circ_v1_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## circ/v1/errors.proto


 


<a name="circ-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| CREATE_CIRC_FAILED | 0 | 创建流转记录失败 |
| CIRC_NOT_FOUND | 1 | 流转记录不存在 |
| INVALID_PARAMS | 2 | 参数错误 |
| CIRC_ALREADY_EXIST | 3 | 流转记录已存在 |
| NOT_OWNER | 4 | 当前操作者非产品所有者 |


 

 

 



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

