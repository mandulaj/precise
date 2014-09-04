package precise

import (
    "testing"
    "math"
)

func closeTo(a, b float64) bool {
    if math.Abs(a-b) < 0.000001 {
        return true
    } else {
        return false
    }
}

func TestAdd(t *testing.T) {
    t.Parallel()

    f := NewFloatU(25.5,0.2,5)
    d := NewFloatU(10.5,0.1,3)
    f.Add(d)

    if f.value != 36.0 {
        t.Errorf("expected value 36.0, got %f", f.value)
    }
    if !closeTo(f.uncertainty, 0.3) {
        t.Errorf("expected uncertainty 0.3, got %f", f.uncertainty)
    }
    if f.sigFig != 3 {
        t.Errorf("expected sigFig to be 3, got %d", f.sigFig)
    }
}

func TestSub(t *testing.T) {
    t.Parallel()

    f := NewFloatU(25.5,0.2,5)
    d := NewFloatU(10.5,0.1,3)
    f.Sub(d)

    if f.value != 15.0 {
        t.Errorf("expected value 15.0, got %f", f.value)
    }
    if !closeTo(f.uncertainty, 0.3) {
        t.Errorf("expected uncertainty 0.3, got %f", f.uncertainty)
    }
    if f.sigFig != 3 {
        t.Errorf("expected sigFig to be 3, got %d", f.sigFig)
    }

}

func TestMul(t *testing.T) {
    t.Parallel()

    f := NewFloatU(25.5,0.2,5)
    d := NewFloatU(10.5,0.1,3)
    f.Mul(d)

    if f.value != 267.75 {
        t.Errorf("expected value 267.75, got %f", f.value)
    }
    if !closeTo(f.uncertainty, 4.65) {
        t.Errorf("expected uncertainty 4.65, got %f", f.uncertainty)
    }
    if f.sigFig != 3 {
        t.Errorf("expected sigFig to be 3, got %d", f.sigFig)
    }

}

func TestDiv(t *testing.T) {
    t.Parallel()

    f := NewFloatU(25.5,0.2,5)
    d := NewFloatU(10.5,0.1,3)
    f.Div(d)

    if f.value != 2.438571429 {
        t.Errorf("expected value 36.0, got %f", f.value)
    }
    if !closeTo(f.uncertainty, 0.3) {
        t.Errorf("expected uncertainty 0.3, got %f", f.uncertainty)
    }
    if f.sigFig != 3 {
        t.Errorf("expected sigFig to be 3, got %d", f.sigFig)
    }

}
