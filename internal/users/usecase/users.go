package usecase

type usersUsecaseItf interface {}

type usersUsecase struct {}

func NewusersUsecase() usersUsecaseItf {
    return &usersUsecase{}
}
