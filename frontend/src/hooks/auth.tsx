import React, { useContext, useEffect, useState } from "react";

export type Nullable<T> = T|null;

export type Auth = {
    username: Nullable<string>;
    token: Nullable<string>;
    refreshToken: Nullable<string>;

    isAnon: boolean;
    loaded: boolean;
};

export type AuthCtx = Auth & {
    login: (username: string, password: string) => void;
    anonLogin: (username: string) => void;
};

const initialState: Auth = {
    username: null,
    token: null,
    refreshToken: null,
    isAnon: true,
    loaded: false,
}

const AuthContext = React.createContext<AuthCtx>({
    ...initialState,
    login: (username, password) => {},
    anonLogin: (username) => {},
});

export function AuthProvider({children}: {children: React.ReactNode}) {
    const [state, setState] = useState<Auth>(initialState);

    const login = (username: string, password: string) => {
        localStorage.setItem('last_username', username);
        localStorage.setItem('token', 'token');
        localStorage.setItem('refresh_token', 'refresh_token');
    };

    const anonLogin = (username: string) => {
        localStorage.setItem('last_username', username);
        localStorage.removeItem('token');
        localStorage.removeItem('refresh_token');

        setState({...state, username, token: null, refreshToken: null});
    };

    useEffect(() => {
        setState({
            ...state,
            username: localStorage.getItem('last_username'),
            token: localStorage.getItem('token'),
            refreshToken: localStorage.getItem('refresh_token'),
            loaded: true,
        })
    }, []);

    return <AuthContext.Provider value={{
        ...state,
        login,
        anonLogin,
    }}>
        {children}
    </AuthContext.Provider>
}

export function useAuth() {
    return useContext<AuthCtx>(AuthContext);
}