import { html } from "../../shared/html.js";
import { loadingAlbums } from "./articles.js";
import { Http } from "../../lib/hyperapp-fx.js";
import { API_ROOT } from "../../config.js";
import { LogError } from "./forms.js";

export const SetAlbums = (state, { items }) => ({ ...state, albums:items });

export const FetchAlbums = Http({
  url: API_ROOT + "/albums",
  action: SetAlbums,
  error: LogError,
  errorResponse: "json",
});

export const ListPagination = ({ pages }) => {
  if (pages.length < 2) {
    return "";
  }
  return html`
    <nav>
      <ul class="pagination">
        ${pages.map(
          (page) =>
            html`
              <li class=${page.isCurrent ? "page-item active" : "page-item"}>
                <a
                  class="page-link"
                  href=""
                  onclick=${(_, event) => {
                      event.preventDefault() 
                      return [ChangePage, { currentPageIndex: page.index }]
                    }
                  }
                >
                  ${page.humanDisplay}
                </a>
              </li>
            `
        )}
      </ul>
    </nav>
  `;
};

export const pages = ({ count, currentPageIndex, maxItem }) =>
  Array.from({ length: Math.ceil(count / maxItem) }).map((e, i) => ({
    index: i,
    isCurrent: i === currentPageIndex,
    humanDisplay: i + 1,
  }));

export const ChangePage = (state, { currentPageIndex }) => {
  const newState = {
    ...state,
    ...loadingAlbums,
    currentPageIndex,
  };

  return [newState, [FetchAlbums(newState)]];
};

export const Grid = (props, children) => {
  
  if (props.isLoading) {
    return html` <div class="article-preview">Loading...</div> `;
  }
  if (props.datas.length === 0) {
    return html` <div class="article-preview">No datas are here... yet.</div> `;
  }

  return html`
    <div><table class="table"><thead><tr>
        ${props.columns.map((column) => html`<th>${column}</th>`)}
        <th colspan="3">action</th>
        </tr></thead>
      <tbody>

        ${props.datas.slice(-(props.maxItem)).map((data) => html`<tr> 
          ${props.columns.map((column) => html`<td>${data[column]}</td>`)}
            <td><a href="#${data["id"]}" 
            onclick="delete ${data["id"]}"
            >delete</a></td>
            <td><a href="#${data["id"]}" 
            onclick="update ${data["id"]}"
            >update</a></td>
            <td><a href="#${data["id"]}" 
            onclick="view ${data["id"]}"
            >view</a></td>
            </tr>`)}
      </tbody>
      <tfoot>
        <tr>
        ${props.columns.map((column) => html`<td><input type="text" placeholder=${column} oninput=${props.oninput}/></td>`)}
        <td colspan="3"><input type="submit" onclick=${props.onclick} value="add"/></td>
        </tr>
      </tfoot>
      </table> ${children}
    </div>
  `;
};