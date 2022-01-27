// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// block_time_threshold is a safety to accept only blocktime as much as it and give inflation even if the chain stops
	// or the last blocktime is huge by manipulating blocktime.
	BlockTimeThreshold time.Duration `protobuf:"bytes,2,opt,name=block_time_threshold,json=blockTimeThreshold,proto3,stdduration" json:"block_time_threshold,omitempty" yaml:"block_time_threshold"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_982b7510344c3451, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlockTimeThreshold() time.Duration {
	if m != nil {
		return m.BlockTimeThreshold
	}
	return 0
}

// InflationSchedule defines the start and end time of the inflation period, and the amount of inflation during that
// period.
type InflationSchedule struct {
	// start_time is a start date time of the inflation period
	StartTime time.Time `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// end_time is a start date time of the inflation period
	EndTime time.Time `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	// amount is the amount of inflation during that period.
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *InflationSchedule) Reset()         { *m = InflationSchedule{} }
func (m *InflationSchedule) String() string { return proto.CompactTextString(m) }
func (*InflationSchedule) ProtoMessage()    {}
func (*InflationSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_982b7510344c3451, []int{1}
}
func (m *InflationSchedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationSchedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationSchedule.Merge(m, src)
}
func (m *InflationSchedule) XXX_Size() int {
	return m.Size()
}
func (m *InflationSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_InflationSchedule proto.InternalMessageInfo

func (m *InflationSchedule) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *InflationSchedule) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Params)(nil), "squad.mint.v1beta1.Params")
	proto.RegisterType((*InflationSchedule)(nil), "squad.mint.v1beta1.InflationSchedule")
}

func init() { proto.RegisterFile("squad/mint/v1beta1/mint.proto", fileDescriptor_982b7510344c3451) }

var fileDescriptor_982b7510344c3451 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xcf, 0x01, 0x05, 0xe2, 0x0e, 0xa8, 0xa7, 0x0e, 0x21, 0x55, 0xec, 0xea, 0x06, 0xd4,
	0x81, 0xd8, 0x6a, 0xd9, 0x3a, 0x46, 0x15, 0x28, 0x1b, 0x0a, 0x19, 0x10, 0x4b, 0xe4, 0x8b, 0xdd,
	0xcb, 0xa9, 0x67, 0x3b, 0x9c, 0x7d, 0x88, 0x7c, 0x04, 0xb6, 0x4e, 0xa8, 0x23, 0x9f, 0x06, 0x75,
	0xec, 0x88, 0x18, 0x0e, 0x94, 0x6c, 0x8c, 0xfd, 0x04, 0xc8, 0xf6, 0x9d, 0x90, 0x68, 0xa4, 0x4e,
	0x77, 0x7e, 0xff, 0xf7, 0xff, 0xfd, 0x9f, 0xad, 0x07, 0x87, 0xe6, 0x63, 0xc5, 0x38, 0x95, 0xb9,
	0xb2, 0xf4, 0xd3, 0x49, 0x2a, 0x2c, 0x3b, 0xf1, 0x07, 0xb2, 0x2a, 0xb5, 0xd5, 0x71, 0xec, 0x65,
	0xe2, 0x2b, 0x8d, 0x3c, 0x38, 0xc8, 0x74, 0xa6, 0xbd, 0x4c, 0xdd, 0x5f, 0xe8, 0x1c, 0xa0, 0x4c,
	0xeb, 0xac, 0x10, 0xd4, 0x9f, 0xd2, 0xea, 0x82, 0xf2, 0xaa, 0x64, 0x36, 0xd7, 0xaa, 0xd1, 0xf1,
	0xff, 0xba, 0xcd, 0xa5, 0x30, 0x96, 0xc9, 0x55, 0x68, 0x48, 0xbe, 0x03, 0xd8, 0x7d, 0xcb, 0x4a,
	0x26, 0x4d, 0x3c, 0x84, 0xd0, 0x25, 0xce, 0xb9, 0x50, 0x5a, 0xf6, 0xc1, 0x11, 0x38, 0xee, 0x4d,
	0x7b, 0xae, 0x72, 0xee, 0x0a, 0xf1, 0x57, 0x00, 0x0f, 0xd2, 0x42, 0x2f, 0x2e, 0xe7, 0x8e, 0x31,
	0xb7, 0xcb, 0x52, 0x98, 0xa5, 0x2e, 0x78, 0xbf, 0x73, 0x04, 0x8e, 0xf7, 0x4e, 0x9f, 0x93, 0x10,
	0x45, 0xda, 0x28, 0x72, 0xde, 0x8c, 0x32, 0x9e, 0xdc, 0xd4, 0x38, 0xfa, 0x53, 0x63, 0xb4, 0xcb,
	0xfe, 0x52, 0xcb, 0xdc, 0x0a, 0xb9, 0xb2, 0xeb, 0xbb, 0x1a, 0x1f, 0xae, 0x99, 0x2c, 0xce, 0x92,
	0x5d, 0x7d, 0xc9, 0xf5, 0x2f, 0x0c, 0xa6, 0xb1, 0x97, 0x66, 0xb9, 0x14, 0xb3, 0x56, 0x38, 0x7b,
	0x7c, 0xfd, 0x0d, 0x47, 0xc9, 0x97, 0x0e, 0xdc, 0x9f, 0xa8, 0x8b, 0xc2, 0x47, 0xbe, 0x5b, 0x2c,
	0x05, 0xaf, 0x0a, 0x11, 0xbf, 0x87, 0xd0, 0x58, 0x56, 0x5a, 0x0f, 0xf3, 0x77, 0xda, 0x3b, 0x1d,
	0xdc, 0x9b, 0x74, 0xd6, 0x3e, 0xca, 0x78, 0xe8, 0x46, 0xbd, 0xab, 0xf1, 0x7e, 0x18, 0xe4, 0x9f,
	0x37, 0xb9, 0x72, 0xf1, 0x3d, 0x5f, 0x70, 0xed, 0xf1, 0x14, 0x3e, 0x15, 0x8a, 0x07, 0x6e, 0xe7,
	0x41, 0xee, 0x61, 0xc3, 0x7d, 0x16, 0xb8, 0xad, 0x33, 0x50, 0x9f, 0x08, 0xc5, 0x3d, 0xf3, 0x35,
	0xec, 0x32, 0xa9, 0x2b, 0x65, 0xfb, 0x8f, 0xdc, 0xeb, 0x8f, 0x89, 0x73, 0xfd, 0xac, 0xf1, 0x8b,
	0x2c, 0xb7, 0xcb, 0x2a, 0x25, 0x0b, 0x2d, 0xe9, 0x42, 0x1b, 0xa9, 0x4d, 0xf3, 0x19, 0x19, 0x7e,
	0x49, 0xed, 0x7a, 0x25, 0x0c, 0x99, 0x28, 0x3b, 0x6d, 0xdc, 0xe3, 0x37, 0x37, 0x1b, 0x04, 0x6e,
	0x37, 0x08, 0xfc, 0xde, 0x20, 0x70, 0xb5, 0x45, 0xd1, 0xed, 0x16, 0x45, 0x3f, 0xb6, 0x28, 0xfa,
	0x30, 0xba, 0x47, 0x72, 0x9b, 0x36, 0x2a, 0x58, 0x6a, 0x68, 0xd8, 0xc9, 0xcf, 0x61, 0x2b, 0x3d,
	0x34, 0xed, 0xfa, 0xab, 0xbc, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x62, 0xb5, 0x47, 0xb0,
	0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.BlockTimeThreshold, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.BlockTimeThreshold):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMint(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InflationSchedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationSchedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationSchedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintMint(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintMint(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.BlockTimeThreshold)
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *InflationSchedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovMint(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovMint(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimeThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.BlockTimeThreshold, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *InflationSchedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: InflationSchedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationSchedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
