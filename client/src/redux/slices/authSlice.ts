import { Tuple, createSlice } from '@reduxjs/toolkit';
import { isEmailValid } from '../../helpers/emailValidator';

interface IUser {
    email: string,
    id: string
}

interface AuthState {
    user: IUser
    token: null
    isAuth: boolean
    isLoading: boolean
}

const initialState: AuthState = {
    user: {
        email: '',
        id: ''
    },
    token: null,
    isAuth: false,
    isLoading: false
}

const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        setLoading: (state, action) => {
            const { setLoading } = action.payload;
            state.isLoading = setLoading;
        },
        setUser: (state, action) => {
            const { user } = action.payload;
            state.user = user;
        },
        setAuth: (state, action) => {
            const { isAuth } = action.payload;
            state.isAuth = isAuth
        },
        setCreadentials: (state, action) => {
            const { user, accessToken } = action.payload;
            state.user = user;
            state.token = accessToken;
            state.isAuth = true
        },
        logOut: (state, action) => {
            state.user = {} as IUser;
            state.token = null
        }
    }
  });

export const { setCreadentials, logOut } = authSlice.actions;

export default authSlice.reducer;
