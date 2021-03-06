package userUcase

import (
	"bytes"
	clients "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/clients/interfaces"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/app/user"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/pkg/errors"
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

type UserUsecase struct {
	userRep          user.Repository
	freelancerClient clients.ClientFreelancer
	managerClient    clients.ManagerClient
	companyClient    clients.CompanyClient
}

func NewUserUsecase(u user.Repository, fClient clients.ClientFreelancer, mClient clients.ManagerClient,
	cClient clients.CompanyClient) user.Usecase {
	return &UserUsecase{
		userRep:          u,
		freelancerClient: fClient,
		managerClient:    mClient,
		companyClient:    cClient,
	}
}

func (u *UserUsecase) CreateUser(data *model.User) *model.HttpError {
	if logErr := data.Validate(); logErr != nil {
		httpErr := &model.HttpError {
			HttpCode: 400,
			LogErr: errors.Wrap(logErr, "user.Validate()"),
			ClientErr: errors.New( "Неправильная форма данных, пароль должен быть не менее 6 символов.\n " +
			"Email должен быть в формате title@gmail.com"),
		}

		return httpErr
	}

	if err := data.BeforeCreate(); err != nil {
		httpErr := &model.HttpError {
			HttpCode: 500,
			LogErr: errors.Wrap(err, "user.Validate()"),
			ClientErr: errors.New( "Внутрення ошибка на сервере"),
		}
		return httpErr
	}

	if err := u.userRep.Create(data); err != nil {
		httpErr := &model.HttpError {
			HttpCode: 400,
			LogErr: errors.Wrap(err, "userRep.Create()"),
			ClientErr: errors.New("Пользователь с таким email уже существует"),
		}

		return httpErr
	}

	return nil
}

func (u *UserUsecase) EditUser(new *model.User, old *model.User) error {
	new.ID = old.ID

	if old.Email != new.Email {
		return errors.Wrap(errors.New("can't change email"), "EditUser")
	}

	new.UserType = old.UserType
	new.EncryptPassword = old.EncryptPassword

	if err := u.userRep.Edit(new); err != nil {
		return errors.Wrap(err, "userRep.Edit()")
	}
	return nil
}

func (u *UserUsecase) EditUserPassword(passwords *model.BodyPassword, user *model.User) error {
	if passwords.NewPassword != passwords.NewPasswordConfirmation {
		return errors.New("Введеные новые пароли не совпадают")
	}

	if !user.ComparePassword(passwords.Password) {
		err := errors.New("Неверный пароль от аккаунта")
		return errors.Wrapf(err, "model.user.ComparePassword")
	}

	newEncryptPassword, err := model.EncryptString(passwords.NewPassword)
	if err != nil {
		return errors.Wrap(err, "model.EncryptString")
	}
	user.EncryptPassword = newEncryptPassword

	if err := u.userRep.Edit(user); err != nil {
		return errors.Wrapf(err, "userRep.Edit")
	}
	return nil
}

func (u *UserUsecase) GetAvatar(user *model.User) ([]byte, error) {
	if user.Avatar != nil && len(user.Avatar) > 3 {
		return user.Avatar, nil
	}

	response, _ := http.Get("https://sun9-40.userapi.com/c858036/v858036437/12d222/TpWDmhunUUg.jpg")
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	im, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "image.Decode()")
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, im, nil); err != nil {
		return nil, errors.Wrap(err, "image.Encode()")
	}

	return buf.Bytes(), nil
}

func (u *UserUsecase) Find(id int64) (*model.User, error) {
	user, err := u.userRep.Find(id)
	if err != nil {
		return nil, errors.Wrap(err, "userRep.Find()")
	}

	currFreelancer, err := u.freelancerClient.GetFreelancerByUserFromServer(id)
	if err != nil {
		return nil, errors.Wrap(err, "clients.GetFreelancerByUserFromServer()")
	}

	currManager, err := u.managerClient.GetManagerByUserFromServer(id)
	if err != nil {
		return nil, errors.Wrap(err, "clients.GetManagerByUserFromServer()")
	}

	user.FreelancerId = currFreelancer.ID
	user.HireManagerId = currManager.ID
	user.CompanyId = currManager.CompanyId

	return user, nil
}

func (u *UserUsecase) SetUserType(user *model.User, userType string) error {
	if err := user.SetUserType(userType); err != nil {
		return errors.Wrap(err, "SetUserType()")
	}

	if err := u.userRep.Edit(user); err != nil {
		return errors.Wrap(err, "userRep.Edit()")
	}
	return nil
}

func (u *UserUsecase) VerifyUser(currUser *model.User) (int64, *model.HttpError) {
	us, err := u.userRep.FindByEmail(currUser.Email)
	if err != nil {
		httpErr := &model.HttpError {
			HttpCode: 400,
			LogErr: errors.Wrapf(err, "userRep.FindByEmail()"),
			ClientErr: errors.New("Аккаунта с данным email не существует"),
		}
		return 0, httpErr
	}

	if !us.ComparePassword(currUser.Password) {
		httpErr := &model.HttpError {
			HttpCode: 400,
			LogErr: errors.New("us.ComparePassword"),
			ClientErr: errors.New("Неверный пароль"),
		}
		return 0, httpErr
	}

	return us.ID, nil
}

func (u *UserUsecase) GetRoles(user *model.User) ([]*model.Role, error) {
	currManager, err := u.managerClient.GetManagerByUserFromServer(user.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "clients.GetManagerByUserFromServer()")
	}

	currCompany, err := u.companyClient.GetCompanyFromServer(currManager.CompanyId)
	if err != nil {
		return nil, errors.Wrap(err, "getCompanyFromServer()")
	}

	var roles []*model.Role

	// TODO: rewrite avatar in Role struct
	clientRole := &model.Role{
		Role:  "client",
		Label: currCompany.CompanyName,
	}

	freelanceRole := &model.Role{
		Role:  "freelancer",
		Label: user.FirstName + " " + user.SecondName,
	}

	roles = append(roles, clientRole)
	roles = append(roles, freelanceRole)

	return roles, nil
}

func (u *UserUsecase) GetNames() ([]string, error) {
	names, err := u.userRep.GetNames()
	if err != nil {
		return nil, errors.Wrap(err, "userRep.GetNames()")
	}
	return names, nil
}