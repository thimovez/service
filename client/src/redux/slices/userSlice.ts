import { createSlice } from '@reduxjs/toolkit'

export const counterSlice = createSlice({
    name: 'user',
    initialState: {
        loading: false,
        user: null,
        error: null
    },
    reducers: {
        setUser: (state, action) => {
            const { user } = action.payload;
            state.user = user;
        },
    }
})

export default counterSlice.reducer;
