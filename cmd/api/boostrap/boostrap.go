package boostrap

import (
	"database/sql"
	"fmt"
	"jagch/backend/internal/auth/usecase"
	"jagch/backend/internal/platform/config"
	"jagch/backend/internal/platform/server"
	psqlauth "jagch/backend/internal/platform/storage/psql/auth"
	psqlcliente "jagch/backend/internal/platform/storage/psql/cliente"
	psqlpemaritimo "jagch/backend/internal/platform/storage/psql/planentregamaritimo"
	psqlpeterrestre "jagch/backend/internal/platform/storage/psql/planentregaterrestre"
	"log"

	_ "github.com/lib/pq"
)

/* const (
	host = "localhost"
	port = 8181

	dbUser = "postgres"
	dbPass = "123456"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "docucenter"
) */

func Run() error {

	psqlURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_USER"), config.Config("DB_PASS"), config.Config("DB_NAME"))
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Established a successful connection!")

	authStorage := psqlauth.NewAuthStorage(db)
	authUsecase := usecase.NewUsecase(authStorage)

	clienteStorage := psqlcliente.NewClienteStorage(db)

	peTerrestreStorage := psqlpeterrestre.NewPETerrestreStorage(db)

	peMaritimoStorage := psqlpemaritimo.NewPEMaritimoStorage(db)

	srv := server.New(config.Config("HOST"), config.Config("PORT"), authUsecase, clienteStorage, peTerrestreStorage, peMaritimoStorage)

	return srv.Run()
}
