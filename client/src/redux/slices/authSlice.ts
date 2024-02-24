import { createSlice, createAsyncThunk, PayloadAction, UnknownAction } from '@reduxjs/toolkit';
import { IUser } from '../../models/IUser';
import { IAuthState } from '../../models/IAuth';
import { setUser } from './userSlice';
import AuthService from '../../services/AuthService';
import axios, { Axios, AxiosError } from 'axios';


const initialState: IAuthState = {
    user: {} as IUser,
    isAuth: false,
    loading: "",
    error: ""
}

interface RegistrationDataRequest {
    email: string,
    password: string
}

export const registration = createAsyncThunk<IUser, RegistrationDataRequest, { rejectValue: string }>(
    'auth/registration',
    async (userData, { rejectWithValue }) => {
        try {
            // const response = await AuthService.registration(userData.email, userData.password);
            // if (response.status != 201) {
            //     throw rejectWithValue("Some error")
            // }
            // localStorage.setItem('token', response.data.accessToken);
            const user = {} as IUser
            // return response.data.user;
            return user
        } catch (error: any) {
            if (axios.isAxiosError(error)) {
                // do something
                // or just re-throw the error
                return rejectWithValue(error.response?.statusText!)
            }
            return rejectWithValue(error);
        }
    }
);

const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        // setLoading: (state, action) => {
        //     const { setLoading } = action.payload;
        //     state.isLoading = setLoading;
        // },
        // setAuth: (state, action) => {
        //     const { isAuth } = action.payload;
        //     state.isAuth = isAuth
        // },
        // setCreadentials: (state, action) => {
        //     const { user, accessToken } = action.payload;
        //     state.user = user;
        //     state.token = accessToken;
        //     state.isAuth = true
        // },
        // logOut: (state) => {
        //     state.user = {} as IUser;
        //     state.token = null
        // }
    },
    extraReducers: (builder) => {   
        builder
            .addCase(registration.fulfilled, (state, action) => {
                state.isAuth = true;
                state.user = action.payload;
                state.loading = "success"    
            })
            .addCase(registration.pending, (state) => {
                state.loading = "pending";  
            })
            .addCase(registration.rejected, (state, action) => {
                state.isAuth = false
                state.loading = "rejected";
                state.error = action.payload
            });
            
    }
  });

// export const { setLoading, setAuth, setCreadentials, logOut } = authSlice.actions;

export default authSlice.reducer;
