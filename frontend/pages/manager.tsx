

import { MasterPage } from '../components/MasterPage'
import { AppBar, Button, Grid, Paper, Tab, Tabs } from '@mui/material'
import { TabPanel } from '../components/panels/TabPanel'
import { useState } from 'react'
import { SecureContextProvider } from '../contexts/SecureContext'
import { FormGridUsers } from '../components/forms/FormGridUsers'

const Page = () => {

  const [indexTab, setIndexTab] = useState<number>(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setIndexTab(newValue);
  };

  return <MasterPage pageTitle='Manager'>

    <Paper sx={{ maxWidth: 936, margin: 'auto', overflow: 'hidden' }}>

      <AppBar component="div" position="static" elevation={0} sx={{ zIndex: 0 }}>
        <Tabs value={indexTab} textColor="inherit" onChange={handleChange}>
          <Tab label="Salesmen" />
          <Tab label="Products" />
          <Tab label="Customers" />
        </Tabs>
      </AppBar>


      <TabPanel value={indexTab} index={0}>
        <FormGridUsers profileType={1} label="Salesman" />
      </TabPanel>
      <TabPanel value={indexTab} index={1}>
        <Grid container spacing={2} alignItems="center">
          <Grid item xs>
          </Grid>
          <Grid item>
            <Button variant="contained" sx={{ mr: 1 }}>
              Add Product
            </Button>
          </Grid>
        </Grid>
      </TabPanel>
      <TabPanel value={indexTab} index={2}>
        <FormGridUsers profileType={2} label="Customer" />
      </TabPanel>
    </Paper>
  </MasterPage >
}

export default SecureContextProvider(Page);
