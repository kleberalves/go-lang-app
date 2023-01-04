import { Button } from "@mui/material";
import { GridColDef, GridValueGetterParams } from "@mui/x-data-grid";
import { useContext } from "react";
import { MasterContext } from "../contexts/MasterContext";
import { User } from "../scheme/User";
import useUserService from "./user.service";
import { formatDate } from "./util.service";


const useUserGridService = () => {

    const { messageBox } = useContext(MasterContext);
    const { deleteUser } = useUserService();

    const execEdit = (user: User) => {

        messageBox?.info(`Edit ${user.FirstName}?`);

    };

    const execRemove = (firstName: string, userId: number) => {

        console.log(firstName, messageBox);

        messageBox?.showConfirm(`Do really want to remove ${firstName}?`, async () => {
            await deleteUser(userId);
        });

    };

    const userGridColumns: GridColDef[] = [
        { field: 'ID', headerName: 'ID', width: 90 },
        {
            field: 'FirstName',
            headerName: 'First name',
            width: 150,
            editable: true,
        },
        {
            field: 'LastName',
            headerName: 'Last name',
            width: 110,
            editable: true,
        },
        {
            field: 'Email',
            headerName: 'E-mail',
            width: 240,
            editable: true,
        },
        {
            field: 'CreatedAt',
            headerName: 'Created date',
            sortable: false,
            width: 160,
            valueGetter: (params: GridValueGetterParams) => formatDate(params.row.CreatedAt),
        },
        {
            field: "action",
            headerName: "Action",
            sortable: false,
            width: 190,
            renderCell: (params) => {
                const onClickEdit = (e: any) => {
                    e.stopPropagation(); // don't select this row after clicking
                    execEdit(params.row);
                };

                const onClickRemove = (e: any) => {
                    e.stopPropagation(); // don't select this row after clicking
                    execRemove(params.row.FirstName, params.row.ID);
                };

                return <>
                    <Button onClick={onClickEdit}>Edit</Button>  <br />
                    <Button onClick={onClickRemove}>Remove</Button>
                </>;
            }
        }
    ];
    return {
        userGridColumns
    }
}


export default useUserGridService;