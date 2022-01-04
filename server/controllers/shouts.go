package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mwelwankuta/shout/server/database"
	"github.com/mwelwankuta/shout/server/models"
)

type PostSuccess struct {
	Message string
	Post    models.Post
}

type LikeType struct {
	Id     uint `json:"id"`
	UserId uint `json:"userid"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	var posts models.Post

	database.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func Post(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("")
		return
	}

	post := models.Post{
		Text:     data["text"],
		Author:   data["author"],
		Comments: []models.Comment{},
		Likes:    []uint{},
	}

	database.DB.Create(&post)
	json.NewEncoder(w).Encode(&PostSuccess{
		Message: "created post",
		Post:    post,
	})
}

func Like(w http.ResponseWriter, r *http.Request) {
	var data LikeType

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("")
		return
	}

	var post models.Post

	var like = &LikeType{
		Id:     post.Id,
		UserId: data.UserId,
	}

	database.DB.Where("id =?", data.Id).First(&post)

	database.DB.Where("id = ?", data.Id).Update("likes", append(post.Likes, like.UserId))

}

func Comment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("")
		return
	}

	var comments []models.Comment

	database.DB.Where("id = ?", comment.PostId).First(&comments)

	database.DB.Where("id = ?", comment.PostId).Update("comments", append(comments, comment))

	json.NewEncoder(w).Encode(map[string]string{
		"message": "post created",
	})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var post map[string]string

	json.NewDecoder(r.Body).Decode(&post)

	database.DB.Where("id = ?", post["id"]).Delete(&post)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted post",
	})
}
