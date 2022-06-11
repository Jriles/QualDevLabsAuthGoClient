# \DefaultApi

All URIs are relative to *https://qualdevlabsauth.uc.r.appspot.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /orgs/{orgId}/apps/{appId}/users | Creates a new user.
[**CreateUserSession**](DefaultApi.md#CreateUserSession) | **Post** /orgs/{orgId}/apps/{appId}/sessions | Login user
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /orgs/{orgId}/apps/{appId}/users/{userId} | Delete a user
[**DeleteUserSession**](DefaultApi.md#DeleteUserSession) | **Delete** /orgs/{orgId}/apps/{appId}/sessions/{userId} | Logout user
[**ForgotPasswordReset**](DefaultApi.md#ForgotPasswordReset) | **Patch** /orgs/{orgId}/apps/{appId}/password_reset | resets an unathenticated user&#39;s password using a token sent to their email
[**RequestForgottenUsername**](DefaultApi.md#RequestForgottenUsername) | **Post** /orgs/{orgId}/apps/{appId}/forgot_username | requests an email get sent to the user with their username
[**RequestPasswordResetToken**](DefaultApi.md#RequestPasswordResetToken) | **Post** /orgs/{orgId}/apps/{appId}/password_reset_tokens | requests a password reset token for users who have forgotten their password
[**UpdateImportantUserDetails**](DefaultApi.md#UpdateImportantUserDetails) | **Patch** /orgs/{orgId}/apps/{appId}/users/{userId} | Updates a user&#39;s important details
[**ValidateSession**](DefaultApi.md#ValidateSession) | **Get** /orgs/{orgId}/apps/{appId}/sessions/{userId} | Verifies if session is valid



## CreateUser

> CreateUser(ctx, orgId, appId).UserSchema(userSchema).Execute()

Creates a new user.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    userSchema := *openapiclient.NewUserSchema("Username_example", "Password_example", "RepeatPassword_example", "Email_example") // UserSchema | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUser(context.Background(), orgId, appId).UserSchema(userSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userSchema** | [**UserSchema**](UserSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserSession

> SessionTokenAndUserId CreateUserSession(ctx, orgId, appId).LoginSchema(loginSchema).Execute()

Login user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    loginSchema := *openapiclient.NewLoginSchema("Username_example", "Password_example", false) // LoginSchema | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUserSession(context.Background(), orgId, appId).LoginSchema(loginSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUserSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserSession`: SessionTokenAndUserId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUserSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loginSchema** | [**LoginSchema**](LoginSchema.md) |  | 

### Return type

[**SessionTokenAndUserId**](SessionTokenAndUserId.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, orgId, appId, userId).Execute()

Delete a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    userId := "userId_example" // string | the user's UUID (unique)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUser(context.Background(), orgId, appId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 
**userId** | **string** | the user&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [passwordInHeader](../README.md#passwordInHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserSession

> DeleteUserSession(ctx, orgId, appId, userId).Execute()

Logout user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    userId := "userId_example" // string | the user's UUID (unique)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUserSession(context.Background(), orgId, appId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 
**userId** | **string** | the user&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [tokenHeader](../README.md#tokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPasswordReset

> ForgotPasswordReset(ctx, orgId, appId).PasswordResetSchema(passwordResetSchema).Execute()

resets an unathenticated user's password using a token sent to their email

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    passwordResetSchema := *openapiclient.NewPasswordResetSchema("Password_example", "RepeatPassword_example", "Email_example") // PasswordResetSchema | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ForgotPasswordReset(context.Background(), orgId, appId).PasswordResetSchema(passwordResetSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ForgotPasswordReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **passwordResetSchema** | [**PasswordResetSchema**](PasswordResetSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [tokenHeader](../README.md#tokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestForgottenUsername

> RequestForgottenUsername(ctx, orgId, appId).UserEmailSchema1(userEmailSchema1).Execute()

requests an email get sent to the user with their username

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    userEmailSchema1 := *openapiclient.NewUserEmailSchema1("Email_example") // UserEmailSchema1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RequestForgottenUsername(context.Background(), orgId, appId).UserEmailSchema1(userEmailSchema1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RequestForgottenUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestForgottenUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userEmailSchema1** | [**UserEmailSchema1**](UserEmailSchema1.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestPasswordResetToken

> RequestPasswordResetToken(ctx, orgId, appId).UserEmailSchema(userEmailSchema).Execute()

requests a password reset token for users who have forgotten their password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    userEmailSchema := *openapiclient.NewUserEmailSchema("Email_example") // UserEmailSchema | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RequestPasswordResetToken(context.Background(), orgId, appId).UserEmailSchema(userEmailSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RequestPasswordResetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestPasswordResetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userEmailSchema** | [**UserEmailSchema**](UserEmailSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImportantUserDetails

> UpdateImportantUserDetails(ctx, orgId, appId, userId).UserUpdatesSchema(userUpdatesSchema).Execute()

Updates a user's important details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    userId := "userId_example" // string | the user's UUID (unique)
    userUpdatesSchema := *openapiclient.NewUserUpdatesSchema() // UserUpdatesSchema | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateImportantUserDetails(context.Background(), orgId, appId, userId).UserUpdatesSchema(userUpdatesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateImportantUserDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 
**userId** | **string** | the user&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImportantUserDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userUpdatesSchema** | [**UserUpdatesSchema**](UserUpdatesSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [passwordInHeader](../README.md#passwordInHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSession

> ValidateSession(ctx, orgId, appId, userId).Execute()

Verifies if session is valid

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | the org's UUID (unique)
    appId := "appId_example" // string | the app's UUID (unique)
    userId := "userId_example" // string | the user's UUID (unique)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ValidateSession(context.Background(), orgId, appId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ValidateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | the org&#39;s UUID (unique) | 
**appId** | **string** | the app&#39;s UUID (unique) | 
**userId** | **string** | the user&#39;s UUID (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [tokenHeader](../README.md#tokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

