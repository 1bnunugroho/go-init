//import { Http } from "@kwasniew/hyperapp-fx";
import { Http } from "../../lib/hyperapp-fx.js";
import { API_ROOT } from "../../config.js";
import { API_ROOT2 } from "../../config.js";
import { authHeader } from "../../shared/authHeader.js";
import { LogError } from "./forms.js";

// Actions & Effects
const SetArticle = (state, { article }) => ({ ...state, ...article });

// Views
export const FetchArticle = ({ slug, token }) =>
  Http({
    url: API_ROOT + "/articles/" + slug,
    options: { headers: authHeader(token) },
    action: SetArticle,
    error: LogError,
  });
