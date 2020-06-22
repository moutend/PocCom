package types

type TreeWalkerDirection uintptr

const (
	TW_NONE        TreeWalkerDirection = 0
	TW_NEXT        TreeWalkerDirection = 1
	TW_PREVIOUS    TreeWalkerDirection = 2
	TW_FIRST_CHILD TreeWalkerDirection = 3
	TW_LAST_CHILD  TreeWalkerDirection = 4
	TW_PARENT      TreeWalkerDirection = 5
)
