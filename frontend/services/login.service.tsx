import Router from 'next/router'
import cookie from 'js-cookie'
import { useContext, useState } from 'react';
import useRequestService from './request.service';
import { validateEmail } from './util.service';
import { MasterContext } from '../components/MasterPage';
export interface UserCredential {
    Email: string;
    Password: string;
};

const useLoginService = () => {

    const [resquestStatus, setRequestStatus] = useState<number>(0);
    const { messageBox } = useContext(MasterContext);

    const { post, get, put } = useRequestService();

    const login = async (pCredential: UserCredential, pUrlRedirect: string): Promise<void> => {

        if (!validateEmail(pCredential.Email)) {
            messageBox?.error("Verifique o formato do seu e-mail.");
            return;
        }

        const request = await post(`/credential/login`, pCredential);

        if (request) {
            if (request.status == 200) {
                const resposta = await request.json();
                cookie.set('token', JSON.stringify(resposta), { expires: 30 })

                if (pUrlRedirect) {
                    Router.push(pUrlRedirect);
                } else {
                    Router.push('/products');
                }
            }

            setRequestStatus(request.status);
        }

    }

    const logout = () => {
        cookie.remove('token');
        window.localStorage.setItem('logout', Date.now().toString());
        Router.push('/login');
    }


    return {
        login,
        logout,
        resquestStatus
    }
}

export default useLoginService;



