export async function fetchConfig() {
  const res = await window.$http.get('config')
  const authScopes = res.data?.auth_scope

  if (authScopes) {
    localStorage.vortexnotes_auth_scope = authScopes
  }

  return res.data
}
