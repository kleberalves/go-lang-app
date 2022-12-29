import { useContext } from "react";
import { MasterPageContext } from "../components/MasterPage";
import useLoginService from "./login.service";
import useRequestMethodsService from "./request-methods.service";

//This hook encapsulates client-side request methods because of Contexts. These do not work on the server side.
const useRequestService = () => {

    const { messageBox } = useContext(MasterPageContext);
    const { logout } = useLoginService();

    const { post, get, put, patch, upload, del } = useRequestMethodsService(messageBox && messageBox.error, logout);

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