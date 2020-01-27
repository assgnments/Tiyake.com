package service

import (
	"teyake/entity"
	"teyake/upvote"
)

// UpVoteService implements menu.UpVoteService interface
type UpVoteService struct {
	UpVoteRepo upvote.UpVoteRepository
}

// NewUpVoteService returns a new UpVoteService object
func NewUpVoteService(commRepo upvote.UpVoteRepository) upvote.UpVoteService {
	return &UpVoteService{UpVoteRepo: commRepo}
}

// UpVotes returns all stored UpVotes
func (as *UpVoteService) UpVotes() ([]entity.UpVote, []error) {
	return as.UpVoteRepo.UpVotes()
}

// UpVote retrieves stored UpVote by its id
func (as *UpVoteService) UpVote(id uint) (*entity.UpVote, []error) {
	return as.UpVoteRepo.UpVote(id)
}

// DeleteUpVote deletes a given UpVote
func (as *UpVoteService) DeleteUpVote(id uint) (*entity.UpVote, []error) {
	return as.UpVoteRepo.DeleteUpVote(id)
}

// StoreUpVote stores a given UpVote
func (as *UpVoteService) StoreUpVote(UpVote *entity.UpVote) (*entity.UpVote, []error) {
	return as.UpVoteRepo.StoreUpVote(UpVote)

}
