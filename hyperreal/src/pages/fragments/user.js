//import { ReadFromStorage, RemoveFromStorage, WriteToStorage } from "@kwasniew/hyperapp-fx";
import { ReadFromStorage, RemoveFromStorage, WriteToStorage } from "../../lib/hyperapp-fx.js";
import { Redirect } from "../../lib/router.js";
import { HOME } from "../links.js";

const SESSION = "session";

// Actions & Effects
const parseJwt = (token) => {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

const genUser = (token) => {
  const payload = parseJwt(token);
  //const user = {};
  const user = parseJwt(token);
  user.token = token;
  return user;
}


const SetUser = (state, { value }) => (value ? { ...state, user: value } : state);

const SaveUser = (user) => WriteToStorage({ key: SESSION, value: user });

export const ReadUser = ReadFromStorage({ key: SESSION, action: SetUser });

export const UserSuccess = (state, { token }) => [{ ...state, user: genUser(token) }, [SaveUser(genUser(token)), Redirect({ path: HOME })]];

export const Logout = (state) => [
  { ...state, user: {} },
  [RemoveFromStorage({ key: SESSION }), Redirect({ path: HOME })],
];
