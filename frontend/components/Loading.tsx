import { Fragment, FunctionComponent, useEffect } from "react";
import { usePromiseTracker } from "react-promise-tracker";
import NProgress from 'nprogress'
import { Router } from "next/router";

Router.events.on('routeChangeStart', url => {
    NProgress.start()
});
Router.events.on('routeChangeComplete', (url) => {
});
Router.events.on('routeChangeError', () => NProgress.done());

export const Loading: FunctionComponent = (props) => {

    const { promiseInProgress } = usePromiseTracker();
    useEffect(() => {

        if (promiseInProgress) {
            NProgress.start()
        } else {
            NProgress.done()
        }

    }, [promiseInProgress]);

    return <Fragment></Fragment>
}