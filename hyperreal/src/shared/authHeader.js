export const authHeader = (token) => (token ? { Authorization: `Bearer ${token}` } : {});
