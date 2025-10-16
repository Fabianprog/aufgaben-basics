package forms

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang des entsprechenden Rechtecks.
func PerimeterRectangle(a, b float64) float64 {
	c := 2*a + 2*b
	return c
}
