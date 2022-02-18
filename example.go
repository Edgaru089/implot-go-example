package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"

	"github.com/Edgaru089/implot-go"
	"github.com/inkyblackness/imgui-go/v4"
)

var plotSize = imgui.Vec2{X: -1, Y: 200}

// floatRange returns a float64 in range [min, max).
func floatRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func showLine() {
	if implot.BeginPlotV("Lines", plotSize, 0) {
		implot.PlotLine("Numbers", []float64{7, 2, 4, 9, 1, 2, 4, 0, 2, 5})
		implot.PlotLineP("Coords",
			[]implot.Point{
				{X: 3, Y: -1},
				{X: 4, Y: 3},
				{X: 5, Y: 4},
				{X: 5, Y: 0},
				{X: 6, Y: -2},
				{X: 7, Y: 8},
			},
		)
		implot.PlotLineG("Sine", func(userData interface{}, idx int) implot.Point {
			return implot.Point{X: float64(idx) / 20, Y: math.Sin(float64(idx)/5)*4 + 4}
		}, nil, 200)
		implot.EndPlot()
	}
}

var (
	shadedShowLines = true
	shadedShowFills = true
	shadedFillRef   = float32(0)
)

func showShaded() {
	rand.Seed(0)
	var xs, s1, s2, s3 [101]float64
	for i := 0; i < 101; i++ {
		xs[i] = float64(i)
		s1[i] = floatRange(400, 450)
		s2[i] = floatRange(275, 350)
		s3[i] = floatRange(150, 225)
	}
	imgui.Checkbox("Lines", &shadedShowLines)
	imgui.SameLine()
	imgui.Checkbox("Fills", &shadedShowFills)
	if shadedShowFills {
		imgui.SameLine()
		imgui.SetNextItemWidth(100)
		imgui.DragFloatV("##Ref", &shadedFillRef, 1, -100, 500, "%.2f", 0)
	}

	if implot.BeginPlotV("Stock Prices", plotSize, 0) {
		if shadedShowFills {
			implot.PlotShadedRefXY("Stock 1", xs, s1, float64(shadedFillRef))
			implot.PlotShadedRefXY("Stock 2", xs, s2, float64(shadedFillRef))
			implot.PlotShadedRefXY("Stock 3", xs, s3, float64(shadedFillRef))
		}
		implot.EndPlot()
	}
}

func showShadedLines() {
	var xs, ys, ys1, ys2, ys3, ys4 [1001]float64
	rand.Seed(0)
	for i := 0; i < 1001; i++ {
		xs[i] = float64(i) * 0.001
		ys[i] = 0.25 + 0.25*math.Sin(25*xs[i])*math.Sin(5*xs[i]) + floatRange(-0.01, 0.01)
		ys1[i] = ys[i] + floatRange(0.1, 0.12)
		ys2[i] = ys[i] - floatRange(0.1, 0.12)
		ys3[i] = 0.75 + 0.2*math.Sin(25*xs[i])
		ys4[i] = 0.75 + 0.1*math.Cos(25*xs[i])
	}

	if implot.BeginPlotV("Shaded Plots", plotSize, 0) {
		implot.PlotShadedLinesXY("Uncertain Data", xs, ys1, ys2)
		implot.PlotLineXY("Uncertain Data", xs, ys)
		implot.PlotShadedLinesXY("Overlapping", xs, ys3, ys4)
		implot.PlotLineXY("Overlapping", xs, ys3)
		implot.PlotLineXY("Overlapping", xs, ys4)
		implot.EndPlot()
	}
}

func showScatter() {
	rand.Seed(0)
	var s1 [100]implot.Point
	var s2 [50]implot.Point
	for i := 0; i < 100; i++ {
		s1[i].X = float64(i) * 0.01
		s1[i].Y = s1[i].X + 0.1*rand.Float64()
	}
	for i := 0; i < 50; i++ {
		s2[i].X = 0.25 + 0.2*rand.Float64()
		s2[i].Y = 0.75 + 0.2*rand.Float64()
	}

	if implot.BeginPlotV("Scatter", plotSize, 0) {
		implot.PlotScatterP("Data 1", s1[:])
		implot.PlotScatterP("Data 2", s2[:])
		implot.EndPlot()
	}
}

func showStairs() {
	var s1, s2 [101]float64
	for i := 0; i < 101; i++ {
		s1[i] = 0.5 + 0.4*math.Sin(50*float64(i)*0.01)
		s2[i] = 0.5 + 0.2*math.Sin(25*float64(i)*0.01)
	}
	if implot.BeginPlotV("Stairstep Plot", plotSize, 0) {
		implot.PlotStairsV("Signal 1", s1, 0.01, 0)
		implot.PlotStairsV("Signal 2", s2, 0.01, 0)
		implot.EndPlot()
	}
}

var (
	showTickLabelsCustomFmt    = true
	showTickLabelsCustomTicks  = false
	showTickLabelsCustomLabels = true
)

func showLogAxes() {
	var xs, ys1, ys2, ys3 [1001]float64
	for i := 0; i < 1001; i++ {
		xs[i] = (float64)(i) * 0.1
		ys1[i] = math.Sin(xs[i]) + 1
		ys2[i] = math.Log(xs[i])
		ys3[i] = math.Pow(10, xs[i])
	}
	imgui.Bullet()
	imgui.Text("Open the plot context menu (right click) to change scales.")
	if implot.BeginPlotV("Log Plot", plotSize, 0) {
		implot.SetupAxis(implot.Axis_X1, "", implot.AxisFlags_LogScale)
		implot.SetupAxesLimits(0.1, 100, 0, 10, implot.Condition_Once)
		implot.PlotLineXY("f(x) = x", xs, xs)
		implot.PlotLineXY("f(x) = sin(x)+1", xs, ys1)
		implot.PlotLineXY("f(x) = log(x)", xs, ys2)
		implot.PlotLineXY("f(x) = 10^x", xs, ys3)
		implot.EndPlot()
	}
}

func showTickLabels() {
	var metric implot.Formatter = func(val float64, userData interface{}) string {
		unit := userData.(string)
		v := []float64{1000000000, 1000000, 1000, 1, 0.001, 0.000001, 0.000000001}
		p := []string{"G", "M", "k", "", "m", "u", "n"}
		if val == 0 {
			return "0 " + unit
		}
		for i := 0; i < 7; i++ {
			if math.Abs(val) >= v[i] {
				return fmt.Sprintf("%.6g %s%s", val/v[i], p[i], unit)
			}
		}
		return fmt.Sprintf("%.6g %s%s", val/v[6], p[6], unit)
	}

	imgui.Checkbox("Show Custom Format", &showTickLabelsCustomFmt)
	imgui.SameLine()
	imgui.Checkbox("Show Custom Ticks", &showTickLabelsCustomTicks)
	if showTickLabelsCustomTicks {
		imgui.SameLine()
		imgui.Checkbox("Show Custom Labels", &showTickLabelsCustomLabels)
	}

	yticks := []float64{100, 300, 700, 900}
	ylabels := []string{"One", "Three", "Seven", "Nine"}
	yticks2 := []float64{0.2, 0.4, 0.6}
	ylabels2 := []string{"A", "B", "C", "D", "E", "F"}

	if implot.BeginPlot("##Ticks") {
		implot.SetupAxesLimits(2.5, 5.0, 0, 1000, implot.Condition_Once)
		implot.SetupAxis(implot.Axis_Y2, "", implot.AxisFlags_AuxDefault)
		implot.SetupAxis(implot.Axis_Y3, "", implot.AxisFlags_AuxDefault)
		if showTickLabelsCustomFmt {
			implot.SetupAxisFormat(implot.Axis_X1, "%.3g ms")
			implot.SetupAxisFormatCallback(implot.Axis_Y1, metric, "Hz")
			implot.SetupAxisFormat(implot.Axis_Y2, "%.3g dB")
			implot.SetupAxisFormatCallback(implot.Axis_Y3, metric, "m")
		}
		if showTickLabelsCustomTicks {
			if showTickLabelsCustomLabels {
				implot.SetupAxisTickValues(implot.Axis_X1, []float64{3.14}, []string{"Pi"}, true)
				implot.SetupAxisTickValues(implot.Axis_Y1, yticks, ylabels, false)
				implot.SetupAxisTickValues(implot.Axis_Y2, yticks2, ylabels2, false)
				implot.SetupAxisTickRange(implot.Axis_Y3, 0, 1, 6, ylabels2, false)
			} else {
				implot.SetupAxisTickValues(implot.Axis_X1, []float64{3.14}, nil, true)
				implot.SetupAxisTickValues(implot.Axis_Y1, yticks, nil, false)
				implot.SetupAxisTickValues(implot.Axis_Y2, yticks2, nil, false)
				implot.SetupAxisTickRange(implot.Axis_Y3, 0, 1, 6, nil, false)
			}
		}
		implot.EndPlot()
	}
}

func example() {
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 400, Y: 600}, imgui.ConditionAppearing)
	if imgui.Begin("ImPlot-Go example") {

		imgui.Text(fmt.Sprintf("ImPlot-Go says hello. (%s)\ncompiled by %s/%s [%s/%s]", implot.Version(), runtime.Compiler, runtime.Version(), runtime.GOOS, runtime.GOARCH))

		if imgui.BeginTabBar("MainTab") {
			if imgui.BeginTabItem("Plots") {

				if imgui.CollapsingHeader("Line Plots") {
					showLine()
				}
				if imgui.CollapsingHeader("Shaded Plots") {
					showShaded()
					showShadedLines()
				}
				if imgui.CollapsingHeader("Scatter Plots") {
					showScatter()
				}
				if imgui.CollapsingHeader("Stair Plots") {
					showStairs()
				}
				imgui.EndTabItem()
			}
			if imgui.BeginTabItem("Axes") {
				if imgui.CollapsingHeader("Log Scale Axes") {
					showLogAxes()
				}
				if imgui.CollapsingHeader("Tick Labels") {
					showTickLabels()
				}
				imgui.EndTabItem()
			}
			imgui.EndTabBar()
		}
	}
	imgui.End()
}
