import React, { useState } from 'react';

import '../App.css';
import { isEmailValid } from '../helpers/emailValidator';
import { registration } from '../redux/slices/authSlice';
import {useAppDispatch, useAppSelector} from '../redux/hook';

const LoginForm: React.FC = () => {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    // TODO change type of generic, and add logic for error
    const [error1, setError] = useState<string>('');

    const {isAuth, loading, user, error} = useAppSelector((state) => state.auth)
    const dispatch = useAppDispatch();
    
    // TODO remove validation logic from this file
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
        {loading ? (
          <div>
          <p>Loading...</p>
          <p>{error}</p>
          </div>
        ) : isAuth ? (
          <>
            <p>Welcome, {user.email}</p>
            {/* <button onClick={() => dispatch(logOut())}>Logout</button> */}
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
            {<p className='error-input-form'>{error1}</p>}
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
