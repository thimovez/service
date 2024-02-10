import $authHost from "../http";
import {AxiosResponse} from 'axios';
import {AuthResponse} from "../models/response/AuthResponse";

export default class AuthService {
    static async login(email: string, password: string): Promise<AxiosResponse<AuthResponse>> {
        return $authHost.post<AuthResponse>('/login', {email, password})
    }

    static async registration(email: string, password: string): Promise<AxiosResponse<AuthResponse>> {
        return $authHost.post<AuthResponse>('/registration', {email, password})
    }

    static async logout(): Promise<void> {
        return $authHost.post('/logout')
    }

}

