package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"patients.com/Asads-APIs/controllers"
	"patients.com/Asads-APIs/services"
)

var (
	server            *gin.Engine
	patientservice    services.PatientService
	patientcontroller controllers.PatientController
	ctx               context.Context
	patientcollection *mongo.Collection
	mongoclient       *mongo.Client
	err               error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongoconn has established")

	patientcollection = (*mongo.Collection)(mongoclient.Database("patientdb").Collection("patients"))
	patientservice = services.NewPatientService(patientcollection, ctx)
	patientcontroller = controllers.New(patientservice)
	server = gin.Default()
}
func main() {

	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	patientcontroller.RegisterPatientRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
