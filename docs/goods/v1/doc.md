# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [goods/v1/model.proto](#goods_v1_model-proto)
    - [Class](#goods-v1-Class)
    - [Goods](#goods-v1-Goods)
    - [Serial](#goods-v1-Serial)
  
- [goods/v1/errors.proto](#goods_v1_errors-proto)
    - [Error](#goods-v1-Error)
  
- [goods/v1/goods.proto](#goods_v1_goods-proto)
    - [BatchCreateGoodsRequest](#goods-v1-BatchCreateGoodsRequest)
    - [BatchCreateGoodsResponse](#goods-v1-BatchCreateGoodsResponse)
    - [CreateClassRequest](#goods-v1-CreateClassRequest)
    - [CreateClassResponse](#goods-v1-CreateClassResponse)
    - [CreateSerialRequest](#goods-v1-CreateSerialRequest)
    - [CreateSerialResponse](#goods-v1-CreateSerialResponse)
    - [GetClassRequest](#goods-v1-GetClassRequest)
    - [GetClassResponse](#goods-v1-GetClassResponse)
    - [GetGoodsRequest](#goods-v1-GetGoodsRequest)
    - [GetGoodsResponse](#goods-v1-GetGoodsResponse)
    - [GetOrgOfXRequest](#goods-v1-GetOrgOfXRequest)
    - [GetOrgOfXResponse](#goods-v1-GetOrgOfXResponse)
    - [GetSerialRequest](#goods-v1-GetSerialRequest)
    - [GetSerialResponse](#goods-v1-GetSerialResponse)
    - [ListGoodsClassRequest](#goods-v1-ListGoodsClassRequest)
    - [ListGoodsClassResponse](#goods-v1-ListGoodsClassResponse)
    - [ListGoodsRequest](#goods-v1-ListGoodsRequest)
    - [ListGoodsResponse](#goods-v1-ListGoodsResponse)
    - [ListGoodsSerialRequest](#goods-v1-ListGoodsSerialRequest)
    - [ListGoodsSerialResponse](#goods-v1-ListGoodsSerialResponse)
    - [UpdateGoodsClassRequest](#goods-v1-UpdateGoodsClassRequest)
    - [UpdateGoodsClassResponse](#goods-v1-UpdateGoodsClassResponse)
    - [UpdateGoodsRequest](#goods-v1-UpdateGoodsRequest)
    - [UpdateGoodsResponse](#goods-v1-UpdateGoodsResponse)
    - [UpdateGoodsSerialRequest](#goods-v1-UpdateGoodsSerialRequest)
    - [UpdateGoodsSerialResponse](#goods-v1-UpdateGoodsSerialResponse)
  
    - [GetOrgOfXRequest.X](#goods-v1-GetOrgOfXRequest-X)
  
    - [GoodsService](#goods-v1-GoodsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="goods_v1_model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## goods/v1/model.proto



<a name="goods-v1-Class"></a>

### Class



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| name | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 产品类别名称 |
| des | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 产品类别描述(json) |
| state | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 状态 |
| creator | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 创建者 |
| org_id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 生产企业 |
| tm | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 商标 |
| material_id | [google.protobuf.Int32Value](#google-protobuf-Int32Value) |  | 原料 |






<a name="goods-v1-Goods"></a>

### Goods



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| state | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 状态 |
| creator | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 创建者 |
| serial | [Serial](#goods-v1-Serial) |  | 批次 |






<a name="goods-v1-Serial"></a>

### Serial



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 生产日期 |
| state | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 状态 |
| creator | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 创建者 |
| class | [Class](#goods-v1-Class) |  | 类别 |





 

 

 

 



<a name="goods_v1_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## goods/v1/errors.proto


 


<a name="goods-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| GOODS_NOT_FOUND | 0 | 未找到产品 |
| GOODS_CLASS_NOT_FOUND | 1 | 产品类未找到 |
| GOODS_SERIAL_NOT_FOUND | 2 | 产品批次未找到 |
| GOODS_CLASS_CREATE_FAILED | 3 | 创建产品类别失败 |
| GOODS_SERIAL_CREATE_FAILED | 4 | 创建产品批次失败 |
| GOODS_CREATE_FAILED | 5 | 创建产品失败 |
| PERMISSION_DENY | 6 | 无权限 |


 

 

 



<a name="goods_v1_goods-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## goods/v1/goods.proto



<a name="goods-v1-BatchCreateGoodsRequest"></a>

### BatchCreateGoodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial_id | [int32](#int32) |  | 产品批次id |
| sum | [int32](#int32) |  | 本批次产品数量 |






<a name="goods-v1-BatchCreateGoodsResponse"></a>

### BatchCreateGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [google.protobuf.Int32Value](#google-protobuf-Int32Value) | repeated | 产品编号列表 |






<a name="goods-v1-CreateClassRequest"></a>

### CreateClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_name | [string](#string) |  | 产品类型名称 |
| goods_des | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 产品类型说明(json) |
| material | [int32](#int32) |  | 产品原材料类型 |
| org_id | [int32](#int32) |  | 生产企业编号 |
| tm | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 商标编号(optional) |






<a name="goods-v1-CreateClassResponse"></a>

### CreateClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [google.protobuf.Int32Value](#google-protobuf-Int32Value) |  | 产品类别号 |






<a name="goods-v1-CreateSerialRequest"></a>

### CreateSerialRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 生产日期 |
| class_id | [int32](#int32) |  | 产品种类号 |






<a name="goods-v1-CreateSerialResponse"></a>

### CreateSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial_id | [google.protobuf.Int32Value](#google-protobuf-Int32Value) |  | 产品批次号 |






<a name="goods-v1-GetClassRequest"></a>

### GetClassRequest
获取类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int32](#int32) |  | 产品类别号 |






<a name="goods-v1-GetClassResponse"></a>

### GetClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [Class](#goods-v1-Class) |  | 产品类别号 |






<a name="goods-v1-GetGoodsRequest"></a>

### GetGoodsRequest
获取商品


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int32](#int32) |  |  |






<a name="goods-v1-GetGoodsResponse"></a>

### GetGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| good | [Goods](#goods-v1-Goods) |  |  |






<a name="goods-v1-GetOrgOfXRequest"></a>

### GetOrgOfXRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| x | [GetOrgOfXRequest.X](#goods-v1-GetOrgOfXRequest-X) |  |  |
| id | [int32](#int32) |  |  |






<a name="goods-v1-GetOrgOfXResponse"></a>

### GetOrgOfXResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [google.protobuf.Int32Value](#google-protobuf-Int32Value) |  | 企业id |






<a name="goods-v1-GetSerialRequest"></a>

### GetSerialRequest
获取批次


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial_id | [int32](#int32) |  | 产品种类号 |






<a name="goods-v1-GetSerialResponse"></a>

### GetSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial | [Serial](#goods-v1-Serial) |  | 产品批次号 |






<a name="goods-v1-ListGoodsClassRequest"></a>

### ListGoodsClassRequest
列出产品类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  | 页偏移 |
| limit | [int32](#int32) |  | 页容量 |
| org_id | [int32](#int32) |  | 企业id |






<a name="goods-v1-ListGoodsClassResponse"></a>

### ListGoodsClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_classes | [Class](#goods-v1-Class) | repeated | 产品类型列表 |






<a name="goods-v1-ListGoodsRequest"></a>

### ListGoodsRequest
列出产品


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  | 页偏移 |
| limit | [int32](#int32) |  | 页容量 |
| org_id | [int32](#int32) |  | 企业id |






<a name="goods-v1-ListGoodsResponse"></a>

### ListGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods | [Goods](#goods-v1-Goods) | repeated |  |






<a name="goods-v1-ListGoodsSerialRequest"></a>

### ListGoodsSerialRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  | 页偏移 |
| limit | [int32](#int32) |  | 页容量 |
| org_id | [int32](#int32) |  | 企业id |






<a name="goods-v1-ListGoodsSerialResponse"></a>

### ListGoodsSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial | [Serial](#goods-v1-Serial) | repeated | 产品编号列表 |






<a name="goods-v1-UpdateGoodsClassRequest"></a>

### UpdateGoodsClassRequest
更新产品类型信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [Class](#goods-v1-Class) |  | 待更新信息 |






<a name="goods-v1-UpdateGoodsClassResponse"></a>

### UpdateGoodsClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="goods-v1-UpdateGoodsRequest"></a>

### UpdateGoodsRequest
更新产品


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods | [Goods](#goods-v1-Goods) |  | 待更新商品 |






<a name="goods-v1-UpdateGoodsResponse"></a>

### UpdateGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="goods-v1-UpdateGoodsSerialRequest"></a>

### UpdateGoodsSerialRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial | [Serial](#goods-v1-Serial) |  | 待更新产品批次 |






<a name="goods-v1-UpdateGoodsSerialResponse"></a>

### UpdateGoodsSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |





 


<a name="goods-v1-GetOrgOfXRequest-X"></a>

### GetOrgOfXRequest.X


| Name | Number | Description |
| ---- | ------ | ----------- |
| class | 0 |  |
| serial | 1 |  |
| goods | 2 |  |


 

 


<a name="goods-v1-GoodsService"></a>

### GoodsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateClass | [CreateClassRequest](#goods-v1-CreateClassRequest) | [CreateClassResponse](#goods-v1-CreateClassResponse) | CreateClass 创建产品类型 |
| CreateSerial | [CreateSerialRequest](#goods-v1-CreateSerialRequest) | [CreateSerialResponse](#goods-v1-CreateSerialResponse) | CreateSerial 创建产品批次 |
| BatchCreateGoods | [BatchCreateGoodsRequest](#goods-v1-BatchCreateGoodsRequest) | [BatchCreateGoodsResponse](#goods-v1-BatchCreateGoodsResponse) | CreateGoods 批量创建产品 |
| GetGoods | [GetGoodsRequest](#goods-v1-GetGoodsRequest) | [GetGoodsResponse](#goods-v1-GetGoodsResponse) | GetGoods 根据id获取商品信息 |
| GetClass | [GetClassRequest](#goods-v1-GetClassRequest) | [GetClassResponse](#goods-v1-GetClassResponse) | GetClass 根据id获取商品类型 |
| GetSerial | [GetSerialRequest](#goods-v1-GetSerialRequest) | [GetSerialResponse](#goods-v1-GetSerialResponse) | GetSerial 根据id获取商品批次 |
| ListGoodsClass | [ListGoodsClassRequest](#goods-v1-ListGoodsClassRequest) | [ListGoodsClassResponse](#goods-v1-ListGoodsClassResponse) | ListGoodsClass 列出产品类型 |
| ListGoodsSerial | [ListGoodsSerialRequest](#goods-v1-ListGoodsSerialRequest) | [ListGoodsSerialResponse](#goods-v1-ListGoodsSerialResponse) | ListGoodsSerial 列出产品批次 |
| ListGoods | [ListGoodsRequest](#goods-v1-ListGoodsRequest) | [ListGoodsResponse](#goods-v1-ListGoodsResponse) | ListGoods 列出产品 |
| UpdateGoodsClass | [UpdateGoodsClassRequest](#goods-v1-UpdateGoodsClassRequest) | [UpdateGoodsClassResponse](#goods-v1-UpdateGoodsClassResponse) | 更新产品类型信息 |
| UpdateGoodsSerial | [UpdateGoodsSerialRequest](#goods-v1-UpdateGoodsSerialRequest) | [UpdateGoodsSerialResponse](#goods-v1-UpdateGoodsSerialResponse) | UpdateGoodsSerial 更新产品批次信息 |
| UpdateGoods | [UpdateGoodsRequest](#goods-v1-UpdateGoodsRequest) | [UpdateGoodsResponse](#goods-v1-UpdateGoodsResponse) | UpdateGoods 更新产品 |
| GetOrgOfX | [GetOrgOfXRequest](#goods-v1-GetOrgOfXRequest) | [GetOrgOfXResponse](#goods-v1-GetOrgOfXResponse) | GetOrgOfX 获取产品/类型/批次所属企业 |

 



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

