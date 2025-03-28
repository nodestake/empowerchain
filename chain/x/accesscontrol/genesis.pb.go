// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/accesscontrol/genesis.proto

package accesscontrol

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

// GenesisState defines the accesscontrol module's genesis state.
type GenesisState struct {
	PermStores []ModulePermStore `protobuf:"bytes,1,rep,name=perm_stores,json=permStores,proto3" json:"perm_stores"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_12aa109cac8f21f4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetPermStores() []ModulePermStore {
	if m != nil {
		return m.PermStores
	}
	return nil
}

// All accesses for a given module
type ModulePermStore struct {
	// Name - will be used as a name for a PermStore
	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	// List of accesses
	Accesses []Access `protobuf:"bytes,2,rep,name=accesses,proto3" json:"accesses"`
}

func (m *ModulePermStore) Reset()         { *m = ModulePermStore{} }
func (m *ModulePermStore) String() string { return proto.CompactTextString(m) }
func (*ModulePermStore) ProtoMessage()    {}
func (*ModulePermStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_12aa109cac8f21f4, []int{1}
}
func (m *ModulePermStore) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModulePermStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModulePermStore.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModulePermStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModulePermStore.Merge(m, src)
}
func (m *ModulePermStore) XXX_Size() int {
	return m.Size()
}
func (m *ModulePermStore) XXX_DiscardUnknown() {
	xxx_messageInfo_ModulePermStore.DiscardUnknown(m)
}

var xxx_messageInfo_ModulePermStore proto.InternalMessageInfo

func (m *ModulePermStore) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *ModulePermStore) GetAccesses() []Access {
	if m != nil {
		return m.Accesses
	}
	return nil
}

// Single access consisting of permissioned address and msgType
// of the message it can invoke
type Access struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MsgType string `protobuf:"bytes,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
}

func (m *Access) Reset()         { *m = Access{} }
func (m *Access) String() string { return proto.CompactTextString(m) }
func (*Access) ProtoMessage()    {}
func (*Access) Descriptor() ([]byte, []int) {
	return fileDescriptor_12aa109cac8f21f4, []int{2}
}
func (m *Access) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Access) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Access.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Access) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Access.Merge(m, src)
}
func (m *Access) XXX_Size() int {
	return m.Size()
}
func (m *Access) XXX_DiscardUnknown() {
	xxx_messageInfo_Access.DiscardUnknown(m)
}

var xxx_messageInfo_Access proto.InternalMessageInfo

func (m *Access) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Access) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "empowerchain.accesscontrol.GenesisState")
	proto.RegisterType((*ModulePermStore)(nil), "empowerchain.accesscontrol.ModulePermStore")
	proto.RegisterType((*Access)(nil), "empowerchain.accesscontrol.Access")
}

func init() {
	proto.RegisterFile("empowerchain/accesscontrol/genesis.proto", fileDescriptor_12aa109cac8f21f4)
}

var fileDescriptor_12aa109cac8f21f4 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4b, 0x02, 0x41,
	0x18, 0xdd, 0xb5, 0x50, 0x1b, 0x83, 0x60, 0xf1, 0xb0, 0x7a, 0x58, 0x65, 0x4f, 0x42, 0xb4, 0x0b,
	0x06, 0xdd, 0x95, 0xa2, 0x53, 0x21, 0x6b, 0x10, 0x74, 0x59, 0xc6, 0xf1, 0x63, 0x5c, 0x70, 0x76,
	0x96, 0xf9, 0x46, 0xd2, 0x7f, 0xd1, 0x8f, 0xe9, 0x47, 0x78, 0x94, 0x4e, 0x9d, 0x22, 0xf4, 0x8f,
	0x84, 0x33, 0x1a, 0x19, 0xd4, 0x6d, 0xbe, 0xf7, 0xbd, 0x37, 0xef, 0xf1, 0x3d, 0xd2, 0x01, 0x51,
	0xc8, 0x67, 0x50, 0x6c, 0x42, 0xb3, 0x3c, 0xa6, 0x8c, 0x01, 0x22, 0x93, 0xb9, 0x56, 0x72, 0x1a,
	0x73, 0xc8, 0x01, 0x33, 0x8c, 0x0a, 0x25, 0xb5, 0xf4, 0x9a, 0x3f, 0x99, 0xd1, 0x01, 0xb3, 0xd9,
	0x60, 0x12, 0x85, 0xc4, 0xd4, 0x30, 0x63, 0x3b, 0x58, 0x59, 0xb3, 0xce, 0x25, 0x97, 0x16, 0xdf,
	0xbe, 0x2c, 0x1a, 0x8e, 0xc8, 0xe9, 0xad, 0xfd, 0x7d, 0xa8, 0xa9, 0x06, 0x2f, 0x21, 0xb5, 0x02,
	0x94, 0x48, 0x51, 0x4b, 0x05, 0xe8, 0xbb, 0xed, 0xa3, 0x4e, 0xad, 0x7b, 0x1e, 0xfd, 0x6d, 0x19,
	0xdd, 0xc9, 0xf1, 0x6c, 0x0a, 0x03, 0x50, 0x62, 0xb8, 0xd5, 0xf4, 0x8f, 0x97, 0x1f, 0x2d, 0x27,
	0x21, 0xc5, 0x1e, 0xc0, 0x70, 0x4e, 0xce, 0x7e, 0x91, 0xbc, 0x16, 0xa9, 0x09, 0x03, 0xa5, 0x39,
	0x15, 0xe0, 0xbb, 0x6d, 0xb7, 0x73, 0x92, 0x10, 0x0b, 0xdd, 0x53, 0x01, 0xde, 0x35, 0xa9, 0x5a,
	0x1b, 0x40, 0xbf, 0x64, 0x42, 0x84, 0xff, 0x85, 0xe8, 0x99, 0x69, 0xe7, 0xfd, 0xad, 0x0c, 0x1f,
	0x49, 0xd9, 0x6e, 0xbc, 0x2e, 0xa9, 0xd0, 0xf1, 0x58, 0x01, 0xa2, 0x35, 0xeb, 0xfb, 0x6f, 0xaf,
	0x17, 0xf5, 0xdd, 0x81, 0x7a, 0x76, 0x33, 0xd4, 0x2a, 0xcb, 0x79, 0xb2, 0x27, 0x7a, 0x0d, 0x52,
	0x15, 0xc8, 0x53, 0xbd, 0x28, 0xc0, 0x2f, 0x99, 0x84, 0x15, 0x81, 0xfc, 0x61, 0x51, 0x40, 0x7f,
	0xb0, 0x5c, 0x07, 0xee, 0x6a, 0x1d, 0xb8, 0x9f, 0xeb, 0xc0, 0x7d, 0xd9, 0x04, 0xce, 0x6a, 0x13,
	0x38, 0xef, 0x9b, 0xc0, 0x79, 0xba, 0xe2, 0x99, 0x9e, 0xcc, 0x46, 0x11, 0x93, 0x22, 0xbe, 0xb1,
	0x81, 0x07, 0x53, 0x8a, 0x3a, 0x63, 0xf1, 0x41, 0xc3, 0xf3, 0xc3, 0x8e, 0x47, 0x65, 0xd3, 0xc7,
	0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x4b, 0xd5, 0x6c, 0x08, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PermStores) > 0 {
		for iNdEx := len(m.PermStores) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PermStores[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ModulePermStore) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModulePermStore) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModulePermStore) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accesses) > 0 {
		for iNdEx := len(m.Accesses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accesses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Access) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Access) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Access) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgType) > 0 {
		i -= len(m.MsgType)
		copy(dAtA[i:], m.MsgType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MsgType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PermStores) > 0 {
		for _, e := range m.PermStores {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ModulePermStore) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Accesses) > 0 {
		for _, e := range m.Accesses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Access) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.MsgType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermStores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PermStores = append(m.PermStores, ModulePermStore{})
			if err := m.PermStores[len(m.PermStores)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ModulePermStore) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ModulePermStore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModulePermStore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accesses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accesses = append(m.Accesses, Access{})
			if err := m.Accesses[len(m.Accesses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Access) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Access: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Access: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
