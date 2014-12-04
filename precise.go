package precise


import "fmt"

/* FloatU struct - a float with uncertainty

  Pointers to this structure are used to represent floating point values with uncertainty
  use NewFloatU(value, uncertainty, sigFig) to create new floating point numbers
*/
type FloatU struct {
    value float64
    uncertainty float64
    sigFig uint
}

/* NewFloatU()

  Used to create new float64 with an uncertainty
  value float64        - self-explanatory
  uncertainty float64  - the uncertainty of your value
  sigFig uint          - number of significant figures
*/
func NewFloatU(value, uncertainty float64, sigFig uint) (f *FloatU) {
    return &FloatU{value, uncertainty, sigFig}
}

// Add u into f with respect to the significant figures and uncertainty
func (f *FloatU) Add(u *FloatU)  {
    f.value += u.value
    f.uncertainty += u.uncertainty // uncertainties add

    if (f.sigFig > u.sigFig) { // significat figures can not be more then any of the operands
        f.sigFig = u.sigFig
    }
}

// Subtract u from f with respect to the significant figures and uncertainty
func (f *FloatU) Sub(u *FloatU) {
    f.value -= u.value
    f.uncertainty += u.uncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

// Multiply f times u with respect to the significant figures and uncertainty
func (f *FloatU) Mul(u *FloatU) {
    fRelUncertainty := f.Relative() // relative uncertainties add
    fRelUncertainty += u.Relative()

    f.value *= u.value
    f.uncertainty = f.value * fRelUncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

// Divide f by u with respect to the significant figures and uncertainty
func (f *FloatU) Div(u* FloatU) {
    fRelUncertainty := f.Relative()
    fRelUncertainty += u.Relative()

    f.value /= u.value
    f.uncertainty = f.value * fRelUncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

// Raise f to the power of p with respect to the significant figures and uncertainty
func (f *FloatU)Pow(p int) {
    temp := new(FloatU)
    *temp = *f

    if p == 0 {
        f.value = 1.0
        return
    }

    for i:=1; i< p; i++ {
        f.Mul(temp)
    }
}

// Interface for scalars
func (f *FloatU)AddS(u float64) {
    f.value += u
}
func (f *FloatU)SubS(u float64) {
    f.value -= u
}
func (f *FloatU)MulS(u float64) {
    f.value *= u
}
func (f *FloatU)DivS(u float64) {
    f.value /= u
}

// FloatU implements the Stringer interface so variables of type FloatU can be used in fmt.Print()
func (f FloatU) String() string {
    format := fmt.Sprintf("%%.%df +-%%.%df",f.sigFig, f.sigFig)
    return fmt.Sprintf(format, f.value, f.uncertainty)
}

// Return the relative uncertainty of f
func (f FloatU) Relative() float64 {
    return f.uncertainty/f.value
}



// Test f == u within the limits of both uncertainties
func (f *FloatU) Eql(u *FloatU) bool {
    if  (f.value + f.uncertainty) >= (u.value - u.uncertainty) &&
        (f.value - f.uncertainty) <= (u.value + u.uncertainty) {
        return true
    } else {
        return false
    }
}

// Test f < u within the limits of both uncertainties
func (f *FloatU) Lt(u *FloatU) bool {
    return f.value + f.uncertainty < u.value - u.uncertainty
}

// Test f > u within the limits of both uncertainties
func (f *FloatU) Gt(u *FloatU) bool {
    return f.value - f.uncertainty > u.value + u.uncertainty
}

// Test f <= u within the limits of both uncertainties
func (f *FloatU) LtE(u *FloatU) bool {
    return f.value - f.uncertainty <= u.value + u.uncertainty
}

// Test f >= u within the limits of both uncertainties
func (f *FloatU) GtE(u *FloatU) bool {
    return f.value + f.uncertainty >= u.value - u.uncertainty
}


// Test f == Scalar within the limits of uncertainty
func (f *FloatU) EqlS(u float64) bool {
    if (f.value + f.uncertainty) >= u &&
       (f.value - f.uncertainty) <= u {
       return true
    } else {
        return false
    }
}

// Test f > Scalar within the limits of uncertainty
func (f *FloatU) GtS(u float64) bool {
    return f.value - f.uncertainty > u
}

// Test f < Scalar within the limits of uncertainty
func (f *FloatU) LtS(u float64) bool {
    return f.value + f.uncertainty < u
}

// Test f >= Scalar within the limits of uncertainty
func (f *FloatU) GtES(u float64) bool {
    return f.value + f.uncertainty >= u
}

// Test f <= Scalar within the limits of uncertainty
func (f *FloatU) LtES(u float64) bool {
    return f.value - f.uncertainty <= u
}

// Round value to i sigFig
func (f FloatU)roundSigFig(i float64) {

}
