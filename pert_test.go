package main

import (
	"testing"
)

func TestPertOne(t *testing.T) {
	const o, m, p = 1.0, 1.0, 1.0
	const out = 2.0
	const outsd = 0.0

	if estimate, sd := pert(o, m, p); estimate != out && sd != outsd {
		t.Errorf("Pert(%v, %v, %v) = %v, want %v %v", o, m, p, estimate, out, sd)
	}
}

func TestPertOther(t *testing.T) {
	const o, m, p = 2, 4, 6
	const out = 4.00
	const outsd = 0.6666666666666666

	if estimate, sd := pert(o, m, p); estimate != out || sd != outsd {
		t.Errorf("Pert(%v, %v, %v) = %v - %v, want %v - %v", o, m, p, estimate, sd, out, outsd)
	}
}
