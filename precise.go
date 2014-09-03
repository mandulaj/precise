package precise


import "fmt"

type FloatU struct {
    value float64
    uncertainty float64
    sigFig uint
}

func NewFloatU(value, uncertainty float64, sigFig uint) (f* FloatU) {
    return &FloatU{value, uncertainty, sigFig}
}

func (f* FloatU) Add(u* FloatU)  {
    f.value += u.value
    f.uncertainty += u.uncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

func (f* FloatU) Sub(u* FloatU) {
    f.value -= u.value
    f.uncertainty += u.uncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}


func (f* FloatU) Mul(u* FloatU) {
    fRelUncertainty := f.Relative()
    fRelUncertainty += u.Relative()

    f.value *= u.value
    f.uncertainty = f.value * fRelUncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

func (f* FloatU) Div(u* FloatU) {
    fRelUncertainty := f.Relative()
    fRelUncertainty += u.Relative()

    f.value /= u.value
    f.uncertainty = f.value * fRelUncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

func (f* FloatU)Pow(p int) {
    temp := new(FloatU)
    temp* = f*

    for (i := 0; i< p; i++) {
        f.Mul(temp*)
    }
}

func (f FloatU) String() string {
    format := fmt.Sprintf("%%.%df +-%%.%df",f.sigFig, f.sigFig)
    return fmt.Sprintf(format, f.value, f.uncertainty)
}


func (f FloatU) Relative() float64 {
    return f.uncertainty/f.value
}
