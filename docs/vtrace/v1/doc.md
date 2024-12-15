# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [vtrace/v1/message.proto](#vtrace_v1_message-proto)
    - [BatchCircRequest](#vtrace-v1-BatchCircRequest)
    - [BatchCircResponse](#vtrace-v1-BatchCircResponse)
    - [BatchCircResponse.GoodsId2transIdEntry](#vtrace-v1-BatchCircResponse-GoodsId2transIdEntry)
    - [BatchCreateGoodsRequest](#vtrace-v1-BatchCreateGoodsRequest)
    - [BatchCreateGoodsResponse](#vtrace-v1-BatchCreateGoodsResponse)
    - [ContractNameRequest](#vtrace-v1-ContractNameRequest)
    - [ContractNameResponse](#vtrace-v1-ContractNameResponse)
    - [CreateCircRequest](#vtrace-v1-CreateCircRequest)
    - [CreateCircResponse](#vtrace-v1-CreateCircResponse)
    - [CreateGoodsClassRequest](#vtrace-v1-CreateGoodsClassRequest)
    - [CreateGoodsClassResponse](#vtrace-v1-CreateGoodsClassResponse)
    - [CreateGoodsSerialRequest](#vtrace-v1-CreateGoodsSerialRequest)
    - [CreateGoodsSerialResponse](#vtrace-v1-CreateGoodsSerialResponse)
    - [DeleteUserRequest](#vtrace-v1-DeleteUserRequest)
    - [DeleteUserResponse](#vtrace-v1-DeleteUserResponse)
    - [GetCircRequest](#vtrace-v1-GetCircRequest)
    - [GetCircResponse](#vtrace-v1-GetCircResponse)
    - [GetClassRequest](#vtrace-v1-GetClassRequest)
    - [GetClassResponse](#vtrace-v1-GetClassResponse)
    - [GetGoodsRequest](#vtrace-v1-GetGoodsRequest)
    - [GetGoodsResponse](#vtrace-v1-GetGoodsResponse)
    - [GetGoodsSerialRequest](#vtrace-v1-GetGoodsSerialRequest)
    - [GetGoodsSerialResponse](#vtrace-v1-GetGoodsSerialResponse)
    - [GetOrgOfUserRequest](#vtrace-v1-GetOrgOfUserRequest)
    - [GetOrgOfUserResponse](#vtrace-v1-GetOrgOfUserResponse)
    - [IdentityAuthRequest](#vtrace-v1-IdentityAuthRequest)
    - [IdentityAuthResponse](#vtrace-v1-IdentityAuthResponse)
    - [ListClassRequest](#vtrace-v1-ListClassRequest)
    - [ListClassResponse](#vtrace-v1-ListClassResponse)
    - [ListGoodsRequest](#vtrace-v1-ListGoodsRequest)
    - [ListGoodsResponse](#vtrace-v1-ListGoodsResponse)
    - [ListGoodsSerialRequest](#vtrace-v1-ListGoodsSerialRequest)
    - [ListGoodsSerialResponse](#vtrace-v1-ListGoodsSerialResponse)
    - [ListMaterialRequest](#vtrace-v1-ListMaterialRequest)
    - [ListMaterialResponse](#vtrace-v1-ListMaterialResponse)
    - [ListModelsRequest](#vtrace-v1-ListModelsRequest)
    - [ListModelsResponse](#vtrace-v1-ListModelsResponse)
    - [ListOrgMemberRequest](#vtrace-v1-ListOrgMemberRequest)
    - [ListOrgMemberResponse](#vtrace-v1-ListOrgMemberResponse)
    - [MemberRequest](#vtrace-v1-MemberRequest)
    - [MemberResponse](#vtrace-v1-MemberResponse)
    - [OrgAuthRequest](#vtrace-v1-OrgAuthRequest)
    - [OrgAuthResponse](#vtrace-v1-OrgAuthResponse)
    - [OrgRemoveMemberRequest](#vtrace-v1-OrgRemoveMemberRequest)
    - [OrgRemoveMemberResponse](#vtrace-v1-OrgRemoveMemberResponse)
    - [PredictRequest](#vtrace-v1-PredictRequest)
    - [PredictResponse](#vtrace-v1-PredictResponse)
    - [ProfileRequest](#vtrace-v1-ProfileRequest)
    - [ProfileResponse](#vtrace-v1-ProfileResponse)
    - [RefreshTokenRequest](#vtrace-v1-RefreshTokenRequest)
    - [RefreshTokenResponse](#vtrace-v1-RefreshTokenResponse)
    - [RegisterRequest](#vtrace-v1-RegisterRequest)
    - [RegisterResponse](#vtrace-v1-RegisterResponse)
    - [RemoveRoleResponse](#vtrace-v1-RemoveRoleResponse)
    - [TokenRequest](#vtrace-v1-TokenRequest)
    - [TokenResponse](#vtrace-v1-TokenResponse)
    - [UpdateClassRequest](#vtrace-v1-UpdateClassRequest)
    - [UpdateClassResponse](#vtrace-v1-UpdateClassResponse)
    - [UpdateGoodsRequest](#vtrace-v1-UpdateGoodsRequest)
    - [UpdateGoodsResponse](#vtrace-v1-UpdateGoodsResponse)
    - [UpdateGoodsSerialRequest](#vtrace-v1-UpdateGoodsSerialRequest)
    - [UpdateGoodsSerialResponse](#vtrace-v1-UpdateGoodsSerialResponse)
    - [UpdateInfoRequest](#vtrace-v1-UpdateInfoRequest)
    - [UpdateOrgRequest](#vtrace-v1-UpdateOrgRequest)
    - [UpdateOrgResponse](#vtrace-v1-UpdateOrgResponse)
    - [UpdateUserRequest](#vtrace-v1-UpdateUserRequest)
    - [UpdateUserResponse](#vtrace-v1-UpdateUserResponse)
    - [UploadRequest](#vtrace-v1-UploadRequest)
    - [UploadRequest.metadata](#vtrace-v1-UploadRequest-metadata)
    - [UploadResponse](#vtrace-v1-UploadResponse)
  
- [vtrace/v1/app.proto](#vtrace_v1_app-proto)
    - [VTraceInterface](#vtrace-v1-VTraceInterface)
  
- [vtrace/v1/errors.proto](#vtrace_v1_errors-proto)
    - [Error](#vtrace-v1-Error)
  
- [Scalar Value Types](#scalar-value-types)



<a name="vtrace_v1_message-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## vtrace/v1/message.proto



<a name="vtrace-v1-BatchCircRequest"></a>

### BatchCircRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_ids | [int32](#int32) | repeated | 商品编号列表 |
| circ_type | [circ.v1.CircType](#circ-v1-CircType) |  | 流转类型 |
| operator | [string](#string) |  | 流转参与者 |
| from | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | from, 当流转类型为&#34;Produce&#34;时为空 |
| to | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | to, 当流转类型为Transfer时不为空 |
| form_info | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 表单json数据 |






<a name="vtrace-v1-BatchCircResponse"></a>

### BatchCircResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goodsId2transId | [BatchCircResponse.GoodsId2transIdEntry](#vtrace-v1-BatchCircResponse-GoodsId2transIdEntry) | repeated | goodsId =&gt; transId |






<a name="vtrace-v1-BatchCircResponse-GoodsId2transIdEntry"></a>

### BatchCircResponse.GoodsId2transIdEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [string](#string) |  |  |






<a name="vtrace-v1-BatchCreateGoodsRequest"></a>

### BatchCreateGoodsRequest
批量创建产品


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial_id | [int32](#int32) |  | 产品批次id |
| sum | [int32](#int32) |  | 本批次产品数量 |






<a name="vtrace-v1-BatchCreateGoodsResponse"></a>

### BatchCreateGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [google.protobuf.Int64Value](#google-protobuf-Int64Value) | repeated | 产品编号列表 |
| successes | [int32](#int32) |  | 成功的数量 |






<a name="vtrace-v1-ContractNameRequest"></a>

### ContractNameRequest







<a name="vtrace-v1-ContractNameResponse"></a>

### ContractNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_name | [google.protobuf.StringValue](#google-protobuf-StringValue) |  |  |






<a name="vtrace-v1-CreateCircRequest"></a>

### CreateCircRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int32](#int32) |  | 商品id |
| circ_type | [circ.v1.CircType](#circ-v1-CircType) |  | 流转类型 |
| from | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 实体来源用户 流转类型为&#34;produce&#34;时，此字段为空 流转类型为&#34;process&#34;、&#34;exam&#34;或&#34;transfer&#34;时，此字段为操作者本身用户名 |
| to | [string](#string) |  | 实体目标用户 流转类型为&#34;produce&#34;、&#34;process&#34;, &#34;exam&#34;时，此字段为操作者 流转类型为&#34;transfer&#34;时，此字段为要转交的用户的用户名 |
| form_value | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 表单数据 |






<a name="vtrace-v1-CreateCircResponse"></a>

### CreateCircResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| circ_id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 流转信息id |
| trans_id | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | transId |






<a name="vtrace-v1-CreateGoodsClassRequest"></a>

### CreateGoodsClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_name | [string](#string) |  | 产品类型名称 |
| goods_des | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 产品类型说明(json) |
| material_id | [int32](#int32) |  | 产品类型配料id |
| tm | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 商标编号(optional) |






<a name="vtrace-v1-CreateGoodsClassResponse"></a>

### CreateGoodsClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 产品类别号 |






<a name="vtrace-v1-CreateGoodsSerialRequest"></a>

### CreateGoodsSerialRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 生产日期 |
| class_id | [int32](#int32) |  | 产品种类号 |






<a name="vtrace-v1-CreateGoodsSerialResponse"></a>

### CreateGoodsSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial_id | [google.protobuf.Int64Value](#google-protobuf-Int64Value) |  | 产品批次id |






<a name="vtrace-v1-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="vtrace-v1-DeleteUserResponse"></a>

### DeleteUserResponse







<a name="vtrace-v1-GetCircRequest"></a>

### GetCircRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int32](#int32) |  | 商品id |






<a name="vtrace-v1-GetCircResponse"></a>

### GetCircResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_info | [goods.v1.Goods](#goods-v1-Goods) |  | 商品信息 |
| circ_records | [circ.v1.CircRecord](#circ-v1-CircRecord) | repeated | 流转记录 |






<a name="vtrace-v1-GetClassRequest"></a>

### GetClassRequest
GetClass 获取类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int32](#int32) |  | 产品类别号 |






<a name="vtrace-v1-GetClassResponse"></a>

### GetClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [goods.v1.Class](#goods-v1-Class) |  | 产品类别号 |






<a name="vtrace-v1-GetGoodsRequest"></a>

### GetGoodsRequest
获得产品


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_id | [int32](#int32) |  |  |






<a name="vtrace-v1-GetGoodsResponse"></a>

### GetGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| good | [goods.v1.Goods](#goods-v1-Goods) |  |  |






<a name="vtrace-v1-GetGoodsSerialRequest"></a>

### GetGoodsSerialRequest
GetGoodsSerial 获取批次


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial_id | [int32](#int32) |  | 产品种类号 |






<a name="vtrace-v1-GetGoodsSerialResponse"></a>

### GetGoodsSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial | [goods.v1.Serial](#goods-v1-Serial) |  | 产品批次号 |






<a name="vtrace-v1-GetOrgOfUserRequest"></a>

### GetOrgOfUserRequest
GetOrgUser 查询用户所属企业


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 唯一用户名 |






<a name="vtrace-v1-GetOrgOfUserResponse"></a>

### GetOrgOfUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [iam.v1.Org](#iam-v1-Org) |  |  |






<a name="vtrace-v1-IdentityAuthRequest"></a>

### IdentityAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| realname | [string](#string) |  | 真实姓名 |
| idcard | [string](#string) |  | 身份证号 |
| username | [string](#string) |  | 唯一用户名 |






<a name="vtrace-v1-IdentityAuthResponse"></a>

### IdentityAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [iam.v1.User](#iam-v1-User) |  |  |






<a name="vtrace-v1-ListClassRequest"></a>

### ListClassRequest
列出产品类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  | 页偏移 |
| limit | [int32](#int32) |  | 页容量 |






<a name="vtrace-v1-ListClassResponse"></a>

### ListClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods_classes | [goods.v1.Class](#goods-v1-Class) | repeated | 产品类型列表 |






<a name="vtrace-v1-ListGoodsRequest"></a>

### ListGoodsRequest
列出产品


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  | 页偏移 |
| limit | [int32](#int32) |  | 页容量 |






<a name="vtrace-v1-ListGoodsResponse"></a>

### ListGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods | [goods.v1.Goods](#goods-v1-Goods) | repeated |  |






<a name="vtrace-v1-ListGoodsSerialRequest"></a>

### ListGoodsSerialRequest
列出产品批次


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  | 页偏移 |
| limit | [int32](#int32) |  | 页容量 |






<a name="vtrace-v1-ListGoodsSerialResponse"></a>

### ListGoodsSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial | [goods.v1.Serial](#goods-v1-Serial) | repeated | 产品编号列表 |






<a name="vtrace-v1-ListMaterialRequest"></a>

### ListMaterialRequest







<a name="vtrace-v1-ListMaterialResponse"></a>

### ListMaterialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| materials | [algo.v1.Material](#algo-v1-Material) | repeated | 原料列表 |






<a name="vtrace-v1-ListModelsRequest"></a>

### ListModelsRequest







<a name="vtrace-v1-ListModelsResponse"></a>

### ListModelsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| models | [algo.v1.Model](#algo-v1-Model) | repeated | 算法模型列表 |






<a name="vtrace-v1-ListOrgMemberRequest"></a>

### ListOrgMemberRequest
企业查询成员列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| offset | [int32](#int32) |  | 页偏移 |
| limit | [int32](#int32) |  | 页容量 |






<a name="vtrace-v1-ListOrgMemberResponse"></a>

### ListOrgMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [iam.v1.User](#iam-v1-User) | repeated |  |






<a name="vtrace-v1-MemberRequest"></a>

### MemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  | org id |
| username | [string](#string) |  | username |






<a name="vtrace-v1-MemberResponse"></a>

### MemberResponse







<a name="vtrace-v1-OrgAuthRequest"></a>

### OrgAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| legal_name | [string](#string) |  | 法人姓名 |
| org_name | [string](#string) |  | 企业全名 |
| code | [string](#string) |  | 社会信用代码 |
| legal_phone | [string](#string) |  | 法人手机号 |
| org_info | [google.protobuf.BytesValue](#google-protobuf-BytesValue) |  | 企业其他信息 |






<a name="vtrace-v1-OrgAuthResponse"></a>

### OrgAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [iam.v1.Org](#iam-v1-Org) |  |  |






<a name="vtrace-v1-OrgRemoveMemberRequest"></a>

### OrgRemoveMemberRequest
RemoveOrgMember 企业删除成员


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  | org id |
| username | [string](#string) |  | 用户名 |






<a name="vtrace-v1-OrgRemoveMemberResponse"></a>

### OrgRemoveMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="vtrace-v1-PredictRequest"></a>

### PredictRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| model_name | [string](#string) |  | 模型名 |
| data | [bytes](#bytes) |  | 用于预测的数据 |






<a name="vtrace-v1-PredictResponse"></a>

### PredictResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| material | [algo.v1.Material](#algo-v1-Material) |  | 预测结果 |






<a name="vtrace-v1-ProfileRequest"></a>

### ProfileRequest







<a name="vtrace-v1-ProfileResponse"></a>

### ProfileResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [iam.v1.User](#iam-v1-User) |  |  |






<a name="vtrace-v1-RefreshTokenRequest"></a>

### RefreshTokenRequest







<a name="vtrace-v1-RefreshTokenResponse"></a>

### RefreshTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | token |






<a name="vtrace-v1-RegisterRequest"></a>

### RegisterRequest
====================================================================
Register
FIXME: 增加验证


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | 电话 |
| nickname | [string](#string) |  | 暱称(nickname) |
| passwd | [string](#string) |  | 密码 |






<a name="vtrace-v1-RegisterResponse"></a>

### RegisterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [iam.v1.User](#iam-v1-User) |  | 唯一用户名 |
| token | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | token |






<a name="vtrace-v1-RemoveRoleResponse"></a>

### RemoveRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="vtrace-v1-TokenRequest"></a>

### TokenRequest
====================================================================
Token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | 电话号码 |
| password | [string](#string) |  | 密码 |






<a name="vtrace-v1-TokenResponse"></a>

### TokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | google.protobuf.Duration expiration = 2; |






<a name="vtrace-v1-UpdateClassRequest"></a>

### UpdateClassRequest
更新产品类型信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [goods.v1.Class](#goods-v1-Class) |  | 待更新信息 |






<a name="vtrace-v1-UpdateClassResponse"></a>

### UpdateClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="vtrace-v1-UpdateGoodsRequest"></a>

### UpdateGoodsRequest
更新产品


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| goods | [goods.v1.Goods](#goods-v1-Goods) |  | 待更新商品 |






<a name="vtrace-v1-UpdateGoodsResponse"></a>

### UpdateGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="vtrace-v1-UpdateGoodsSerialRequest"></a>

### UpdateGoodsSerialRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial | [goods.v1.Serial](#goods-v1-Serial) |  | 待更新产品批次 |






<a name="vtrace-v1-UpdateGoodsSerialResponse"></a>

### UpdateGoodsSerialResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="vtrace-v1-UpdateInfoRequest"></a>

### UpdateInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 暱称 |
| email | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 邮箱 |
| password | [google.protobuf.StringValue](#google-protobuf-StringValue) |  | 密码 |






<a name="vtrace-v1-UpdateOrgRequest"></a>

### UpdateOrgRequest
UpdateOrg 企业信息更新


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [iam.v1.Org](#iam-v1-Org) |  | 待更新企业 |






<a name="vtrace-v1-UpdateOrgResponse"></a>

### UpdateOrgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [iam.v1.Org](#iam-v1-Org) |  |  |






<a name="vtrace-v1-UpdateUserRequest"></a>

### UpdateUserRequest
UpdateUser 用户信息更新


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [iam.v1.User](#iam-v1-User) |  |  |






<a name="vtrace-v1-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [iam.v1.User](#iam-v1-User) |  |  |






<a name="vtrace-v1-UploadRequest"></a>

### UploadRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [bytes](#bytes) |  | 文件内容(base64) |
| mds | [UploadRequest.metadata](#vtrace-v1-UploadRequest-metadata) | repeated | 文件元信息(可选) |






<a name="vtrace-v1-UploadRequest-metadata"></a>

### UploadRequest.metadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="vtrace-v1-UploadResponse"></a>

### UploadResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | 文件url |





 

 

 

 



<a name="vtrace_v1_app-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## vtrace/v1/app.proto


 

 

 


<a name="vtrace-v1-VTraceInterface"></a>

### VTraceInterface


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterRequest](#vtrace-v1-RegisterRequest) | [RegisterResponse](#vtrace-v1-RegisterResponse) | Register |
| Token | [TokenRequest](#vtrace-v1-TokenRequest) | [TokenResponse](#vtrace-v1-TokenResponse) | Token |
| RefreshToken | [RefreshTokenRequest](#vtrace-v1-RefreshTokenRequest) | [RefreshTokenResponse](#vtrace-v1-RefreshTokenResponse) | RefreshToken 更新令牌 |
| Profile | [ProfileRequest](#vtrace-v1-ProfileRequest) | [ProfileResponse](#vtrace-v1-ProfileResponse) | Profile 获取用户信息 |
| DeleteUser | [DeleteUserRequest](#vtrace-v1-DeleteUserRequest) | [DeleteUserResponse](#vtrace-v1-DeleteUserResponse) | DeleteUser 删除用户 |
| IdentityAuth | [IdentityAuthRequest](#vtrace-v1-IdentityAuthRequest) | [IdentityAuthResponse](#vtrace-v1-IdentityAuthResponse) | IdentityAuth 实名认证 |
| OrgAuth | [OrgAuthRequest](#vtrace-v1-OrgAuthRequest) | [OrgAuthResponse](#vtrace-v1-OrgAuthResponse) | 企业认证 |
| Member | [MemberRequest](#vtrace-v1-MemberRequest) | [MemberResponse](#vtrace-v1-MemberResponse) | 企业增加成员 |
| CreateGoodsClass | [CreateGoodsClassRequest](#vtrace-v1-CreateGoodsClassRequest) | [CreateGoodsClassResponse](#vtrace-v1-CreateGoodsClassResponse) | 创建商品种类 |
| CreateGoodsSerial | [CreateGoodsSerialRequest](#vtrace-v1-CreateGoodsSerialRequest) | [CreateGoodsSerialResponse](#vtrace-v1-CreateGoodsSerialResponse) | 创建商品批次 |
| BatchCreateGoods | [BatchCreateGoodsRequest](#vtrace-v1-BatchCreateGoodsRequest) | [BatchCreateGoodsResponse](#vtrace-v1-BatchCreateGoodsResponse) | 批量创建商品 |
| ListClass | [ListClassRequest](#vtrace-v1-ListClassRequest) | [ListClassResponse](#vtrace-v1-ListClassResponse) | ListClass 列出产品类型 |
| ListGoodsSerial | [ListGoodsSerialRequest](#vtrace-v1-ListGoodsSerialRequest) | [ListGoodsSerialResponse](#vtrace-v1-ListGoodsSerialResponse) | ListGoodsSerial 列出产品批次 |
| ListGoods | [ListGoodsRequest](#vtrace-v1-ListGoodsRequest) | [ListGoodsResponse](#vtrace-v1-ListGoodsResponse) | ListGoods 列出商品 |
| UpdateClass | [UpdateClassRequest](#vtrace-v1-UpdateClassRequest) | [UpdateClassResponse](#vtrace-v1-UpdateClassResponse) | 更新商品种类 |
| UpdateGoodsSerial | [UpdateGoodsSerialRequest](#vtrace-v1-UpdateGoodsSerialRequest) | [UpdateGoodsSerialResponse](#vtrace-v1-UpdateGoodsSerialResponse) | 更新产品批次 |
| UpdateGoods | [UpdateGoodsRequest](#vtrace-v1-UpdateGoodsRequest) | [UpdateGoodsResponse](#vtrace-v1-UpdateGoodsResponse) | 更新商品 |
| ListMaterial | [ListMaterialRequest](#vtrace-v1-ListMaterialRequest) | [ListMaterialResponse](#vtrace-v1-ListMaterialResponse) | 获取原材料列表 |
| CreateCirc | [CreateCircRequest](#vtrace-v1-CreateCircRequest) | [CreateCircResponse](#vtrace-v1-CreateCircResponse) | 商品流转 |
| BatchCirc | [BatchCircRequest](#vtrace-v1-BatchCircRequest) | [BatchCircResponse](#vtrace-v1-BatchCircResponse) | 产品批量流传 |
| GetCirc | [GetCircRequest](#vtrace-v1-GetCircRequest) | [GetCircResponse](#vtrace-v1-GetCircResponse) | 获取商品流转历史 |
| ListModels | [ListModelsRequest](#vtrace-v1-ListModelsRequest) | [ListModelsResponse](#vtrace-v1-ListModelsResponse) | ListModels 列出所有的算法模型 |
| Predict | [PredictRequest](#vtrace-v1-PredictRequest) | [PredictResponse](#vtrace-v1-PredictResponse) | Predict 使用算法模型进行预测 |
| Upload | [UploadRequest](#vtrace-v1-UploadRequest) stream | [UploadResponse](#vtrace-v1-UploadResponse) | Upload 上传文件 |
| OrgRemoveMember | [OrgRemoveMemberRequest](#vtrace-v1-OrgRemoveMemberRequest) | [OrgRemoveMemberResponse](#vtrace-v1-OrgRemoveMemberResponse) | OrgRemoveMember 企业删除成员 |
| ListOrgMember | [ListOrgMemberRequest](#vtrace-v1-ListOrgMemberRequest) | [ListOrgMemberResponse](#vtrace-v1-ListOrgMemberResponse) | ListOrgMember 企业查询成员列表 |
| UpdateOrg | [UpdateOrgRequest](#vtrace-v1-UpdateOrgRequest) | [UpdateOrgResponse](#vtrace-v1-UpdateOrgResponse) | UpdateOrg 企业信息更新 |
| UpdateUser | [UpdateUserRequest](#vtrace-v1-UpdateUserRequest) | [UpdateUserResponse](#vtrace-v1-UpdateUserResponse) | UpdateUser 用户信息更新 |
| GetOrgUser | [GetOrgOfUserRequest](#vtrace-v1-GetOrgOfUserRequest) | [GetOrgOfUserResponse](#vtrace-v1-GetOrgOfUserResponse) | GetOrgUser 查询用户所属企业 |

 



<a name="vtrace_v1_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## vtrace/v1/errors.proto


 


<a name="vtrace-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| PERMISSION_DENIED | 0 | 不允许的操作 |
| SERVICE_OFFLINE | 1 | 服务未上线 |


 

 

 



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

