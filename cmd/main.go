package main

import (
	"log"
	"os"

	"github.com/akshaybt001/DatingApp_Payment_Service/db"
	"github.com/akshaybt001/DatingApp_Payment_Service/initializer"
	"github.com/joho/godotenv"
)

func main() {
	if err:=godotenv.Load("../.env");err!=nil{
		log.Fatal("error loading env file")
	}
	addr :=os.Getenv("DB_KEY")
	DB,err:=db.InitDB(addr)
	if err!=nil{
		log.Fatal("error initialising database")
	}
	initializer.Initializer(DB).Start(":8090")
}