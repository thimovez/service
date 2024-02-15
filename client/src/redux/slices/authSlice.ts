import { createSlice } from '@reduxjs/toolkit';
import { IUser } from '../../models/IUser';
import { IAuthState } from '../../models/IAuth';

const initialState: IAuthState = {
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
