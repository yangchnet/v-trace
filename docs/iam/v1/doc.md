# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [iam/v1/error.proto](#iam_v1_error-proto)
    - [Error](#iam-v1-Error)
  
- [iam/v1/model.proto](#iam_v1_model-proto)
    - [Org](#iam-v1-Org)
    - [User](#iam-v1-User)
  
    - [Role](#iam-v1-Role)
    - [Status](#iam-v1-Status)
  
- [iam/v1/iam.proto](#iam_v1_iam-proto)
    - [CreateIdentityRequest](#iam-v1-CreateIdentityRequest)
    - [CreateIdentityResponse](#iam-v1-CreateIdentityResponse)
    - [CreateOrgRequest](#iam-v1-CreateOrgRequest)
    - [CreateOrgResponse](#iam-v1-CreateOrgResponse)
    - [CreateUserRequest](#iam-v1-CreateUserRequest)
    - [CreateUserResponse](#iam-v1-CreateUserResponse)
    - [DeleteOrgMemberRequest](#iam-v1-DeleteOrgMemberRequest)
    - [DeleteOrgMemberResponse](#iam-v1-DeleteOrgMemberResponse)
    - [DeleteUserRequest](#iam-v1-DeleteUserRequest)
    - [DeleteUserResponse](#iam-v1-DeleteUserResponse)
    - [GetOrgOfUserRequest](#iam-v1-GetOrgOfUserRequest)
    - [GetOrgOfUserResponse](#iam-v1-GetOrgOfUserResponse)
    - [GetOrgRequest](#iam-v1-GetOrgRequest)
    - [GetOrgResponse](#iam-v1-GetOrgResponse)
    - [GetRoleRequest](#iam-v1-GetRoleRequest)
    - [GetRoleResponse](#iam-v1-GetRoleResponse)
    - [GetUserRequest](#iam-v1-GetUserRequest)
    - [GetUserResponse](#iam-v1-GetUserResponse)
    - [ListOrgMemberRequest](#iam-v1-ListOrgMemberRequest)
    - [ListOrgMemberResponse](#iam-v1-ListOrgMemberResponse)
    - [OrgAddMemberRequest](#iam-v1-OrgAddMemberRequest)
    - [OrgAddMemberResponse](#iam-v1-OrgAddMemberResponse)
    - [RefreshTokenRequest](#iam-v1-RefreshTokenRequest)
    - [RefreshTokenResponse](#iam-v1-RefreshTokenResponse)
    - [TokenRequest](#iam-v1-TokenRequest)
    - [TokenResponse](#iam-v1-TokenResponse)
    - [UpdateOrgRequest](#iam-v1-UpdateOrgRequest)
    - [UpdateOrgResponse](#iam-v1-UpdateOrgResponse)
    - [UpdateUserRequest](#iam-v1-UpdateUserRequest)
    - [UpdateUserResponse](#iam-v1-UpdateUserResponse)
  
    - [OperationResult](#iam-v1-OperationResult)
  
    - [IamService](#iam-v1-IamService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="iam_v1_error-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iam/v1/error.proto


 


<a name="iam-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| PASSWD_HASH_FAILED | 0 | 密码hash错误 |
| CREATE_USER_FAILED | 1 | 创建用户失败 |
| USER_NOT_FOUND | 2 | 用户未找到 |
| PASSWD_ERR | 3 | 登录密码错误 |
| PARAMS_ERR | 4 | 参数错误 |
| ROLE_NOT_EXIST | 5 | 不存在的角色 |
| ORG_NOT_EXIST | 6 | 企业不存在 |
| ENTRY_DUPLICATE | 7 | 数据已存在 |
| TOKEN_EXPIRED | 8 | token过期 |
| REFRESH_USER_FAILED | 9 | 更新用户信息失败 |


 

 

 



<a name="iam_v1_model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iam/v1/model.proto



<a name="iam-v1-Org"></a>

### Org



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| org_name | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 企业名称 |
| org_code | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 企业社会信用代码 |
| legal_person_name | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 企业法人姓名 |
| legal_person_phone | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 企业法人手机号 |
| owner | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 企业所有者或管理员 |
| org_info | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 企业信息(json) |






<a name="iam-v1-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 唯一用户名 |
| nickname | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 用户暱称 |
| role | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 用户角色 |
| phone | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 手机 |
| email | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 邮箱 |
| create_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 注册时间 |
| realname | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 实际姓名 |
| idcard | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 身份证号 |
| id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | id |
| avatar | [string](#string) |  | 头像 |
| status | [Status](#iam-v1-Status) |  | 用户状态 （正常/true|已删除/false） |





 


<a name="iam-v1-Role"></a>

### Role


| Name | Number | Description |
| ---- | ------ | ----------- |
| admin | 0 |  |
| producer | 1 |  |
| transporter | 2 |  |
| examiner | 3 |  |
| normal | 4 |  |
| boss | 5 |  |



<a name="iam-v1-Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| enable | 0 |  |
| deleted | 1 |  |


 

 

 



<a name="iam_v1_iam-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iam/v1/iam.proto



<a name="iam-v1-CreateIdentityRequest"></a>

### CreateIdentityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| real_name | [string](#string) |  |  |
| id_card | [string](#string) |  |  |






<a name="iam-v1-CreateIdentityResponse"></a>

### CreateIdentityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#iam-v1-User) |  |  |






<a name="iam-v1-CreateOrgRequest"></a>

### CreateOrgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | 企业所有者或管理员 |
| org_name | [string](#string) |  | 企业名称 |
| org_code | [string](#string) |  | 企业社会信用代码 |
| legal_name | [string](#string) |  | 企业法人姓名 |
| legal_phone | [string](#string) |  | 企业法人手机号 |
| can_produce | [bool](#bool) |  | 是否可进行产品生产 |
| org_info | [bytes](#bytes) |  | 企业信息(json) |






<a name="iam-v1-CreateOrgResponse"></a>

### CreateOrgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#iam-v1-Org) |  |  |






<a name="iam-v1-CreateUserRequest"></a>

### CreateUserRequest
User Manage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) |  | 暱称, 必填 |
| password | [string](#string) |  | 密码, 必填 |
| phone | [string](#string) |  | 手机号码，要求符合11位数字格式, 必填 |
| email | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 邮箱（选填） |
| avatar | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 头像 （选填） |






<a name="iam-v1-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#iam-v1-User) |  |  |






<a name="iam-v1-DeleteOrgMemberRequest"></a>

### DeleteOrgMemberRequest
RemoveOrgMember 企业删除成员


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int64](#int64) |  | org id |
| username | [string](#string) |  | 用户名 |






<a name="iam-v1-DeleteOrgMemberResponse"></a>

### DeleteOrgMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="iam-v1-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 唯一用户名 |






<a name="iam-v1-DeleteUserResponse"></a>

### DeleteUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [OperationResult](#iam-v1-OperationResult) |  | 操作结果 |






<a name="iam-v1-GetOrgOfUserRequest"></a>

### GetOrgOfUserRequest
查询用户所属企业


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 唯一用户名 |






<a name="iam-v1-GetOrgOfUserResponse"></a>

### GetOrgOfUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#iam-v1-Org) |  |  |






<a name="iam-v1-GetOrgRequest"></a>

### GetOrgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int64](#int64) |  |  |






<a name="iam-v1-GetOrgResponse"></a>

### GetOrgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#iam-v1-Org) |  |  |






<a name="iam-v1-GetRoleRequest"></a>

### GetRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="iam-v1-GetRoleResponse"></a>

### GetRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [google.protobuf.StringValue](#google-protobuf-StringValue) |  |  |






<a name="iam-v1-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="iam-v1-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#iam-v1-User) |  | 用户信息 |






<a name="iam-v1-ListOrgMemberRequest"></a>

### ListOrgMemberRequest
企业查询成员列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int64](#int64) |  |  |
| offset | [int64](#int64) |  | 页偏移 |
| limit | [int64](#int64) |  | 页容量 |






<a name="iam-v1-ListOrgMemberResponse"></a>

### ListOrgMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#iam-v1-User) | repeated |  |






<a name="iam-v1-OrgAddMemberRequest"></a>

### OrgAddMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int64](#int64) |  | org id |
| username | [string](#string) |  | 用户名 |






<a name="iam-v1-OrgAddMemberResponse"></a>

### OrgAddMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="iam-v1-RefreshTokenRequest"></a>

### RefreshTokenRequest







<a name="iam-v1-RefreshTokenResponse"></a>

### RefreshTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | token |






<a name="iam-v1-TokenRequest"></a>

### TokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | 手机号 |
| passwd | [string](#string) |  | 密码 |






<a name="iam-v1-TokenResponse"></a>

### TokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | token |






<a name="iam-v1-UpdateOrgRequest"></a>

### UpdateOrgRequest
企业信息更新


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#iam-v1-Org) |  | 待更新企业 |






<a name="iam-v1-UpdateOrgResponse"></a>

### UpdateOrgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#iam-v1-Org) |  |  |






<a name="iam-v1-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#iam-v1-User) |  |  |






<a name="iam-v1-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#iam-v1-User) |  |  |





 


<a name="iam-v1-OperationResult"></a>

### OperationResult


| Name | Number | Description |
| ---- | ------ | ----------- |
| SUCCESS | 0 |  |
| FAIL | 1 |  |


 

 


<a name="iam-v1-IamService"></a>

### IamService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Token | [TokenRequest](#iam-v1-TokenRequest) | [TokenResponse](#iam-v1-TokenResponse) | Token |
| RefreshToken | [RefreshTokenRequest](#iam-v1-RefreshTokenRequest) | [RefreshTokenResponse](#iam-v1-RefreshTokenResponse) | RefreshToken 更新令牌 |
| CreateUser | [CreateUserRequest](#iam-v1-CreateUserRequest) | [CreateUserResponse](#iam-v1-CreateUserResponse) | CreateUser 创建一个新用户 |
| DeleteUser | [DeleteUserRequest](#iam-v1-DeleteUserRequest) | [DeleteUserResponse](#iam-v1-DeleteUserResponse) | DeleteUser 删除用户 1. delete user entity 2. delete user&#39;s role relation 3. delete member from org |
| GetUser | [GetUserRequest](#iam-v1-GetUserRequest) | [GetUserResponse](#iam-v1-GetUserResponse) | GetUser 根据用户名获取用户信息 |
| UpdateUser | [UpdateUserRequest](#iam-v1-UpdateUserRequest) | [UpdateUserResponse](#iam-v1-UpdateUserResponse) | UpdateUser 用户信息更新 |
| GetRole | [GetRoleRequest](#iam-v1-GetRoleRequest) | [GetRoleResponse](#iam-v1-GetRoleResponse) | GetRole 根据用户名获取用户角色 |
| CreateOrg | [CreateOrgRequest](#iam-v1-CreateOrgRequest) | [CreateOrgResponse](#iam-v1-CreateOrgResponse) | CreateOrg 创建一个企业组织 this api should do: 1. create an org 2. add caller as org member 3. grant caller boss role |
| OrgAddMember | [OrgAddMemberRequest](#iam-v1-OrgAddMemberRequest) | [OrgAddMemberResponse](#iam-v1-OrgAddMemberResponse) | OrgAddMember 企业增加成员 this api should do: 1. add member to org 2. grant producer role to member |
| GetOrg | [GetOrgRequest](#iam-v1-GetOrgRequest) | [GetOrgResponse](#iam-v1-GetOrgResponse) | GetOrg 获取企业信息 |
| DeleteOrgMember | [DeleteOrgMemberRequest](#iam-v1-DeleteOrgMemberRequest) | [DeleteOrgMemberResponse](#iam-v1-DeleteOrgMemberResponse) | DeleteOrgMember 企业删除成员 1. delete member 2. remove producer role |
| ListOrgMember | [ListOrgMemberRequest](#iam-v1-ListOrgMemberRequest) | [ListOrgMemberResponse](#iam-v1-ListOrgMemberResponse) | ListOrgMember 企业查询成员列表 |
| UpdateOrg | [UpdateOrgRequest](#iam-v1-UpdateOrgRequest) | [UpdateOrgResponse](#iam-v1-UpdateOrgResponse) | UpdateOrg 企业信息更新 |
| CreateIdentity | [CreateIdentityRequest](#iam-v1-CreateIdentityRequest) | [CreateIdentityResponse](#iam-v1-CreateIdentityResponse) | 记录实名信息 1. 记录实名信息 2. 授予transporter权限 |
| GetOrgOfUser | [GetOrgOfUserRequest](#iam-v1-GetOrgOfUserRequest) | [GetOrgOfUserResponse](#iam-v1-GetOrgOfUserResponse) | 查询用户所属企业 |

 



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

