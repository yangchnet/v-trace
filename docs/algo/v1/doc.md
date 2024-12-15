# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [algo/v1/model.proto](#algo_v1_model-proto)
    - [Material](#algo-v1-Material)
    - [Model](#algo-v1-Model)
    - [Model.MetadataEntry](#algo-v1-Model-MetadataEntry)
  
- [algo/v1/algo.proto](#algo_v1_algo-proto)
    - [ListMaterialsRequest](#algo-v1-ListMaterialsRequest)
    - [ListMaterialsResponse](#algo-v1-ListMaterialsResponse)
    - [ListModelsRequest](#algo-v1-ListModelsRequest)
    - [ListModelsResponse](#algo-v1-ListModelsResponse)
    - [PredictRequest](#algo-v1-PredictRequest)
    - [PredictResponse](#algo-v1-PredictResponse)
  
    - [AlgoService](#algo-v1-AlgoService)
  
- [algo/v1/errors.proto](#algo_v1_errors-proto)
    - [Error](#algo-v1-Error)
  
- [Scalar Value Types](#scalar-value-types)



<a name="algo_v1_model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## algo/v1/model.proto



<a name="algo-v1-Material"></a>

### Material



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 原料编号 |
| name | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 原料正式名称 |
| alias | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 原料别名 |
| des | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 原料描述 |
| available_models | [Model](#algo-v1-Model) | repeated | 可用于此材料的模型，可能为空 |






<a name="algo-v1-Model"></a>

### Model



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| name | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 模型名称 |
| version | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 模型版本 |
| state | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 模型状态 |
| des | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 模型描述(text) |
| metadata | [Model.MetadataEntry](#algo-v1-Model-MetadataEntry) | repeated | metadata |






<a name="algo-v1-Model-MetadataEntry"></a>

### Model.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |





 

 

 

 



<a name="algo_v1_algo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## algo/v1/algo.proto



<a name="algo-v1-ListMaterialsRequest"></a>

### ListMaterialsRequest







<a name="algo-v1-ListMaterialsResponse"></a>

### ListMaterialsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| materials | [Material](#algo-v1-Material) | repeated | 原料列表 |






<a name="algo-v1-ListModelsRequest"></a>

### ListModelsRequest







<a name="algo-v1-ListModelsResponse"></a>

### ListModelsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| models | [Model](#algo-v1-Model) | repeated | 模型列表 |






<a name="algo-v1-PredictRequest"></a>

### PredictRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| model_name | [string](#string) |  | 模型名 |
| data | [bytes](#bytes) |  | 用于预测的数据 |






<a name="algo-v1-PredictResponse"></a>

### PredictResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| material | [Material](#algo-v1-Material) |  | 预测结果类别 |





 

 

 


<a name="algo-v1-AlgoService"></a>

### AlgoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListMaterials | [ListMaterialsRequest](#algo-v1-ListMaterialsRequest) | [ListMaterialsResponse](#algo-v1-ListMaterialsResponse) | 获取算法模型支持的材料 |
| ListModels | [ListModelsRequest](#algo-v1-ListModelsRequest) | [ListModelsResponse](#algo-v1-ListModelsResponse) | 列出所有算法模型 |
| Predict | [PredictRequest](#algo-v1-PredictRequest) | [PredictResponse](#algo-v1-PredictResponse) | 进行预测 |

 



<a name="algo_v1_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## algo/v1/errors.proto


 


<a name="algo-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| MODEL_NOT_FOUND | 0 | 模型不存在 |
| MODEL_PREDICT_FAILED | 1 | 调用模型失败 |
| PREDICT_RESULT_UNAVAILABLE | 2 | 分类结果无法识别 |
| MATERIAL_NOT_FOUND | 3 | 原料不存在 |


 

 

 



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

