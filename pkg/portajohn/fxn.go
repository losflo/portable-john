package portajohn

import (
	"regexp"
	"strings"
)

func FormatPhoneNumber(s string) string {
	s = strings.ToUpper(strings.TrimSpace(s))

	rgx1 := regexp.MustCompile(`[;:_\(\)\\\/\.-]+|FAX`)
	s = rgx1.ReplaceAllString(s, "")

	rgx2 := regexp.MustCompile(`\s+`)
	s = rgx2.ReplaceAllString(s, "")

	rgxPhone := regexp.MustCompile(`[0-9]{10}`)
	pn := rgxPhone.FindString(s)

	rgxExt := regexp.MustCompile(`(X|EXT)[0-9]*`)
	extDirty := rgxExt.FindString(s)

	if len(extDirty) > 0 {
		rgx := regexp.MustCompile(`[0-9]*`)
		ext := rgx.FindString(extDirty)
		pn += " X" + ext
	}
	return pn
} // ./FormatPhoneNumber
