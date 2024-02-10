import axios from 'axios';
import {AuthResponse} from "../models/response/AuthResponse";
import {store} from "../index";
import {IUser} from "../models/IUser";

export const API_URL = process.env.API_URL  

// Created new instace of axios for authorization requests
const $authHost = axios.create({
    withCredentials: true,
    baseURL: API_URL
})

// Interceptor receives the Bearer header and assigns 
// token received from local storage for every request 
// made from this instace
$authHost.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
    return config;
})

$authHost.interceptors.response.use((config) => {
    return config;
},async (error) => {
    const originalRequest = error.config;
    if (error.response.status == process.env.StatusUnauthorized && error.config && !error.config._isRetry) {
        originalRequest._isRetry = true;
        try {
            const response = await axios.get<AuthResponse>(`${API_URL}/refresh`, {withCredentials: true})
            localStorage.setItem('token', response.data.accessToken);
            return $authHost.request(originalRequest);
        } catch (e) {
            console.log('Not Authorized')
        }
    }
    throw error;
})

export default $authHost;
