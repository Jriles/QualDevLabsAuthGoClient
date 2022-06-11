# UserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**RepeatPassword** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewUserSchema

`func NewUserSchema(username string, password string, repeatPassword string, email string, ) *UserSchema`

NewUserSchema instantiates a new UserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSchemaWithDefaults

`func NewUserSchemaWithDefaults() *UserSchema`

NewUserSchemaWithDefaults instantiates a new UserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSchema) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserSchema) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRepeatPassword

`func (o *UserSchema) GetRepeatPassword() string`

GetRepeatPassword returns the RepeatPassword field if non-nil, zero value otherwise.

### GetRepeatPasswordOk

`func (o *UserSchema) GetRepeatPasswordOk() (*string, bool)`

GetRepeatPasswordOk returns a tuple with the RepeatPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatPassword

`func (o *UserSchema) SetRepeatPassword(v string)`

SetRepeatPassword sets RepeatPassword field to given value.


### GetEmail

`func (o *UserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


