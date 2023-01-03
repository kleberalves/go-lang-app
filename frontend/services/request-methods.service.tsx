import { getenv } from "./util.service";
import cookie from 'js-cookie';
import { trackPromise } from "react-promise-tracker";

//This hook can be use in server also client side.
const useRequestMethodsService = (showErrorDialog?: ((msg: string) => void) | undefined,
    logout?: () => void) => {

    const post = async (url: string, body: any): Promise<Response | undefined> => {

        try {
            return await request(url, "POST", JSON.stringify(body));
        } catch (message) {
            showError(message);
        }
    }

    const get = async (url: string): Promise<Response | undefined> => {

        try {
            return await request(url, "GET", null);
        } catch (message) {
            showError(message);
        }
    }

    const del = async (url: string): Promise<Response | undefined> => {

        try {
            return await request(url, "DELETE", null);
        } catch (message) {
            showError(message);
        }
    }

    const upload = async (url: string, file: any): Promise<Response | undefined> => {
        try {
            let formData = new FormData();
            await formData.append('image', file);
            return await request(url, "POST", formData, "");

        } catch (error) {
            showError(error);
        }

    }

    const put = async (url: string, body: any): Promise<Response | undefined> => {
        try {
            return await request(url, "PUT", JSON.stringify(body));

        } catch (error) {
            showError(error);
        }
    }

    const patch = async (url: string, body: any): Promise<Response | undefined> => {
        try {
            return await request(url, "PATCH", JSON.stringify(body));

        } catch (error) {
            showError(error);
        }
    }

    const showError = (obj: any) => {

        if (obj && showErrorDialog) {

            let msg = "";
            if (obj.error) {
                msg = obj.error.message;
            } else {
                msg = obj;
            }

            showErrorDialog(msg);

            if (obj.error && logout) {
                if (obj.error.statusCode == 401 || obj.error.statusCode == 403) {
                    logout();
                }
            }
        }
    }

    const request = async (url: string,
        method: string,
        body: any = null,
        contentType: string = 'application/json'): Promise<Response> => {

        return new Promise<any>(async (resolveTrack, rejectTrack) => {

            trackPromise(new Promise(async (resolve, reject) => {

                let tokenCookie = cookie.get('token');
                let token = undefined;

                if (tokenCookie) {
                    token = JSON.parse(tokenCookie).JwToken;
                }

                let config: RequestInit = {
                    method: method,
                    headers: {},
                }

                if (token) {
                    config.headers = {
                        'Authorization': "Bearer " + token
                    }
                }

                if (contentType !== "") {
                    config.headers = {
                        'Content-Type': contentType
                    }
                }

                if (method !== "GET") {
                    config["body"] = body;
                }

                fetch(`${getenv('BASE_URL_SERVER')}${url}`, config)
                    .then((response) => {

                        if (response.status >= 400
                            && response.status != 412
                            && response.status != 451
                            && response.status != 404) {
                            //Retorna um json do erro para que o próximo "then"
                            //possa rejeitar a promise.
                            return response.json();
                        } else {
                            //Retornos abaixo de 400 são considerados "ok"
                            resolve(response);
                        }
                    })
                    .then((object: any) => {
                        reject(object);
                    }).catch((object) => {
                        //TODO: Send client errors to server
                        reject("Check your internet connection and try again.");
                    });

            })).then((ok) => {
                resolveTrack(ok);
            }).catch((err) => {
                rejectTrack(err);
            });
        });
    }
    return {
        del,
        post,
        get,
        put,
        patch,
        upload,
        showError
    }
}

export default useRequestMethodsService;