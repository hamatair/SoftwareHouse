package usecase

type invoicesUsecaseItf interface {}

type invoicesUsecase struct {}

func NewinvoicesUsecase() invoicesUsecaseItf {
    return &invoicesUsecase{}
}
