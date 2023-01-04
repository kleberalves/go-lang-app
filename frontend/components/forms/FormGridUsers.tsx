import { Box, Button, Grid } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { FunctionComponent, useEffect, useState } from "react";
import { User } from "../../scheme/User";
import useUserGridService from "../../services/user.grid.service";
import useUserService from "../../services/user.service";

interface FormGridUsersProps {
    profileType: number;
    label: string;
}

export const FormGridUsers: FunctionComponent<FormGridUsersProps> = ({ profileType, label }) => {

    const [users, setUsers] = useState<User[]>([]);
    const { getUsers } = useUserService();
    const { userGridColumns } = useUserGridService();
    const init = async () => {
        setUsers(await getUsers({ "ProfileType": profileType }));
    };

    useEffect(() => {
        init();
    }, []);


    return <>
        <Grid container spacing={2}>
            <Grid item xs>

            </Grid>
            <Grid item>
                <Button variant="contained" sx={{ mr: 1 }}>
                    Add {label}
                </Button>

            </Grid>
        </Grid>

        {users.length > 0 && <Box sx={{ height: 400, width: '100%' }}>
            <DataGrid
                rows={users}
                columns={userGridColumns}
                pageSize={5}
                getRowId={(row) => row.ID}
                rowsPerPageOptions={[5]}
                checkboxSelection
                disableSelectionOnClick
                experimentalFeatures={{ newEditingApi: true }}
            />
        </Box>}</>


}