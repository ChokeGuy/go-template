package dto

type ProductsListResponse struct {
	TotalCount int64              `json:"totalCount" bson:"totalCount"`
	TotalPages int64              `json:"totalPages" bson:"totalPages"`
	Page       int64              `json:"page" bson:"page"`
	Size       int64              `json:"size" bson:"size"`
	HasMore    bool               `json:"hasMore" bson:"hasMore"`
	Products   []*ProductResponse `json:"products" bson:"products"`
}
