// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v2/health_check.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/health_check/v2/health_check.proto

	It has these top-level messages:
		HealthCheck
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
import envoy_type1 "github.com/envoyproxy/go-control-plane/envoy/type"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheck struct {
	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *google_protobuf1.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode" json:"pass_through_mode,omitempty"`
	// Specifies the incoming HTTP endpoint that should be considered the
	// health check endpoint. For example */healthcheck*.
	// Note that this field is deprecated in favor of
	// :ref:`headers <envoy_api_field_config.filter.http.health_check.v2.HealthCheck.headers>`.
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *time.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,stdduration" json:"cache_time,omitempty"`
	// If operating in non-pass-through mode, specifies a set of upstream cluster
	// names and the minimum percentage of servers in each of those clusters that
	// must be healthy in order for the filter to return a 200.
	ClusterMinHealthyPercentages map[string]*envoy_type1.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// Specifies a set of health check request headers to match on. The health check filter will
	// check a request’s headers against all the specified headers. To specify the health check
	// endpoint, set the ``:path`` header to match on. Note that if the
	// :ref:`endpoint <envoy_api_field_config.filter.http.health_check.v2.HealthCheck.endpoint>`
	// field is set, it will overwrite any ``:path`` header to match.
	Headers []*envoy_api_v2_route.HeaderMatcher `protobuf:"bytes,5,rep,name=headers" json:"headers,omitempty"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptorHealthCheck, []int{0} }

func (m *HealthCheck) GetPassThroughMode() *google_protobuf1.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *HealthCheck) GetCacheTime() *time.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*envoy_type1.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*envoy_api_v2_route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck")
}
func (m *HealthCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PassThroughMode != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(m.PassThroughMode.Size()))
		n1, err := m.PassThroughMode.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Endpoint) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(len(m.Endpoint)))
		i += copy(dAtA[i:], m.Endpoint)
	}
	if m.CacheTime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(types.SizeOfStdDuration(*m.CacheTime)))
		n2, err := types.StdDurationMarshalTo(*m.CacheTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.ClusterMinHealthyPercentages) > 0 {
		for k, _ := range m.ClusterMinHealthyPercentages {
			dAtA[i] = 0x22
			i++
			v := m.ClusterMinHealthyPercentages[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovHealthCheck(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovHealthCheck(uint64(len(k))) + msgSize
			i = encodeVarintHealthCheck(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintHealthCheck(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintHealthCheck(dAtA, i, uint64(v.Size()))
				n3, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n3
			}
		}
	}
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintHealthCheck(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintHealthCheck(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HealthCheck) Size() (n int) {
	var l int
	_ = l
	if m.PassThroughMode != nil {
		l = m.PassThroughMode.Size()
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	if m.CacheTime != nil {
		l = types.SizeOfStdDuration(*m.CacheTime)
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	if len(m.ClusterMinHealthyPercentages) > 0 {
		for k, v := range m.ClusterMinHealthyPercentages {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovHealthCheck(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovHealthCheck(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovHealthCheck(uint64(mapEntrySize))
		}
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovHealthCheck(uint64(l))
		}
	}
	return n
}

func sovHealthCheck(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealthCheck(x uint64) (n int) {
	return sovHealthCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PassThroughMode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PassThroughMode == nil {
				m.PassThroughMode = &google_protobuf1.BoolValue{}
			}
			if err := m.PassThroughMode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CacheTime == nil {
				m.CacheTime = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.CacheTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterMinHealthyPercentages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterMinHealthyPercentages == nil {
				m.ClusterMinHealthyPercentages = make(map[string]*envoy_type1.Percent)
			}
			var mapkey string
			var mapvalue *envoy_type1.Percent
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHealthCheck
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHealthCheck
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthHealthCheck
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHealthCheck
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthHealthCheck
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthHealthCheck
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &envoy_type1.Percent{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipHealthCheck(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthHealthCheck
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ClusterMinHealthyPercentages[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &envoy_api_v2_route.HeaderMatcher{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
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
func skipHealthCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealthCheck
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
					return 0, ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowHealthCheck
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHealthCheck
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealthCheck
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipHealthCheck(dAtA[start:])
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
	ErrInvalidLengthHealthCheck = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealthCheck   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/health_check/v2/health_check.proto", fileDescriptorHealthCheck)
}

var fileDescriptorHealthCheck = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0xa5, 0x33, 0x33, 0xea, 0xf4, 0x1c, 0x1c, 0xa3, 0x60, 0x1c, 0x24, 0x3b, 0xeb, 0x69, 0xbc,
	0x74, 0x43, 0xbc, 0x88, 0x0b, 0x1e, 0xb2, 0x0a, 0x7b, 0x19, 0x90, 0xb0, 0x28, 0x78, 0x09, 0xbd,
	0x49, 0x4d, 0xd2, 0x6c, 0x26, 0xdd, 0x74, 0x3a, 0x91, 0x7c, 0x80, 0x17, 0xbf, 0xc0, 0x93, 0x1f,
	0xe2, 0xc9, 0xa3, 0x47, 0xff, 0x40, 0x99, 0x9b, 0x7f, 0x21, 0xdd, 0x9d, 0x51, 0x97, 0x01, 0xdd,
	0x4b, 0xa8, 0xd4, 0xab, 0x57, 0xf5, 0x78, 0xaf, 0xf1, 0x09, 0xd4, 0x9d, 0xe8, 0x69, 0x26, 0xea,
	0x0d, 0x2f, 0xe8, 0x86, 0x57, 0x1a, 0x14, 0x2d, 0xb5, 0x96, 0xb4, 0x04, 0x56, 0xe9, 0x32, 0xcd,
	0x4a, 0xc8, 0x2e, 0x69, 0x17, 0x5d, 0xf9, 0x27, 0x52, 0x09, 0x2d, 0xfc, 0x95, 0x25, 0x13, 0x47,
	0x26, 0x8e, 0x4c, 0x0c, 0x99, 0x5c, 0x19, 0xee, 0xa2, 0x45, 0x58, 0x08, 0x51, 0x54, 0x40, 0x2d,
	0xef, 0xa2, 0xdd, 0xd0, 0xbc, 0x55, 0x4c, 0x73, 0x51, 0xbb, 0x4d, 0x87, 0xf8, 0x3b, 0xc5, 0xa4,
	0x04, 0xd5, 0xec, 0x71, 0x27, 0x93, 0x49, 0x6e, 0xa4, 0x28, 0xd1, 0x6a, 0x70, 0xdf, 0x01, 0x0f,
	0x1c, 0xae, 0x7b, 0x09, 0x54, 0x82, 0xca, 0xa0, 0xd6, 0x03, 0x72, 0xbf, 0x63, 0x15, 0xcf, 0x99,
	0x06, 0xba, 0x2f, 0x06, 0xe0, 0x5e, 0x21, 0x0a, 0x61, 0x4b, 0x6a, 0x2a, 0xd7, 0x7d, 0xf4, 0x7e,
	0x8c, 0x67, 0x67, 0x56, 0xfc, 0xa9, 0xd1, 0xee, 0x27, 0xf8, 0x8e, 0x64, 0x4d, 0x93, 0xea, 0x52,
	0x89, 0xb6, 0x28, 0xd3, 0xad, 0xc8, 0x21, 0x40, 0x4b, 0xb4, 0x9a, 0x45, 0x0b, 0xe2, 0x44, 0x93,
	0xbd, 0x68, 0x12, 0x0b, 0x51, 0xbd, 0x66, 0x55, 0x0b, 0x31, 0xfe, 0xfc, 0xf3, 0xcb, 0x68, 0xf2,
	0x01, 0x79, 0x73, 0x94, 0xdc, 0x36, 0x0b, 0xce, 0x1d, 0x7f, 0x2d, 0x72, 0xf0, 0x43, 0x7c, 0x0b,
	0xea, 0x5c, 0x0a, 0x5e, 0xeb, 0xc0, 0x5b, 0xa2, 0xd5, 0x34, 0xf6, 0x02, 0x94, 0xfc, 0xee, 0xf9,
	0xcf, 0x31, 0xce, 0x58, 0x56, 0x42, 0xaa, 0xf9, 0x16, 0x82, 0x91, 0x3d, 0xf6, 0xe0, 0xe0, 0xd8,
	0x8b, 0xc1, 0xc1, 0x78, 0xfc, 0xf1, 0xfb, 0x11, 0x4a, 0xa6, 0x96, 0x72, 0xce, 0xb7, 0xe0, 0x7f,
	0x42, 0xf8, 0x28, 0xab, 0xda, 0x46, 0x83, 0x4a, 0xb7, 0xbc, 0x4e, 0x5d, 0x18, 0x7d, 0x3a, 0x18,
	0xc3, 0x0a, 0x68, 0x82, 0xf1, 0x72, 0xb4, 0x9a, 0x45, 0x6f, 0xc8, 0x75, 0x13, 0x24, 0x7f, 0x99,
	0x42, 0x4e, 0xdd, 0xf2, 0x35, 0xaf, 0x5d, 0xb7, 0x7f, 0xf5, 0x67, 0xf3, 0xcb, 0x5a, 0xab, 0x3e,
	0x79, 0x98, 0xfd, 0x63, 0xc4, 0x3f, 0xc1, 0x37, 0x4b, 0x60, 0x39, 0xa8, 0x26, 0x98, 0x58, 0x1d,
	0xc7, 0x83, 0x0e, 0x26, 0xb9, 0xb9, 0xe5, 0x92, 0x3d, 0xb3, 0x23, 0x6b, 0xa6, 0xb3, 0x12, 0x54,
	0xb2, 0x67, 0x2c, 0x72, 0x7c, 0xfc, 0xdf, 0xfb, 0xfe, 0x1c, 0x8f, 0x2e, 0xa1, 0xb7, 0x41, 0x4d,
	0x13, 0x53, 0xfa, 0x8f, 0xf1, 0xa4, 0x33, 0xd1, 0x58, 0xc7, 0x67, 0xd1, 0xdd, 0xe1, 0xa2, 0x79,
	0x31, 0x64, 0xa0, 0x27, 0x6e, 0xe2, 0x99, 0xf7, 0x14, 0xc5, 0xf3, 0xaf, 0xbb, 0x10, 0x7d, 0xdb,
	0x85, 0xe8, 0xc7, 0x2e, 0x44, 0x6f, 0xbd, 0x2e, 0xba, 0xb8, 0x61, 0x9d, 0x7f, 0xf2, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0xf7, 0x33, 0xc6, 0x32, 0x03, 0x00, 0x00,
}
