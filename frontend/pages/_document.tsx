import { Html, Head, Main, NextScript } from 'next/document'
import { Loading } from '../components/Loading'

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <link type="text/css" rel="stylesheet" href="/nprogress.css" />
      </Head>
      <body>
        <Loading />

        <Main />
        <NextScript />
      </body>
    </Html>
  )
}
