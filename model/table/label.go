package table

type Label struct {
	Content   string `gorm:"type:varchar(20);not null" `
	OwnerType string `gorm:"type:varchar(20);not null" json:"-"`
	OwnerID   uint   `gorm:",check:owner_id > 0" json:"-"`
	Model
}
