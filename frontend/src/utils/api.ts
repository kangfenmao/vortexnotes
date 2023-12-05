export async function fetchConfig() {
  const res = await window.$http.get('config')
  const authScopes = res.data?.auth_scope
  const authType = res.data?.auth_type

  if (authType) {
    localStorage.vortexnotes_auth_type = authType
  }

  if (authScopes) {
    localStorage.vortexnotes_auth_scope = authScopes
  }

  return res.data
}
