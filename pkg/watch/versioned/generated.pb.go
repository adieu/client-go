// +build !appengine

/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/watch/versioned/generated.proto
// DO NOT EDIT!

/*
	Package versioned is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/watch/versioned/generated.proto

	It has these top-level messages:
		Event
*/
package versioned

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *Event) Reset()                    { *m = Event{} }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*Event)(nil), "k8s.io.client-go.pkg.watch.versioned.Event")
}
func (m *Event) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Event) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Type)))
	i += copy(data[i:], m.Type)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Object.Size()))
	n1, err := m.Object.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Object.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Object:` + strings.Replace(strings.Replace(this.Object.String(), "RawExtension", "k8s_io_kubernetes_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Event) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Object.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x1b, 0x99, 0x83, 0x55, 0xf1, 0xd0, 0xd3, 0xe8, 0x21, 0x2b, 0x82, 0x30, 0x18, 0x4b,
	0x50, 0x10, 0x3c, 0x17, 0x76, 0x16, 0xaa, 0x27, 0x6f, 0x6d, 0xf7, 0x99, 0xc5, 0xba, 0xa4, 0xa4,
	0x5f, 0x37, 0x77, 0x11, 0x7f, 0x82, 0x3f, 0xab, 0xc7, 0x1d, 0x3d, 0x0d, 0x5b, 0xff, 0x88, 0x98,
	0x95, 0x09, 0xa3, 0xb7, 0xbc, 0x24, 0xcf, 0xc3, 0xfb, 0xc6, 0xbd, 0xcd, 0xee, 0x0a, 0x26, 0x35,
	0xcf, 0xca, 0x04, 0x8c, 0x02, 0x84, 0x82, 0xe7, 0x99, 0xe0, 0xeb, 0x18, 0xd3, 0x05, 0x5f, 0x81,
	0x29, 0xa4, 0x56, 0x30, 0xe7, 0x02, 0x14, 0x98, 0x18, 0x61, 0xce, 0x72, 0xa3, 0x51, 0x7b, 0x57,
	0x7b, 0x8c, 0xfd, 0x63, 0x2c, 0xcf, 0x04, 0xb3, 0x18, 0x3b, 0x60, 0xfe, 0x54, 0x48, 0x5c, 0x94,
	0x09, 0x4b, 0xf5, 0x92, 0x0b, 0x2d, 0x34, 0xb7, 0x74, 0x52, 0x3e, 0xdb, 0x64, 0x83, 0x3d, 0xed,
	0xad, 0xfe, 0xb4, 0xbb, 0x8c, 0x29, 0x15, 0xca, 0x25, 0x1c, 0x97, 0xf0, 0xaf, 0xbb, 0x9f, 0x97,
	0x28, 0x5f, 0xb9, 0x54, 0x58, 0xa0, 0x39, 0x46, 0x2e, 0xdf, 0xdd, 0xd3, 0xd9, 0x0a, 0x14, 0x7a,
	0x81, 0xdb, 0xc3, 0x4d, 0x0e, 0x43, 0x12, 0x90, 0xf1, 0x20, 0x3c, 0xaf, 0x76, 0x23, 0xa7, 0xd9,
	0x8d, 0x7a, 0x8f, 0x9b, 0x1c, 0x22, 0x7b, 0xe3, 0x3d, 0xb8, 0x7d, 0x9d, 0xbc, 0x40, 0x8a, 0xc3,
	0x93, 0x80, 0x8c, 0xcf, 0x6e, 0x26, 0xac, 0x7b, 0x73, 0xdb, 0x8e, 0x45, 0xf1, 0x7a, 0xf6, 0x86,
	0xa0, 0xfe, 0xa6, 0x87, 0x17, 0xad, 0xb0, 0x7f, 0x6f, 0x15, 0x51, 0xab, 0x0a, 0x27, 0x55, 0x4d,
	0x9d, 0x6d, 0x4d, 0x9d, 0xaf, 0x9a, 0x3a, 0x1f, 0x0d, 0x25, 0x55, 0x43, 0xc9, 0xb6, 0xa1, 0xe4,
	0xbb, 0xa1, 0xe4, 0xf3, 0x87, 0x3a, 0x4f, 0x83, 0xc3, 0xef, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x23, 0x3d, 0x7b, 0x7e, 0x9c, 0x01, 0x00, 0x00,
}
