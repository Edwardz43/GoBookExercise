package weightconv

func MToI(k KiloGram) Pound { return Pound(k * 2.20462262) }

func IToM(p Pound) KiloGram { return KiloGram(p * 0.45359237) }
