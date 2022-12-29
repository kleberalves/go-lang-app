import { useState, useEffect, useReducer } from "react";


interface MessageBoxHook {
    messageBox: MessageBox;
}

export class MessageBoxTypes {
    public static Info: string = "info";
    public static Error: string = "error";
    public static Success: string = "success";
    public static Confirm: string = "confirm";
}

export interface MessageBox {
    state: MessageBoxState,
    element: any,
    success: (msg: string) => void,
    info: (msg: string) => void,
    error: (msg: string) => void,
    confirm: (msg: string, pOnConfirm: Function) => void,
    confirmOk: () => void,
    close: Function
}

interface MessageBoxState {
    confirmOk: () => void;
    message: string;
    window: string;
    type: MessageBoxTypes
}

const useMessageBox = (): MessageBoxHook => {

    const [messageEl, setContainerEl] = useState<any>();

    const reducerMessage = (state: any, action: any) => {
        switch (action.type) {

            case 'CONFIRM':
                return {
                    ...state,
                    window: "show",
                    type: MessageBoxTypes.Confirm,
                    message: action.value,
                    confirmOk: action.confirmOk
                };

            case 'SUCCESS':
                return {
                    ...state,
                    window: "show",
                    type: MessageBoxTypes.Success,
                    message: action.value
                };

            case 'INFO':
                return {
                    ...state,
                    window: "show",
                    type: MessageBoxTypes.Info,
                    message: action.value
                };

            case 'ERROR':
                return {
                    ...state,
                    window: "show",
                    type: MessageBoxTypes.Error,
                    message: action.value
                };

            case 'CLOSE':
                return {
                    ...state,
                    window: "close",
                    message: ""
                };

            default: return state;
        };
    }

    const initial = {
        confirmOk: () => { },
        message: "",
        window: "",
        type: MessageBoxTypes.Info
    }

    const [stateMessage, dispatchMessage] = useReducer(reducerMessage, initial);

    useEffect(() => {
        setContainerEl(document.createElement("div"));
    }, []);

    useEffect(() => {
        if (messageEl) {
            document.body.appendChild(messageEl);

        }
        return () => {
            if (messageEl) {
                document.body.removeChild(messageEl);
            }
        }
    }, [messageEl]);

    const success = async (msg: string) => {
        dispatchMessage({
            type: 'SUCCESS',
            value: msg
        });
    }

    const info = async (msg: string) => {
        dispatchMessage({
            type: 'INFO',
            value: msg
        });
    }

    const close = async () => {
        dispatchMessage({
            type: 'CLOSE'
        });
    }

    const error = async (msg: string) => {
        dispatchMessage({
            type: 'ERROR',
            value: msg
        });
    }

    const confirm = async (msg: string, pOnConfirm: Function) => {
        dispatchMessage({
            type: 'CONFIRM',
            value: msg,
            confirmOk: pOnConfirm
        });
    }

    const confirmOk = async () => {

        stateMessage.confirmOk();

        dispatchMessage({
            type: 'CLOSE',
        });
    }

    return {
        messageBox: {
            state: stateMessage,
            element: messageEl,
            success: success,
            confirm: confirm,
            confirmOk: confirmOk,
            info: info,
            error: error,
            close: close
        }
    }

}

export default useMessageBox;