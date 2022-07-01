// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package unitconv

// FToM converts feet to meters.
func FToM(f Feet) Meters { return Meters(f / 3.281) }

// PToK converts meters to feet.
func MtoF(m Meters) Feet { return Feet(m * 3.281) }

//!-
