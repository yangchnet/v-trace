# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [ca/v1/ca.proto](#ca_v1_ca-proto)
    - [GenCertRequest](#ca-v1-GenCertRequest)
    - [GenCertResponse](#ca-v1-GenCertResponse)
    - [GetCertRequest](#ca-v1-GetCertRequest)
    - [GetCertResponse](#ca-v1-GetCertResponse)
  
    - [CertUsage](#ca-v1-CertUsage)
    - [UserType](#ca-v1-UserType)
  
    - [CAService](#ca-v1-CAService)
  
- [ca/v1/errors.proto](#ca_v1_errors-proto)
    - [Error](#ca-v1-Error)
  
- [Scalar Value Types](#scalar-value-types)



<a name="ca_v1_ca-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ca/v1/ca.proto



<a name="ca-v1-GenCertRequest"></a>

### GenCertRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [string](#string) |  | 组织id |
| username | [string](#string) |  | 唯一用户名 |
| user_type | [UserType](#ca-v1-UserType) |  | 用户类型(root, ca, admin, client, consensus, common) |
| cert_usage | [CertUsage](#ca-v1-CertUsage) | repeated | 证书用途(sign, tls, tls-enc) |
| private_key_pwd | [string](#string) |  | 密钥密码 |
| country | [string](#string) |  | 证书字段-国家 |
| locality | [string](#string) |  | 证书字段-城市 |
| province | [string](#string) |  | 证书字段-省份 |






<a name="ca-v1-GenCertResponse"></a>

### GenCertResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cert | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 证书内容 |
| private_key | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 密钥内容 |
| username | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 用户名 |
| tls_cert | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | tls证书 |
| tls_key | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | tls密钥 |






<a name="ca-v1-GetCertRequest"></a>

### GetCertRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [string](#string) |  | 组织id |
| username | [string](#string) |  | 唯一用户名 |
| user_type | [UserType](#ca-v1-UserType) |  | 用户类型(root, ca, admin, client, consensus, common) |
| cert_usage | [CertUsage](#ca-v1-CertUsage) | repeated | 证书用途(sign, tls, tls-enc) |






<a name="ca-v1-GetCertResponse"></a>

### GetCertResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cert | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 证书内容 |
| private_key | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 密钥内容 |
| username | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 用户名 |
| tls_cert | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | tls证书 |
| tls_key | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | tls密钥 |





 


<a name="ca-v1-CertUsage"></a>

### CertUsage


| Name | Number | Description |
| ---- | ------ | ----------- |
| sign | 0 |  |
| tls | 1 |  |
| tls_enc | 2 |  |



<a name="ca-v1-UserType"></a>

### UserType


| Name | Number | Description |
| ---- | ------ | ----------- |
| root | 0 |  |
| ca | 1 |  |
| admin | 2 |  |
| client | 3 |  |
| consensus | 4 |  |
| common | 5 |  |


 

 


<a name="ca-v1-CAService"></a>

### CAService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GenCert | [GenCertRequest](#ca-v1-GenCertRequest) | [GenCertResponse](#ca-v1-GenCertResponse) |  |
| GetCert | [GetCertRequest](#ca-v1-GetCertRequest) | [GetCertResponse](#ca-v1-GetCertResponse) |  |

 



<a name="ca_v1_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ca/v1/errors.proto


 


<a name="ca-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| CERT_NOT_FOUND | 0 | 证书不存在 |
| INVALID_PARAMS | 1 | 参数错误 |
| GEN_CERT_ERR | 2 | 证书生成失败 |


 

 

 



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

