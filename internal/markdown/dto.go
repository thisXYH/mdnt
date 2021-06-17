package markdown

type docInfo struct {
	absPath string
	name    string
	id      string
	content string

	docRefs []docRef
}

type docRef struct {
	original string //[alter](xxx/x.md?id)
	alter    string //alter
	id       string //id
}
