# V2

## Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthLoginResponse">V2AuthLoginResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthLogoutResponse">V2AuthLogoutResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthRefreshTokenResponse">V2AuthRefreshTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthGetUserDetailsResponse">V2AuthGetUserDetailsResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthRevokeTokenResponse">V2AuthRevokeTokenResponse</a>

Methods:

- <code title="post /v2/auth/login">client.V2.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthLoginParams">V2AuthLoginParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthLoginResponse">V2AuthLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/auth/logout">client.V2.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthLogoutResponse">V2AuthLogoutResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/auth/refresh">client.V2.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthRefreshTokenResponse">V2AuthRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/auth/userdetails">client.V2.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthService.GetUserDetails">GetUserDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthGetUserDetailsResponse">V2AuthGetUserDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/auth/revoke">client.V2.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthService.RevokeToken">RevokeToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2AuthRevokeTokenResponse">V2AuthRevokeTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Encrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2EncryptEncryptDataResponseUnion">V2EncryptEncryptDataResponseUnion</a>

Methods:

- <code title="post /v2/encrypt">client.V2.Encrypt.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2EncryptService.EncryptData">EncryptData</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2EncryptEncryptDataParams">V2EncryptEncryptDataParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2EncryptEncryptDataResponseUnion">V2EncryptEncryptDataResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Decrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2DecryptDecryptPayloadResponseUnion">V2DecryptDecryptPayloadResponseUnion</a>

Methods:

- <code title="post /v2/decrypt">client.V2.Decrypt.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2DecryptService.DecryptPayload">DecryptPayload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2DecryptDecryptPayloadParams">V2DecryptDecryptPayloadParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2DecryptDecryptPayloadResponseUnion">V2DecryptDecryptPayloadResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2APIKeyRevokeResponse">V2APIKeyRevokeResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2APIKeyRotateResponse">V2APIKeyRotateResponse</a>

Methods:

- <code title="patch /v2/api-keys/{apiKey}/revoke">client.V2.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2APIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2APIKeyRevokeResponse">V2APIKeyRevokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/api-keys/{apiKey}/rotate">client.V2.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2APIKeyService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V2APIKeyRotateResponse">V2APIKeyRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
