import { MasterPage } from '../components/MasterPage'
import { Box, Grid } from '@mui/material'

import { FormLogin } from '../components/forms/FormLogin'


export default function Login() {
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
        padding: 5
      }}>
        <h2>Welcome</h2>
        <FormLogin urlRedirect="/manager" />
      </Grid>
    </Box>

  </MasterPage >
}
