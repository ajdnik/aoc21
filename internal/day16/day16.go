// Package day16 solves AoC 2021 day 16: Packet Decoder.
// Parse nested BITS packets from hex and evaluate their expression tree.
package day16

import (
	"math"
)

type bitReader struct {
	bits []byte
	pos  int
}

func newBitReader(hex string) *bitReader {
	bits := make([]byte, len(hex)*4)
	for i, ch := range hex {
		var val int
		if ch >= '0' && ch <= '9' {
			val = int(ch - '0')
		} else {
			val = int(ch-'A') + 10
		}
		for b := 3; b >= 0; b-- {
			bits[i*4+(3-b)] = byte((val >> b) & 1)
		}
	}
	return &bitReader{bits: bits}
}

func (r *bitReader) read(n int) int {
	val := 0
	for i := 0; i < n; i++ {
		val = val<<1 | int(r.bits[r.pos])
		r.pos++
	}
	return val
}

type packet struct {
	version int
	typeID  int
	value   int
	sub     []packet
}

func parsePacket(r *bitReader) packet {
	p := packet{
		version: r.read(3),
		typeID:  r.read(3),
	}
	if p.typeID == 4 {
		val := 0
		for {
			group := r.read(5)
			val = val<<4 | (group & 0xF)
			if group>>4 == 0 {
				break
			}
		}
		p.value = val
		return p
	}
	lengthTypeID := r.read(1)
	if lengthTypeID == 0 {
		totalBits := r.read(15)
		end := r.pos + totalBits
		for r.pos < end {
			p.sub = append(p.sub, parsePacket(r))
		}
	} else {
		numSub := r.read(11)
		for i := 0; i < numSub; i++ {
			p.sub = append(p.sub, parsePacket(r))
		}
	}
	return p
}

func sumVersions(p packet) int {
	total := p.version
	for _, s := range p.sub {
		total += sumVersions(s)
	}
	return total
}

// evaluate recursively computes the packet value.
// Type IDs: 0=sum, 1=product, 2=min, 3=max, 4=literal, 5=gt, 6=lt, 7=eq.
func evaluate(p packet) int {
	switch p.typeID {
	case 4:
		return p.value
	case 0:
		sum := 0
		for _, s := range p.sub {
			sum += evaluate(s)
		}
		return sum
	case 1:
		prod := 1
		for _, s := range p.sub {
			prod *= evaluate(s)
		}
		return prod
	case 2:
		min := math.MaxInt
		for _, s := range p.sub {
			if v := evaluate(s); v < min {
				min = v
			}
		}
		return min
	case 3:
		max := 0
		for _, s := range p.sub {
			if v := evaluate(s); v > max {
				max = v
			}
		}
		return max
	case 5:
		if evaluate(p.sub[0]) > evaluate(p.sub[1]) {
			return 1
		}
		return 0
	case 6:
		if evaluate(p.sub[0]) < evaluate(p.sub[1]) {
			return 1
		}
		return 0
	case 7:
		if evaluate(p.sub[0]) == evaluate(p.sub[1]) {
			return 1
		}
		return 0
	}
	return 0
}

func Part1(lines []string) int {
	r := newBitReader(lines[0])
	p := parsePacket(r)
	return sumVersions(p)
}

func Part2(lines []string) int {
	r := newBitReader(lines[0])
	p := parsePacket(r)
	return evaluate(p)
}
