package data

import (
	"errors"
	"log"
	"sirloinapi/features/user"

	"gorm.io/gorm"
)

type userQry struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQry{
		db: db,
	}
}

func (uq *userQry) CheckUser(newUser user.Core) error {
	u := User{}
	uq.db.Where("email = ? OR phone_number = ?", newUser.Email, newUser.PhoneNumber).First(&u)
	if u.ID != 0 {
		if u.Email == newUser.Email {
			return errors.New("duplicate users.email")
		} else {
			return errors.New("duplicate users.phone_number")
		}
	}
	return nil
}

func (uq *userQry) Register(newUser user.Core) (user.Core, error) {
	// Chek User
	if err := uq.CheckUser(newUser); err != nil {
		log.Println("error create new user: ", err.Error())
		return user.Core{}, err
	}

	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("error create query: ", err.Error())
		return user.Core{}, err
	}

	newUser.ID = cnv.ID

	return newUser, nil
}

func (uq *userQry) Login(email string) (user.Core, error) {
	res := User{}

	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error: ", err.Error())
		return user.Core{}, errors.New("user not found")
	}
	return ToCore(res), nil
}

func (uq *userQry) Profile(id uint) (user.Core, error) {
	res := User{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.Core{}, err
	}

	return ToCore(res), nil
}

func (uq *userQry) Update(id uint, updateData user.Core) (user.Core, error) {
	// Chek User
	res, err := uq.Profile(id)
	if err != nil {
		log.Println("\tupdate user query error: ", err.Error())
		return user.Core{}, errors.New("not found")
	}
	if res.Email == updateData.Email {
		updateData.Email = ""
	}
	if res.PhoneNumber == updateData.PhoneNumber {
		updateData.PhoneNumber = ""
	}

	if err := uq.CheckUser(updateData); err != nil {
		log.Println("error update user: ", err.Error())
		return user.Core{}, err
	}
	cnvUpd := CoreToData(updateData)
	qry := uq.db.Where("id = ?", id).Updates(cnvUpd)
	if qry.RowsAffected <= 0 {
		log.Println("\tupdate user query error: data not found")
		return user.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("\tupdate user query error: ", err.Error())
		return user.Core{}, errors.New("not found")
	}
	return ToCore(cnvUpd), nil
}

func (uq *userQry) Delete(id uint) error {
	user := User{}
	qry := uq.db.Where("id = ?", id).Delete(&user)
	if affrows := qry.RowsAffected; affrows <= 0 {
		return errors.New("user doesn't exist")
	}
	if err := qry.Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (uq *userQry) RegisterDevice(id uint, dvcToken string) error {
	dt := DeviceToken{}
	uq.db.First(&dt, "token = ?", dvcToken)
	if dt.Token != "" {
		return errors.New("duplicated")
	}

	dt = DeviceToken{UserId: id, Token: dvcToken}
	err := uq.db.Create(&dt).Error
	if err != nil {
		log.Println("error registering device token: ", err.Error())
		return errors.New("error registering device token: " + err.Error())
	}

	return nil
}
func (uq *userQry) UnregDevice(id uint) error {
	dt := DeviceToken{}
	err := uq.db.Where("user_id = ?", id).Delete(dt).Error
	if err != nil {
		log.Println("error delete device token: ", err)
		return errors.New("error delete device token: ")
	}

	return nil
}
