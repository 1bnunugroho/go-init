import { html } from "../shared/html.js";
//import { Http } from "@kwasniew/hyperapp-fx";
import { Http } from "../lib/hyperapp-fx.js";
import { API_ROOT } from "../config.js";
import { authHeader } from "../shared/authHeader.js";
import { LogError } from "./fragments/forms.js";
import { Grid } from "./fragments/grid.js";

// Actions & Effects
export const loadingUsers = {
  users: [],
  usersCount: 0,
  isLoading: true,
};

export const loadingAlbums = {
  albums: [],
  albumsCount: 0,
  isLoading: true,
};

export const SetUsers = (state, { items, total_count }) => ({
  ...state,
  isLoading: false,
  users: items,
  usersCount: total_count,
});

export const SetAlbums = (state, { items, total_count }) => ({
  ...state,
  isLoading: false,
  albums: items,
  albumCount: total_count,
});

const isEmptyList = (list) => list.length === 0;

const NewUser = (state, props) =>
  isEmptyList(state.users)
    ? state
    : {
        ...state,
        users: state.users.concat({
          username: props.username,
          email: props.email,
          created_at: props.created_at ,
        }),
      };

const NewAlbum = (state, props) =>
  isEmptyList(state.albums)
    ? state
    : {
        ...state,
        albums: state.albums.concat({
          name: props.name,
          updated_at: props.updated_at,
          created_at: props.created_at ,
        }),
      };

// Views
export const FetchUsers = Http({
  url: API_ROOT + "/users",
  action: SetUsers,
  error: LogError,
  errorResponse: "json",
});

export const FetchAlbums = Http({
  url: API_ROOT + "/albums",
  action: SetAlbums,
  error: LogError,
  errorResponse: "json",
});

export const LoadUsersPage = (page) => (state) => {
  const user_columns = ["username", "email", "created_at", "updated_at"];
  const album_columns = ["name", "created_at", "updated_at"];
  const newState = {
    page,
    user_columns,
    album_columns,
    user: state.user,
    album: state.album,
    ...state.user,
    ...state.album,
    ...loadingUsers,
    ...loadingAlbums,
  };

  return [newState, [FetchUsers, FetchAlbums]];
};

// Views

const isLoggedIn = ({ user }) => user.token;
const isOwnProfile = ({ user, profile }) => user.username === profile.username;
/*
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
        ${props.datas.map((data) => html`<tr> 
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
        ${props.columns.map((column) => html`<td><input type="text" placeholder=${column}/></td>`)}
        <td colspan="3"><input type="submit" onclick=${[NewUser, {username:"jun", email:"jun@local.host", created_at:"now"}]} value="add"/></td>
        </tr>
      </tfoot>
      </table> ${children}
    </div>
  `;
};
*/

const ListPagination = ({ pages }) => {
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
                  onclick=${[preventDefault(ChangePage), eventWith({ currentPageIndex: page.index })]}
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

const pages = ({ count, currentPageIndex }) =>
  Array.from({ length: Math.ceil(count / 10) }).map((e, i) => ({
    index: i,
    isCurrent: i === currentPageIndex,
    humanDisplay: i + 1,
  }));

export const UsersPage = ({ page, user, profile, isLoading, users, usersCount, user_columns, album, albums, album_columns, albumsCount, currentPageIndex }) => html`
  <div class="profile-page">
    <div>
      <div class="container">
        <div class="row">
          <div class="col-xs-12 col-md-10 offset-md-1">
            <div class="article-preview">
              
            ${Grid( 
              {datas: users, columns: user_columns, isLoading: isLoading
                , onclick: [NewUser, {username:"jun", email:"jun@local.host", created_at:"now", update_at:"now"}]},
              ListPagination({
                pages: pages({ count: usersCount, currentPageIndex }),
              })
            )}

            <hr/>
            ${Grid( 
              {datas: albums, columns: album_columns, isLoading: isLoading
                , onclick: [NewAlbum, {name:"jun", created_at:"now", updated_at:"now"}]},
              ListPagination({
                pages: pages({ count: albumsCount, currentPageIndex }),
              })
            )}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
`;
