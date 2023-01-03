import { useContext } from "react";
import { MasterContext } from "../contexts/MasterContext";
import { SecureContext } from "../contexts/SecureContext";
import useRequestMethodsService from "./request-methods.service";

//This hook encapsulates client-side request methods because of Contexts (These do not work on the server side).
const useRequestService = () => {

    const { logout } = useContext(SecureContext);
    const { messageBox } = useContext(MasterContext);

    const { post, get, put, patch, upload, del } = useRequestMethodsService(messageBox?.showError, logout);

    return {
        del,
        post,
        get,
        put,
        patch,
        upload
    }
}

export default useRequestService;