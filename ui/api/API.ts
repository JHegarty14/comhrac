import BaseAPI from "./BaseAPI";

class API extends BaseAPI {
    app = {
        getConstants: (): Promise<any> => (
            this.fetchData(this.getUri("constants/"))
        ),
    };

    user = {
        auth: (): Promise<any> => (
            this.fetchData(this.getUri("users/auth/"))
        ),
        logout: (): Promise<any> => (
            this.fetchData(this.getUri("users/auth/"), this.requestInit("DELETE"))
        ),
    };
}

export default API;
