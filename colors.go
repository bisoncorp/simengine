package simengine

type Color struct {
	Red, Green, Blue, Alpha float32
}

var (
	Black  = Color{Red: 0.207, Green: 0.288, Blue: 0.373, Alpha: 0}
	Blue   = Color{Red: 0.207, Green: 0.596, Blue: 0.858, Alpha: 0}
	Gray   = Color{Red: 0.587, Green: 0.651, Blue: 0.654, Alpha: 0}
	Green  = Color{Red: 0.180, Green: 0.796, Blue: 0.443, Alpha: 0}
	Orange = Color{Red: 0.902, Green: 0.502, Blue: 0.131, Alpha: 0}
	Purple = Color{Red: 0.607, Green: 0.354, Blue: 0.718, Alpha: 0}
	Red    = Color{Red: 0.906, Green: 0.301, Blue: 0.235, Alpha: 0}
	White  = Color{Red: 0.925, Green: 0.941, Blue: 0.942, Alpha: 0}
	Yellow = Color{Red: 0.957, Green: 0.772, Blue: 0.059, Alpha: 0}
)
