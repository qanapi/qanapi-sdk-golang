# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLoginResponse">AuthLoginResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLogoutResponse">AuthLogoutResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthGetUserDetailsResponse">AuthGetUserDetailsResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRevokeTokenResponse">AuthRevokeTokenResponse</a>

Methods:

- <code title="post /auth/login">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLoginParams">AuthLoginParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLoginResponse">AuthLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/logout">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthLogoutResponse">AuthLogoutResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/refresh">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /auth/userdetails">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.GetUserDetails">GetUserDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthGetUserDetailsResponse">AuthGetUserDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/revoke">client.Auth.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthService.RevokeToken">RevokeToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#AuthRevokeTokenResponse">AuthRevokeTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Encrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptEncryptDataResponseUnion">EncryptEncryptDataResponseUnion</a>

Methods:

- <code title="post /encrypt">client.Encrypt.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptService.EncryptData">EncryptData</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptEncryptDataParams">EncryptEncryptDataParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#EncryptEncryptDataResponseUnion">EncryptEncryptDataResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Decrypt

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptDecryptPayloadResponseUnion">DecryptDecryptPayloadResponseUnion</a>

Methods:

- <code title="post /decrypt">client.Decrypt.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptService.DecryptPayload">DecryptPayload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptDecryptPayloadParams">DecryptDecryptPayloadParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#DecryptDecryptPayloadResponseUnion">DecryptDecryptPayloadResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRevokeResponse">APIKeyRevokeResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRotateResponse">APIKeyRotateResponse</a>

Methods:

- <code title="patch /api-keys/{apiKey}/revoke">client.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRevokeResponse">APIKeyRevokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api-keys/{apiKey}/rotate">client.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKeyRotateResponse">APIKeyRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
