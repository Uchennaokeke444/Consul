// Code generated by protoc-gen-grpc-inmem. DO NOT EDIT.

package pbacl

import (
	"context"

	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ ACLServiceClient = CloningACLServiceClient{}

// IsCloningACLServiceClient is an interface that can be used to detect
// that a ACLServiceClient is using the in-memory transport and has already
// been wrapped with a with a CloningACLServiceClient.
type IsCloningACLServiceClient interface {
	IsCloningACLServiceClient() bool
}

// CloningACLServiceClient implements the ACLServiceClient interface by wrapping
// another implementation and copying all protobuf messages that pass through the client.
// This is mainly useful to wrap the an in-process client to insulate users of that
// client from having to care about potential immutability of data they receive or having
// the server implementation mutate their internal memory.
type CloningACLServiceClient struct {
	ACLServiceClient
}

func NewCloningACLServiceClient(client ACLServiceClient) ACLServiceClient {
	if cloner, ok := client.(IsCloningACLServiceClient); ok && cloner.IsCloningACLServiceClient() {
		// prevent a double clone if the underlying client is already the cloning client.
		return client
	}

	return CloningACLServiceClient{
		ACLServiceClient: client,
	}
}

// IsCloningACLServiceClient implements the IsCloningACLServiceClient interface. This
// is only used to detect wrapped clients that would be double cloning data and prevent that.
func (c CloningACLServiceClient) IsCloningACLServiceClient() bool {
	return true
}

func (c CloningACLServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	in = proto.Clone(in).(*LoginRequest)

	out, err := c.ACLServiceClient.Login(ctx, in)
	if err != nil {
		return nil, err
	}

	return proto.Clone(out).(*LoginResponse), nil
}

func (c CloningACLServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	in = proto.Clone(in).(*LogoutRequest)

	out, err := c.ACLServiceClient.Logout(ctx, in)
	if err != nil {
		return nil, err
	}

	return proto.Clone(out).(*LogoutResponse), nil
}