package upvote

import "teyake/entity"

// UpVoteService specifies customer Answer related service
type UpVoteService interface {
	UpVotes() ([]entity.UpVote, []error)
	UpVote(id uint) (*entity.UpVote, []error) //needed
	DeleteUpVote(id uint) (*entity.UpVote, []error)
	StoreUpVote(UpVote *entity.UpVote) (*entity.UpVote, []error)
}
