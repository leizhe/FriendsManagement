package dtos

type GetAllFriendsOutput struct {
	BaseResult
	Count   int64
	Friends []string
}
