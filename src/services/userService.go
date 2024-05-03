package services

import (
	"encoding/json"
	"net/http"

	"github.com/kylerequez/go-crud-api/src/models"
	"github.com/kylerequez/go-crud-api/src/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	ur *repositories.UserRepository
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	results, err := us.ur.GetAllUsers()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]any{
		"users": results,
	})
}

func (us *UserService) GetUserById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var id string = req.PathValue("id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{
			"error": "Id must not empty",
		})
		return
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	result, err := us.ur.GetUserById(oid)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]any{
		"user": &result,
	})
}

func (us *UserService) InsertUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)
	var newUser models.User
	if err := decoder.Decode(&newUser); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	_, err := us.ur.InsertUser(newUser)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func (us *UserService) PatchUpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var id string = req.PathValue("id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{
			"error": "Id must not empty",
		})
		return
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	decoder := json.NewDecoder(req.Body)
	var updatedUser models.User
	if err := decoder.Decode(&updatedUser); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	_, err = us.ur.GetUserById(oid)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	_, err = us.ur.UpdateUser(oid, updatedUser)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (us *UserService) PutUpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var id string = req.PathValue("id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{
			"error": "Id must not empty",
		})
		return
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	decoder := json.NewDecoder(req.Body)
	var updatedUser models.User
	if err := decoder.Decode(&updatedUser); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	isExists, _ := us.ur.GetUserById(oid)
	// if err != nil {
	// 	res.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(res).Encode(map[string]string{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	if isExists != nil {
		_, err := us.ur.UpdateUser(oid, updatedUser)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		res.WriteHeader(http.StatusOK)
	} else {
		_, err := us.ur.InsertUser(updatedUser)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		res.WriteHeader(http.StatusCreated)
	}

}

func (us *UserService) DeleteUserById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var id string = req.PathValue("id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{
			"error": "Id must not empty",
		})
		return
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	result, err := us.ur.DeleteUserById(oid)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	if result.DeletedCount <= 0 {
		res.WriteHeader(http.StatusNotFound)
	} else {
		res.WriteHeader(http.StatusOK)
	}
}
