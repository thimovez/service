import React, {FC, useContext, useState} from 'react';
import {Context} from "../index";
import {observer} from "mobx-react-lite";
import '../App.css';

const LoginForm: FC = () => {
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const {store} = useContext(Context);

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
            <button className='login-button' onClick={() => store.login(email, password)}>
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
