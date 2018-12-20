package tempconv

import "fmt"

type Celsius float64    // 攝氏溫度
type Fahrenheit float64 // 華氏溫度
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15 // 絶對零度(C)
	AbsoluteZeroK Kelvin  = 0       // 絶對零度(K)
	FreezingC     Celsius = 0       // 結冰點溫度
	BoilingC      Celsius = 100     // 沸水溫度
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
