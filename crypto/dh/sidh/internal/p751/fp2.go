// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package p751

import (
	"github.com/MaxflowO2/mfc/cryptodh/sidh/internal/common"
)

// Montgomery multiplication. Input values must be already
// in Montgomery domain.
func mulP(dest, lhs, rhs *common.Fp) {
	var ab common.FpX2
	mulP751(&ab, lhs, rhs) // = a*b*R*R
	rdcP751(dest, &ab)     // = a*b*R mod p
}

// Set dest = x^((p-3)/4).  If x is square, this is 1/sqrt(x).
// Uses variation of sliding-window algorithm from with window size
// of 5 and least to most significant bit sliding (left-to-right)
// See HAC 14.85 for general description.
//
// Allowed to overlap x with dest.
// All values in Montgomery domains
// Set dest = x^(2^k), for k >= 1, by repeated squarings.
func p34(dest, x *common.Fp) {
	var lookup [16]common.Fp

	// This performs sum(powStrategy) + 1 squarings and len(lookup) + len(mulStrategy)
	// multiplications.
	powStrategy := []uint8{5, 7, 6, 2, 10, 4, 6, 9, 8, 5, 9, 4, 7, 5, 5, 4, 8, 3, 9, 5, 5, 4, 10, 4, 6, 6, 6, 5, 8, 9, 3, 4, 9, 4, 5, 6, 6, 2, 9, 4, 5, 5, 5, 7, 7, 9, 4, 6, 4, 8, 5, 8, 6, 6, 2, 9, 7, 4, 8, 8, 8, 4, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 2}
	mulStrategy := []uint8{15, 11, 10, 0, 15, 3, 3, 3, 4, 4, 9, 7, 11, 11, 5, 3, 12, 2, 10, 8, 5, 2, 8, 3, 5, 4, 11, 4, 0, 9, 2, 1, 12, 7, 5, 14, 15, 0, 14, 5, 6, 4, 5, 13, 6, 9, 7, 15, 1, 14, 11, 15, 12, 5, 0, 10, 9, 7, 7, 10, 14, 6, 11, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 1}
	initialMul := uint8(13)

	// Precompute lookup table of odd multiples of x for window
	// size k=5.
	var xx common.Fp
	mulP(&xx, x, x)
	lookup[0] = *x
	for i := 1; i < 16; i++ {
		mulP(&lookup[i], &lookup[i-1], &xx)
	}

	// Now lookup = {x, x^3, x^5, ... }
	// so that lookup[i] = x^{2*i + 1}
	// so that lookup[k/2] = x^k, for odd k
	*dest = lookup[initialMul]
	for i := uint8(0); i < uint8(len(powStrategy)); i++ {
		mulP(dest, dest, dest)
		for j := uint8(1); j < powStrategy[i]; j++ {
			mulP(dest, dest, dest)
		}
		mulP(dest, dest, &lookup[mulStrategy[i]])
	}
}

func add(dest, lhs, rhs *common.Fp2) {
	addP751(&dest.A, &lhs.A, &rhs.A)
	addP751(&dest.B, &lhs.B, &rhs.B)
}

func sub(dest, lhs, rhs *common.Fp2) {
	subP751(&dest.A, &lhs.A, &rhs.A)
	subP751(&dest.B, &lhs.B, &rhs.B)
}

func mul(dest, lhs, rhs *common.Fp2) {
	var bMinA, cMinD common.Fp
	var ac, bd common.FpX2
	var adPlusBc common.FpX2
	var acMinBd common.FpX2

	// Let (a,b,c,d) = (lhs.a,lhs.b,rhs.a,rhs.b).
	//
	// (a + bi)*(c + di) = (a*c - b*d) + (a*d + b*c)i
	//
	// Use Karatsuba's trick: note that
	//
	// (b - a)*(c - d) = (b*c + a*d) - a*c - b*d
	//
	// so (a*d + b*c) = (b-a)*(c-d) + a*c + b*d.
	mulP751(&ac, &lhs.A, &rhs.A)       // = a*c*R*R
	mulP751(&bd, &lhs.B, &rhs.B)       // = b*d*R*R
	subP751(&bMinA, &lhs.B, &lhs.A)    // = (b-a)*R
	subP751(&cMinD, &rhs.A, &rhs.B)    // = (c-d)*R
	mulP751(&adPlusBc, &bMinA, &cMinD) // = (b-a)*(c-d)*R*R
	adlP751(&adPlusBc, &adPlusBc, &ac) // = ((b-a)*(c-d) + a*c)*R*R
	adlP751(&adPlusBc, &adPlusBc, &bd) // = ((b-a)*(c-d) + a*c + b*d)*R*R
	rdcP751(&dest.B, &adPlusBc)        // = (a*d + b*c)*R mod p
	sulP751(&acMinBd, &ac, &bd)        // = (a*c - b*d)*R*R
	rdcP751(&dest.A, &acMinBd)         // = (a*c - b*d)*R mod p
}

// Set dest = 1/x
//
// Allowed to overlap dest with x.
//
// Returns dest to allow chaining operations.
func inv(dest, x *common.Fp2) {
	var e1, e2 common.FpX2
	var f1, f2 common.Fp

	// We want to compute
	//
	//    1          1     (a - bi)	    (a - bi)
	// -------- = -------- -------- = -----------
	// (a + bi)   (a + bi) (a - bi)   (a^2 + b^2)
	//
	// Letting c = 1/(a^2 + b^2), this is
	//
	// 1/(a+bi) = a*c - b*ci.

	mulP751(&e1, &x.A, &x.A) // = a*a*R*R
	mulP751(&e2, &x.B, &x.B) // = b*b*R*R
	adlP751(&e1, &e1, &e2)   // = (a^2 + b^2)*R*R
	rdcP751(&f1, &e1)        // = (a^2 + b^2)*R mod p
	// Now f1 = a^2 + b^2

	mulP(&f2, &f1, &f1)
	p34(&f2, &f2)
	mulP(&f2, &f2, &f2)
	mulP(&f2, &f2, &f1)

	mulP751(&e1, &x.A, &f2)
	rdcP751(&dest.A, &e1)

	subP751(&f1, &common.Fp{}, &x.B)
	mulP751(&e1, &f1, &f2)
	rdcP751(&dest.B, &e1)
}

func sqr(dest, x *common.Fp2) {
	var a2, aPlusB, aMinusB common.Fp
	var a2MinB2, ab2 common.FpX2

	a := &x.A
	b := &x.B

	// (a + bi)*(a + bi) = (a^2 - b^2) + 2abi.
	addP751(&a2, a, a)                   // = a*R + a*R = 2*a*R
	addP751(&aPlusB, a, b)               // = a*R + b*R = (a+b)*R
	subP751(&aMinusB, a, b)              // = a*R - b*R = (a-b)*R
	mulP751(&a2MinB2, &aPlusB, &aMinusB) // = (a+b)*(a-b)*R*R = (a^2 - b^2)*R*R
	mulP751(&ab2, &a2, b)                // = 2*a*b*R*R
	rdcP751(&dest.A, &a2MinB2)           // = (a^2 - b^2)*R mod p
	rdcP751(&dest.B, &ab2)               // = 2*a*b*R mod p
}

// In case choice == 1, performs following swap in constant time:
// 	xPx <-> xQx
//	xPz <-> xQz
// Otherwise returns xPx, xPz, xQx, xQz unchanged
func cswap(xPx, xPz, xQx, xQz *common.Fp2, choice uint8) {
	cswapP751(&xPx.A, &xQx.A, choice)
	cswapP751(&xPx.B, &xQx.B, choice)
	cswapP751(&xPz.A, &xQz.A, choice)
	cswapP751(&xPz.B, &xQz.B, choice)
}

// Converts in.A and in.B to Montgomery domain and stores
// in 'out'
// out.A = in.A * R mod p
// out.B = in.B * R mod p
// Performs v = v*R^2*R^(-1) mod p, for both in.A and in.B
func ToMontgomery(out, in *common.Fp2) {
	var aRR common.FpX2

	// a*R*R
	mulP751(&aRR, &in.A, &P751R2)
	// a*R mod p
	rdcP751(&out.A, &aRR)
	mulP751(&aRR, &in.B, &P751R2)
	rdcP751(&out.B, &aRR)
}

// Converts in.A and in.B from Montgomery domain and stores
// in 'out'
// out.A = in.A mod p
// out.B = in.B mod p
//
// After returning from the call 'in' is not modified.
func FromMontgomery(out, in *common.Fp2) {
	var aR common.FpX2

	// convert from montgomery domain
	copy(aR[:], in.A[:])
	rdcP751(&out.A, &aR) // = a mod p in [0, 2p)
	modP751(&out.A)      // = a mod p in [0, p)
	for i := range aR {
		aR[i] = 0
	}
	copy(aR[:], in.B[:])
	rdcP751(&out.B, &aR)
	modP751(&out.B)
}