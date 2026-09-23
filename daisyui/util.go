package daisyui

import "strings"

// classFor returns the daisyUI modifier class "<prefix>-<value>" (e.g.
// "btn-primary" for prefix "btn", value "primary"). An empty value produces
// no class (returns "").
func classFor(prefix, val string) string {
	val = strings.TrimSpace(val)
	if val == "" {
		return ""
	}
	return prefix + "-" + val
}

// joinClasses trims and joins the given classes with a single space, skipping
// empty entries. Used to build the final class attribute list.
func joinClasses(classes ...string) string {
	var out []string
	for _, c := range classes {
		c = strings.TrimSpace(c)
		if c != "" {
			out = append(out, c)
		}
	}
	return strings.Join(out, " ")
}

// glassClass returns "glass" when glass is true, otherwise "". Components that
// support a translucent, blurred surface (the glass effect) include it in
// their class list. The .glass class is defined in assets/styles.css.
func glassClass(glass bool) string {
	if glass {
		return "glass"
	}
	return ""
}
