package models

import (
	"errors"
	"time"

	"github.com/koding/bongo"
)

type Interaction struct {
	// unique identifier of the Interaction
	Id int64 `json:"id"`

	// Id of the interacted message
	MessageId int64 `json:"messageId"             sql:"NOT NULL"`

	// Id of the actor
	AccountId int64 `json:"accountId"             sql:"NOT NULL"`

	// Type of the interaction
	TypeConstant string `json:"typeConstant"      sql:"NOT NULL;TYPE:VARCHAR(100);"`

	// Creation of the interaction
	CreatedAt time.Time `json:"createdAt"         sql:"NOT NULL"`
}

var AllowedInteractions = map[string]struct{}{
	"like":     struct{}{},
	"upvote":   struct{}{},
	"downvote": struct{}{},
}

const (
	Interaction_TYPE_LIKE     = "like"
	Interaction_TYPE_UPVOTE   = "upvote"
	Interaction_TYPE_DONWVOTE = "downvote"
)

func (i *Interaction) GetId() int64 {
	return i.Id
}

func (i *Interaction) TableName() string {
	return "interaction"
}

func NewInteraction() *Interaction {
	return &Interaction{}
}

func (i *Interaction) Fetch() error {
	return bongo.B.Fetch(i)
}

func (i *Interaction) Create() error {
	return bongo.B.Create(i)
}

func (i *Interaction) AfterCreate() {
	bongo.B.AfterCreate(i)
}

func (i *Interaction) AfterUpdate() {
	bongo.B.AfterUpdate(i)
}

func (i *Interaction) AfterDelete() {
	bongo.B.AfterDelete(i)
}

func (i *Interaction) Delete() error {
	if err := bongo.B.DB.
		Where("message_id = ? and account_id = ?", i.MessageId, i.AccountId).
		Delete(NewInteraction()).Error; err != nil {
		return err
	}
	return nil
}

func (c *Interaction) List(interactionType string) ([]int64, error) {
	var interactions []int64

	if c.MessageId == 0 {
		return interactions, errors.New("ChannelId is not set")
	}

	if err := bongo.B.DB.Table(c.TableName()).
		Where("message_id = ?", c.MessageId).
		Pluck("account_id", &interactions).
		Error; err != nil {
		return nil, err
	}

	if interactions == nil {
		return make([]int64, 0), nil
	}

	// change this part to use c.m.some

	// selector := map[string]interface{}{
	// 	"message_id": c.MessageId,
	// }

	// pluck := map[string]interface{}{
	// 	"account_id": true,
	// }

	// err := c.m.Some(c, &interactions, selector, nil, pluck)
	// if err != nil && err != gorm.RecordNotFound {
	// 	return nil, err
	// }

	return interactions, nil
}
