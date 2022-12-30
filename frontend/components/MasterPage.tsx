import * as React from 'react';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Link from '@mui/material/Link';
import Navigator from './Navigator';
import Header from './Header';
import { Loading } from './Loading';
import Head from 'next/head';
import useMessageBox, { MessageBox } from '../services/message-box.service';
import { MessageBoxComp } from './feedback/MessageBox';


function Copyright() {
    return (
        <Typography variant="body2" color="text.secondary" align="center">
            {'Copyright © '}
            <Link color="inherit" href="https://github.com/kleberalves/problemCompanyApp">
                Problem Company App Test
            </Link>{' '}
            {new Date().getFullYear()}.
        </Typography>
    );
}



/**
 *  palette: {
       
    },
    typography: {
        h5: {
            fontWeight: 500,
            fontSize: 26,
            letterSpacing: 0.5,
        },
    },
    shape: {
        borderRadius: 8,
    },
    components: {
        MuiTab: {
            defaultProps: {
                disableRipple: true,
            },
        },
    },
 */
let theme = createTheme({
    palette: {
        mode: "dark",
        background: {
            paper: '#081627'
        },
        primary: {
            light: '#63ccff',
            main: '#009be5',
            dark: '#006db3',
        },
    },
    mixins: {
        toolbar: {
            minHeight: 48,
        },
    },
});

theme = {
    ...theme,
    components: {
        MuiDrawer: {
            styleOverrides: {
                paper: {
                    backgroundColor: '#081627',
                },
            },
        },

    },
};

const drawerWidth = 256;

interface MasterPageProps {
    hideNavigation?: boolean;
    pageTitle?: string;
}
export const MasterPage: React.FunctionComponent<React.PropsWithChildren<MasterPageProps>> = ({ children, hideNavigation, pageTitle }) => {
    const [mobileOpen, setMobileOpen] = React.useState(false);
    const isSmUp = useMediaQuery(theme.breakpoints.up('sm'));


    //For global message 
    const { messageBox } = useMessageBox();
    
    const handleDrawerToggle = () => {
        setMobileOpen(!mobileOpen);
    };

    return <MasterContext.Provider value={{
        messageBox: messageBox
    }}>
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Loading />

            <Head>
                <title>Problem Company App Test - {pageTitle}</title>
                <meta name="description" content="Application by Kleber Alves for the Problem Company test." />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.svg" />
            </Head>
            <Box sx={{ display: 'flex', minHeight: '100vh' }}>

                {hideNavigation ? null : (<Box
                    component="nav"
                    sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 } }}
                >
                    <Navigator
                        PaperProps={{ style: { width: drawerWidth } }}
                        variant="temporary"
                        open={mobileOpen}
                        onClose={handleDrawerToggle}
                    />

                    <Navigator
                        PaperProps={{ style: { width: drawerWidth } }}
                        sx={{ display: { sm: 'block', xs: 'none' } }}
                    />
                </Box>)}
                <Box sx={{ flex: 1, display: 'flex', flexDirection: 'column' }}>
                    {hideNavigation ? null : (<Header onDrawerToggle={handleDrawerToggle} />)}
                    <Box component="main" sx={{ flex: 1, py: 6, px: 4, bgcolor: '#eaeff1' }}>
                        {children}
                    </Box>
                    <Box component="footer" sx={{ p: 2, bgcolor: '#eaeff1' }}>
                        <Copyright />
                    </Box>
                </Box>
            </Box>
            <MessageBoxComp messageBox={messageBox} />
        </ThemeProvider>
    </MasterContext.Provider >
}

export interface MasterPageContextProps {
    messageBox?: MessageBox
}

export const MasterContext = React.createContext<MasterPageContextProps>({ messageBox: undefined })