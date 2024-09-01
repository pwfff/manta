package manta

import (
	"fmt"
	"math"
)

type fieldDecoder func(*reader) interface{}
type fieldFactory func(*field) fieldDecoder

var fieldTypeFactories = map[string]fieldFactory{
	"float32":                  floatFactory,
	"CNetworkedQuantizedFloat": quantizedFactory,
	"Vector":                   vectorFactory(3),
	"Vector2D":                 vectorFactory(2),
	"Vector4D":                 vectorFactory(4),
	"uint64":                   unsigned64Factory,
	"MatchID_t":                unsigned64Factory,
	"QAngle":                   qangleFactory,
	"CHandle":                  unsignedFactory,
	"CStrongHandle":            unsigned64Factory,
	"CEntityHandle":            unsignedFactory,
	"GameTime_t":               floatFactory,
}

var fieldNameDecoders = map[string]fieldDecoder{}

var fieldTypeDecoders = map[string]fieldDecoder{
	"bool":    booleanDecoder,
	"char":    stringDecoder,
	"color32": unsignedDecoder,
	"int16":   signedDecoder,
	"int32":   signedDecoder,
	"int64":   signedDecoder,
	"int8":    signedDecoder,
	"uint16":  unsignedDecoder,
	"uint32":  unsignedDecoder,
	"uint8":   unsignedDecoder,

	//"AttachmentHandle_t":                unsignedDecoder,
	//"ShardSolid_t":                      unsignedDecoder,
	//"DoorState_t":                       unsignedDecoder,
	//"PointWorldTextReorientMode_t":      unsignedDecoder,
	//"PointWorldTextJustifyVertical_t":   unsignedDecoder,
	//"PointWorldTextJustifyHorizontal_t": unsignedDecoder,

	//"style_index_t": byteDecoder,
	//"itemid_t": unsignedDecoder,

	"CBodyComponent":       componentDecoder,
	"CGameSceneNodeHandle": unsignedDecoder,
	"Color":                unsignedDecoder,
	"CPhysicsComponent":    componentDecoder,
	"CRenderComponent":     componentDecoder,
	"CLightComponent":      componentDecoder,
	"CUtlString":           stringDecoder,
	"CUtlStringToken":      unsignedDecoder,
	"CUtlSymbolLarge":      stringDecoder,
}

func unsignedFactory(f *field) fieldDecoder {
	return unsignedDecoder
}

func unsigned64Factory(f *field) fieldDecoder {
	switch f.encoder {
	case "fixed64":
		return fixed64Decoder
	}
	return unsigned64Decoder
}

func floatFactory(f *field) fieldDecoder {
	//fmt.Println("picking decoder for float", f.getName(), f.encoder, derefStr(f.bitCount))
	switch f.encoder {
	case "coord":
		return floatCoordDecoder
	case "simtime":
		return simulationTimeDecoder
	case "runetime":
		return runeTimeDecoder
	}

	if f.bitCount == nil || (*f.bitCount <= 0 || *f.bitCount >= 32) {
		return noscaleDecoder
	}

	return quantizedFactory(f)
}

func derefStr[T any](v *T) string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", *v)
}

func quantizedFactory(f *field) fieldDecoder {
	//fmt.Println("quantized factory", f.getName(), derefStr(f.bitCount), derefStr(f.encodeFlags), derefStr(f.lowValue), derefStr(f.highValue))
	qfd := newQuantizedFloatDecoder(f.bitCount, f.encodeFlags, f.lowValue, f.highValue)
	return func(r *reader) interface{} {
		return qfd.decode(r)
	}
}

func vectorFactory(n int) fieldFactory {
	return func(f *field) fieldDecoder {
		if n == 3 && f.encoder == "normal" {
			return vectorNormalDecoder
		}

		d := floatFactory(f)
		return func(r *reader) interface{} {
			x := make([]float32, n)
			for i := 0; i < n; i++ {
				x[i] = d(r).(float32)
			}
			return x
		}
	}
}

func vectorNormalDecoder(r *reader) interface{} {
	return r.read3BitNormal()
}

func byteDecoder(r *reader) interface{} {
	return r.readByte()
}

func fixed64Decoder(r *reader) interface{} {
	return r.readLeUint64()
}

func handleDecoder(r *reader) interface{} {
	return r.readVarUint32()
}

func booleanDecoder(r *reader) interface{} {
	return r.readBoolean()
}

func stringDecoder(r *reader) interface{} {
	return r.readString()
}

func defaultDecoder(r *reader) interface{} {
	//fmt.Println("in default decoder")
	return r.readVarUint32()
}

func signedDecoder(r *reader) interface{} {
	return r.readVarInt32()
}

func floatCoordDecoder(r *reader) interface{} {
	return r.readCoord()
}

func noscaleDecoder(r *reader) interface{} {
	return math.Float32frombits(r.readBits(32))
}

func runeTimeDecoder(r *reader) interface{} {
	return math.Float32frombits(r.readBits(4))
}

func simulationTimeDecoder(r *reader) interface{} {
	return float32(r.readVarUint32()) * (1.0 / 30)
}

func qangleFactory(f *field) fieldDecoder {
	if f.encoder == "qangle_pitch_yaw" {
		n := uint32(*f.bitCount)
		return func(r *reader) interface{} {
			return []float32{
				r.readAngle(n),
				r.readAngle(n),
				0.0,
			}
		}
	}

	if f.encoder == "qangle_precise" {
		var n uint32 = 20
		return func(r *reader) interface{} {
			rx := r.readBoolean()
			ry := r.readBoolean()
			rz := r.readBoolean()

			ret := []float32{
				0.0,
				0.0,
				0.0,
			}

			if rx {
				ret[0] = r.readAngle(n)
			}
			if ry {
				ret[1] = r.readAngle(n)
			}
			if rz {
				ret[2] = r.readAngle(n)
			}

			return ret
		}
	}

	if f.bitCount != nil && *f.bitCount != 0 {
		n := uint32(*f.bitCount)
		return func(r *reader) interface{} {
			return []float32{
				r.readAngle(n),
				r.readAngle(n),
				r.readAngle(n),
			}
		}
	}

	return func(r *reader) interface{} {
		ret := make([]float32, 3)
		rX := r.readBoolean()
		rY := r.readBoolean()
		rZ := r.readBoolean()
		if rX {
			ret[0] = r.readCoord()
		}
		if rY {
			ret[1] = r.readCoord()
		}
		if rZ {
			ret[2] = r.readCoord()
		}
		return ret
	}
}

func vector2Decoder(r *reader) interface{} {
	return []float32{r.readFloat(), r.readFloat()}
}

func unsignedDecoder(r *reader) interface{} {
	return uint64(r.readVarUint32())
}

func unsigned64Decoder(r *reader) interface{} {
	return r.readVarUint64()
}

func componentDecoder(r *reader) interface{} {
	return r.readBits(1)
}

func findDecoder(f *field) fieldDecoder {
	if v, ok := fieldTypeFactories[f.fieldType.baseType]; ok {
		return v(f)
	}

	if v, ok := fieldNameDecoders[f.varName]; ok {
		return v
	}

	if v, ok := fieldTypeDecoders[f.fieldType.baseType]; ok {
		return v
	}

	//fmt.Println("default decoder for", f.varName, f.fieldType.baseType)

	return defaultDecoder
}

func findDecoderByBaseType(f *field, baseType string) fieldDecoder {
	if v, ok := fieldTypeDecoders[baseType]; ok {
		//fmt.Println("useful decoder for baseType", baseType)
		return v
	}

	if v, ok := fieldTypeFactories[baseType]; ok {
		return v(f)
	}

	//fmt.Println("default decoder for baseType", baseType)

	return defaultDecoder
}
