package pkg

type SearchUsersParams struct {
	Page  *int // 页码
	Limit *int // 每页条数

	Account   *string // 账号
	Name      *string // 姓名
	Direction *int    // 方向
}
