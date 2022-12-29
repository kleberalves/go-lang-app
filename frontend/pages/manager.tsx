import Head from 'next/head'
import { Inter } from '@next/font/google'
import { MasterPage } from '../components/MasterPage'
import { AppBar, Button, Grid, Paper, Tab, Tabs, TextField, Toolbar, Typography } from '@mui/material'
import { TabPanel } from '../components/panels/TabPanel'
import { useState } from 'react'

const inter = Inter({ subsets: ['latin'] })

export default function Manager() {

  const [indexTab, setIndexTab] = useState<number>(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setIndexTab(newValue);
  };
  return <MasterPage>
    <Head>
      <title>Problem Company App Test - Manager</title>
      <meta name="viewport" content="width=device-width, initial-scale=1" />
      <link rel="icon" href="/favicon.svg" />
    </Head>

    <Paper sx={{ maxWidth: 936, margin: 'auto', overflow: 'hidden' }}>

      <AppBar component="div" position="static" elevation={0} sx={{ zIndex: 0 }}>
        <Tabs value={indexTab} textColor="inherit" onChange={handleChange}>
          <Tab label="Salesmen" />
          <Tab label="Products" />
          <Tab label="Customers" />
        </Tabs>
      </AppBar>


      <TabPanel value={indexTab} index={0}>
       
          <Toolbar>
            <Grid container spacing={2} alignItems="center">
              <Grid item xs>
               
              </Grid>
              <Grid item>
                <Button variant="contained" sx={{ mr: 1 }}>
                  Add Salesman
                </Button>

              </Grid>
            </Grid>
          </Toolbar>
      


      </TabPanel>
      <TabPanel value={indexTab} index={1}>

          <Toolbar>
            <Grid container spacing={2} alignItems="center">
              <Grid item xs>
               
              </Grid>
              <Grid item>
                <Button variant="contained" sx={{ mr: 1 }}>
                  Add Product
                </Button>

              </Grid>
            </Grid>
          </Toolbar>
      </TabPanel>
      <TabPanel value={indexTab} index={2}>
    
          <Toolbar>
            <Grid container spacing={2} alignItems="center">
              <Grid item xs>
               
              </Grid>
              <Grid item>
                <Button variant="contained" sx={{ mr: 1 }}>
                  Add Customer
                </Button>

              </Grid>
            </Grid>
          </Toolbar>
      </TabPanel>
    </Paper>



  </MasterPage>
}
