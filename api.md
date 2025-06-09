# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthLoginResponse">AuthLoginResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthLogoutResponse">AuthLogoutResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthGetUserDetailsResponse">AuthGetUserDetailsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthRevokeTokenResponse">AuthRevokeTokenResponse</a>

Methods:

- <code title="post /auth/login">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthLoginParams">AuthLoginParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthLoginResponse">AuthLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/logout">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthLogoutResponse">AuthLogoutResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/refresh">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /auth/userdetails">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthService.GetUserDetails">GetUserDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthGetUserDetailsResponse">AuthGetUserDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/revoke">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthService.RevokeToken">RevokeToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#AuthRevokeTokenResponse">AuthRevokeTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Encrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#EncryptEncryptDataResponseUnion">EncryptEncryptDataResponseUnion</a>

Methods:

- <code title="post /encrypt">client.Encrypt.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#EncryptService.EncryptData">EncryptData</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#EncryptEncryptDataParams">EncryptEncryptDataParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#EncryptEncryptDataResponseUnion">EncryptEncryptDataResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Decrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#DecryptDecryptPayloadResponseUnion">DecryptDecryptPayloadResponseUnion</a>

Methods:

- <code title="post /decrypt">client.Decrypt.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#DecryptService.DecryptPayload">DecryptPayload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#DecryptDecryptPayloadParams">DecryptDecryptPayloadParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#DecryptDecryptPayloadResponseUnion">DecryptDecryptPayloadResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyRevokeResponse">APIKeyRevokeResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyRotateResponse">APIKeyRotateResponse</a>

Methods:

- <code title="patch /api-keys/{apiKey}/revoke">client.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyRevokeResponse">APIKeyRevokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api-keys/{apiKey}/rotate">client.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyRotateResponse">APIKeyRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Scopes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeGetResponse">APIKeyScopeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeAttachResponse">APIKeyScopeAttachResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeDetachResponse">APIKeyScopeDetachResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeSyncResponse">APIKeyScopeSyncResponse</a>

Methods:

- <code title="get /api-keys/{apiKey}/scopes">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeGetResponse">APIKeyScopeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api-keys/{apiKey}/scopes/attach">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeAttachParams">APIKeyScopeAttachParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeAttachResponse">APIKeyScopeAttachResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api-keys/{apiKey}/scopes/detach">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeService.Detach">Detach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeDetachParams">APIKeyScopeDetachParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeDetachResponse">APIKeyScopeDetachResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api-keys/{apiKey}/scopes/sync">client.APIKeys.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeService.Sync">Sync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeSyncParams">APIKeyScopeSyncParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#APIKeyScopeSyncResponse">APIKeyScopeSyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scopes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeNewResponse">ScopeNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeGetResponse">ScopeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeUpdateResponse">ScopeUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeListResponse">ScopeListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeDeleteResponse">ScopeDeleteResponse</a>

Methods:

- <code title="post /scopes">client.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeNewParams">ScopeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeNewResponse">ScopeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /scopes/{id}">client.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeGetResponse">ScopeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /scopes/{id}">client.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeUpdateParams">ScopeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeUpdateResponse">ScopeUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /scopes">client.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeListResponse">ScopeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /scopes/{id}">client.Scopes.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go">qanapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/qanapi-go#ScopeDeleteResponse">ScopeDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
