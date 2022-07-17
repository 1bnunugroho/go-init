import { html } from "../shared/html.js";
//import { Http } from "@kwasniew/hyperapp-fx";
import { Http } from "../lib/hyperapp-fx.js";
import { API_ROOT } from "../config.js";
import { authHeader } from "../shared/authHeader.js";
import { LogError } from "./fragments/forms.js";
import { Grid, ListPagination, ChangePage, pages } from "./fragments/grid.js";
import { preventDefault } from "../lib/events.js";
import { eventWith } from "../lib/events.js";
//import { ArticleList, loadingArticles } from "./fragments/articles.js";
//import { FetchArticles } from "./fragments/articles.js";
//import { profile, profileFavorited, SETTINGS } from "./links.js";

// Actions & Effects
export const loadingAlbums = {
  albums: [],
  albumsCount: 0,
  maxItem: 5,
  isLoading: true,
  album_columns: ['name', 'updated_at', 'created_at'],
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
      albumsCount: state.albumsCount + 1
    };

const isEmptyList = (list) => list.length === 0;

export const SetAlbums = (state, { items, total_count }) => ({
  ...state,
  isLoading: false,
  albums: items,
  albumsCount: total_count,
});
//export const SetAlbums = (state, { items }) => ({ ...state, albums:items });

// Views
export const FetchAlbums = Http({
  url: API_ROOT + "/albums",
  action: SetAlbums,
  error: LogError,
  errorResponse: "json",
});

export const LoadAlbumsPage = (page) => (state) => {
  const newState = {
    name: '',
    page,
    user: state.user,
    currentPageIndex: 0,
    ...state.user,
    ...loadingAlbums,
  };

  return [newState, [FetchAlbums]];
};

const NewValue = (state, ev) => ({...state, name: event.target.value});

// Views

const isLoggedIn = ({ user }) => user.token;
const isOwnProfile = ({ user, profile }) => user.username === profile.username;

export const AlbumList = ({ isLoading, albums }, children) => {
  if (isLoading) {
    return html` <div class="article-preview">Loading...</div> `;
  }
  if (albums.length === 0) {
    return html` <div class="article-preview">No albums are here... yet.</div> `;
  }
  return html`
    <div>
      ${albums.map((album) => html`<li> ${album.name}</li>`)} ${children}
    </div>
  `;
};

export const AlbumsPage = ({ page, user, name, profile, isLoading, albums, albumsCount, album_columns, currentPageIndex, maxItem }) => html`
  <div class="profile-page">
    <div>
      <div class="container">
        <div class="row">
          <div class="col-xs-12 col-md-10 offset-md-1">
          <h1>Albums</h1>
            <div class="article-preview">
              ${Grid( 
                {datas: albums, columns: album_columns, isLoading: isLoading, maxItem: maxItem
                  , oninput: [NewValue, name]
                  , onclick: [NewAlbum, {name, created_at:"now", updated_at:"now"}]},
                ListPagination({
                  pages: pages({ count: albumsCount, currentPageIndex, maxItem }),
                })
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
`;
