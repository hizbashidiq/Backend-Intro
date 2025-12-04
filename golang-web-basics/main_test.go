package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

var (
  cube Cube = Cube{4}
  expectedVolume float64 = 64
  expectedLuas float64 = 96
  expectedKeliling float64 = 48
)

// in Go, the test function should be started with "Test", parameter *testing.T
// there's a lot of method in testing package like Fail(), Log(), Logf(), etc

func TestVolume( t *testing.T){
  t.Log(cube.Volume())

  if cube.Volume() != expectedVolume{
    t.Errorf("Want: %.2f Got: %.2f", expectedVolume, cube.Volume())
  }
}
func TestKeliling( t *testing.T){
  t.Log(cube.Keliling())

  if cube.Keliling() != expectedKeliling{
    t.Errorf("Want: %.2f Got: %.2f", expectedKeliling, cube.Keliling())
  }
}
func TestLuas( t *testing.T){
  t.Log(cube.Luas())

  if cube.Luas() != expectedLuas{
    t.Errorf("Want: %.2f Got: %.2f", expectedLuas, cube.Luas())
  }
}

// in Go, the benchmark function should be started with "Benchmark", parameter *testing.B
func BenchmarkLuas(b *testing.B){
  for i:=0;i<b.N;i++{
    cube.Luas()
  }
}

func TestAssertVolume(t *testing.T){
  assert.Equal(t, cube.Volume(), expectedVolume, "WRONG VOLUME!")
}
func TestAssertKeliling(t *testing.T){
  assert.Equal(t, cube.Keliling(), expectedKeliling, "WRONG KELILING!")
}
func TestAssertLuas(t *testing.T){
  assert.Equal(t, cube.Luas(), expectedLuas, "WRONG LUAS!")
}