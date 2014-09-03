package precise


import "fmt"

/* FloatU struct - a float with uncertainty
 * 
 * Pointers to this structure are used to represent floating point values with uncertainty
 * use NewFloatU(value, uncertainty, sigFig) to create new floating point numbers
*/
type FloatU struct {
    value float64
    uncertainty float64
    sigFig uint
}

/* NewFloatU()
 *
 * Used to create new float64 with an uncertainty
 * value float64        - self-explanatory
 * uncertainty float64  - the uncertainty of your value
 * sigFig uint          - number of significant figures
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

// Subtract u from f
func (f *FloatU) Sub(u *FloatU) {
    f.value -= u.value
    f.uncertainty += u.uncertainty 

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

// Multiply f times u
func (f *FloatU) Mul(u *FloatU) {
    fRelUncertainty := f.Relative() // relative uncertainties add
    fRelUncertainty += u.Relative()

    f.value *= u.value
    f.uncertainty = f.value * fRelUncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

// Divide f by u
func (f *FloatU) Div(u* FloatU) {
    fRelUncertainty := f.Relative()
    fRelUncertainty += u.Relative()

    f.value /= u.value
    f.uncertainty = f.value * fRelUncertainty

    if (f.sigFig > u.sigFig) {
        f.sigFig = u.sigFig
    }
}

// Raise f to the power of p
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

// FloatU implements the Stringer interface so variables of type FloatU can be used in fmt.Print()
func (f FloatU) String() string {
    format := fmt.Sprintf("%%.%df +-%%.%df",f.sigFig, f.sigFig)
    return fmt.Sprintf(format, f.value, f.uncertainty)
}

// What is the relative uncertainty of f
func (f FloatU) Relative() float64 {
    return f.uncertainty/f.value
}
