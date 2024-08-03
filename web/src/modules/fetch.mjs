/**
* @return {Promise<Response>}
*/
export function fetchFromApi(url, options) {
    if (!options) { options = {} }
    if (!options.headers) { options.headers = {} }
    if (!options.headers["Accept-Language"]) {
      options.headers["Accept-Language"] = "en"
    }
    if (!options.headers["AuthToken"]) {
      options.headers["AuthToken"] = import.meta.env.VITE_API_AUTH_TOKEN
    }
    return fetch(url, options);
}