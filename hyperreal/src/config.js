const me = typeof window === "undefined" ? global.location : window.location;
//export const API_ROOT = me.apiUrl || "https://conduit.productionready.io/api";
//export const API_ROOT = "http://localhost:8080/v1";
//export const API_ROOT = `${me.protocol}://${me.hostname}/v1`;
export const API_ROOT = "/v1";
