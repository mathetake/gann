package item

// Vector ... vector
type Vector []float32

// Item ... item
type Item struct {
	ID  int64  `json:"id"`
	Vec Vector `json:"vec"`
}
