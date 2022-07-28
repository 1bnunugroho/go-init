import html from "hyperlit";
import { pages } from "./index.js";
import { HOME, LOGIN, NEW_EDITOR, REGISTER, SETTINGS, profile } from "./links.js";

// Views
const NavItem = ({ page, path }, children) => html`
  
    <a href="${path}" class=${{ "btn btn-link": true, active: page === path }}>
      ${children}
    </a>
  
`;

const UserImage = ({ user }) => html` <img src="${user.image}" class="user-pic" alt="${user.username}" /> `;

const UserLink = ({ user }) => html`
  ${user.image ? UserImage({ user }) : ""} <span data-test="profile">${user.username}</span>
`;

const P404Page = ({state}) => html`
  <div class="404 container page" key="404">
    <h1>404.</h1>
    <p>Page not found.</p>
    <a href="/">Go back to home page</a>
  </div>
`;

const Header = ({ page, user }) =>
  html`
    <header class="navbar container grid-lg">
    
      <section class="navbar-section">
        <a class="navbar-brand mr-2" href=${HOME}>
          Skill Hub
        </a>
      </section>
      <section class="navbar-section">
        ${NavItem({ page, path: HOME }, "Home")}
        ${user.token && NavItem({ page, path: NEW_EDITOR }, html` <i class="ion-compose" /> New Post `)}
        ${user.token ? 
          NavItem({ page, path: SETTINGS }, html` <i class="ion-gear-a" /> Settings `) 
          : NavItem({ page, path: LOGIN }, "Sign in")}
        ${user.token ? 
          NavItem({ page, path: profile(user.username) }, UserLink({ user })) 
          : NavItem({ page, path: REGISTER }, "Sign up")}
      </section>
    
    </header>
  `;

export const view = (state) => html`
  <div>
    ${Header({ page: state.page, user: state.user })} ${state.page ? pages[state.page](state) : ""}
  </div>
`;
