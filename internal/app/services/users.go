package services

import (
	"net/http"

	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/pkg/logger"
)

type IUserService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type UserService struct {
	UserUsecase usecases.IUserUsecase
	log         logger.Ilogger
}

func NewUserService(u usecases.IUserUsecase, l logger.Ilogger) IUserService {
	l.Info("permission service is created")
	return &UserService{
		log:         l,
		UserUsecase: u,
	}
}
func (u *UserService) Delete(w http.ResponseWriter, r *http.Request) {}
func (u *UserService) GetAll(w http.ResponseWriter, r *http.Request) {

	// //get users

	// users, err := u.UserUsecase.GetAll()
	// if err != nil {
	// 	http.Error(w, "faild to ger users", http.StatusBadRequest)
	// 	return
	// }

	// //convert to json
	// jsonResponse, errMarshal := json.Marshal(users)
	// if errMarshal != nil {
	// 	http.Error(w, "faild to parse json to serve", http.StatusBadRequest)
	// 	return
	// }

	// //set headers
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// //send response
	// w.Write(jsonResponse)
}

func (u *UserService) Create(w http.ResponseWriter, r *http.Request) {

	// //get data from the request body and convert to json
	// var requestUser dto.CreateUserRequest
	// err := json.NewDecoder(r.Body).Decode(&requestUser)
	// if err != nil {
	// 	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 	return
	// }

	// //validate the inputs
	// uv := validation.NewUserValidator()
	// errors := uv.ValidateCreateUser(requestUser)
	// if errors != nil {
	// 	http.Error(w, errors.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // var roles []models.Role
	// // if len(requestUser.RoleIDs) > 0 {
	// // 	// Fetch roles from the database using role IDs
	// // 	rolesOFDB, err := u.UserRepo.FindByRoleIdes(requestUser.RoleIDs)
	// // 	if err != nil {
	// // 		http.Error(w, "there is no role with this id", http.StatusBadRequest)
	// // 		return
	// // 	}

	// // 	// Check if all role IDs are valid
	// // 	if len(*rolesOFDB) != len(requestUser.RoleIDs) {
	// // 		http.Error(w, "invalid role ids provided", http.StatusBadRequest)
	// // 		return
	// // 	}

	// // 	roles = *rolesOFDB
	// // }

	// //hash the password
	// hashedPass, err := utils.HashPassword(requestUser.Password)
	// if err != nil {
	// 	http.Error(w, "encrypting password failed", http.StatusBadRequest)
	// 	return
	// }

	// //map the inputs to the user
	// var newUser = models.User{
	// 	Name:        requestUser.Name,
	// 	LastName:    requestUser.LastName,
	// 	PhoneNumber: requestUser.PhoneNumber,
	// 	Email:       requestUser.Email,
	// 	Password:    hashedPass,
	// }
	// // if len(roles) > 0 {
	// // 	newUser.Roles = roles
	// // }

	// //create the user in db
	// user, er := u.UserRepo.Create(&newUser)
	// if er != nil {
	// 	http.Error(w, er.Error(), http.StatusBadRequest)
	// }

	// //convert to json
	// jsonResponse, errMarshal := json.Marshal(user)
	// if errMarshal != nil {
	// 	fmt.Println("fail to marshal user")
	// }

	// //set headers response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// //send response
	// w.Write(jsonResponse)
}

func (u *UserService) GetById(w http.ResponseWriter, r *http.Request) {

	// //get the id from uri
	// userId := mux.Vars(r)["id"]
	// uId, err := strconv.Atoi(userId)
	// if err != nil {
	// 	http.Error(w, "invalid user id", http.StatusBadRequest)
	// 	return
	// }

	// //get the user from db

	// user, errGetUser := u.UserRepo.GetById(uId)
	// if errGetUser != nil {
	// 	http.Error(w, "can't get the user", http.StatusBadRequest)
	// 	return
	// }

	// //convert the user to json
	// jsonResponse, errMarshal := json.Marshal(&user)
	// if errMarshal != nil {
	// 	http.Error(w, "failed to parse json to serve", http.StatusBadRequest)
	// 	return
	// }

	// //set response header
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// //send the response
	// w.Write(jsonResponse)
}

func (u *UserService) Update(w http.ResponseWriter, r *http.Request) {

	// 	//get the id from uri
	// 	userId := mux.Vars(r)["id"]
	// 	uId, err := strconv.Atoi(userId)
	// 	if err != nil {
	// 		http.Error(w, "invalid user id", http.StatusBadRequest)
	// 		return
	// 	}

	// 	//does user exist
	// 	_, errFInd := u.UserRepo.GetById(uId)
	// 	if errFInd != nil {
	// 		if err == gorm.ErrRecordNotFound {
	// 			http.Error(w, "there is no such user", http.StatusBadRequest)
	// 			return
	// 		}
	// 	}

	// 	//get data from the request body and convert to json
	// 	var requestUser dto.UserUpdateRequest
	// 	err = json.NewDecoder(r.Body).Decode(&requestUser)
	// 	if err != nil {
	// 		//fmt.Println("fail to decode the body")
	// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 		return
	// 	}

	// 	//validate the inputs
	// 	uv := validation.NewUserValidator()
	// 	errors := uv.ValidateUpdateUser(requestUser)
	// 	if errors != nil {
	// 		http.Error(w, errors.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	//hash the password
	// 	var hashedPass = requestUser.Password
	// 	if requestUser.Password != "" {
	// 		hashedPass, err = utils.HashPassword(requestUser.Password)
	// 		if err != nil {
	// 			http.Error(w, "encrypting password faild", http.StatusBadRequest)
	// 			return
	// 		}
	// 	}

	// 	//map the inputs to the user
	// 	var updatedUser = models.User{
	// 		Name:        requestUser.Name,
	// 		LastName:    requestUser.LastName,
	// 		PhoneNumber: requestUser.PhoneNumber,
	// 		Email:       requestUser.Email,
	// 		Password:    hashedPass,
	// 	}

	// 	// set the new sata
	// 	_, errUpdate := u.UserRepo.Update(&updatedUser)
	// 	if errUpdate != nil {
	// 		http.Error(w, "can't update the user", http.StatusBadRequest)
	// 		return
	// 	}

	// 	//set response header
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)

	// 	//send the response
	// 	json.NewEncoder(w).Encode(updatedUser)
	// }

	// func (u *UserService) Delete(w http.ResponseWriter, r *http.Request) {

	// 	//get the id from uri
	// 	userId := mux.Vars(r)["id"]
	// 	uId, err := strconv.Atoi(userId)
	// 	if err != nil {
	// 		http.Error(w, "invalid user id", http.StatusBadRequest)
	// 		return
	// 	}

	// 	//does user exist
	// 	user, err := u.UserRepo.GetById(uId)
	// 	if err != nil {
	// 		if err == gorm.ErrRecordNotFound {
	// 			http.Error(w, "there is no such user", http.StatusBadRequest)
	// 			return
	// 		}
	// 	}
	// 	//delete the user
	// 	err = u.UserRepo.Delete(int(user.ID))
	// 	if err != nil {
	// 		http.Error(w, "cant delete the user", http.StatusBadRequest)
	// 		return
	// 	}

	// 	//set response header
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)

	// 	//send the response
	// 	json.NewEncoder(w).Encode(models.User{})

}
