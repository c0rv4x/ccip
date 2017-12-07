package services

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-go/models"
)

type Store struct {
	models.ORM
	Scheduler *Scheduler
	Config    Config
}

func NewStore(config Config) Store {
	orm := models.InitORM(config.RootDir)
	return Store{
		ORM:       orm,
		Scheduler: NewScheduler(orm),
		Config:    config,
	}
}

func (self Store) Start() error {
	return self.Scheduler.Start()
}

func (self Store) Close() {
	self.Scheduler.Stop()
	self.ORM.Close()
}

func (self Store) AddJob(job models.Job) error {
	err := job.Validate()
	if err != nil {
		return err
	}

	err = self.Save(&job)
	if err != nil {
		return err
	}

	self.Scheduler.AddJob(job)
	return nil
}

func (self Store) AddPassword(password models.Password) error {
	return self.Save(&password)
}

func (self Store) JobRunsFor(job models.Job) ([]models.JobRun, error) {
	var runs []models.JobRun
	err := self.Where("JobID", job.ID, &runs)
	return runs, err
}

func (self Store) CreateKey(password string) (common.Address, error) {
	return keystore.StoreKey(
		self.Config.KeysDir(),
		password,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
}
