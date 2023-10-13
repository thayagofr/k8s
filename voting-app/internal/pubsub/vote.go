package pubsub

type category string

func (c category) string() string {
	return string(c)
}

const (
	cat category = "Cat"
	dog category = "Dog"
)

type Vote struct {
	Category category `json:"category"`
}

func NewVote(vote string) Vote {
	if vote == cat.string() {
		return Vote{Category: cat}
	}
	return Vote{Category: dog}
}
