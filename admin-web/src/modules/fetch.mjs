/**
* @return {Promise<Response>}
*/
export function fetchFromApi(url, options) {
  if (!options) { options = {} }
  if (!options.headers) { options.headers = {} }
  if (!options.headers["Accept-Language"]) {
    var lang = localStorage.getItem('lang'); 
    if(lang == null){
      lang = "hu"
    }
    options.headers["Accept-Language"] = lang
  }
  if (!options.headers["AuthToken"]) {
    options.headers["AuthToken"] = import.meta.env.VITE_API_AUTH_TOKEN
  }
  return fetch(url, options);
}