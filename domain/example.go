package domain

import "time"

type Example struct {
	Str  string
	Int  int64
	Flt  float64
	Byte []byte
	Tm   time.Time
}

