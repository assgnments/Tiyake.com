package upvote

import "teyake/entity"

// UpVoteRepository specifies customer UpVote related database operations
type UpVoteRepository interface {
	UpVotes() ([]entity.UpVote, []error)
	UpVote(id uint) (*entity.UpVote, []error)
	UpVoteByAnswer(id uint) (*[]entity.UpVote, []error)
	DeleteUpVote(id uint) (*entity.UpVote, []error)
	StoreUpVote(UpVote *entity.UpVote) (*entity.UpVote, []error)
}
