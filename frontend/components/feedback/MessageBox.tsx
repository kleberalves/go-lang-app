import React, { FunctionComponent, useEffect } from "react";
import { MessageBox, MessageBoxTypes } from "../../services/message-box.service";
import { Box, Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Divider, Modal, Typography } from "@mui/material";

type MessageBoxProps = {
    messageBox: MessageBox
}

export const MessageBoxComp: FunctionComponent<MessageBoxProps> = ({ messageBox }) => {

    const [open, setOpen] = React.useState(false);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    useEffect(() => {

        console.log(messageBox.state);

        if (messageBox.state.message !== "") {
            setOpen(true);
        } else if (open) {
            setOpen(false);
        }

    }, [messageBox.state.message])

    return <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
    >
        <DialogContent>
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
