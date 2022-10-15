package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TODO struct {
	TodoId      string `bson:"todoId" json:"todoId"`
	Title       string `bson:"title" json:"title"`
	IsCompleted bool   `bson:"isCompleted" json:"isCompleted"`
	IsDeleted   bool   `bson:"isDeleted" json:"isDeleted"`
	CreatedAt   int64  `bson:"createdAt" json:"createdAt"`
	UpdateAt    int64  `bson:"updateAt" json:"updateAt"`
}

type CreateTodoDTO struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoDTO struct {
	IsCompleted bool `json:"isCompleted"`
}

func InitHandler(router *gin.Engine, client *mongo.Client) {
	todoCollection := client.Database(config.MongoDB).Collection("todo")

	router.POST("todo", func(ctx *gin.Context) {

		var createTodoDTO CreateTodoDTO

		if len(strings.TrimSpace(createTodoDTO.Title)) > 35 {
			ctx.JSON(http.StatusBadRequest, Error{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   "title length should be less than 35 charactor",
			})
			return
		}

		if err := ctx.BindJSON(&createTodoDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, Error{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			return
		}

		todo := TODO{
			TodoId:      ksuid.New().String(),
			Title:       createTodoDTO.Title,
			IsCompleted: false,
			IsDeleted:   false,
			CreatedAt:   time.Now().Unix(),
			UpdateAt:    time.Now().Unix(),
		}

		_, err := todoCollection.InsertOne(ctx.Request.Context(), todo)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Error{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, Resposne{
			Status:  http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
			Result:  todo,
		})
	})

	router.GET("todo", func(ctx *gin.Context) {
		filter := bson.M{
			"isDeleted": false,
		}

		sort := bson.M{
			"_id": -1,
		}

		cur, err := todoCollection.Find(ctx.Request.Context(), filter, options.Find().SetSort(sort))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Error{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			})
			return
		}

		var todos []TODO = []TODO{}

		err = cur.All(ctx.Request.Context(), &todos)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Error{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, Resposne{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Result:  todos,
		})
	})

	router.PATCH("todo/:todoId", func(ctx *gin.Context) {

		todoId := ctx.Param("todoId")

		var updateTodoDTO UpdateTodoDTO

		if err := ctx.BindJSON(&updateTodoDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, Error{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			return
		}

		filter := bson.M{
			"todoId": todoId,
		}

		update := bson.M{
			"$set": bson.M{
				"isCompleted": updateTodoDTO.IsCompleted,
				"updatedAt":   time.Now().Unix(),
			},
		}

		var todo TODO = TODO{}

		err := todoCollection.FindOneAndUpdate(ctx.Request.Context(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&todo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Error{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, Resposne{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Result:  todo,
		})

	})

	router.DELETE("todo/:todoId", func(ctx *gin.Context) {

		todoId := ctx.Param("todoId")

		filter := bson.M{
			"todoId": todoId,
		}

		update := bson.M{
			"$set": bson.M{
				"isDeleted": true,
				"updatedAt": time.Now().Unix(),
			},
		}

		todoCollection.FindOneAndUpdate(ctx.Request.Context(), filter, update)

		ctx.JSON(http.StatusOK, Resposne{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Result:  "deleted successfully",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Resposne{
			Status:  http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
			Result:  "no route defined",
		})
	})
}
