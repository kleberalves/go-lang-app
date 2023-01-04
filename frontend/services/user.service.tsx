
import { useContext, useState } from 'react';
import useRequestService from './request.service';
import { MasterContext } from '../contexts/MasterContext';
import { UserFilter } from '../scheme/UserFilter';
import { User } from '../scheme/User';

export interface UserCredential {
    Email: string;
    Password: string;
};

const useUserService = () => {

    const [resquestStatus, setRequestStatus] = useState<number>(0);
    const { messageBox } = useContext(MasterContext);

    const { post, put, del } = useRequestService();

    const getUsers = async (userFilter: UserFilter): Promise<User[]> => {

        setRequestStatus(-1);

        const request = await post(`/users/get`, userFilter);

        if (request) {
            if (request.status == 200) {
                return await request.json();
            }

            setRequestStatus(request.status);
        } else {
            setRequestStatus(0);
        }

        return []
    }

    const deleteUser = async (userId: number): Promise<void> => {

        setRequestStatus(-1);

        const request = await del(`/users/${userId}`);

        if (request) {
            if (request.status == 200) {
                messageBox?.info("User deleted successfully.");
            }

            setRequestStatus(request.status);
        } else {
            setRequestStatus(0);
        }

    }

    return {
        getUsers,
        deleteUser,
        resquestStatus
    }
}

export default useUserService;



