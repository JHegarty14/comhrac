import APIResponseException from "../exceptions/APIResponseException";
import Config from 'react-native-config'

class BaseAPI {
    csrftoken = "";

    apiUrl = Config.API_URL;

    apiVersion = Config.API_VERSION;

    safeMethods = ["GET", "HEAD", "OPTIONS"];

    setCsrfToken(cookie: string) {
        const hasToken = cookie && cookie.indexOf("csrftoken=") >= 0;
        if (hasToken) {
            const cookieArr = cookie.split(";");
            const csrfToken = cookieArr.find(item => item.indexOf("csrftoken=") >= 0);
            if (csrfToken) this.csrftoken = csrfToken.replace("csrftoken=", "");
        }
        else {
            this.csrftoken = "";
        }
    }

    requestInit(requestType: string = "GET", body: object = {}, cors: boolean = true): object {
        const init: {
            method: string,
            headers: {
                "Content-Type": string,
                "X-CSRFTOKEN"?: string,
            }
            mode?: string,
            credentials?: string,
            body?: string,
        } = {
            method: requestType,
            headers: { "Content-Type": "application/json" },
        };

        if (cors) {
            init.mode = "cors";
            init.credentials = "include";
        }

        if (Object.keys(body).length > 0) {
            init.body = JSON.stringify(body);
        }

        if (!this.safeMethods.includes(requestType)) {
            init.headers["X-CSRFTOKEN"] = this.csrftoken;
        }

        return init;
    }

    async fetchData(input: string, init = this.requestInit()) {
        const response = await fetch(input, init);
        // if no content in response
        if (response.status === 204) {
            return null;
        }
        const json = await response.json();

        this.setCsrfToken(document.cookie);

        if (response.ok) {
            return json;
        }
        throw new APIResponseException(response, json);
    }

    getUri(location: string): string {
        return `${this.apiUrl}/v${this.apiVersion}/${location}`;
    }
}

export default BaseAPI;
