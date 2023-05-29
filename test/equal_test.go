package test

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringEqual(b *testing.B) {
	s1 := strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10)
	s2 := strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10)
	a := 0
	for i := 0; i < b.N; i++ {
		if s1 == s2 {
			a = 1
		}
	}
	_ = a
}

func BenchmarkStringCompare(b *testing.B) {
	s1 := strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10)
	s2 := strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10)
	a := 0
	for i := 0; i < b.N; i++ {
		if 1 == strings.Compare(s1, s2) {
			a = 1
		}
	}
	_ = a
}

func BenchmarkBytesCompare(b *testing.B) {
	s1 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	s2 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	a := 0
	for i := 0; i < b.N; i++ {
		if 1 == bytes.Compare(s1, s2) {
			a = 1
		}
	}
	_ = a
}

func BenchmarkBytesToStringCompare(b *testing.B) {
	s1 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	s2 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	a := 0
	for i := 0; i < b.N; i++ {
		if 1 == bytes.Compare(s1, s2) {
			a = 1
		}
	}
	_ = a
}

func BenchmarkBytesEqual(b *testing.B) {
	s1 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	s2 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	a := 0
	for i := 0; i < b.N; i++ {
		if bytes.Equal(s1, s2) {
			a = 1
		}
	}
	_ = a
}

func BenchmarkBytesIterate(b *testing.B) {
	s1 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	s2 := []byte(strings.Repeat("adfweoiuHKjkjfdklajio342987439jkjfdsapoifuHjkfdkasj2349&*&^&*^(*^&*&)*&&*^&*^&*(&(*&))", 10))
	a := 0
	for i := 0; i < b.N; i++ {
		if len(s1) != len(s2) {
			break
		}
		for i, b := range s1 {
			if s2[i] != b {
				break
			}
		}
		a = 1
	}
	_ = a
}
