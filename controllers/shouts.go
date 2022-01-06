package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mwelwankuta/shout/database"
	"github.com/mwelwankuta/shout/models"
)

type PostSuccess struct {
	Message string
	Post    models.Post
}

type LikeType struct {
	Id     uint `json:"id"`
	UserId uint `json:"userid"`
}

// get all posts
func Home(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post

	database.DB.Find(&models.Post{}).Order("id RAND()")

	json.NewEncoder(w).Encode(posts)
}

// creat a post
func Post(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

// like post
func Like(w http.ResponseWriter, r *http.Request) {
	var data LikeType

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

// create comment
func Comment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("")
		return
	}

	var postComment models.Comment
	database.DB.First(&postComment, "post_id=?", comment.PostId)

	var postComments []models.Comment
	database.DB.Find(&postComments, "post_id=?", comment.PostId)

	// combine request comment to database comments
	postComments = append(postComments, comment)

	database.DB.Model(&postComment).Select("comments").Updates(postComments)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "created created",
	})
}

// delete post
func Delete(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&post)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There was a problem deleting the post"))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted post",
	})
}
