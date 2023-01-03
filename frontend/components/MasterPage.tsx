import * as React from 'react';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';
import Navigator from './Navigator';
import Header from './Header';
import { Loading } from './Loading';
import Head from 'next/head';
import useMessageBox from '../services/message-box.service';
import { MessageBoxDialog } from './feedback/MessageBox';
import { Footer } from './Footer';
import { MasterContext } from '../contexts/MasterContext';

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
        MuiInputBase: {
            styleOverrides: {
                input: {
                    '&:-webkit-autofill': {
                        transitionDelay: '9999s',
                        transitionProperty: 'background-color, color',
                    }
                }
            },
        },
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
    const [validPage, setValidPage] = React.useState(false);

    //For global messages dialog.
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
                <title>{`Problem Company App Test - ${pageTitle}`}</title>
                <meta name="description" content="Application by Kleber Alves for the Problem Company test." />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.svg" />
                <link type="text/css" href="/nprogress.css" as="style" />
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
                    <Footer />
                </Box>
            </Box>
            <MessageBoxDialog messageBox={messageBox} />
        </ThemeProvider>
    </MasterContext.Provider>
}
