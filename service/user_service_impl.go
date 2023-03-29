package service

import (
	"errors"
	"example.com/auction-api/entity"
	"example.com/auction-api/model"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type UserServiceImpl struct {
	db    *gorm.DB
	redis RedisService
}

func NewUserService(db *gorm.DB, redis RedisService) UserService {
	return &UserServiceImpl{
		db:    db,
		redis: redis,
	}
}

func (u *UserServiceImpl) Register(userRegisterRequest *model.UserRegister) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), 10)
	if err != nil {
		return err
	}
	user := entity.User{
		Email:     userRegisterRequest.Email,
		Password:  string(hash),
		FirstName: userRegisterRequest.FirstName,
		LastName:  userRegisterRequest.LastName,
	}
	result := u.db.Create(&user)
	if result.Error != nil {
		return errors.New("failed to create userRegisterRequest")
	}
	return err
}

func (u *UserServiceImpl) Login(userLoginRequest *model.UserLogin) (*entity.User, *string, error) {

	var user entity.User
	u.db.First(&user, "email = ?", userLoginRequest.Email)
	if user.ID == 0 {
		return nil, nil, errors.New("invalid email")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return nil, nil, errors.New("invalid password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, nil, err
	}
	err = u.redis.Put(strconv.Itoa(int(user.ID)), tokenString)
	if err != nil {
		return nil, nil, err
	}
	user.Password = ""

	return &user, &tokenString, nil

}

func (u *UserServiceImpl) Logout(id string) error {
	err := u.redis.Remove(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) Validate() error {

	return nil
}

/*
func (u *UserServiceImpl) GetUser(name *string) (*entity.User, error) {
	//var user *entity.User
	//get user
	return nil, nil
}

*/
