package repository

import (
	"teyake/entity"
	"teyake/upvote"

	"github.com/jinzhu/gorm"
)

// UpVoteGormRepo implements menu.UpVoteRepository interface
type UpVoteGormRepo struct {
	conn *gorm.DB
}

// NewUpVoteGormRepo returns new object of UpVoteGormRepo
func NewUpVoteGormRepo(db *gorm.DB) upvote.UpVoteRepository {
	return &UpVoteGormRepo{conn: db}
}

// UpVotes returns all customer UpVotes stored in the database
func (UpVoteRepo *UpVoteGormRepo) UpVotes() ([]entity.UpVote, []error) {
	cmnts := []entity.UpVote{}
	errs := UpVoteRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// UpVote retrieves a customer UpVote from the database by its id
func (UpVoteRepo *UpVoteGormRepo) UpVote(id uint) (*entity.UpVote, []error) {
	cmnt := entity.UpVote{}
	errs := UpVoteRepo.conn.Set("gorm:auto_preload", true).First(&cmnt, id).GetErrors()
	return &cmnt, errs
}

// DeleteUpVote deletes a given customer UpVote from the database
func (UpVoteRepo *UpVoteGormRepo) DeleteUpVote(id uint) (*entity.UpVote, []error) {
	cmnt, errs := UpVoteRepo.UpVote(id)

	if len(errs) > 0 {
		return nil, errs
	}
	errs = UpVoteRepo.conn.Delete(cmnt, id).GetErrors()
	return cmnt, errs
}

// StoreUpVote stores a given customer UpVote in the database
func (UpVoteRepo *UpVoteGormRepo) StoreUpVote(UpVote *entity.UpVote) (*entity.UpVote, []error) {
	errs := UpVoteRepo.conn.Create(UpVote).GetErrors()
	return UpVote, errs
}
