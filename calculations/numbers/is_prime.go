package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	div := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			div++
		}

	}
	return div == 2
}
