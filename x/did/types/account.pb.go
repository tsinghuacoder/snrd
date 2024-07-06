// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/v1/account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DIDDocument defines a DID document
type DIDDocument struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VerificationMethods  []*VerificationMethod `protobuf:"bytes,2,rep,name=verification_methods,json=verificationMethods,proto3" json:"verification_methods,omitempty"`
	Authentication       []string              `protobuf:"bytes,4,rep,name=authentication,proto3" json:"authentication,omitempty"`
	AssertionMethod      []string              `protobuf:"bytes,5,rep,name=assertion_method,json=assertionMethod,proto3" json:"assertion_method,omitempty"`
	CapabilityDelegation []string              `protobuf:"bytes,7,rep,name=capability_delegation,json=capabilityDelegation,proto3" json:"capability_delegation,omitempty"`
	CapabilityInvocation []string              `protobuf:"bytes,8,rep,name=capability_invocation,json=capabilityInvocation,proto3" json:"capability_invocation,omitempty"`
}

func (m *DIDDocument) Reset()         { *m = DIDDocument{} }
func (m *DIDDocument) String() string { return proto.CompactTextString(m) }
func (*DIDDocument) ProtoMessage()    {}
func (*DIDDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1c7ebf7d0ca3c72, []int{0}
}
func (m *DIDDocument) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DIDDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DIDDocument.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DIDDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DIDDocument.Merge(m, src)
}
func (m *DIDDocument) XXX_Size() int {
	return m.Size()
}
func (m *DIDDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_DIDDocument.DiscardUnknown(m)
}

var xxx_messageInfo_DIDDocument proto.InternalMessageInfo

func (m *DIDDocument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DIDDocument) GetVerificationMethods() []*VerificationMethod {
	if m != nil {
		return m.VerificationMethods
	}
	return nil
}

func (m *DIDDocument) GetAuthentication() []string {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *DIDDocument) GetAssertionMethod() []string {
	if m != nil {
		return m.AssertionMethod
	}
	return nil
}

func (m *DIDDocument) GetCapabilityDelegation() []string {
	if m != nil {
		return m.CapabilityDelegation
	}
	return nil
}

func (m *DIDDocument) GetCapabilityInvocation() []string {
	if m != nil {
		return m.CapabilityInvocation
	}
	return nil
}

// VerificationMethod defines a verification method
type VerificationMethod struct {
	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Controller string     `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	PublicKey  *PublicKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (m *VerificationMethod) Reset()         { *m = VerificationMethod{} }
func (m *VerificationMethod) String() string { return proto.CompactTextString(m) }
func (*VerificationMethod) ProtoMessage()    {}
func (*VerificationMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1c7ebf7d0ca3c72, []int{1}
}
func (m *VerificationMethod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerificationMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerificationMethod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerificationMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationMethod.Merge(m, src)
}
func (m *VerificationMethod) XXX_Size() int {
	return m.Size()
}
func (m *VerificationMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationMethod.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationMethod proto.InternalMessageInfo

func (m *VerificationMethod) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VerificationMethod) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *VerificationMethod) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*DIDDocument)(nil), "did.v1.DIDDocument")
	proto.RegisterType((*VerificationMethod)(nil), "did.v1.VerificationMethod")
}

func init() { proto.RegisterFile("did/v1/account.proto", fileDescriptor_f1c7ebf7d0ca3c72) }

var fileDescriptor_f1c7ebf7d0ca3c72 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xb1, 0x8e, 0x9b, 0x40,
	0x10, 0x86, 0x0d, 0x4e, 0x9c, 0x78, 0x2d, 0x39, 0xc9, 0x86, 0x48, 0xc4, 0x05, 0x42, 0x2e, 0x22,
	0xd2, 0x40, 0x6c, 0xb7, 0xa9, 0x22, 0x1a, 0x2b, 0xb2, 0x14, 0x51, 0xa4, 0x48, 0x83, 0x60, 0x77,
	0x03, 0xab, 0x83, 0x1d, 0xc4, 0x2e, 0xdc, 0xf1, 0x16, 0xf7, 0x1a, 0xf7, 0x26, 0x57, 0xba, 0xbc,
	0xf2, 0x64, 0xbf, 0xc8, 0xc9, 0x80, 0x7d, 0x96, 0xdd, 0x0d, 0xdf, 0xff, 0xcf, 0xcf, 0xec, 0x0c,
	0x32, 0x28, 0xa7, 0x5e, 0xbd, 0xf0, 0x22, 0x42, 0xa0, 0x12, 0xca, 0x2d, 0x4a, 0x50, 0x80, 0x47,
	0x94, 0x53, 0xb7, 0x5e, 0xcc, 0xbe, 0x12, 0x90, 0x39, 0xc8, 0xb0, 0xa5, 0x5e, 0xf7, 0xd1, 0x59,
	0x66, 0x46, 0x02, 0x09, 0x74, 0xfc, 0x50, 0x1d, 0x69, 0x1f, 0x97, 0x30, 0xc1, 0x24, 0xef, 0xbd,
	0xf3, 0x07, 0x1d, 0x4d, 0xfc, 0xb5, 0xef, 0x03, 0xa9, 0x72, 0x26, 0x14, 0x9e, 0x22, 0x9d, 0x53,
	0x53, 0xb3, 0x35, 0x67, 0x1c, 0xe8, 0x9c, 0xe2, 0x0d, 0x32, 0x6a, 0x56, 0xf2, 0xff, 0x9c, 0x44,
	0x8a, 0x83, 0x08, 0x73, 0xa6, 0x52, 0xa0, 0xd2, 0xd4, 0xed, 0xa1, 0x33, 0x59, 0xce, 0xdc, 0x6e,
	0x1a, 0xf7, 0xef, 0x99, 0x67, 0xd3, 0x5a, 0x82, 0xcf, 0xf5, 0x15, 0x93, 0xf8, 0x1b, 0x9a, 0x46,
	0x95, 0x4a, 0x99, 0x50, 0xbd, 0x60, 0xbe, 0xb1, 0x87, 0xce, 0x38, 0xb8, 0xa0, 0xf8, 0x3b, 0xfa,
	0x18, 0x49, 0xc9, 0xca, 0xb3, 0x7f, 0x9a, 0x6f, 0x5b, 0xe7, 0x87, 0x13, 0xef, 0x32, 0xf1, 0x0a,
	0x7d, 0x21, 0x51, 0x11, 0xc5, 0x3c, 0xe3, 0xaa, 0x09, 0x29, 0xcb, 0x58, 0xd2, 0x25, 0xbf, 0x6b,
	0xfd, 0xc6, 0xab, 0xe8, 0x9f, 0xb4, 0x8b, 0x26, 0x2e, 0x6a, 0xe8, 0xc7, 0x79, 0x7f, 0xd9, 0xb4,
	0x3e, 0x69, 0xf3, 0x1a, 0xe1, 0xeb, 0x77, 0x5e, 0x6d, 0xcc, 0x42, 0x88, 0x80, 0x50, 0x25, 0x64,
	0x19, 0x2b, 0x4d, 0xbd, 0xe5, 0x67, 0x04, 0xff, 0x40, 0xa8, 0xa8, 0xe2, 0x8c, 0x93, 0xf0, 0x86,
	0x35, 0xe6, 0xd0, 0xd6, 0x9c, 0xc9, 0xf2, 0xd3, 0x71, 0x8f, 0x7f, 0x5a, 0xe5, 0x37, 0x6b, 0x82,
	0x71, 0x71, 0x2c, 0x7f, 0xfd, 0x7c, 0xdc, 0x59, 0xda, 0x76, 0x67, 0x69, 0xcf, 0x3b, 0x4b, 0xbb,
	0xdf, 0x5b, 0x83, 0xed, 0xde, 0x1a, 0x3c, 0xed, 0xad, 0xc1, 0xbf, 0x79, 0xc2, 0x55, 0x5a, 0xc5,
	0x2e, 0x81, 0xdc, 0x03, 0x21, 0x41, 0x94, 0x5e, 0x7a, 0x1b, 0x35, 0xde, 0x9d, 0x77, 0x38, 0xb6,
	0x6a, 0x0a, 0x26, 0xe3, 0x51, 0x7b, 0xe8, 0xd5, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x7a,
	0x87, 0xb1, 0x4f, 0x02, 0x00, 0x00,
}

func (m *DIDDocument) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DIDDocument) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DIDDocument) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CapabilityInvocation) > 0 {
		for iNdEx := len(m.CapabilityInvocation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityInvocation[iNdEx])
			copy(dAtA[i:], m.CapabilityInvocation[iNdEx])
			i = encodeVarintAccount(dAtA, i, uint64(len(m.CapabilityInvocation[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for iNdEx := len(m.CapabilityDelegation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityDelegation[iNdEx])
			copy(dAtA[i:], m.CapabilityDelegation[iNdEx])
			i = encodeVarintAccount(dAtA, i, uint64(len(m.CapabilityDelegation[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AssertionMethod) > 0 {
		for iNdEx := len(m.AssertionMethod) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AssertionMethod[iNdEx])
			copy(dAtA[i:], m.AssertionMethod[iNdEx])
			i = encodeVarintAccount(dAtA, i, uint64(len(m.AssertionMethod[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authentication[iNdEx])
			copy(dAtA[i:], m.Authentication[iNdEx])
			i = encodeVarintAccount(dAtA, i, uint64(len(m.Authentication[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.VerificationMethods) > 0 {
		for iNdEx := len(m.VerificationMethods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerificationMethods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VerificationMethod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerificationMethod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerificationMethod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DIDDocument) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if len(m.VerificationMethods) > 0 {
		for _, e := range m.VerificationMethods {
			l = e.Size()
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	if len(m.Authentication) > 0 {
		for _, s := range m.Authentication {
			l = len(s)
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	if len(m.AssertionMethod) > 0 {
		for _, s := range m.AssertionMethod {
			l = len(s)
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for _, s := range m.CapabilityDelegation {
			l = len(s)
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for _, s := range m.CapabilityInvocation {
			l = len(s)
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	return n
}

func (m *VerificationMethod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DIDDocument) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DIDDocument: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DIDDocument: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethods = append(m.VerificationMethods, &VerificationMethod{})
			if err := m.VerificationMethods[len(m.VerificationMethods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentication = append(m.Authentication, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionMethod = append(m.AssertionMethod, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityDelegation = append(m.CapabilityDelegation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityInvocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityInvocation = append(m.CapabilityInvocation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerificationMethod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VerificationMethod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerificationMethod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)
