package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"

	"github.com/Edgaru089/imgui-go/v4"
)

var plotSize = imgui.Vec2{X: -1, Y: 200}

// floatRange returns a float64 in range [min, max).
func floatRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func showLine() {
	if imgui.BeginPlotV("Lines", plotSize, 0) {
		imgui.PlotLine("Numbers", []float64{7, 2, 4, 9, 1, 2, 4, 0, 2, 5})
		imgui.PlotLineP("Coords",
			[]imgui.Point{
				{X: 3, Y: -1},
				{X: 4, Y: 3},
				{X: 5, Y: 4},
				{X: 5, Y: 0},
				{X: 6, Y: -2},
				{X: 7, Y: 8},
			},
		)
		imgui.PlotLineG("Sine", func(userData interface{}, idx int) imgui.Point {
			return imgui.Point{X: float64(idx) / 20, Y: math.Sin(float64(idx)/5)*4 + 4}
		}, nil, 200)
		imgui.EndPlot()
	}
}

var (
	shadedShowLines = true
	shadedShowFills = true
	shadedFillRef   = float32(0)
)

var shadedAlpha float32 = 0.25

func showShaded() {
	rand.Seed(0)
	var xs, s1, s2, s3 [101]float64
	for i := 0; i < 101; i++ {
		xs[i] = float64(i)
		s1[i] = floatRange(400, 450)
		s2[i] = floatRange(275, 350)
		s3[i] = floatRange(150, 225)
	}
	imgui.DragFloatV("Alpha", &shadedAlpha, 0.01, 0, 1, "%.2g", 0)
	imgui.Checkbox("Lines", &shadedShowLines)
	imgui.SameLine()
	imgui.Checkbox("Fills", &shadedShowFills)
	if shadedShowFills {
		imgui.SameLine()
		imgui.SetNextItemWidth(100)
		imgui.DragFloatV("##Ref", &shadedFillRef, 1, -100, 500, "%.2f", 0)
	}

	if imgui.BeginPlotV("Stock Prices", plotSize, 0) {
		if shadedShowFills {
			imgui.PushPlotStyleVar(imgui.PlotStyleVar_FillAlpha, shadedAlpha)
			imgui.PlotShadedRefXY("Stock 1", xs, s1, float64(shadedFillRef))
			imgui.PlotShadedRefXY("Stock 2", xs, s2, float64(shadedFillRef))
			imgui.PlotShadedRefXY("Stock 3", xs, s3, float64(shadedFillRef))
			imgui.PopPlotStyleVar()
		}
		if shadedShowLines {
			imgui.PlotLineXY("Stock 1", xs, s1)
			imgui.PlotLineXY("Stock 2", xs, s2)
			imgui.PlotLineXY("Stock 3", xs, s3)
		}
		imgui.EndPlot()
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

	if imgui.BeginPlotV("Shaded Plots", plotSize, 0) {
		imgui.PushPlotStyleVar(imgui.PlotStyleVar_FillAlpha, shadedAlpha)
		imgui.PlotShadedLinesXY("Uncertain Data", xs, ys1, ys2)
		imgui.PlotLineXY("Uncertain Data", xs, ys)
		imgui.PlotShadedLinesXY("Overlapping", xs, ys3, ys4)
		imgui.PlotLineXY("Overlapping", xs, ys3)
		imgui.PlotLineXY("Overlapping", xs, ys4)
		imgui.PopPlotStyleVar()
		imgui.EndPlot()
	}
}

func showScatter() {
	rand.Seed(0)
	var s1 [100]imgui.Point
	var s2 [50]imgui.Point
	for i := 0; i < 100; i++ {
		s1[i].X = float64(i) * 0.01
		s1[i].Y = s1[i].X + 0.1*rand.Float64()
	}
	for i := 0; i < 50; i++ {
		s2[i].X = 0.25 + 0.2*rand.Float64()
		s2[i].Y = 0.75 + 0.2*rand.Float64()
	}

	if imgui.BeginPlotV("Scatter", plotSize, 0) {
		imgui.PlotScatterP("Data 1", s1[:])
		imgui.PushPlotStyleVar(imgui.PlotStyleVar_FillAlpha, 0.25)
		imgui.SetNextMarkerStyle(imgui.Marker_Square, 6, imgui.AutoColor, imgui.Auto, imgui.AutoColor)
		imgui.PlotScatterP("Data 2", s2[:])
		imgui.PopPlotStyleVar()
		imgui.EndPlot()
	}
}

func showStairs() {
	var s1, s2 [101]float64
	for i := 0; i < 101; i++ {
		s1[i] = 0.5 + 0.4*math.Sin(50*float64(i)*0.01)
		s2[i] = 0.5 + 0.2*math.Sin(25*float64(i)*0.01)
	}
	if imgui.BeginPlotV("Stairstep Plot", plotSize, 0) {
		imgui.PlotStairsV("Signal 1", s1, 0.01, 0)
		imgui.SetNextMarkerStyle(imgui.Marker_Square, 2, imgui.AutoColor, imgui.Auto, imgui.AutoColor)
		imgui.PlotStairsV("Signal 2", s2, 0.01, 0)
		imgui.EndPlot()
	}
}

func showBars() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if imgui.BeginPlotV("Bar Plot", plotSize, 0) {
		imgui.PlotBarsV("Bars", data, 0.7, 1)
		imgui.PlotBarsHV("BarsH", data, 0.4, 1)
		imgui.EndPlot()
	}
}

var (
	showBarGroupsStacked    = false
	showBarGroupsHorizontal = false
)

func showBarGroups() {
	data := [][]float64{
		{83, 67, 23, 89, 83, 78, 91, 82, 85, 90},
		{80, 62, 56, 99, 55, 78, 88, 78, 90, 100},
		{80, 69, 52, 92, 72, 78, 75, 76, 89, 95},
	}
	ilabels := []string{"Midterm", "Final", "Course"}
	glabels := []string{"S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S10"}
	positions := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	imgui.Checkbox("Stacked", &showBarGroupsStacked)
	imgui.SameLine()
	imgui.Checkbox("Horizontal", &showBarGroupsHorizontal)

	var flags imgui.BarGroupsFlags
	if showBarGroupsStacked {
		flags |= imgui.BarGroupsFlags_Stacked
	}

	if imgui.BeginPlotV("##BarGroups", plotSize, 0) {
		imgui.SetupLegend(imgui.Location_East, imgui.LegendFlags_Outside)
		if showBarGroupsHorizontal {
			imgui.SetupAxes("Score", "Student", imgui.AxisFlags_AutoFit, imgui.AxisFlags_AutoFit)
			imgui.SetupAxisTickValues(imgui.Axis_Y1, positions, glabels, false)
			imgui.PlotBarGroupsH(ilabels, data, 0.67, 0, flags)
		} else {
			imgui.SetupAxes("Student", "Score", imgui.AxisFlags_AutoFit, imgui.AxisFlags_AutoFit)
			imgui.SetupAxisTickValues(imgui.Axis_X1, positions, glabels, false)
			imgui.PlotBarGroups(ilabels, data, 0.67, 0, flags)
		}
		imgui.EndPlot()
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
	if imgui.BeginPlotV("Log Plot", plotSize, 0) {
		imgui.SetupAxis(imgui.Axis_X1, "", imgui.AxisFlags_LogScale)
		imgui.SetupAxesLimits(0.1, 100, 0, 10, imgui.Cond(imgui.ConditionOnce))
		imgui.PlotLineXY("f(x) = x", xs, xs)
		imgui.PlotLineXY("f(x) = sin(x)+1", xs, ys1)
		imgui.PlotLineXY("f(x) = log(x)", xs, ys2)
		imgui.PlotLineXY("f(x) = 10^x", xs, ys3)
		imgui.EndPlot()
	}
}

func showTickLabels() {
	var metric imgui.Formatter = func(val float64, userData interface{}) string {
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

	if imgui.BeginPlot("##Ticks") {
		imgui.SetupAxesLimits(2.5, 5.0, 0, 1000, imgui.Cond(imgui.ConditionOnce))
		imgui.SetupAxis(imgui.Axis_Y2, "", imgui.AxisFlags_AuxDefault)
		imgui.SetupAxis(imgui.Axis_Y3, "", imgui.AxisFlags_AuxDefault)
		if showTickLabelsCustomFmt {
			imgui.SetupAxisFormat(imgui.Axis_X1, "%.3g ms")
			imgui.SetupAxisFormatCallback(imgui.Axis_Y1, metric, "Hz")
			imgui.SetupAxisFormat(imgui.Axis_Y2, "%.3g dB")
			imgui.SetupAxisFormatCallback(imgui.Axis_Y3, metric, "m")
		}
		if showTickLabelsCustomTicks {
			if showTickLabelsCustomLabels {
				imgui.SetupAxisTickValues(imgui.Axis_X1, []float64{3.14}, []string{"Pi"}, true)
				imgui.SetupAxisTickValues(imgui.Axis_Y1, yticks, ylabels, false)
				imgui.SetupAxisTickValues(imgui.Axis_Y2, yticks2, ylabels2, false)
				imgui.SetupAxisTickRange(imgui.Axis_Y3, 0, 1, 6, ylabels2, false)
			} else {
				imgui.SetupAxisTickValues(imgui.Axis_X1, []float64{3.14}, nil, true)
				imgui.SetupAxisTickValues(imgui.Axis_Y1, yticks, nil, false)
				imgui.SetupAxisTickValues(imgui.Axis_Y2, yticks2, nil, false)
				imgui.SetupAxisTickRange(imgui.Axis_Y3, 0, 1, 6, nil, false)
			}
		}
		imgui.EndPlot()
	}
}

func example() {
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 400, Y: 600}, imgui.ConditionAppearing)
	if imgui.Begin("ImPlot-Go example") {

		imgui.Text(fmt.Sprintf("ImPlot-Go says hello. (%s)\ncompiled by %s/%s [%s/%s]", imgui.PlotVersion(), runtime.Compiler, runtime.Version(), runtime.GOOS, runtime.GOARCH))

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
				if imgui.CollapsingHeader("Bar Plots") {
					showBars()
				}
				if imgui.CollapsingHeader("Bar Groups") {
					showBarGroups()
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
