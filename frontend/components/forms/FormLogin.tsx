import { Box, Button, Grid } from "@mui/material";
import { FormContainer, PasswordElement, TextFieldElement } from "react-hook-form-mui";
import useLoginService, { UserCredential } from "../../services/login.service";

import { useContext, useEffect } from "react";
import { MasterContext } from "../../contexts/MasterContext";

interface FormLoginProps {
    urlRedirect: string;
}
export const FormLogin: React.FunctionComponent<FormLoginProps> = () => {

    const { messageBox } = useContext(MasterContext);
    const { login, resquestStatus } = useLoginService();

    const onSubmit = (data: UserCredential) => {
        login(data, "/manager");
    }


    return <FormContainer onSuccess={onSubmit}

        FormProps={{
            'aria-autocomplete': 'none',
            autoComplete: 'new-password'
        }}>
        <Grid container>
            <Box width={"100%"} >
                <TextFieldElement
                    required
                    style={{
                        width: "100%"
                    }}
                    type={'email'}
                    margin={'dense'}
                    label={'Email'}
                    name={'Email'}
                    autoComplete="new-password"
                    inputProps={{
                        autoComplete: "dont-fill-me",
                    }}
                />
            </Box>
            <Box width={"100%"}>
                <PasswordElement margin={'dense'}
                    label={'Password'}
                    required
                    name={'Password'}
                    autoComplete="new-password"
                    inputProps={{
                        autoComplete: "dont-fill-me",
                    }}
                />
            </Box>
            <Box width={"100%"} display="flex"
                justifyContent="right">
                <Button type={'submit'} color={'primary'} variant={'contained'}>
                    {resquestStatus == -1 ?
                        <>Loading...</> :
                        <>Login</>}
                </Button>
            </Box>
        </Grid>
    </FormContainer>
}
