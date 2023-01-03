import { Box, Link, Typography } from "@mui/material";

export const Footer = () => {
    return (
        <Box component="footer" sx={{ p: 2, bgcolor: '#eaeff1' }}>
            <Typography variant="body2" color="text.secondary" align="center">
                {'Copyright © '}
                <Link color="inherit" href="https://github.com/kleberalves/problemCompanyApp">
                    Problem Company App Test
                </Link>{' '}
                {new Date().getFullYear()}.
            </Typography>
        </Box>
    );
}
