package dtos

type GetCommonFriendsOutput struct {
	BaseResult
	Count   int
	Friends []string
}
