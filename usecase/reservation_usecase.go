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
	CreateReservation(r entity.Reservation) (*entity.Reservation, error)
	CreateReservationWithUser(r dto.CreateReservationRequest) (*dto.CreateReservationResponse, error)
	GetReservationByBookingCode(code string) (*dto.ReservationDetail, error)
	UpdateStatusReservation(id int, statusID int) (*entity.Reservation, error)
	GetReservationById(id int) (*entity.Reservation, error)
	GetReservationByUserId(userId int) (*dto.ReservationList, error)
}


type ReservationUsecaseImplementation struct {
	repository    repository.ReservationRepository
	userUsecase   UserUsecase
	pickupUsecase PickupUsecase
}

type ReservationUsecaseImplementationConfig struct {
	Repository    repository.ReservationRepository
	UserUsecase   UserUsecase
	PickupUsecase PickupUsecase
}

func NewReservationUseCase(c ReservationUsecaseImplementationConfig) ReservationUsecase {
	return &ReservationUsecaseImplementation{
		repository:    c.Repository,
		userUsecase:   c.UserUsecase,
		pickupUsecase: c.PickupUsecase,
	}
}

func (u *ReservationUsecaseImplementation) CreateReservation(r entity.Reservation) (*entity.Reservation, error) {

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

	return reservation, nil
}

func (u *ReservationUsecaseImplementation) CreateReservationWithUser(r dto.CreateReservationRequest) (*dto.CreateReservationResponse, error) {

	code := uuid.New()

	user, isExist := u.userUsecase.IsUserExistByEmail(r.Email)

	fmt.Println(user, isExist)

	if !isExist {
		userCreated, err := u.userUsecase.CreateUser(entity.User{
			Fullname: r.Fullname,
			Email:    r.Email,
			Role:     "guest",
			CityID:   r.CityID,
		})
		if err != nil {
			return nil, err
		}

		reservationUnregisteredAcc := entity.Reservation{
			CheckIn:     r.CheckIn,
			CheckOut:    r.CheckOut,
			TotalPrice:  r.TotalPrice,
			HouseID:     r.HouseID,
			Expired:     time.Now().Add(1 * time.Hour),
			StatusID:    1,
			UserID:      int(userCreated.ID),
			BookingCode: code.String(),
		}

		reservation, err := u.CreateReservation(reservationUnregisteredAcc)
		if err != nil {
			return nil, err
		}

		res := (&dto.CreateReservationResponse{}).BuildResponse(*reservation)

		fmt.Println("hihi")

		if !r.IsRequestPickup {
			return res, nil
		}

		_, err = u.pickupUsecase.CreatePickup(dto.CreatePickupRequest{
			ReservationID: int(reservation.ID),
			UserID:        int(userCreated.ID),
		})
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	reservation := entity.Reservation{
		CheckIn:     r.CheckIn,
		CheckOut:    r.CheckOut,
		TotalPrice:  r.TotalPrice,
		HouseID:     r.HouseID,
		Expired:     time.Now().Add(1 * time.Hour),
		StatusID:    1,
		UserID:      int(user.ID),
		BookingCode: code.String(),
	}
	reservationCreated, err := u.CreateReservation(reservation)
	if err != nil {
		return nil, err
	}
	res := (&dto.CreateReservationResponse{}).BuildResponse(*reservationCreated)

	if !r.IsRequestPickup {
		return res, nil
	}

	_, err = u.pickupUsecase.CreatePickup(dto.CreatePickupRequest{
		ReservationID: int(reservationCreated.ID),
		UserID:        int(user.ID),
	})
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

	res := (&dto.ReservationDetail{}).BuildResponse(*reservation)
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

func (u *ReservationUsecaseImplementation) GetReservationByUserId(userId int) (*dto.ReservationList, error) {
	reservations, err := u.repository.GetReservationByUserID(userId)
	if err != nil {
		return nil, err
	}

	res:= (&dto.ReservationList{}).BuildResponse(reservations)
	
	return res, nil
}