package tools

import (
	"fmt"
	"math"
)

//Gcd 最大公约数:(辗转相除法)
func Gcd(x, y int64) int64 {
	x = int64(math.Abs(float64(x)))
	y = int64(math.Abs(float64(y)))

	var tmp int64
	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}

//Lcm 最小公倍数:((x*y)/最大公约数)
func Lcm(x, y int64) int64 {
	return (x * y) / Gcd(x, y)
}

// FAL Fractional operation correlation
type FAL struct {
	Nume int64 // Nume numerator.分子
	Deno int64 // denominator (must not be zero).分母 (一定不为0)
}

// Format output
func (f FAL) String() string { // 格式化输出
	return fmt.Sprintf("%v/%v", f.Nume, f.Deno)
}

// Model Create a score (molecular, denominator) with a denominator default of 1
func Model(nd ...int64) FAL { // 创建一个分数(分子，分母)，分母默认为1
	var f FAL
	if len(nd) == 1 {
		f.Nume = nd[0]
		f.Deno = 1
	} else if len(nd) == 2 {
		f.Nume = nd[0]
		f.Deno = nd[1]
	}

	if f.Deno == 0 { // denominator is 0 .分母为0
		panic(fmt.Sprintf("fractional init error. denominator can't zero."))
	}

	return f
}

// Broadsheet  .阔张
func (s *FAL) broad(lcm int64) {
	s.Nume = s.Nume * (lcm / s.Deno)
	s.Deno = lcm
}

// Compression Finishing .压缩 整理
func (s *FAL) offset() {
	lcm := Gcd(s.Nume, s.Deno)

	s.Nume /= lcm
	s.Deno /= lcm
}

// Add Fraction addition
func (s *FAL) Add(f FAL) *FAL {
	// Getting the Minimum Common Multiplier 获取最小公倍数
	lcm := Lcm(f.Deno, s.Deno)
	s.broad(lcm)
	f.broad(lcm)

	s.Nume += f.Nume
	s.offset()
	return s
}

// Sub fraction subtraction
func (s *FAL) Sub(f FAL) *FAL {
	// Getting the Minimum Common Multiplier 获取最小公倍数
	lcm := Lcm(s.Deno, f.Deno)
	s.broad(lcm)
	f.broad(lcm)

	s.Nume -= f.Nume
	s.offset()
	return s
}

// Mul multiplication
func (s *FAL) Mul(f FAL) *FAL { // 乘法
	s.Deno *= f.Deno
	s.Nume *= f.Nume
	s.offset()
	return s
}

// Div division
func (s *FAL) Div(f FAL) *FAL { // 除法
	s.Mul(Model(f.Deno, f.Nume))
	s.offset()
	return s
}

// Verdict Calculation results
func (s *FAL) Verdict() float64 { // 计算结果
	return float64(s.Nume) / float64(s.Deno)
}

func (s *FAL)Simplify() *FAL {
	tmp:=Model(s.Nume, int64(s.Deno))
	return tmp.Mul(Model(1, 1))
}