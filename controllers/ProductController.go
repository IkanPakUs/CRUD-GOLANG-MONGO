package controllers

import (
	"belajar-golang/cmd/ProductService/databases"
	"belajar-golang/cmd/ProductService/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection_use = "products"

func GetProducts(ctx *gin.Context) {
	db := databases.Database.Client.Database(databases.Database.DatabaseName).Collection(collection_use)

	var products []models.Product
	filter := bson.D{}

	cursor, err := db.Find(context.TODO(), filter)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &products); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	ctx.JSON(http.StatusOK, products)
}

func GetProductById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	db := databases.Database.Client.Database(databases.Database.DatabaseName).Collection(collection_use)

	var product models.Product
	filter := bson.D{{Key: "_id", Value: id}}

	err = db.FindOne(context.TODO(), filter).Decode(&product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	ctx.JSON(http.StatusOK, product)
}

func CreateProduct(ctx *gin.Context) {
	var new_product models.InputProduct

	if err := ctx.BindJSON(&new_product); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	db := databases.Database.Client.Database(databases.Database.DatabaseName).Collection(collection_use)

	product, err := db.InsertOne(context.TODO(), new_product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	ctx.JSON(http.StatusOK, product.InsertedID)
}

func UpdateProduct(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	var update_product models.InputProduct

	if err := ctx.BindJSON(&update_product); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: update_product.Name},
			{Key: "stock", Value: update_product.Stock},
			{Key: "price", Value: update_product.Price},
		}},
	}

	db := databases.Database.Client.Database(databases.Database.DatabaseName).Collection(collection_use)

	product, err := db.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	ctx.JSON(http.StatusOK, product.UpsertedID)
}

func DeleteProduct(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	db := databases.Database.Client.Database(databases.Database.DatabaseName).Collection(collection_use)

	filter := bson.D{{Key: "_id", Value: id}}

	product, err := db.DeleteOne(context.TODO(), filter)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		panic(err)
	}

	ctx.JSON(http.StatusOK, product.DeletedCount)
}
