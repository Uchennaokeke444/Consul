// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package meshv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using MeshGateway within kubernetes types, where deepcopy-gen is used.
func (in *MeshGateway) DeepCopyInto(out *MeshGateway) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshGateway. Required by controller-gen.
func (in *MeshGateway) DeepCopy() *MeshGateway {
	if in == nil {
		return nil
	}
	out := new(MeshGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MeshGateway. Required by controller-gen.
func (in *MeshGateway) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MeshGatewayListener within kubernetes types, where deepcopy-gen is used.
func (in *MeshGatewayListener) DeepCopyInto(out *MeshGatewayListener) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshGatewayListener. Required by controller-gen.
func (in *MeshGatewayListener) DeepCopy() *MeshGatewayListener {
	if in == nil {
		return nil
	}
	out := new(MeshGatewayListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MeshGatewayListener. Required by controller-gen.
func (in *MeshGatewayListener) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}