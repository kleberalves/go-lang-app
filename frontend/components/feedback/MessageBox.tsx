import React, { FunctionComponent, useEffect } from "react";
import { MessageBox, MessageBoxTypes } from "../../services/message-box.service";
import { Box, Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Divider, Modal, Typography } from "@mui/material";

type MessageBoxProps = {
    messageBox: MessageBox
}

export const MessageBoxDialog: FunctionComponent<MessageBoxProps> = ({ messageBox }) => {

    const [open, setOpen] = React.useState(false);

    const handleClose = () => {
        messageBox.close();
    };

    useEffect(() => {

        if (messageBox.state.window !== "close" &&
            messageBox.state.window !== "") {
            setOpen(true);
        } else if (messageBox.state.window == "close") {
            setOpen(false);
        }

    }, [messageBox.state.window])

    return <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"

    >
        <DialogContent sx={{
            borderWidth: 1,
            borderColor: "red"
        }}>
            <DialogContentText id="alert-dialog-description">
                <span dangerouslySetInnerHTML={{ __html: messageBox.state.message }}></span>
            </DialogContentText>
        </DialogContent>
        <DialogActions>
            {messageBox.state.type == MessageBoxTypes.Confirm &&
                <>
                    <Button onClick={() => messageBox.close()}>Cancel</Button>
                    <Button onClick={() => messageBox.confirmOk()} autoFocus>
                        Ok
                    </Button>
                </>}
            {messageBox.state.type != MessageBoxTypes.Confirm &&
                <Button value="Ok" autoFocus onClick={() => messageBox.close()} />}

        </DialogActions>
    </Dialog >
}
