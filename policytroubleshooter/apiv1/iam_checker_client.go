// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package policytroubleshooter

import (
	"context"
	"math"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	policytroubleshooterpb "google.golang.org/genproto/googleapis/cloud/policytroubleshooter/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newIamCheckerClientHook clientHook

// IamCheckerCallOptions contains the retry settings for each method of IamCheckerClient.
type IamCheckerCallOptions struct {
	TroubleshootIamPolicy []gax.CallOption
}

func defaultIamCheckerClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("policytroubleshooter.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIamCheckerCallOptions() *IamCheckerCallOptions {
	return &IamCheckerCallOptions{
		TroubleshootIamPolicy: []gax.CallOption{},
	}
}

// IamCheckerClient is a client for interacting with Policy Troubleshooter API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type IamCheckerClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	iamCheckerClient policytroubleshooterpb.IamCheckerClient

	// The call options for this service.
	CallOptions *IamCheckerCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewIamCheckerClient creates a new iam checker client.
//
// IAM Policy Troubleshooter service.
//
// This service helps you troubleshoot access issues for Google Cloud resources.
func NewIamCheckerClient(ctx context.Context, opts ...option.ClientOption) (*IamCheckerClient, error) {
	clientOpts := defaultIamCheckerClientOptions()

	if newIamCheckerClientHook != nil {
		hookOpts, err := newIamCheckerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &IamCheckerClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultIamCheckerCallOptions(),

		iamCheckerClient: policytroubleshooterpb.NewIamCheckerClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *IamCheckerClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IamCheckerClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IamCheckerClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// TroubleshootIamPolicy checks whether a member has a specific permission for a specific resource,
// and explains why the member does or does not have that permission.
func (c *IamCheckerClient) TroubleshootIamPolicy(ctx context.Context, req *policytroubleshooterpb.TroubleshootIamPolicyRequest, opts ...gax.CallOption) (*policytroubleshooterpb.TroubleshootIamPolicyResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.TroubleshootIamPolicy[0:len(c.CallOptions.TroubleshootIamPolicy):len(c.CallOptions.TroubleshootIamPolicy)], opts...)
	var resp *policytroubleshooterpb.TroubleshootIamPolicyResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamCheckerClient.TroubleshootIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
