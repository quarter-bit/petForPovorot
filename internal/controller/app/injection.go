package app

import (
	"Pet_store/internal/controller/db"
	"Pet_store/internal/controller/handlers"
	"Pet_store/internal/repo"
	"Pet_store/internal/service"
)

type App struct {
	UH *handlers.UserHandlers
	PH *handlers.PetHandlers
	SH *handlers.StoreHandler
}

func Inject() *App {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	//Иньекция юзера
	repoUser := repo.NewUserRepo(db.Db)
	serviceUser := service.NewUserService(repoUser)
	HAndlerUser := handlers.NewUserHandlers(serviceUser)

	//Иньекция питомца
	repoPet := repo.NewPetRepo(db.Db)
	servicePet := service.NewPetService(repoPet)
	HAndlerPet := handlers.NewPetHandlers(servicePet)

	//Иньекция магазина
	repoStore := repo.NewStoreRepo(db.Db)
	serviceStore := service.NewStoreService(repoStore)
	HAndlerStore := handlers.NewStoreHandler(serviceStore)

	return &App{
		UH: HAndlerUser,
		PH: HAndlerPet,
		SH: HAndlerStore,
	}
}
