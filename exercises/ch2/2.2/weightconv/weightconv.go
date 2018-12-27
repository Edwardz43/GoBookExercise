package weightconv

import "fmt"

type KiloGram float64

type Pound float64

func (k KiloGram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glbs", p) }
