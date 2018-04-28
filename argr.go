package argr

import "strings"

// Tokenize tokenizes string as command line args
func Tokenize(str string) []string {
	inq := false
	var last rune
	parts := strings.FieldsFunc(str, func(r rune) bool {
		if r == '"' && last != '\\' {
			inq = !inq
		}
		last = r
		return !inq && r == ' '
	})
	for k, v := range parts {
		sx := strings.HasSuffix(v, `"`)
		if strings.HasPrefix(v, `"`) && sx {
			v = v[1 : len(v)-1]
			v = strings.Replace(v, `\"`, `"`, -1)
		} else if sx {
			ix := strings.IndexRune(v, '"')
			if ix != -1 {
				v = v[0:ix] + v[ix+1:len(v)-1]
				v = strings.Replace(v, `\"`, `"`, -1)
			}
		}
		if strings.HasPrefix(v, `""`) && strings.HasSuffix(v, `""`) {
			v = strings.Trim(v, `"`)
		}
		if v == "" {
			continue
		}
		parts[k] = v
	}
	return parts
}
