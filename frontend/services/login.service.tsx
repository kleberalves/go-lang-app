import Router from 'next/router'
import cookie from 'js-cookie'
import { useContext, useState } from 'react';
import useRequestService from './request.service';
import { validateEmail } from './util.service';
import { MasterContext } from '../contexts/MasterContext';

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
            messageBox?.showError("Verifique o formato do seu e-mail.");
            return;
        }

        setRequestStatus(-1);

        const request = await post(`/credential/login`, pCredential);

        if (request) {
            if (request.status == 200) {
                const resposta = await request.json();
                cookie.set('token', JSON.stringify(resposta), { expires: 30 })

                if (pUrlRedirect) {
                    Router.push(pUrlRedirect);
                } else {
                    Router.push('/manager');
                }
            }

            setRequestStatus(request.status);
        } else {
            setRequestStatus(0);
        }

    }

    return {
        login,
        resquestStatus
    }
}

export default useLoginService;



