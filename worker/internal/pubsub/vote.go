package pubsub

type category string

const (
	cat category = "Cat"
	dog category = "Dog"
)

type Vote struct {
	Category category `json:"category"`
}
