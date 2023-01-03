import React from "react";
import { useEffect } from 'react'
import Router from 'next/router'
import cookie from 'js-cookie'
import nextCookies from 'next-cookies'

export const checkSecurity = async (ctx: any) => {
    let { token } = nextCookies(ctx) as any;

    if (!ctx.req) {
        //Client side
        let cookieToken: string | undefined = cookie.get('token');

        if (!token && cookieToken) {
            token = JSON.parse(cookieToken);
        }

        if (!token) {
            Router.push(`/login?url=${ctx.asPath}`);
        }
    } else {
        //Server side
        if (!token && ctx.res) {
            ctx.res.writeHead(301, {
                Location: `/login?url=${ctx.asPath}`
            });
            ctx.res.end();
        }
    }

    return token;
}

export const SecureContextProvider = (WrappedComponent: any) => {

    const SecureContextWrapper = (props: any) => {

        const logout = () => {
            cookie.remove('token');
            Router.push('/login');
        }

        return <SecureContext.Provider value={{
            logout: logout
        }}>
            <WrappedComponent {...props} />
        </SecureContext.Provider >
    }

    return SecureContextWrapper
}

type SecureContextProps = {
    logout: () => void;
}

export const SecureContext = React.createContext<SecureContextProps>({
    logout: () => { }
});
