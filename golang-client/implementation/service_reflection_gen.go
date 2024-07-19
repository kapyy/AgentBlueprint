package implementation

func GetDataIDFromFunctionID(functionID uint64) uint64 {
	switch functionID {
	case 110000001:
		return 1001
	case 310000002:
		return 4001
	default:
		return 1
	}
}
