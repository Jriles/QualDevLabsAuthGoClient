# Go API client for qualdevlabs_auth_go_client

API for the QualDevLabs Auth library.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./qualdevlabs_auth_go_client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://qualdevlabsauth.uc.r.appspot.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateUser**](docs/DefaultApi.md#createuser) | **Post** /orgs/{orgId}/apps/{appId}/users | Creates a new user.
*DefaultApi* | [**CreateUserSession**](docs/DefaultApi.md#createusersession) | **Post** /orgs/{orgId}/apps/{appId}/sessions | Login user
*DefaultApi* | [**DeleteUser**](docs/DefaultApi.md#deleteuser) | **Delete** /orgs/{orgId}/apps/{appId}/users/{userId} | Delete a user
*DefaultApi* | [**DeleteUserSession**](docs/DefaultApi.md#deleteusersession) | **Delete** /orgs/{orgId}/apps/{appId}/sessions/{userId} | Logout user
*DefaultApi* | [**ForgotPasswordReset**](docs/DefaultApi.md#forgotpasswordreset) | **Patch** /orgs/{orgId}/apps/{appId}/password_reset | resets an unathenticated user&#39;s password using a token sent to their email
*DefaultApi* | [**RequestForgottenUsername**](docs/DefaultApi.md#requestforgottenusername) | **Post** /orgs/{orgId}/apps/{appId}/forgot_username | requests an email get sent to the user with their username
*DefaultApi* | [**RequestPasswordResetToken**](docs/DefaultApi.md#requestpasswordresettoken) | **Post** /orgs/{orgId}/apps/{appId}/password_reset_tokens | requests a password reset token for users who have forgotten their password
*DefaultApi* | [**UpdateImportantUserDetails**](docs/DefaultApi.md#updateimportantuserdetails) | **Patch** /orgs/{orgId}/apps/{appId}/users/{userId} | Updates a user&#39;s important details
*DefaultApi* | [**ValidateSession**](docs/DefaultApi.md#validatesession) | **Get** /orgs/{orgId}/apps/{appId}/sessions/{userId} | Verifies if session is valid


## Documentation For Models

 - [LoginSchema](docs/LoginSchema.md)
 - [PasswordResetSchema](docs/PasswordResetSchema.md)
 - [SessionTokenAndUserId](docs/SessionTokenAndUserId.md)
 - [UserEmailSchema](docs/UserEmailSchema.md)
 - [UserEmailSchema1](docs/UserEmailSchema1.md)
 - [UserSchema](docs/UserSchema.md)
 - [UserUpdatesSchema](docs/UserUpdatesSchema.md)


## Documentation For Authorization



### apiKeyHeader

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-api-key and passed in as the auth context for each request.


### passwordInHeader

- **Type**: API key
- **API key parameter name**: password
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: password and passed in as the auth context for each request.


### tokenHeader

- **Type**: API key
- **API key parameter name**: token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: token and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



