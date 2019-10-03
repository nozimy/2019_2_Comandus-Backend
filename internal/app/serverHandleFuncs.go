package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const clientUrl = "https://comandus.now.sh"

func (s *server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	defer func() {
		// TODO: handle err
		r.Body.Close()
	}()
	decoder := json.NewDecoder(r.Body)
	newUserInput := new(model.UserInput)
	err := decoder.Decode(newUserInput)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	err = newUserInput.CheckEmail()
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	s.userType = newUserInput.UserType
	if s.userType != userFreelancer && s.userType != userCustomer {
		s.userType = userFreelancer
	}
	fmt.Println(s.userType)
	fmt.Println(newUserInput)

	s.usersdb.Mu.Lock()
	var id int
	var idf int
	var idc int
	var idCompany int

	if len(s.usersdb.Users) > 0 {
		id = s.usersdb.Users[len(s.usersdb.Users)-1].ID + 1
		idf = s.usersdb.Freelancers[len(s.usersdb.Freelancers)-1].ID + 1
		idc = s.usersdb.HireManagers[len(s.usersdb.HireManagers)-1].ID + 1
		idCompany = s.usersdb.Companies[len(s.usersdb.Companies)-1].ID + 1
	}

	user := model.User{
		ID:        id,
		FirstName: newUserInput.Name,
		SecondName: newUserInput.Surname,
		Email:     newUserInput.Email,
		Password:  newUserInput.Password,
	}

	err = user.BeforeCreate()
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newUserInput)
	}
	s.usersdb.Users = append(s.usersdb.Users, user)

	s.usersdb.Freelancers = append(s.usersdb.Freelancers, model.Freelancer{
		ID:        idf,
		AccountId: id,
	})

	s.usersdb.HireManagers = append(s.usersdb.HireManagers, model.HireManager{
		ID:        idc,
		AccountID: id,
	})

	s.usersdb.Companies = append(s.usersdb.Companies, model.Company{
		ID:          idCompany,
		CompanyName: "Company name",
	})

	fmt.Println(s.usersdb.Users[id])
	s.usersdb.Mu.Unlock()

	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	session.Values["user_id"] = user.ID
	session.Values["user_type"] = s.userType
	//session.Values["user_type"] = userFreelancer
	if err := s.sessionStore.Save(r, w, session); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	cookie := http.Cookie{Name: userTypeCookieName, Value: s.userType}
	cookie2 := http.Cookie{Name: hireManagerIdCookieName, Value: strconv.Itoa(idc)}
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)

	s.respond(w, r, http.StatusCreated, newUserInput)
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", clientUrl)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		var u *model.User
		var found bool

		s.usersdb.Mu.Lock()
		for i := 0; i < len(s.usersdb.Users); i++ {
			if id == s.usersdb.Users[i].ID {
				u = &s.usersdb.Users[i]
				found = true
			}
		}
		s.usersdb.Mu.Unlock()

		if !found {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) HandleSessionCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	defer func() {
		// TODO: handle err
		r.Body.Close()
	}()

	decoder := json.NewDecoder(r.Body)
	newUserInput := new(model.UserInput)
	err := decoder.Decode(newUserInput)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	log.Println(newUserInput)

	s.usersdb.Mu.Lock()
	for i := 0; i < len(s.usersdb.Users); i++ {
		if s.usersdb.Users[i].Email == newUserInput.Email &&
			s.usersdb.Users[i].ComparePassword(newUserInput.Password) {

			u := s.usersdb.Users[i]
			session, err := s.sessionStore.Get(r, sessionName)
			if err != nil {
				s.usersdb.Mu.Unlock()
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
			session.Values["user_id"] = u.ID
			session.Values["user_type"] = s.userType
			//session.Values["user_type"] = userFreelancer
			if err := s.sessionStore.Save(r, w, session); err != nil {
				s.usersdb.Mu.Unlock()
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
			s.usersdb.Mu.Unlock()
			s.respond(w, r, http.StatusOK, struct {
			}{})
			return
		}
	}
	s.usersdb.Mu.Unlock()
	s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
}

func (s *server) HandleLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, err)
		return
	}

	if session == nil {
		s.error(w, r, http.StatusNotFound, errors.New("failed to delete session"))
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		s.error(w, r, http.StatusExpectationFailed, err)
	}
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleSetUserType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//s.respond(w,r, http.StatusOK, nil)
	//return
	// TODO check if input user type invalid
	type Input struct {
		UserType string `json:"type"`
	}
	defer func() {
		// TODO: handle err
		r.Body.Close()
	}()

	decoder := json.NewDecoder(r.Body)
	newInput := new(Input)
	err := decoder.Decode(newInput)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	session, err := s.sessionStore.Get(r, sessionName)
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}

	if newInput.UserType != userCustomer && newInput.UserType != userFreelancer {
		s.error(w,r, http.StatusBadRequest, errors.New("user type may be only customer or freelancer"))
	}
	session.Values["user_type"] = newInput.UserType
	session.Save(r, w)
}

func (s *server) HandleShowProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, SendErr, CodeStatus := s.GetUserFromRequest(r)
	if SendErr != nil {
		s.error(w, r, CodeStatus, SendErr)
		return
	}
	s.respond(w, r, http.StatusOK, user)
}

func (s *server) HandleEditProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	user, SendErr, CodeStatus := s.GetUserFromRequest(r)
	if SendErr != nil {
		s.error(w, r, CodeStatus, SendErr)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	log.Println(user)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		SendErr := fmt.Errorf("invalid format of data")
		s.error(w, r, http.StatusBadRequest, SendErr)
		return
	}
	s.respond(w, r, http.StatusOK, struct{}{})

}

func (s *server) HandleEditPassword(w http.ResponseWriter, r *http.Request) {
	type BodyPassword struct {
		Password string
		NewPassword string
		NewPasswordConfirmation string
	}
	w.Header().Set("Content-Type", "application/json")
	var err error
	user, sendErr, codeStatus := s.GetUserFromRequest(r)
	if sendErr != nil {
		s.error(w, r, codeStatus, sendErr)
		return
	}
	bodyPassword := new(BodyPassword)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(bodyPassword)
	if bodyPassword.NewPassword != bodyPassword.NewPasswordConfirmation {
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("new Passwords are different"))
		return
	}

	if !user.ComparePassword(bodyPassword.Password) {
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("wrong password"))
		return
	}

	newEncryptPassword, err := model.EncryptString(bodyPassword.NewPasswordConfirmation)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error in updating password"))
		return
	}
	user.EncryptPassword = newEncryptPassword
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) GetUserFromRequest(r *http.Request) (*model.User, error, int) {
	session, err := s.sessionStore.Get(r, sessionName)
	if err == http.ErrNoCookie {
		SendErr := fmt.Errorf("user isn't authorized")
		return nil, SendErr, http.StatusUnauthorized
	}
	uidInteface := session.Values["user_id"]
	uid := uidInteface.(int)

	s.usersdb.Mu.Lock()
	user := s.usersdb.GetUserByID(uid)
	s.usersdb.Mu.Unlock()

	if user == nil {
		SendErr := fmt.Errorf("can't find user with id:" + strconv.Itoa(int(uid)))
		return nil, SendErr, http.StatusBadRequest
	}
	return user, nil, http.StatusOK
}

// TODO:
func (s *server) HandleEditNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, sendErr, codeStatus := s.GetUserFromRequest(r)
	if sendErr != nil {
		s.error(w, r, codeStatus, sendErr)
		return
	}
	userNotification := s.usersdb.GetNotificationsByUserID(user.ID)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(userNotification)
	log.Println(user)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		SendErr := fmt.Errorf("invalid format of data")
		s.error(w, r, http.StatusBadRequest, SendErr)
		return
	}
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleUploadAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("error retrieving the file"))
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, errors.New("error retrieving the file"))
		return
	}
	defer file.Close()

	/*tempFile, err := ioutil.TempFile("internal/store/avatars", "upload-*.png")
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, errors.New("error creating the file"))
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		s.error(w,r, http.StatusNotFound, err)
	}
	tempFile.Write(fileBytes)*/

	session, err := s.sessionStore.Get(r, sessionName)
	if err == http.ErrNoCookie {
		s.error(w, r, http.StatusUnauthorized, err)
		return
	}

	uidInterface := session.Values["user_id"]
	uid, ok := uidInterface.(int)
	if !ok {
		s.error(w,r, http.StatusInternalServerError, errors.New("cookie value not set"))
	}

	var image []byte
	file.Read(image)

	s.usersdb.Mu.Lock()
	user := s.usersdb.GetUserByID(uid)
	user.Avatar = true
	s.usersdb.ImageStore[uid] = image
	s.usersdb.Mu.Unlock()

	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleDownloadAvatar(w http.ResponseWriter, r *http.Request) {
	session, err := s.sessionStore.Get(r, sessionName)
	if err == http.ErrNoCookie {
		s.error(w,r,http.StatusUnauthorized, err)
		return
	}
	uidInterface := session.Values["user_id"]
	uid, ok := uidInterface.(int)
	if !ok {
		s.error(w,r, http.StatusInternalServerError, errors.New("cookie value not set"))
	}

	s.usersdb.Mu.Lock()
	user := s.usersdb.GetUserByID(uid)
	s.usersdb.Mu.Unlock()

	var Openfile *os.File
	if user.Avatar {
		s.usersdb.Mu.Lock()
		image := s.usersdb.ImageStore[uid]
		FileContentType := http.DetectContentType(image)
		w.Header().Set("Content-Type", FileContentType)
		w.Header().Set("Content-Length", strconv.Itoa(len(image)))
		if _, err := w.Write(image); err != nil {
			s.error(w,r,http.StatusInternalServerError, err)
		}
		s.usersdb.Mu.Unlock()
	} else {
		Filename := "internal/store/avatars/default.png"
		Openfile, err = os.Open(Filename)
		defer Openfile.Close()
		if err != nil {
			s.error(w, r, http.StatusNotFound, errors.New("cant open file"))
			return
		}
		FileHeader := make([]byte, 100000) // max image size!!!
		Openfile.Read(FileHeader)
		FileContentType := http.DetectContentType(FileHeader)

		FileStat, _ := Openfile.Stat()
		FileSize := strconv.FormatInt(FileStat.Size(), 10)

		w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
		w.Header().Set("Content-Type", FileContentType)
		w.Header().Set("Content-Length", FileSize)

		Openfile.Seek(0, 0)
		io.Copy(w, Openfile)
	}
	s.respond(w,r,http.StatusOK, struct{}{})
}

func (s *server) HandleRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	user, sendErr, codeStatus := s.GetUserFromRequest(r)
	if sendErr != nil {
		s.error(w, r, codeStatus, sendErr)
		return
	}
	HireManager := s.usersdb.GetHireManagerByID(user.ID)
	Company := s.usersdb.GetCompanyByID(HireManager.ID)
	var Roles []*model.Role
	clientRole := &model.Role{
		Role:   "client",
		Label:  Company.CompanyName,
		Avatar: "/default.png",
	}
	freelanceRole := &model.Role{
		Role:   "freelancer",
		Label:  user.FirstName + " " + user.SecondName,
		Avatar: "/default.png",
	}
	Roles = append(Roles, clientRole)
	Roles = append(Roles, freelanceRole)
	s.respond(w, r, http.StatusOK, Roles)
}

func (s *server) HandleGetAuthHistory(w http.ResponseWriter, r *http.Request) {
	// TODO: get auth history
}

func (s *server) HandleGetSecQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: get sec question
}

func (s *server) HandleEditSecQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: edit sec question
}

func (s *server) HandleCheckSecQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: check seq question
}

func (s *server) HandleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Methods", "POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Lol")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	s.respond(w, r, http.StatusOK, nil)
}

func (s *server) HandleCreateJob(w http.ResponseWriter, r *http.Request) {
	session, err := s.sessionStore.Get(r, sessionName)
	if err == http.ErrNoCookie {
		s.error(w, r, http.StatusNotFound, err)
		return
	}

	if session == nil {
		s.error(w, r, http.StatusNotFound, errors.New("session is nil"))
		return
	}

	uti := session.Values["user_type"]
	ut, ok := uti.(string)
	if !ok {
		s.error(w,r, http.StatusInternalServerError, errors.New("cookie value not set"))
	}

	log.Println(ut)

	// TODO: add test for this case
	//if ut != userCustomer {
	//	s.error(w, r, http.StatusBadRequest, errors.New("current user is not a manager"))
	//	return
	//}

	defer func() {
		// TODO: handle err
		r.Body.Close()
	}()
	decoder := json.NewDecoder(r.Body)
	newJob := new(model.Job)
	err = decoder.Decode(newJob)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	newJob.ID = len(s.usersdb.Jobs)

	uidi := session.Values["user_id"]
	uid, ok := uidi.(int)
	if !ok {
		s.error(w,r, http.StatusInternalServerError, errors.New("cookie value not set"))
	}

	// TODO write getbyID func for Hire Manager
	for i := 0; i < len(s.usersdb.HireManagers); i++ {
		if s.usersdb.HireManagers[i].AccountID == uid {
			newJob.HireManagerId = s.usersdb.HireManagers[i].ID
			s.usersdb.Jobs = append(s.usersdb.Jobs, *newJob)
			s.respond(w, r, http.StatusOK, newJob)
			return
		}
	}
	s.error(w, r, http.StatusInternalServerError, errors.New("client not found"))
}

func (s *server) HandleGetJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("wrong id"))
	}

	for i := 0; i < len(s.usersdb.Jobs); i++ {
		if id == s.usersdb.Jobs[i].ID {
			s.respond(w, r, http.StatusOK, &s.usersdb.Jobs[i])
			return
		}
	}
	s.error(w, r, http.StatusNotFound, errors.New("job not found"))
}


func (s *server) HandleEditFreelancer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	user, SendErr, CodeStatus := s.GetUserFromRequest(r)
	if SendErr != nil {
		s.error(w, r, CodeStatus, SendErr)
		return
	}
	Profile := s.usersdb.GetFreelancerByUserID(user.ID)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(Profile)
	fmt.Println(user)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		SendErr := fmt.Errorf("invalid format of data")
		s.error(w, r, http.StatusBadRequest, SendErr)
		return
	}
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleGetFreelancer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	vars := mux.Vars(r)
	FreelancerIDStr := vars["freelancerId"]
	FreelancerID, err := strconv.Atoi(FreelancerIDStr)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("id isn't number"))
	}
	Profile, err := s.usersdb.GetFreelancerByID(FreelancerID)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
	}
	s.respond(w, r, http.StatusOK, Profile)
}
func (s *server) HandleGetAvatar(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("wrong id"))
	}

	s.usersdb.Mu.Lock()
	user := s.usersdb.GetUserByID(id)
	if user == nil {
		s.error(w, r, http.StatusNotFound, errors.New("no such user in database"))
	}

	Filename := user.Avatar
	s.usersdb.Mu.Unlock()

	if Filename == "" {
		Filename = "/internal/store/avatars/default.png"
	}
	//fmt.Println("Client requests: " + Filename)

	Openfile, err := os.Open(Filename)
	defer Openfile.Close()
	if err != nil {
		s.error(w, r, http.StatusNotFound, errors.New("cant open file"))
		return
	}

	FileHeader := make([]byte, 100000) // max image size!!!
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)

	FileStat, _ := Openfile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)*/
}