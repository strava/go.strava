package strava

type Polyline string

// Decode will take the polyline which is a string
// in standard Google polyline encoding and convert it to an array.
func (p Polyline) Decode() [][2]float64 {
	var count, index int
	factor := 1.0e5

	line := make([][2]float64, 0)
	tempLatLng := [2]int{0, 0}

	for index < len(p) {
		var result int
		var b int = 0x20
		var shift uint

		for b >= 0x20 {
			b = int(p[index]) - 63
			index++

			result |= (b & 0x1f) << shift
			shift += 5
		}

		// sign dection
		if result&1 != 0 {
			result = ^(result >> 1)
		} else {
			result = result >> 1
		}

		if count%2 == 0 {
			result += tempLatLng[0]
			tempLatLng[0] = result
		} else {
			result += tempLatLng[1]
			tempLatLng[1] = result

			line = append(line, [2]float64{float64(tempLatLng[0]) / factor, float64(tempLatLng[1]) / factor})
		}

		count++
	}

	return line
}
