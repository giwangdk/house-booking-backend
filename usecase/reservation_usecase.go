package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
	"fmt"
	"time"
)

type ReservationUsecase interface {
	CreateReservation (r entity.Reservation) (*dto.CreateReservationResponse, error)
	CreateReservationWithUser(r dto.CreateReservationRequest) (*dto.CreateReservationResponse, error)
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


	reservation, err := u.repository.CreateReservation(r)
	if err != nil {
		return nil, err
	}


	res := (&dto.CreateReservationResponse{}).BuildResponse(*reservation)

	return res, nil
}

func (u *ReservationUsecaseImplementation) CreateReservationWithUser(r dto.CreateReservationRequest) (*dto.CreateReservationResponse, error) {


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
		
	}
	res, err:= u.CreateReservation(reservation)
		if err != nil {
			return nil, err
		}
	return res, nil

}
