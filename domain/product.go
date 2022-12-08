package domain

// domain model을 정의합니다.

type Product struct {
	Code  string  `json:"code" bson:"code"`
	Name  string  `json:"name" bson:"name"`
	Price float32 `json:"price" bson:"price"`
}
