package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/httperror"
	"final-project-backend/repository"
	"time"
)

type HouseUsecase interface {
	GetHouses(page int, limit int, sortBy string, sort string, searchBy string, filterByCity int,checkin string, checkout string) (*dto.HouseLists, error)
	CreateHouse(r dto.CreateHouseRequest) (*dto.CreateHouseResponse, error)
	GetHouseById(houseId int) (*dto.House, error)
	UpdateHouse(r dto.UpdateHouseRequest, houseId int) (*dto.UpdateHouseResponse, error)
	GetHousesHost(userId int, page int, limit int, sortBy string, sort string, searchBy string) (*dto.HouseLists, error)
	DeleteHouse(houseId int) error
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

func (u *HouseUsecaseImplementation) GetHouses(page int, limit int, sortBy string, sort string, searchBy string, filterByCity int, checkin string, checkout string) (*dto.HouseLists, error) {
	houses, total, err := u.repository.GetHouses(0,page, limit, sortBy, sort, searchBy, filterByCity, checkin, checkout)

	if err != nil {
		return nil, err
	}

	resHouses := *(&dto.HouseLists{}).BuildResponse(*houses, page, limit, total)

	return &resHouses, nil
}

func (u *HouseUsecaseImplementation) CreateHouse(r dto.CreateHouseRequest) (*dto.CreateHouseResponse, error) {

	entityHouse := entity.HouseProfile{
		Name:        r.Name,
		Price:       r.Price,
		Description: r.Description,
		CityID:      r.CityID,
		UserID:      r.UserID,
		Location:    r.Location,
	}

	house, err := u.repository.CreateHouse(entityHouse)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateHouseResponse{}).BuildResponse(*house)

	return res, nil

}

func (u *HouseUsecaseImplementation) GetHouseById(houseId int) (*dto.House, error) {
	house, err := u.repository.GetHouseById(houseId)
	if err != nil {
		return nil, err
	}

	res := (&dto.House{}).BuildResponse(*house)

	return res, nil
}

func (u *HouseUsecaseImplementation) UpdateHouse(r dto.UpdateHouseRequest, houseId int) (*dto.UpdateHouseResponse, error) {
	_, err := u.GetHouseById(houseId)
	if err != nil {
		return nil, err
	}

	entity := entity.HouseProfile{
		Name:        r.Name,
		Price:       r.Price,
		Description: r.Description,
		CityID:      r.CityID,
		Location:    r.Location,
	}

	updatedUser, err := u.repository.UpdateHouse(entity, houseId)
	if err != nil {
		return nil, err
	}

	res := (&dto.UpdateHouseResponse{}).BuildResponse(*updatedUser)
	return res, nil
}

func (u *HouseUsecaseImplementation) GetHousesHost(userId int, page int, limit int, sortBy string, sort string, searchBy string) (*dto.HouseLists, error) {
	
	houses, total, err := u.repository.GetHouses(userId, page, limit, sortBy, sort, searchBy, 0, "", "")

	if err != nil {
		return nil, err
	}

	resHouses := *(&dto.HouseLists{}).BuildResponse(*houses, page, limit, total)

	return &resHouses, nil
}

func (u *HouseUsecaseImplementation) DeleteHouse(houseId int) error {
	_, err := u.GetHouseById(houseId)
	if err != nil {
		return err
	}

	err= u.repository.IsBooked(houseId, time.Now())
	if err != nil  {
		return httperror.BadRequestError("There is reservation ongoing!", "FAILED_DELETE_HOUSE")
	}

	err = u.repository.DeleteHouse(houseId)
	if err != nil {
		return err
	}

	return nil
}