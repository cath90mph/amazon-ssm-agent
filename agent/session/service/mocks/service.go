// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by mockery v1.0.0
package mocks

import (
	log "github.com/aws/amazon-ssm-agent/agent/log"
	service "github.com/aws/amazon-ssm-agent/agent/session/service"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateControlChannel provides a mock function with given fields: _a0, createControlChannelInput, channelId
func (_m *Service) CreateControlChannel(_a0 log.T, createControlChannelInput *service.CreateControlChannelInput, channelId string) (*service.CreateControlChannelOutput, error) {
	ret := _m.Called(_a0, createControlChannelInput, channelId)

	var r0 *service.CreateControlChannelOutput
	if rf, ok := ret.Get(0).(func(log.T, *service.CreateControlChannelInput, string) *service.CreateControlChannelOutput); ok {
		r0 = rf(_a0, createControlChannelInput, channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateControlChannelOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(log.T, *service.CreateControlChannelInput, string) error); ok {
		r1 = rf(_a0, createControlChannelInput, channelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDataChannel provides a mock function with given fields: _a0, createDataChannelInput, sessionId
func (_m *Service) CreateDataChannel(_a0 log.T, createDataChannelInput *service.CreateDataChannelInput, sessionId string) (*service.CreateDataChannelOutput, error) {
	ret := _m.Called(_a0, createDataChannelInput, sessionId)

	var r0 *service.CreateDataChannelOutput
	if rf, ok := ret.Get(0).(func(log.T, *service.CreateDataChannelInput, string) *service.CreateDataChannelOutput); ok {
		r0 = rf(_a0, createDataChannelInput, sessionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateDataChannelOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(log.T, *service.CreateDataChannelInput, string) error); ok {
		r1 = rf(_a0, createDataChannelInput, sessionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegion provides a mock function with given fields:
func (_m *Service) GetRegion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetV4Signer provides a mock function with given fields:
func (_m *Service) GetV4Signer() *v4.Signer {
	ret := _m.Called()

	var r0 *v4.Signer
	if rf, ok := ret.Get(0).(func() *v4.Signer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v4.Signer)
		}
	}

	return r0
}
