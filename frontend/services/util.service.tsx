import getConfig from 'next/config'
import 'isomorphic-unfetch'

const {
    publicRuntimeConfig: publicRuntimeConfig,
    serverRuntimeConfig: serverRuntimeConfig
} = getConfig()

export const getenv = (key: string): string | undefined => {

    if (key == "BASE_URL_SERVER") {
        if (serverRuntimeConfig["BASE_URL_SERVER"]) {
            //Production server
            return serverRuntimeConfig["BASE_URL_SERVER"];
        } else {
            //Production client
            return publicRuntimeConfig["BASE_URL"];
        }
    } else {
        return serverRuntimeConfig[key] ? serverRuntimeConfig[key] : publicRuntimeConfig[key];
    }
}
/**
 * Clona um objeto removendo a referência.
 */
export const clone = (obj: any) => {
    return JSON.parse(JSON.stringify(obj));
}

export const validateEmail = (email: string) => {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}



