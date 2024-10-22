package usecase

type userRolesUsecaseItf interface {}

type user_rolesUsecase struct {}

func NewuserRolesUsecase() userRolesUsecaseItf {
    return &user_rolesUsecase{}
}
