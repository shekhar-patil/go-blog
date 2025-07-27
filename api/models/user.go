package models
import(
	"time"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID uint32 					`gorm:"primary_key;auto_increment" json:"id"`
	FirstName string 			`gorm:"size:100;not null;" json:"first_name"`
	LastName string 			`gorm:"size:100;not null;" json:"last_name"`
	Email string 				`gorm:"size:100;not null;unique" json:"email"`
	Password string 			`gorm:"size:100;not null;" json:"password"`
	Posts []Post				`gorm:"foreignkey:AuthorID" json:"posts"`
	CreatedAt time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) FindAllUsers(db *gorm.DB) ([]User, error) {
	users := []User{}

	result := db.Find(&users)

	if result.Error != nil {
		return []User{}, errors.New(result.Error.Error())
	}

	return users, nil
}