import { Button, Grid } from "@mui/material";
import { FormContainer, PasswordElement, TextFieldElement } from "react-hook-form-mui";
import { UserCredential } from "../../services/login.service";
import { MasterContext } from "../../components/MasterPage";
import { useContext } from "react";

interface FormLoginProps {
    urlRedirect: string;
}
export const FormLogin: React.FunctionComponent<FormLoginProps> = () => {

    const { messageBox } = useContext(MasterContext);

    const onSubmit = (data: UserCredential) => {
        //login(data, "/manager");

        messageBox?.confirm(`Deseja realmente remover?`, async () => {
            console.log("Ok!")

        });
    }

    return <FormContainer onSuccess={onSubmit}

        FormProps={{
            'aria-autocomplete': 'none',
            autoComplete: 'off'
        }}>
        <Grid container>
            <Grid item xs>
                <TextFieldElement
                    required
                    type={'email'}
                    margin={'dense'}
                    label={'Email'}
                    name={'email'}
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs>
                <PasswordElement margin={'dense'}
                    label={'Password'}
                    required
                    name={'password'}
                />
            </Grid>
            <Grid item xs
                justifyContent="right">
                <Button type={'submit'} color={'primary'} variant={'contained'}>Login</Button>
            </Grid>
        </Grid>
    </FormContainer>
}
