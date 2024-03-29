/*
QualDevLabsAuth

API for the QualDevLabs Auth library.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qualdevlabs_auth_go_client

import (
	"encoding/json"
)

// LoginSchema struct for LoginSchema
type LoginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RememberMe bool `json:"rememberMe"`
}

// NewLoginSchema instantiates a new LoginSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginSchema(username string, password string, rememberMe bool) *LoginSchema {
	this := LoginSchema{}
	this.Username = username
	this.Password = password
	this.RememberMe = rememberMe
	return &this
}

// NewLoginSchemaWithDefaults instantiates a new LoginSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginSchemaWithDefaults() *LoginSchema {
	this := LoginSchema{}
	return &this
}

// GetUsername returns the Username field value
func (o *LoginSchema) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *LoginSchema) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *LoginSchema) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *LoginSchema) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoginSchema) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoginSchema) SetPassword(v string) {
	o.Password = v
}

// GetRememberMe returns the RememberMe field value
func (o *LoginSchema) GetRememberMe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RememberMe
}

// GetRememberMeOk returns a tuple with the RememberMe field value
// and a boolean to check if the value has been set.
func (o *LoginSchema) GetRememberMeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RememberMe, true
}

// SetRememberMe sets field value
func (o *LoginSchema) SetRememberMe(v bool) {
	o.RememberMe = v
}

func (o LoginSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["rememberMe"] = o.RememberMe
	}
	return json.Marshal(toSerialize)
}

type NullableLoginSchema struct {
	value *LoginSchema
	isSet bool
}

func (v NullableLoginSchema) Get() *LoginSchema {
	return v.value
}

func (v *NullableLoginSchema) Set(val *LoginSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginSchema(val *LoginSchema) *NullableLoginSchema {
	return &NullableLoginSchema{value: val, isSet: true}
}

func (v NullableLoginSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


