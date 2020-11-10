package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func Convert(v []float64) plotter.Values {
	c := make(plotter.Values,len(v))
	c = v
	return c
}


func Draw(v ...[]float64) {

	groupA := Convert(v[0])
	groupB := Convert(v[1])
	groupC := Convert(v[2])

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.BackgroundColor = color.RGBA{R: 12, B: 28, G: 50, A: 55}

	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w

	barsB, err := plotter.NewBarChart(groupB, w)
	if err != nil {
		panic(err)
	}
	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(1)

	barsC, err := plotter.NewBarChart(groupC, w)
	if err != nil {
		panic(err)
	}
	barsC.LineStyle.Width = vg.Length(0)
	barsC.Color = plotutil.Color(2)
	barsC.Offset = w

	p.Add(barsA, barsB, barsC)
	p.Legend.Add("Group A", barsA)
	p.Legend.Add("Group B", barsB)
	p.Legend.Add("Group C", barsC)
	p.Legend.Top = true
	p.NominalX("One", "Two", "Three", "Four", "Five")

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.svg"); err != nil {
		panic(err)
	}
}




func main() {
  a := []float64{3,5,7}
  b := []float64{2,4,8}
  c := []float64{1,7, 3}

  Draw(a,b,c)
}
