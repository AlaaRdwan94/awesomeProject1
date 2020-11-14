package models

import(
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

//this will be the Companies structure
type Company struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Promotion string    `gorm:"size:100;not null;unique" json:"promotion"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}


//this function will hash password before save it
func (u *Company) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

//this function will prepare the Company object before save it
func (u *Company) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

//this function validate the http requeest for Company CRUD
func (u *Company) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

//this function will add Company to database
func (u *Company) Save(db *gorm.DB) error {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

//this function will get first 100 Companies in the database
func (u *Company) FindAll(db *gorm.DB) (*[]Company, error) {
	var err error
	Companies := []Company{}
	err = db.Debug().Model(&Company{}).Limit(100).Find(&Companies).Error
	if err != nil {
		return &[]Company{}, err
	}
	return &Companies, err
}

//this function will get Company by id
func (u *Company) FindByID(db *gorm.DB, uid uint32)  error {
	var err error
	err = db.Debug().Model(Company{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		return  errors.New("Company Not Found")
	}
	return  err
}
//this function will update Company
func (u *Company) UpdatePromotion(db *gorm.DB, uid uint32,promotion string) error {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&Company{}).UpdateColumns(
		map[string]interface{}{
			"promotion":   promotion,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return  db.Error
	}
	// This is the display the updated Company
	err = db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return err
	}
	return  nil
}
//this function will update Company
func (u *Company) Update(db *gorm.DB, uid uint32) error {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&Company{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"name":       u.Name,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return db.Error
	}
	// This is the display the updated Company
	err = db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return  err
	}
	return nil
}

//this will delete Company from the database
func (u *Company) Delete(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&Company{}).Delete(&Company{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
