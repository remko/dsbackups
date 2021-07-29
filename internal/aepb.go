package internal

import (
	"fmt"

	aepb "github.com/remko/dsbackups/internal/aepb"
	pb "google.golang.org/genproto/googleapis/datastore/v1"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UnmarshalAEEntity(b []byte) (*pb.Entity, error) {
	src := aepb.EntityProto{}
	if err := proto.Unmarshal(b, &src); err != nil {
		return nil, err
	}

	props, rawProps := src.Property, src.RawProperty
	outProps := map[string]*pb.Value{}
	for {
		var (
			x       *aepb.Property
			noIndex bool
		)
		if len(props) > 0 {
			x, props = props[0], props[1:]
		} else if len(rawProps) > 0 {
			x, rawProps = rawProps[0], rawProps[1:]
			noIndex = true
		} else {
			break
		}

		if x.Meaning != nil && *x.Meaning == aepb.Property_INDEX_VALUE {
			return nil, fmt.Errorf("index values not supported")
		}
		// value := &pb.Value{ExcludeFromIndexes: noIndex}

		value, err := propValue(x.Value, x.GetMeaning())
		if err != nil {
			return nil, err
		}
		value.ExcludeFromIndexes = noIndex

		prev, ok := outProps[x.GetName()]
		if x.GetMultiple() {
			if !ok {
				outProps[x.GetName()] = &pb.Value{
					ValueType: &pb.Value_ArrayValue{
						ArrayValue: &pb.ArrayValue{
							Values: []*pb.Value{value}},
					},
				}
			} else {
				prev.GetArrayValue().Values = append(prev.GetArrayValue().Values, value)
			}
		} else {
			if ok {
				return nil, fmt.Errorf("multiple properties for %s", x.GetName())
			}
			outProps[x.GetName()] = value
		}
	}

	return &pb.Entity{
		Key:        aeProtoToProtoKey(src.Key),
		Properties: outProps,
	}, nil
}

// Mostly copy from https://github.com/golang/appengine/blob/master/datastore/load.go
func propValue(v *aepb.PropertyValue, m aepb.Property_Meaning) (*pb.Value, error) {
	value := &pb.Value{}
	switch {
	case v.Int64Value != nil:
		if m == aepb.Property_GD_WHEN {
			value.ValueType = &pb.Value_TimestampValue{TimestampValue: fromUnixMicro(*v.Int64Value)}
		} else {
			value.ValueType = &pb.Value_IntegerValue{IntegerValue: *v.Int64Value}
		}
	case v.BooleanValue != nil:
		value.ValueType = &pb.Value_BooleanValue{BooleanValue: *v.BooleanValue}
	case v.StringValue != nil:
		if m == aepb.Property_BLOB {
			value.ValueType = &pb.Value_BlobValue{BlobValue: []byte(*v.StringValue)}
		} else if m == aepb.Property_BLOBKEY {
			return nil, fmt.Errorf("blobkey not supported")
		} else if m == aepb.Property_BYTESTRING {
			value.ValueType = &pb.Value_BlobValue{BlobValue: []byte(*v.StringValue)}
		} else if m == aepb.Property_ENTITY_PROTO {
			return nil, fmt.Errorf("entityproto not supported")
		} else {
			value.ValueType = &pb.Value_StringValue{StringValue: *v.StringValue}
		}
	case v.DoubleValue != nil:
		value.ValueType = &pb.Value_DoubleValue{DoubleValue: *v.DoubleValue}
	case v.Referencevalue != nil:
		value.ValueType = &pb.Value_KeyValue{KeyValue: referenceValueToProtoKey(v.Referencevalue)}
	case v.Pointvalue != nil:
		value.ValueType = &pb.Value_GeoPointValue{GeoPointValue: &latlng.LatLng{
			Latitude:  *v.Pointvalue.X,
			Longitude: *v.Pointvalue.Y,
		}}
	default:
		value.ValueType = &pb.Value_NullValue{}
	}
	return value, nil
}

func fromUnixMicro(t int64) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{Seconds: t / 1e6, Nanos: int32(((t % 1e6) * 1e3))}
}

func aeProtoToProtoKey(r *aepb.Reference) *pb.Key {
	if r == nil {
		return nil
	}
	var path []*pb.Key_PathElement
	for _, e := range r.Path.Element {
		el := &pb.Key_PathElement{Kind: e.GetType()}
		if e.GetId() != 0 {
			el.IdType = &pb.Key_PathElement_Id{Id: e.GetId()}
		} else if e.GetName() != "" {
			el.IdType = &pb.Key_PathElement_Name{Name: e.GetName()}
		}
		path = append(path, el)
	}
	key := &pb.Key{Path: path}
	if r.GetNameSpace() != "" {
		key.PartitionId = &pb.PartitionId{
			NamespaceId: r.GetNameSpace(),
		}
	}
	return key
}

func referenceValueToProtoKey(r *aepb.PropertyValue_ReferenceValue) (k *pb.Key) {
	var path []*pb.Key_PathElement
	for _, e := range r.Pathelement {
		el := &pb.Key_PathElement{Kind: e.GetType()}
		if e.GetId() != 0 {
			el.IdType = &pb.Key_PathElement_Id{Id: e.GetId()}
		} else if e.GetName() != "" {
			el.IdType = &pb.Key_PathElement_Name{Name: e.GetName()}
		}
		path = append(path, el)
	}
	key := &pb.Key{Path: path, PartitionId: &pb.PartitionId{
		ProjectId:   r.GetApp(),
		NamespaceId: r.GetNameSpace(),
	}}
	return key
}
