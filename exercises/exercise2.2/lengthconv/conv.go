package lengthconv

func MToI(m MilliMeter) Inch {
	return Inch(m / 2.54)
}

func IToM(i Inch) MilliMeter {
	return MilliMeter(i * 2.54)
}
