import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IUser } from '../../models/IUser';

const user = {} as IUser

export const counterSlice = createSlice({
    name: 'user',
    initialState: user,
    reducers: {
        setUser: (state, action: PayloadAction<IUser>) => {
            const user  = action.payload;
            state.email = user.email;
            state.id = user.id;
        },
    }
})

export const { setUser } = counterSlice.actions

export default counterSlice.reducer;
