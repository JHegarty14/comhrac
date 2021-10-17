import api from "../api";

describe("test setCsrfToken", () => {
    test("setCsrfToken from document with only csrf token sets token correctly", () => {
        const token = "this is a token";
        const cookie = `csrftoken=${token}`;
        api.setCsrfToken(cookie);
        expect(api.csrftoken).toEqual(token);
    });

    test("setCsrfToken from document with empty string as cookie sets an empty string", () => {
        const cookie = "";
        api.setCsrfToken(cookie);
        expect(api.csrftoken).toEqual("");
    });

    test("test secCsrfToken from document without csrftoken sets an empty string", () => {
        api.setCsrfToken("doesnothavecorrectname");
        expect(api.csrftoken).toEqual("");
    });

    test("setCsrfToken from document with csrf token at beginning of multiple items", () => {
        const token = "this is the middle token";
        const cookie = `csrftoken=${token};path=somepath;extra=somethingelse`;
        api.setCsrfToken(cookie);
        expect(api.csrftoken).toEqual(token);
    });

    test("setCsrfToken from document with csrf token in the middle of multiple items", () => {
        const token = "this is the middle token";
        const cookie = `path=somepath;csrftoken=${token};extra=somethingelse`;
        api.setCsrfToken(cookie);
        expect(api.csrftoken).toEqual(token);
    });

    test("setCsrfToken from document with csrf token at the end of multiple items", () => {
        const token = "this is the middle token";
        const cookie = `path=somepath;extra=somethingelse;csrftoken=${token}`;
        api.setCsrfToken(cookie);
        expect(api.csrftoken).toEqual(token);
    });
});
