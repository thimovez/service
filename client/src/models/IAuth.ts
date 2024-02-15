import { IUser } from "./IUser"

export interface IAuthState {
    user: IUser
    token: null
    isAuth: boolean
    isLoading: boolean
}