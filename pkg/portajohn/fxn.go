package portajohn

import (
	"regexp"
	"strings"
)

func FormatString(s string) string {
	rgx2 := regexp.MustCompile(`\s+`)
	s = rgx2.ReplaceAllString(s, " ")
	s = strings.TrimSpace(strings.ToUpper(s))
	return s
} // ./FormatString

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

// FormatEmails formats email addresses
// if delimiter found, will return multiple emails
func FormatEmails(s string) (string, []string) {
	splt := strings.Split(s, ",")
	if len(splt) == 0 {
		splt := strings.Split(s, ";")
		if len(splt) == 0 {
			return "", nil
		}
		return formatEmails(splt)
	}
	return formatEmails(splt)
} // ./FormatEmails

func formatEmails(splt []string) (string, []string) {
	ss := []string{}
	primaryEmail := ""
	for i, e := range splt {
		e = strings.ToLower(e)
		if i == 0 {
			primaryEmail = e
			continue
		}
		ss = append(ss, e)
		return primaryEmail, ss
	}
	return primaryEmail, ss
} // ./formatEmails
