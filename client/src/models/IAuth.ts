import { IUser } from "./IUser"

export interface IAuthState {
    user: IUser
    isAuth: boolean
    loading: string
    error: string | undefined
}