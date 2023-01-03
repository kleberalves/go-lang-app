import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { checkSecurity } from '../contexts/SecureContext'

const App = ({ Component, pageProps }: AppProps) => {
  return <Component {...pageProps} />
}

App.getInitialProps = async ({ Component, ctx, res }: any) => {
  let props: any = {};

  if (Component.name == "SecureContextWrapper") {
    //Check security.
    props["token"] = await checkSecurity(ctx);;
  }

  return { ...props }
}

export default App;
