# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/proto/auth/v1/auth.proto](#api_proto_auth_v1_auth-proto)
    - [RefreshRequest](#auth-v1-RefreshRequest)
    - [RefreshResponse](#auth-v1-RefreshResponse)
    - [SignInRequest](#auth-v1-SignInRequest)
    - [SignInResponse](#auth-v1-SignInResponse)
    - [SignUpRequest](#auth-v1-SignUpRequest)
    - [SignUpResponse](#auth-v1-SignUpResponse)
  
    - [AuthService](#auth-v1-AuthService)
  
- [api/proto/health/v1/health.proto](#api_proto_health_v1_health-proto)
    - [CheckRequest](#health-v1-CheckRequest)
    - [CheckResponse](#health-v1-CheckResponse)
  
    - [HealthService](#health-v1-HealthService)
  
- [api/proto/article/v1/article.proto](#api_proto_article_v1_article-proto)
    - [AddTagRequest](#article-v1-AddTagRequest)
    - [AddTagResponse](#article-v1-AddTagResponse)
    - [Article](#article-v1-Article)
    - [DeleteRequest](#article-v1-DeleteRequest)
    - [DeleteResponse](#article-v1-DeleteResponse)
    - [ListRequest](#article-v1-ListRequest)
    - [ListResponse](#article-v1-ListResponse)
    - [ListTagRequest](#article-v1-ListTagRequest)
    - [ListTagResponse](#article-v1-ListTagResponse)
    - [ReadRequest](#article-v1-ReadRequest)
    - [ReadResponse](#article-v1-ReadResponse)
    - [ShareRequest](#article-v1-ShareRequest)
    - [ShareResponse](#article-v1-ShareResponse)
  
    - [ArticleService](#article-v1-ArticleService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_proto_auth_v1_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/proto/auth/v1/auth.proto



<a name="auth-v1-RefreshRequest"></a>

### RefreshRequest
リフレッシュリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| refresh_token | [string](#string) |  |  |






<a name="auth-v1-RefreshResponse"></a>

### RefreshResponse
リフレッシュレスポンス


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |






<a name="auth-v1-SignInRequest"></a>

### SignInRequest
サインインリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| login_id | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="auth-v1-SignInResponse"></a>

### SignInResponse
サインインレスポンス


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |






<a name="auth-v1-SignUpRequest"></a>

### SignUpRequest
サインアップリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| login_id | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="auth-v1-SignUpResponse"></a>

### SignUpResponse
サインアップレスポンス





 

 

 


<a name="auth-v1-AuthService"></a>

### AuthService
認証サービス

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SignUp | [SignUpRequest](#auth-v1-SignUpRequest) | [SignUpResponse](#auth-v1-SignUpResponse) | サインアップ |
| SignIn | [SignInRequest](#auth-v1-SignInRequest) | [SignInResponse](#auth-v1-SignInResponse) | サインイン |
| Refresh | [RefreshRequest](#auth-v1-RefreshRequest) | [RefreshResponse](#auth-v1-RefreshResponse) | リフレッシュ Need Authorization Header |

 



<a name="api_proto_health_v1_health-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/proto/health/v1/health.proto



<a name="health-v1-CheckRequest"></a>

### CheckRequest
チェックリクエスト






<a name="health-v1-CheckResponse"></a>

### CheckResponse
チェックレスポンス





 

 

 


<a name="health-v1-HealthService"></a>

### HealthService
ヘルスサービス

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [CheckRequest](#health-v1-CheckRequest) | [CheckResponse](#health-v1-CheckResponse) | チェック |

 



<a name="api_proto_article_v1_article-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/proto/article/v1/article.proto



<a name="article-v1-AddTagRequest"></a>

### AddTagRequest
タグ追加リクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| tag | [string](#string) |  |  |






<a name="article-v1-AddTagResponse"></a>

### AddTagResponse
タグ追加レスポンス






<a name="article-v1-Article"></a>

### Article
記事モデル


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| url | [string](#string) |  |  |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| thumbnail | [string](#string) |  |  |
| tags | [string](#string) | repeated |  |






<a name="article-v1-DeleteRequest"></a>

### DeleteRequest
削除リクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="article-v1-DeleteResponse"></a>

### DeleteResponse
削除レスポンス






<a name="article-v1-ListRequest"></a>

### ListRequest
記事一覧リクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_token | [string](#string) |  |  |
| max_page_size | [uint32](#uint32) |  |  |






<a name="article-v1-ListResponse"></a>

### ListResponse
記事一覧レスポンス


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| articles | [Article](#article-v1-Article) | repeated |  |
| next_page_token | [string](#string) |  |  |






<a name="article-v1-ListTagRequest"></a>

### ListTagRequest
タグ一覧リクエスト






<a name="article-v1-ListTagResponse"></a>

### ListTagResponse
タグ一覧レスポンス


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [string](#string) | repeated |  |






<a name="article-v1-ReadRequest"></a>

### ReadRequest
既読リクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="article-v1-ReadResponse"></a>

### ReadResponse
既読レスポンス






<a name="article-v1-ShareRequest"></a>

### ShareRequest
記事共有リクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| thumbnail | [string](#string) |  |  |






<a name="article-v1-ShareResponse"></a>

### ShareResponse
記事共有レスポンス


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| article | [Article](#article-v1-Article) |  |  |





 

 

 


<a name="article-v1-ArticleService"></a>

### ArticleService
記事サービス

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Share | [ShareRequest](#article-v1-ShareRequest) | [ShareResponse](#article-v1-ShareResponse) | 共有 Need X-API-KEY Header |
| List | [ListRequest](#article-v1-ListRequest) | [ListResponse](#article-v1-ListResponse) | 一覧取得 Need Authorization Header |
| Delete | [DeleteRequest](#article-v1-DeleteRequest) | [DeleteResponse](#article-v1-DeleteResponse) | 削除 Need Authorization Header |
| Read | [ReadRequest](#article-v1-ReadRequest) | [ReadResponse](#article-v1-ReadResponse) | 既読 Need Authorization Header |
| AddTag | [AddTagRequest](#article-v1-AddTagRequest) | [AddTagResponse](#article-v1-AddTagResponse) | タグ追加 |
| ListTag | [ListTagRequest](#article-v1-ListTagRequest) | [ListTagResponse](#article-v1-ListTagResponse) | タグ一覧 |

 



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

