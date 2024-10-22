package usecase

type rolesUsecaseItf interface {}

type rolesUsecase struct {}

func NewrolesUsecase() rolesUsecaseItf {
    return &rolesUsecase{}
}
