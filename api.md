# V3

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Configuration">Configuration</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Permission">Permission</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Role">Role</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#User">User</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Value">Value</a>

## Roles

Methods:

- <code title="get /v3/roles">client.V3.Roles.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3RoleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Role">Role</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Configurations

Methods:

- <code title="post /v3/configurations">client.V3.Configurations.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3ConfigurationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3ConfigurationNewParams">V3ConfigurationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/configurations/{configuration}">client.V3.Configurations.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3ConfigurationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, configuration <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3ConfigurationUpdateParams">V3ConfigurationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/configurations">client.V3.Configurations.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3ConfigurationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/configurations/{configuration}">client.V3.Configurations.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3ConfigurationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, configuration <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v3/configurations/{configuration}">client.V3.Configurations.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3ConfigurationService.Show">Show</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, configuration <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#Configuration">Configuration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Methods:

- <code title="post /v3/users">client.V3.Users.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserNewParams">V3UserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/users">client.V3.Users.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/users/{user}">client.V3.Users.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, user <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v3/users/me">client.V3.Users.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserService.Me">Me</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/users/{user}">client.V3.Users.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, user <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserPatchParams">V3UserPatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/users/{user}/restore">client.V3.Users.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserService.Restore">Restore</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, user <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/users/{user}">client.V3.Users.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3UserService.Show">Show</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, user <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3APIKeyRotateResponse">V3APIKeyRotateResponse</a>

Methods:

- <code title="get /v3/api-keys">client.V3.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/api-keys/{apiKey}/revoke">client.V3.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3APIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/api-keys/{apiKey}/rotate">client.V3.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3APIKeyService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3APIKeyRotateResponse">V3APIKeyRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/api-keys/{apiKey}">client.V3.APIKeys.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3APIKeyService.Show">Show</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKey <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogActivityResponse">V3LogActivityResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogAPIResponse">V3LogAPIResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogQanapiFlowResponse">V3LogQanapiFlowResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogUnifiedResponse">V3LogUnifiedResponse</a>

Methods:

- <code title="get /v3/logs/activity">client.V3.Logs.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogService.Activity">Activity</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogActivityParams">V3LogActivityParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogActivityResponse">V3LogActivityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/logs/api">client.V3.Logs.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogService.API">API</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogAPIParams">V3LogAPIParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogAPIResponse">V3LogAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/logs/qanapi-flow">client.V3.Logs.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogService.QanapiFlow">QanapiFlow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogQanapiFlowParams">V3LogQanapiFlowParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogQanapiFlowResponse">V3LogQanapiFlowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/logs/unified">client.V3.Logs.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogService.Unified">Unified</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogUnifiedParams">V3LogUnifiedParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3LogUnifiedResponse">V3LogUnifiedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Encryption

Response Types:

- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionDecryptResponse">V3EncryptionDecryptResponse</a>
- <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionEncryptResponse">V3EncryptionEncryptResponse</a>

Methods:

- <code title="post /v3/encryption/{proxy}/decrypt">client.V3.Encryption.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionService.Decrypt">Decrypt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, proxy <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionDecryptParams">V3EncryptionDecryptParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionDecryptResponse">V3EncryptionDecryptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/encryption/{proxy}/encrypt">client.V3.Encryption.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionService.Encrypt">Encrypt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, proxy <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionEncryptParams">V3EncryptionEncryptParams</a>) (\*<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang">qanapi</a>.<a href="https://pkg.go.dev/github.com/qanapi/qanapi-sdk-golang#V3EncryptionEncryptResponse">V3EncryptionEncryptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
