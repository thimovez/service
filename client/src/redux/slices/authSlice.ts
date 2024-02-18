import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { IUser } from '../../models/IUser';
import { IAuthState } from '../../models/IAuth';
import { setUser } from './userSlice';
import AuthService from '../../services/AuthService';


const initialState: IAuthState = {
    user: {} as IUser,
    token: null,
    isAuth: false,
    isLoading: false
}

interface RegistrationDataRequest {
    email: string,
    password: string
}

export const registration = createAsyncThunk<IUser, { email: string; password: string }, { rejectValue: string }>(
    'auth/registration',
    async (userData: RegistrationDataRequest, { rejectWithValue, dispatch }) => {
        try {
            const response = await AuthService.registration(userData.email, userData.password);
            
            localStorage.setItem('token', response.data.accessToken);
            dispatch(setAuth(true))
            return response.data.user;
        } catch (error: any) {
            return rejectWithValue(error.response?.data?.message);
        }
    }
);

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
        logOut: (state) => {
            state.user = {} as IUser;
            state.token = null
        }
    },
    extraReducers: (builder) => {   
        builder
            .addCase(registration.fulfilled, (state, action) => {
                state.isAuth = true;
                state.user = action.payload;    
            })
            .addCase(registration.pending, (state) => {
                state.isLoading = true;  
            })
            .addCase(registration.rejected, (state) => {
                state.isLoading = false; 
            });
            
    }
  });

export const { setLoading, setAuth, setCreadentials, logOut } = authSlice.actions;

export default authSlice.reducer;
