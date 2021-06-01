package goldilocks

import (
	"errors"
	"fmt"

	fp "github.com/MaxflowO2/mfc/cryptomath/fp448"
)

// Point is a point on the Goldilocks Curve.
type Point struct{ x, y, z, ta, tb fp.Elt }

func (P Point) String() string {
	return fmt.Sprintf("x: %v\ny: %v\nz: %v\nta: %v\ntb: %v", P.x, P.y, P.z, P.ta, P.tb)
}

// FromAffine creates a point from affine coordinates.
func FromAffine(x, y *fp.Elt) (*Point, error) {
	P := &Point{
		x:  *x,
		y:  *y,
		z:  fp.One(),
		ta: *x,
		tb: *y,
	}
	if !(Curve{}).IsOnCurve(P) {
		return P, errors.New("point not on curve")
	}
	return P, nil
}

// isLessThan returns true if 0 <= x < y, and assumes that slices are of the
// same length and are interpreted in little-endian order.
func isLessThan(x, y []byte) bool {
	i := len(x) - 1
	for i > 0 && x[i] == y[i] {
		i--
	}
	return x[i] < y[i]
}

// FromBytes returns a point from the input buffer.
func FromBytes(in []byte) (*Point, error) {
	if len(in) < fp.Size+1 {
		return nil, errors.New("wrong input length")
	}
	var err = errors.New("invalid decoding")
	P := &Point{}
	signX := in[fp.Size] >> 7
	copy(P.y[:], in[:fp.Size])
	p := fp.P()
	if !isLessThan(P.y[:], p[:]) {
		return nil, err
	}

	u, v := &fp.Elt{}, &fp.Elt{}
	one := fp.One()
	fp.Sqr(u, &P.y)                // u = y^2
	fp.Mul(v, u, &paramD)          // v = dy^2
	fp.Sub(u, u, &one)             // u = y^2-1
	fp.Sub(v, v, &one)             // v = dy^2-1
	isQR := fp.InvSqrt(&P.x, u, v) // x = sqrt(u/v)
	if !isQR {
		return nil, err
	}
	fp.Modp(&P.x) // x = x mod p
	if fp.IsZero(&P.x) && signX == 1 {
		return nil, err
	}
	if signX != (P.x[0] & 1) {
		fp.Neg(&P.x, &P.x)
	}
	P.ta = P.x
	P.tb = P.y
	P.z = fp.One()
	return P, nil
}

// IsIdentity returns true is P is the identity Point.
func (P *Point) IsIdentity() bool {
	return fp.IsZero(&P.x) && !fp.IsZero(&P.y) && !fp.IsZero(&P.z) && P.y == P.z
}

// IsEqual returns true if P is equivalent to Q.
func (P *Point) IsEqual(Q *Point) bool {
	l, r := &fp.Elt{}, &fp.Elt{}
	fp.Mul(l, &P.x, &Q.z)
	fp.Mul(r, &Q.x, &P.z)
	fp.Sub(l, l, r)
	b := fp.IsZero(l)
	fp.Mul(l, &P.y, &Q.z)
	fp.Mul(r, &Q.y, &P.z)
	fp.Sub(l, l, r)
	b = b && fp.IsZero(l)
	fp.Mul(l, &P.ta, &P.tb)
	fp.Mul(l, l, &Q.z)
	fp.Mul(r, &Q.ta, &Q.tb)
	fp.Mul(r, r, &P.z)
	fp.Sub(l, l, r)
	b = b && fp.IsZero(l)
	return b
}

// Neg obtains the inverse of the Point.
func (P *Point) Neg() { fp.Neg(&P.x, &P.x); fp.Neg(&P.ta, &P.ta) }

// ToAffine returns the x,y affine coordinates of P.
func (P *Point) ToAffine() (x, y fp.Elt) {
	fp.Inv(&P.z, &P.z)       // 1/z
	fp.Mul(&P.x, &P.x, &P.z) // x/z
	fp.Mul(&P.y, &P.y, &P.z) // y/z
	fp.Modp(&P.x)
	fp.Modp(&P.y)
	fp.SetOne(&P.z)
	P.ta = P.x
	P.tb = P.y
	return P.x, P.y
}

// ToBytes stores P into a slice of bytes.
func (P *Point) ToBytes(out []byte) error {
	if len(out) < fp.Size+1 {
		return errors.New("invalid decoding")
	}
	x, y := P.ToAffine()
	out[fp.Size] = (x[0] & 1) << 7
	return fp.ToBytes(out[:fp.Size], &y)
}

// MarshalBinary encodes the receiver into a binary form and returns the result.
func (P *Point) MarshalBinary() (data []byte, err error) {
	data = make([]byte, fp.Size+1)
	err = P.ToBytes(data[:fp.Size+1])
	return data, err
}

// UnmarshalBinary must be able to decode the form generated by MarshalBinary.
func (P *Point) UnmarshalBinary(data []byte) error { Q, err := FromBytes(data); *P = *Q; return err }

// Double sets P = 2Q.
func (P *Point) Double() { P.Add(P) }

// Add sets P =P+Q..
func (P *Point) Add(Q *Point) {
	// This is formula (5) from "Twisted Edwards Curves Revisited" by
	// Hisil H., Wong K.KH., Carter G., Dawson E. (2008)
	// https://doi.org/10.1007/978-3-540-89255-7_20
	x1, y1, z1, ta1, tb1 := &P.x, &P.y, &P.z, &P.ta, &P.tb
	x2, y2, z2, ta2, tb2 := &Q.x, &Q.y, &Q.z, &Q.ta, &Q.tb
	x3, y3, z3, E, H := &P.x, &P.y, &P.z, &P.ta, &P.tb
	A, B, C, D := &fp.Elt{}, &fp.Elt{}, &fp.Elt{}, &fp.Elt{}
	t1, t2, F, G := C, D, &fp.Elt{}, &fp.Elt{}
	fp.Mul(t1, ta1, tb1)  // t1 = ta1*tb1
	fp.Mul(t2, ta2, tb2)  // t2 = ta2*tb2
	fp.Mul(A, x1, x2)     // A = x1*x2
	fp.Mul(B, y1, y2)     // B = y1*y2
	fp.Mul(C, t1, t2)     // t1*t2
	fp.Mul(C, C, &paramD) // C = d*t1*t2
	fp.Mul(D, z1, z2)     // D = z1*z2
	fp.Add(F, x1, y1)     // x1+y1
	fp.Add(E, x2, y2)     // x2+y2
	fp.Mul(E, E, F)       // (x1+y1)*(x2+y2)
	fp.Sub(E, E, A)       // (x1+y1)*(x2+y2)-A
	fp.Sub(E, E, B)       // E = (x1+y1)*(x2+y2)-A-B
	fp.Sub(F, D, C)       // F = D-C
	fp.Add(G, D, C)       // G = D+C
	fp.Sub(H, B, A)       // H = B-A
	fp.Mul(z3, F, G)      // Z = F * G
	fp.Mul(x3, E, F)      // X = E * F
	fp.Mul(y3, G, H)      // Y = G * H, T = E * H
}