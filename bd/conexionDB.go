package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://yoandredb:hV9HhwKHDjQcR3uD@cluster0.enacc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD es necesario*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printfln("Conexion exitosa con la BD")
	return client
}

/*CheckConect es la funcion*/
func CheckConect() int {
	err = MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
