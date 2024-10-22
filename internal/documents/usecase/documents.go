package usecase

type documentsUsecaseItf interface {}

type documentsUsecase struct {}

func NewdocumentsUsecase() documentsUsecaseItf {
    return &documentsUsecase{}
}
