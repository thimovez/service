import React, { useEffect, useState} from 'react';
import '../App.css';
import { isEmailValid } from '../helpers/emailValidator';
import { useSelector } from 'react-redux';
import { RootState, AppDispatch } from '../redux/store';
import { setUser } from '../redux/slices/userSlice';
import { logOut, registration } from '../redux/slices/authSlice';
import {useAppDispatch, useAppSelector} from '../redux/hook';

const LoginForm: React.FC = () => {
    const id: string = "1";
    const {isAuth, isLoading, token, user} = useAppSelector((state) => state.authReducer)
    const dispatch = useAppDispatch();
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [error, setError] = useState('')

    
    // Check if user input correct data
    const validateRegistrationForm = () => {
        if (!email || !password) {
          setError("All fields must be filled in");
        } else if (!isEmailValid(email)) {
          setError("Enter valid email"); 
        } else if (password.length <= 3) {
          setError("password must be longer than 5 symbol");
        } else {
          setError('');
        }
      };

    return (
      <div>
        {isLoading ? (
          <p>Loading...</p>
        ) : isAuth ? (
          <>
            <p>Welcome, {user.email}</p>
            <button onClick={() => dispatch(logOut())}>Logout</button>
          </>
        ) : (
          <div className='container'>
        <div className='authorization-form'>
            <input className='email-input'
                onChange={e => setEmail(e.target.value)}
                value={email}
                type="text"
                placeholder='Email'
            />
            <input className='password-input'
                onChange={e => setPassword(e.target.value)}
                value={password}
                type="password"
                placeholder='Password'
            />
            {<p className='error-input-form'>{error}</p>}
            <button className='login-button' onClick={() => {
                validateRegistrationForm()
                }}>
                Sign in
            </button>
            <button className='registration-button' onClick={() => dispatch(registration({email, password}))}>
                Register
            </button>
        </div>
        </div>
        )}
      </div>
    );
};

export default LoginForm;
