import React from "react"
import { MessageBox } from "../services/message-box.service"



export interface MasterPageContextProps {
    messageBox?: MessageBox
}

export const MasterContext = React.createContext<MasterPageContextProps>({
    messageBox: undefined
})