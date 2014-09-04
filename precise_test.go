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

// Test all arithmetic
func TestAdd(t *testing.T) {
    t.Parallel()

    f := NewFloatU(25.5,0.2,5)
    d := NewFloatU(10.5,0.1,3)
    f.Add(d)

    if !f.EqlS(36.0) {
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

    if !f.EqlS(15.0) {
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

    if !f.EqlS(267.75) {
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

    if !f.EqlS( 2.438571429) {
        t.Errorf("expected value 36.0, got %f", f.value)
    }
    if !closeTo(f.uncertainty, 0.042177) {
        t.Errorf("expected uncertainty 0.042177, got %f", f.uncertainty)
    }
    if f.sigFig != 3 {
        t.Errorf("expected sigFig to be 3, got %d", f.sigFig)
    }

}

// Test compairsons
func TestEql(t *testing.T) {
    t.Parallel()

    f := NewFloatU(20.1, 0.5, 5)
    d := NewFloatU(20.3, 0.3, 5)
    e := NewFloatU(20.7, 0.1, 4)
    if !f.Eql(d) {
        t.Errorf("%s is not equla to %s", f, d)
    }
    if !f.Eql(e) {
        t.Errorf("%s is not equal to %s", f, e)
    }
    if !f.GtE(e) || !f.LtE(e) {
        t.Errorf("%s falils the Lt or Gt with %s", f, e)
    }

    e.AddS(0.1)

    if f.Eql(e) {
        t.Errorf("%s equals %s", f, e)
    }

    if f.Gt(e) {
        t.Errorf("%s is greater then %s", f, e)
    }

    if !f.Lt(e) {
        t.Errorf("%s is not less then %s", f, e)
    }

}


// Benchmarks
func BenchmarkNewFloatU(b *testing.B) {
    for n := 0; n < b.N; n++ {
        NewFloatU(20.5,0.3,4)
    }
}

func BenchmarkEql(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.Eql(d)
    }
}

func BenchmarkLt(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.Lt(d)
    }
}

func BenchmarkGt(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.Gt(d)
    }
}


func BenchmarkLtE(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.LtE(d)
    }
}

func BenchmarkGtE(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.GtE(d)
    }
}

func BenchmarkAdd(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.Add(d)
    }
}

func BenchmarkSub(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.Sub(d)
    }
}

func BenchmarkMul(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.Mul(d)
    }
}

func BenchmarkDiv(b *testing.B) {
    f := NewFloatU(10.0,0.2,4)
    d := NewFloatU(10.1, 0.2, 2)

    for n := 0; n < b.N; n++ {
        f.Div(d)
    }
}

func benchPow(i int, b *testing.B) {
    f := NewFloatU(10.0, 0.2, 2)
    for n := 0; n < b.N; n++ {
        f.Pow(i)
    }
}

func BenchmarkPow2(b *testing.B)  { benchPow(2, b);  }
func BenchmarkPow4(b *testing.B)  { benchPow(4, b);  }
func BenchmarkPow8(b *testing.B)  { benchPow(8, b);  }
func BenchmarkPow16(b *testing.B) { benchPow(16, b); }
func BenchmarkPow32(b *testing.B) { benchPow(32, b); }
func BenchmarkPow64(b *testing.B) { benchPow(64, b); }
