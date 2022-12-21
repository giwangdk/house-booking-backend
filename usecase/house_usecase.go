package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type HouseUsecase interface {
	GetHouses(page int, limit int, sortBy string, sort string, searchBy string, filterByCity int) (*dto.HouseLists, error)
	CreateHouse(r dto.CreateHouseRequest) (*dto.CreateHouseResponse, error)
}

type HouseUsecaseImplementation struct {
	repository repository.HouseRepository
}

type HouseUsecaseImplementationConfig struct {
	Repository repository.HouseRepository
}

func NewHouseUseCase(c HouseUsecaseImplementationConfig) HouseUsecase {
	return &HouseUsecaseImplementation{
		repository: c.Repository,
	}
}

func (u *HouseUsecaseImplementation) GetHouses(page int, limit int, sortBy string, sort string, searchBy string, filterByCity int) (*dto.HouseLists, error) {
	houses, total, err := u.repository.GetHouses(page, limit, sortBy, sort, searchBy, filterByCity)
	
	if err != nil {
		return nil, err
	}

	resHouses:= *(&dto.HouseLists{}).BuildResponse(*houses, page, limit, total)

	return &resHouses, nil
}


func (u *HouseUsecaseImplementation) CreateHouse(r dto.CreateHouseRequest) (*dto.CreateHouseResponse, error) {

	entityHouse := entity.House{
		Name: r.Name,
		Price: r.Price,
		Description: r.Description,
		CityID: r.CityID,
		UserID: r.UserID,
		Location: r.Location,
	}

	house, err := u.repository.CreateHouse(entityHouse)
	if err != nil {
		return nil, err
	}

	res:= (&dto.CreateHouseResponse{}).BuildResponse(*house)

	return res, nil

}