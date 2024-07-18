package implementation

func GetDataIDFromFunctionID(functionID uint64) uint64 {
	switch functionID {
	case 100100100:
		return 1001 //SomeFunction>>Actions
	default:
		return 1
	}
}
