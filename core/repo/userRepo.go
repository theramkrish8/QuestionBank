package repo

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/common"
	"QuestionBank/core/mgorepo"
	"QuestionBank/core/models"
)

//UserRepo type handles user collection operations
type UserRepo struct {
}

var userInstance *UserRepo
var userOnce sync.Once

//GetUserRepo returns singleton instance of UserRepo class
func GetUserRepo() *UserRepo {
	userOnce.Do(func() {
		userInstance = &UserRepo{}
	})
	return userInstance
}

//GetAllUsers returns all users list from UserCollection
func (uc UserRepo) GetAllUsers() ([]models.User, error) {
	var results []models.User
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.UserCollection).Find(nil).All(&results)
	return results, err
}

//GetUserByID returns particular User by id
func (uc UserRepo) GetUserByID(oid bson.ObjectId) ([]common.UserResult, error) {
	col1 := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.UserCollection)
	pipe := col1.Pipe([]bson.M{
		bson.M{
			"$lookup": bson.M{
				"from":         "schools",
				"localField":   "school",
				"foreignField": "_id",
				"as":           "school",
			},
		},
		bson.M{
			"$match": bson.M{
				"_id": oid,
			},
		},
	})
	iter := pipe.Iter()
	defer iter.Close()
	res := common.UserResult{}
	var results []common.UserResult
	i := 0
	for iter.Next(&res) {
		results[i] = res
		i++
	}
	return results, iter.Err()
}

//InsertUser insert new user to user collection
func (uc UserRepo) InsertUser(user models.User) (models.User, error) {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.UserCollection).Insert(&user)
	return user, err
}

//DeleteUser deletes user from usercollection
func (uc UserRepo) DeleteUser(oid bson.ObjectId) error {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.UserCollection).RemoveId(oid)
	return err
}

//UpdateUserByID updates user data identified by user
func (uc UserRepo) UpdateUserByID(user models.User) (models.User, error) {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.UserCollection).UpdateId(&user.UserID, &user)
	return user, err
}
