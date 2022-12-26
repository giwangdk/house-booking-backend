package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/httperror"
	"final-project-backend/repository"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ReservationUsecase interface {
	CreateReservation(r entity.Reservation) (*dto.CreateReservationResponse, error)
	CreateReservationWithUser(r dto.CreateReservationRequest) (*dto.CreateReservationResponse, error)
	GetReservationByBookingCode(code string) (*dto.ReservationDetail, error)
	UpdateStatusReservation(id int, statusID int) (*entity.Reservation, error)
	GetReservationById(id int) (*entity.Reservation, error) 
}

type ReservationUsecaseImplementation struct {
	repository repository.ReservationRepository
	userUsecase UserUsecase
}

type ReservationUsecaseImplementationConfig struct {
	Repository repository.ReservationRepository
	UserUsecase UserUsecase
}

func NewReservationUseCase(c ReservationUsecaseImplementationConfig) ReservationUsecase {
	return &ReservationUsecaseImplementation{
		repository: c.Repository,
		userUsecase: c.UserUsecase,
	}
}


func (u *ReservationUsecaseImplementation) CreateReservation (r entity.Reservation) (*dto.CreateReservationResponse, error) {

	isAvailable, err := u.repository.IsHouseAvailable(r.CheckIn, r.CheckOut, r.HouseID)
	fmt.Println(isAvailable, err)
	if !isAvailable && err == nil {
		return nil, httperror.BadRequestError("House is not available", "ERROR_HOUSE_NOT_AVAILABLE")
	}
	if err != nil {
		return nil, err
	}

	reservation, err := u.repository.CreateReservation(r)
	if err != nil {
		return nil, err
	}


	res := (&dto.CreateReservationResponse{}).BuildResponse(*reservation)

	return res, nil
}

func (u *ReservationUsecaseImplementation) CreateReservationWithUser(r dto.CreateReservationRequest) (*dto.CreateReservationResponse, error) {

    code:= uuid.New()

	
	user, isExist := u.userUsecase.IsUserExistByEmail(r.Email)
	
	fmt.Println(user, isExist)

	if !isExist {
		userCreated, err := u.userUsecase.CreateUser(entity.User{
			Fullname: r.Fullname,
			Email: r.Email,
			Role: "guest",
			CityID: r.CityID,
		})
		if err != nil {
			return nil, err
		}

		reservationUnregisteredAcc := entity.Reservation{
			CheckIn: r.CheckIn,
			CheckOut: r.CheckOut,
			TotalPrice: r.TotalPrice,
			HouseID: r.HouseID,
			Expired : time.Now().Add(1 * time.Hour),
			StatusID: 1,
			UserID: int(userCreated.ID),
			BookingCode: code.String(),
		}

		res, err:= u.CreateReservation(reservationUnregisteredAcc)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	reservation := entity.Reservation{
		CheckIn: r.CheckIn,
		CheckOut: r.CheckOut,
		TotalPrice: r.TotalPrice,
		HouseID: r.HouseID,
		Expired : time.Now().Add(1 * time.Hour),
		StatusID: 1,
		UserID: int(user.ID),
		BookingCode: code.String(),
		
	}
	res, err:= u.CreateReservation(reservation)
		if err != nil {
			return nil, err
		}
	return res, nil

}

func (u *ReservationUsecaseImplementation) GetReservationByBookingCode(code string) (*dto.ReservationDetail, error) {
	reservation, err := u.repository.GetReservationByBookingCode(code)
	if err != nil {
		return nil, err
	}

	res:= (&dto.ReservationDetail{}).BuildResponse(*reservation)
	return res, nil
}

func (u *ReservationUsecaseImplementation) GetReservationById(id int) (*entity.Reservation, error) {
	reservation, err := u.repository.GetReservationById(id)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}
func (u *ReservationUsecaseImplementation) UpdateStatusReservation(id int, statusID int) (*entity.Reservation, error) {
	reservation, err := u.repository.UpdateStatus(id, statusID)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}