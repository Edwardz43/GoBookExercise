package lengthconv

import "fmt"

// type CentiMeter float64
// type Meter float64
// type KiloMeter float64
type MilliMeter float64

type Inch float64

// type Foot float64
// type Yard float64
// type Mile float64

func (m MilliMeter) String() string { return fmt.Sprintf("%gmm", m) }
func (i Inch) String() string       { return fmt.Sprintf("%ginch", i) }
