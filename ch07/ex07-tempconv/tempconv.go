package temp

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
