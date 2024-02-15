
// export default class Store {
//     async login(email: string, password: string) {
//         try {
//             const response = await AuthService.login(email, password);
//             console.log(response)
//             localStorage.setItem('token', response.data.accessToken);
//             this.setAuth(true);
//             this.setUser(response.data.user);
//         } catch (e: any) {
//             console.log(e.response?.data?.message);
//         }
//     }

//     async registration(email: string, password: string) {
//         try {
//             const response = await AuthService.registration(email, password);
//             console.log(response)
//             localStorage.setItem('token', response.data.accessToken);
//             this.setAuth(true);
//             this.setUser(response.data.user);
//         } catch (e: any) {
//             console.log(e.response?.data?.message);
//         }
//     }

//     async logout() {
//         try {
//             const response = await AuthService.logout();
//             localStorage.removeItem('token');
//             this.setAuth(false);
//             this.setUser({} as IUser);
//         } catch (e: any) {
//             console.log(e.response?.data?.message);
//         }
//     }

//     async checkAuth() {
//         this.setLoading(true);
//         try {
//             const response = await axios.get<AuthResponse>(`${API_URL}/refresh`, {withCredentials: true})
//             console.log(response);
//             localStorage.setItem('token', response.data.accessToken);
//             this.setAuth(true);
//             this.setUser(response.data.user);
//         } catch (e: any) {
//             console.log(e.response?.data?.message);
//         } finally {
//             this.setLoading(false);
//         }
//     }
// }

import { configureStore } from '@reduxjs/toolkit';
import userReducer from './slices/userSlice';
import authReducer from './slices/authSlice'

export const store = configureStore({
    reducer: {userReducer, authReducer}
});

  // Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;

export default store;