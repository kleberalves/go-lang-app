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




let theme = createTheme({
    palette: {
        primary: {
            light: '#63ccff',
            main: '#009be5',
            dark: '#006db3',
        },
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
}
export const MasterPage: React.FunctionComponent<React.PropsWithChildren<MasterPageProps>> = ({ children, hideNavigation }) => {
    const [mobileOpen, setMobileOpen] = React.useState(false);
    const isSmUp = useMediaQuery(theme.breakpoints.up('sm'));

    const handleDrawerToggle = () => {
        setMobileOpen(!mobileOpen);
    };

    return (
        <ThemeProvider theme={theme}>
                <Loading />
                <Box sx={{ display: 'flex', minHeight: '100vh' }}>
                    <CssBaseline />
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
        </ThemeProvider>
    );
}