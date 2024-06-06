package main

import (
	"fmt"
	"math"
)

// Point представляет структуру с координатами x и y.
type Point struct {
	x float64
	y float64
}

// NewPoint является конструктором для создания нового экземпляра Point.
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance рассчитывает расстояние между двумя точками по формуле Пифагора
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt((other.x-p.x)*(other.x-p.x) + (other.y-p.y)*(other.y-p.y))
}

func main() {
	// Создаем две точки
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	// Вычисляем расстояние между ними
	distance := p1.Distance(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
