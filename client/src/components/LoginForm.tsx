import React, {FC, useContext, useState} from 'react';
import {Context} from "../index";
import {observer} from "mobx-react-lite";
import '../App.css';

const LoginForm: FC = () => {
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [error, setError] = useState('')
    const {store} = useContext(Context);

    const isEmailValid = (email: string) => {
        // Basic email validation using a regular expression
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
      };

    const handleRegistration = () => {
        if (!email || !password) {
          setError("all fields must be filled in");
        } else if (!isEmailValid(email)) {
            setError("Enter valid email"); 
        } else if (password.length <= 3) {
          setError("password must be longer than 5 symbol");
        } else {
          setError('');
        }
      };

    return (
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
                handleRegistration()
                store.login(email, password)}}>
                Sign in
            </button>
            <button className='registration-button' onClick={() => store.registration(email, password)}>
                Registration
            </button>
        </div>
        </div>
    );
};

export default observer(LoginForm);
