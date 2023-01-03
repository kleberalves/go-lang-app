import '../styles/globals.css'
import type { AppProps, AppInitialProps } from 'next/app'
import { checkSecurity } from '../contexts/SecureContext'

const App = ({ Component, pageProps }: AppProps) => {
  return <Component {...pageProps} />
}

App.getInitialProps = async ({ Component, ctx }: any) => {
  let props = {}

  if (Component.name == "LoginWrapper") {
    //Check security at server side for first load.
    await checkSecurity(ctx);
  }

  return { ...props }
}

export default App;
