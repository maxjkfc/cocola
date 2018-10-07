package wrap

// AutoWrapSline -
func AutoWrapSline(text string, limit int) []string {
	ns := make([]string, 0)
	//os := strings.Fields(strings.TrimSpace(text))

	var (
		temp       string
		shiftspace int
		tab        = "  "
	)

	for _, v := range text {

		if v == '\n' {

			ns = append(ns, temp)
			shiftspace = 0
			temp = ""

		} else {

			shiftspace++
			temp += string(v)

			if shiftspace == limit {
				ns = append(ns, temp)
				shiftspace = 0
				temp = tab

			}

		}
	}
	ns = append(ns, temp)

	return ns
}
