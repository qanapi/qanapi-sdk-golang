# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLoginResponse">AuthLoginResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLogoutResponse">AuthLogoutResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthGetUserDetailsResponse">AuthGetUserDetailsResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRevokeTokenResponse">AuthRevokeTokenResponse</a>

Methods:

- <code title="post /auth/login">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLoginParams">AuthLoginParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLoginResponse">AuthLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/logout">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLogoutResponse">AuthLogoutResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/refresh">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /auth/userdetails">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.GetUserDetails">GetUserDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthGetUserDetailsResponse">AuthGetUserDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/revoke">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.RevokeToken">RevokeToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRevokeTokenResponse">AuthRevokeTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Encrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptEncryptDataResponseUnion">EncryptEncryptDataResponseUnion</a>

Methods:

- <code title="post /encrypt">client.Encrypt.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptService.EncryptData">EncryptData</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptEncryptDataParams">EncryptEncryptDataParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptEncryptDataResponseUnion">EncryptEncryptDataResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Decrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptDecryptPayloadResponseUnion">DecryptDecryptPayloadResponseUnion</a>

Methods:

- <code title="post /decrypt">client.Decrypt.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptService.DecryptPayload">DecryptPayload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptDecryptPayloadParams">DecryptDecryptPayloadParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptDecryptPayloadResponseUnion">DecryptDecryptPayloadResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRevokeResponse">APIKeyRevokeResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRotateResponse">APIKeyRotateResponse</a>

Methods:

- <code title="patch /api-keys/{apiKey}/revoke">client.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRevokeResponse">APIKeyRevokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api-keys/{apiKey}/rotate">client.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRotateResponse">APIKeyRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Scopes

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeGetResponse">APIKeyScopeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeAttachResponse">APIKeyScopeAttachResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeDetachResponse">APIKeyScopeDetachResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeSyncResponse">APIKeyScopeSyncResponse</a>

Methods:

- <code title="get /api-keys/{apiKey}/scopes">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>) ([]<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeGetResponse">APIKeyScopeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api-keys/{apiKey}/scopes/attach">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeAttachParams">APIKeyScopeAttachParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeAttachResponse">APIKeyScopeAttachResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api-keys/{apiKey}/scopes/detach">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeService.Detach">Detach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeDetachParams">APIKeyScopeDetachParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeDetachResponse">APIKeyScopeDetachResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api-keys/{apiKey}/scopes/sync">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeService.Sync">Sync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeSyncParams">APIKeyScopeSyncParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyScopeSyncResponse">APIKeyScopeSyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scopes

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeNewResponse">ScopeNewResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeGetResponse">ScopeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeUpdateResponse">ScopeUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeListResponse">ScopeListResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeDeleteResponse">ScopeDeleteResponse</a>

Methods:

- <code title="post /scopes">client.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeNewParams">ScopeNewParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeNewResponse">ScopeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /scopes/{id}">client.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeGetResponse">ScopeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /scopes/{id}">client.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeUpdateParams">ScopeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeUpdateResponse">ScopeUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /scopes">client.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeListResponse">ScopeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /scopes/{id}">client.Scopes.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#ScopeDeleteResponse">ScopeDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
