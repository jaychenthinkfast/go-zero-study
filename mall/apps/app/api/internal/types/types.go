// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package types

type Banner struct {
	ID   int64  `json:"id"`
	Name string `json:"name"` // 名称
	URL  string `json:"url"`  // 图片地址
}

type CartListRequest struct {
	UID int64 `form:"uid"`
}

type CartListResponse struct {
	Products []*CartProduct `json:"products"`
}

type CartProduct struct {
	Product *Product `json:"product"`
	Count   int64    `json:"count"` // 购买数量
}

type CategoryListRequest struct {
	Cursor   int64  `form:"cursor"`        // 分页游标
	Ps       int64  `form:"ps,default=20"` // 每页大小
	Category string `form:"category"`      // 分类
	Sort     string `form:"sort"`          // 排序
}

type CategoryListResponse struct {
	Products []*Product `json:"products"`
	IsEnd    bool       `json:"is_end"`
	LastVal  int64      `json:"last_val"`
}

type Comment struct {
	ID         int64    `json:"id"`          // 评论ID
	ProductID  int64    `json:"product_id"`  // 商品ID
	Content    string   `json:"content"`     // 评论内容
	Images     []*Image `json:"images"`      // 评论图片
	User       *User    `json:"user"`        // 用户信息
	CreateTime int64    `json:"create_time"` // 评论时间
	UpdateTime int64    `json:"update_time"` // 更新时间
}

type FlashSaleResponse struct {
	StartTime int64      `json:"start_time"` // 抢购开始时间
	Products  []*Product `json:"products"`
}

type HomeBannerResponse struct {
	Banners []*Banner `json:"banners"`
}

type Image struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type Order struct {
	OrderID            string  `json:"order_id"`
	Status             int32   `json:"status"`
	Quantity           int64   `json:"quantity"`
	Payment            float64 `json:"payment"`
	TotalPrice         float64 `json:"total_price"`
	CreateTime         int64   `json:"create_time"`
	ProductID          int64   `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductImage       string  `json:"product_image"`
	ProductDescription string  `json:"product_description"`
}

type OrderListRequest struct {
	UID    int64 `form:"uid"`
	Status int32 `form:"status,optional"`
	Cursor int64 `form:"cursor,optional"`
	Ps     int64 `form:"ps,default=20"`
}

type OrderListResponse struct {
	Orders    []*Order `json:"orders"`
	IsEnd     bool     `json:"is_end"` // 是否最后一页
	OrderTime int64    `json:"order_time"`
}

type Product struct {
	ID          uint64  `json:"id"`          // 商品ID
	Name        string  `json:"name"`        // 产品名称
	Description string  `json:"description"` // 商品描述
	Price       float64 `json:"price"`       // 商品价格
	Stock       int64   `json:"stock"`       // 库存
	Category    string  `json:"category"`    // 分类
	Status      int64   `json:"status"`      // 状态：1-正常，2-下架
	CreateTime  int64   `json:"create_time"` // 创建时间
	UpdateTime  int64   `json:"update_time"` // 更新时间
}

type ProductCommentRequest struct {
	ProductID int64 `form:"product_id"`
	Cursor    int64 `form:"cursor"`
	Ps        int64 `form:"ps,default=20"`
}

type ProductCommentResponse struct {
	Comments    []*Comment `json:"comments"`
	IsEnd       bool       `json:"is_end"`       // 是否最后一页
	CommentTime int64      `json:"comment_time"` // 评论列表最后一个评论的时间
}

type RecommendRequest struct {
	Cursor int64 `json:"cursor"`
	Ps     int64 `form:"ps,default=20"` // 每页大小
}

type RecommendResponse struct {
	Products      []*Product `json:"products"`
	IsEnd         bool       `json:"is_end"`         // 是否最后一页
	RecommendTime int64      `json:"recommend_time"` // 商品列表最后一个商品的推荐时间
}

type User struct {
	ID     int64  `json:"id"`     // 用户ID
	Name   string `json:"name"`   // 用户名
	Avatar string `json:"avatar"` // 头像
}
