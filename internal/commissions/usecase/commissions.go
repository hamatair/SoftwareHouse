package usecase

type commissionsUsecaseItf interface {}

type commissionsUsecase struct {}

func NewcommissionsUsecase() commissionsUsecaseItf {
    return &commissionsUsecase{}
}
