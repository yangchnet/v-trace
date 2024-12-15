package pbtools

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"golang.org/x/exp/constraints"
)

func FromProtoTimestamp(t *timestamp.Timestamp) (tt time.Time) {
	tt, err := ptypes.Timestamp(t)
	if err != nil {
		panic(err)
	}
	return
}

func FromProtoStringSlice(pbStrings []*wrappers.StringValue) []string {
	strs := make([]string, 0)
	for _, s := range pbStrings {
		strs = append(strs, s.GetValue())
	}

	return strs
}

func ToProtoStringSlice(strs []string) []*wrappers.StringValue {
	pbStrings := make([]*wrappers.StringValue, 0)
	for _, s := range strs {
		pbStrings = append(pbStrings, ToProtoString(s))
	}

	return pbStrings
}

func ToProtoTimestamp(t time.Time) (tt *timestamp.Timestamp) {
	if t.IsZero() {
		return nil
	}
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func ToProtoString(str string) *wrappers.StringValue {
	return &wrappers.StringValue{Value: str}
}

func ToProtoUInt32(uint32 uint32) *wrappers.UInt32Value {
	return &wrappers.UInt32Value{Value: uint32}
}

func ToProtoInt32[T constraints.Signed](i T) *wrappers.Int32Value {
	_i := int32(i)
	return &wrappers.Int32Value{Value: _i}
}

func ToProtoInt32List(is []int32) []*wrappers.Int32Value {
	ll := make([]*wrappers.Int32Value, 0)
	for _, i := range is {
		ll = append(ll, ToProtoInt32(i))
	}

	return ll
}

func FromProtoInt32Slice(is []*wrappers.Int32Value) []int32 {
	res := make([]int32, 0)
	for _, i := range is {
		res = append(res, i.GetValue())
	}

	return res
}

func FromProtoInt64Slice(is []*wrappers.Int64Value) []int {
	res := make([]int, 0)
	for _, i := range is {
		res = append(res, int(i.GetValue()))
	}

	return res
}

// a try, genertic.
func ToProtoInt64[T constraints.Signed](i T) *wrappers.Int64Value {
	_i := int64(i)
	return &wrappers.Int64Value{Value: _i}
}

func ToProtoInt64List[T constraints.Signed](is []T) []*wrappers.Int64Value {
	ll := make([]*wrappers.Int64Value, 0)
	for _, i := range is {
		ll = append(ll, ToProtoInt64(i))
	}

	return ll
}

func ToProtoBool(bool bool) *wrappers.BoolValue {
	return &wrappers.BoolValue{Value: bool}
}

func ToProtoBytes(bytes []byte) *wrappers.BytesValue {
	return &wrappers.BytesValue{Value: bytes}
}
