package preprocessor

import (
	"bufio"
	"bytes"
	"strings"
)

func Process(text []byte) string {
	sc := bufio.NewScanner(bytes.NewReader(text))

	out := bytes.NewBuffer(nil)

	for sc.Scan() {
		i := strings.Index(sc.Text(), "//")

		if i == -1 {
			out.WriteString(sc.Text() + "\n")
		} else {
			out.WriteString(sc.Text()[:i] + "\n")
		}
	}

	return out.String()
}
