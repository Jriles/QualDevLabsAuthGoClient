# PasswordResetSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**RepeatPassword** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewPasswordResetSchema

`func NewPasswordResetSchema(password string, repeatPassword string, email string, ) *PasswordResetSchema`

NewPasswordResetSchema instantiates a new PasswordResetSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordResetSchemaWithDefaults

`func NewPasswordResetSchemaWithDefaults() *PasswordResetSchema`

NewPasswordResetSchemaWithDefaults instantiates a new PasswordResetSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PasswordResetSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordResetSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordResetSchema) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRepeatPassword

`func (o *PasswordResetSchema) GetRepeatPassword() string`

GetRepeatPassword returns the RepeatPassword field if non-nil, zero value otherwise.

### GetRepeatPasswordOk

`func (o *PasswordResetSchema) GetRepeatPasswordOk() (*string, bool)`

GetRepeatPasswordOk returns a tuple with the RepeatPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatPassword

`func (o *PasswordResetSchema) SetRepeatPassword(v string)`

SetRepeatPassword sets RepeatPassword field to given value.


### GetEmail

`func (o *PasswordResetSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PasswordResetSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PasswordResetSchema) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


