import { MasterPage } from '../components/MasterPage'
import { Box, Button, Grid } from '@mui/material'
import Router from 'next/router'


export default function Home() {
    const onClick = (e: any) => {
        Router.push('/manager');
    }
    return <MasterPage hideNavigation={true} pageTitle={"Welcome login"}>

        <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            minHeight="100vh"
        >
            <Grid item xs={3} sx={{
                bgcolor: 'background.paper',
                boxShadow: 1,
                borderRadius: 2,
                padding: 5,
                justifyContent: "center",
                flexFlow: "column",
                display: "flex"
            }}>
                <h1>Problem Company App</h1>
                <br />
                <Button onClick={onClick}> Login </Button>
            </Grid>
        </Box>

    </MasterPage >
}
