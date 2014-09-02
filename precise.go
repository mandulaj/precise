package precise


import "fmt"

type FloatU struct {
    value float64
    uncertainty float32
    sigFig int
}

func NewFloatU(value float64, uncertainty float32, sigFig int) (f* FloatU) {
    return &FloatU{value, uncertainty, sigFig}
}

func (f* FloatU) Add(u FloatU) {
    f.value += u.value
}

func (f* FloatU) Sub(u FloatU) {
    f.value -= u.value
}

func (f* FloatU) Mul(u FloatU) {
    f.value *= u.value
}

func (f* FloatU) Div(u FloatU) {
    f.value /= u.value
}

func (f* FloatU)Pow(i int) {

}

func (f FloatU) String() string {
    format := fmt.Sprintf("%%.%df +/-%%.%df",f.sigFig, f.sigFig)
    return fmt.Sprintf(format, f.value, f.uncertainty)
}

