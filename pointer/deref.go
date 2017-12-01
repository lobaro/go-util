package pointer

func ToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
