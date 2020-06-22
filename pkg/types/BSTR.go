package types

type BSTR uintptr

// String returns BSTR as string.
func (p BSTR) String() string {
	return bstrString(p)
}
