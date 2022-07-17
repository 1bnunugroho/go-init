import { html } from "../shared/html.js";
export const Load404 = (page) => (state) => {
  const newState = {
    user: state.user,
    page,
  }
  return newState
};

export const P404P = () => html`
  <div class="container page">
    <h1>404.</h1>
    <p>Page not found.</p>
    <a href="/">Go back to home page</a>
  </div>
`;