package denary

func Parse(src interface{}) Result {

	switch casted := src.(type) {

	case Type:
		return casted.Result()
	case Option:
		return casted.Result()
	case Result:
		var result Result = casted
		return result

	case int8:
		return someResult(Int8(casted).String())
	case int16:
		return someResult(Int16(casted).String())
	case int32:
		return someResult(Int32(casted).String())
	case int64:
		return someResult(Int64(casted).String())

	case uint8:
		return someResult(Uint8(casted).String())
	case uint16:
		return someResult(Uint16(casted).String())
	case uint32:
		return someResult(Uint32(casted).String())
	case uint64:
		return someResult(Uint64(casted).String())

	default:
		return Errorf("denary: %T is an unsupported type", src)
	}
}
