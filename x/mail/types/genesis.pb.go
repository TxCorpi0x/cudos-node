// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mail/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the mail module's genesis state.
type GenesisState struct {
	Params              Params    `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	MessageList         []Message `protobuf:"bytes,2,rep,name=messageList,proto3" json:"messageList"`
	MessageCount        uint64    `protobuf:"varint,3,opt,name=messageCount,proto3" json:"messageCount,omitempty"`
	MessageSentList     []uint64  `protobuf:"varint,4,rep,packed,name=messageSentList,proto3" json:"messageSentList,omitempty"`
	MessageReceivedList []uint64  `protobuf:"varint,5,rep,packed,name=messageReceivedList,proto3" json:"messageReceivedList,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_526b2b065b85185a, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetMessageList() []Message {
	if m != nil {
		return m.MessageList
	}
	return nil
}

func (m *GenesisState) GetMessageCount() uint64 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

func (m *GenesisState) GetMessageSentList() []uint64 {
	if m != nil {
		return m.MessageSentList
	}
	return nil
}

func (m *GenesisState) GetMessageReceivedList() []uint64 {
	if m != nil {
		return m.MessageReceivedList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cudoventures.cudosnode.mail.GenesisState")
}

func init() { proto.RegisterFile("mail/genesis.proto", fileDescriptor_526b2b065b85185a) }

var fileDescriptor_526b2b065b85185a = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x33, 0x6d, 0xfe, 0x2e, 0xa6, 0x85, 0x1f, 0x47, 0x17, 0x21, 0xc2, 0x18, 0xaa, 0x8b,
	0x6c, 0x9c, 0x91, 0xfa, 0x04, 0xb6, 0x0b, 0x11, 0x2a, 0x48, 0x0a, 0x2e, 0xdc, 0xa5, 0xc9, 0x25,
	0x06, 0x4c, 0x26, 0x64, 0x26, 0x45, 0xdf, 0xc2, 0x37, 0xf0, 0x75, 0xba, 0xec, 0xd2, 0x95, 0x48,
	0xf2, 0x22, 0x92, 0x99, 0x59, 0xa8, 0x48, 0x77, 0x73, 0xcf, 0x7c, 0xe7, 0x1c, 0xee, 0xc5, 0xa4,
	0x88, 0xf3, 0x27, 0x9e, 0x41, 0x09, 0x32, 0x97, 0xac, 0xaa, 0x85, 0x12, 0xe4, 0x38, 0x69, 0x52,
	0xb1, 0x81, 0x52, 0x35, 0x35, 0x48, 0xd6, 0x0f, 0xb2, 0x14, 0x29, 0xb0, 0x1e, 0xf5, 0x8f, 0x32,
	0x91, 0x09, 0xcd, 0xf1, 0xfe, 0x65, 0x2c, 0xfe, 0x81, 0x8e, 0xa9, 0xe2, 0x3a, 0x2e, 0x6c, 0x8a,
	0x6f, 0x92, 0x0b, 0x90, 0x32, 0xce, 0xc0, 0x68, 0xd3, 0xb7, 0x01, 0x9e, 0x5c, 0x9b, 0xae, 0x95,
	0x8a, 0x15, 0x90, 0x2b, 0x3c, 0x32, 0x26, 0x0f, 0x05, 0x28, 0x1c, 0xcf, 0x4e, 0xd9, 0x9e, 0x6e,
	0x76, 0xa7, 0xd1, 0xb9, 0xbb, 0xfd, 0x38, 0x71, 0x22, 0x6b, 0x24, 0x4b, 0x3c, 0xb6, 0x25, 0xcb,
	0x5c, 0x2a, 0x6f, 0x10, 0x0c, 0xc3, 0xf1, 0xec, 0x6c, 0x6f, 0xce, 0xad, 0xe1, 0x6d, 0xd0, 0x77,
	0x3b, 0x99, 0xe2, 0x89, 0x1d, 0x17, 0xa2, 0x29, 0x95, 0x37, 0x0c, 0x50, 0xe8, 0x46, 0x3f, 0x34,
	0x12, 0xe2, 0xff, 0x76, 0x5e, 0x41, 0xa9, 0x74, 0xab, 0x1b, 0x0c, 0x43, 0x37, 0xfa, 0x2d, 0x93,
	0x0b, 0x7c, 0x68, 0xa5, 0x08, 0x12, 0xc8, 0x37, 0x90, 0x6a, 0xfa, 0x9f, 0xa6, 0xff, 0xfa, 0x9a,
	0xdf, 0x6c, 0x5b, 0x8a, 0x76, 0x2d, 0x45, 0x9f, 0x2d, 0x45, 0xaf, 0x1d, 0x75, 0x76, 0x1d, 0x75,
	0xde, 0x3b, 0xea, 0x3c, 0xf0, 0x2c, 0x57, 0x8f, 0xcd, 0x9a, 0x25, 0xa2, 0xe0, 0x8b, 0x26, 0x15,
	0xf7, 0x76, 0x39, 0xae, 0x97, 0x3b, 0xef, 0xb7, 0xe3, 0xcf, 0x5c, 0x1f, 0x5d, 0xbd, 0x54, 0x20,
	0xd7, 0x23, 0x7d, 0xf3, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x1c, 0xad, 0x0e, 0xe3,
	0x01, 0x00, 0x00,
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
	if len(m.MessageReceivedList) > 0 {
		dAtA2 := make([]byte, len(m.MessageReceivedList)*10)
		var j1 int
		for _, num := range m.MessageReceivedList {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGenesis(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MessageSentList) > 0 {
		dAtA4 := make([]byte, len(m.MessageSentList)*10)
		var j3 int
		for _, num := range m.MessageSentList {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGenesis(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x22
	}
	if m.MessageCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MessageCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.MessageList) > 0 {
		for iNdEx := len(m.MessageList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MessageList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.MessageList) > 0 {
		for _, e := range m.MessageList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MessageCount != 0 {
		n += 1 + sovGenesis(uint64(m.MessageCount))
	}
	if len(m.MessageSentList) > 0 {
		l = 0
		for _, e := range m.MessageSentList {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
	}
	if len(m.MessageReceivedList) > 0 {
		l = 0
		for _, e := range m.MessageReceivedList {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageList", wireType)
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
			m.MessageList = append(m.MessageList, Message{})
			if err := m.MessageList[len(m.MessageList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageCount", wireType)
			}
			m.MessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.MessageSentList = append(m.MessageSentList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.MessageSentList) == 0 {
					m.MessageSentList = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.MessageSentList = append(m.MessageSentList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageSentList", wireType)
			}
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.MessageReceivedList = append(m.MessageReceivedList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.MessageReceivedList) == 0 {
					m.MessageReceivedList = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.MessageReceivedList = append(m.MessageReceivedList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageReceivedList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
