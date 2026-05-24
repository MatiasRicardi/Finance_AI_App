const BASE_URL = import.meta.env.VITE_API_BASE_URL ?? ''

interface ApiError {
  code: string
  message: string
  details?: Record<string, unknown>
}

export class ApiClientError extends Error {
  readonly code: string
  readonly details?: Record<string, unknown>

  constructor(apiError: ApiError) {
    super(apiError.message)
    this.name = 'ApiClientError'
    this.code = apiError.code
    this.details = apiError.details
  }
}

async function request<T>(path: string, init: RequestInit = {}): Promise<T> {
  const token = localStorage.getItem('auth_token')

  const headers: HeadersInit = {
    ...init.headers,
  }

  if (token) {
    (headers as Record<string, string>)['Authorization'] = `Bearer ${token}`
  }

  if (!(init.body instanceof FormData)) {
    (headers as Record<string, string>)['Content-Type'] = 'application/json'
  }

  const res = await fetch(`${BASE_URL}${path}`, { ...init, headers })

  if (!res.ok) {
    let apiError: ApiError = { code: 'unknown_error', message: 'An unexpected error occurred.' }
    try {
      const body = await res.json()
      if (body?.error) apiError = body.error
    } catch {
      // ignore parse error
    }
    throw new ApiClientError(apiError)
  }

  if (res.status === 204) {
    return undefined as unknown as T
  }

  return res.json() as Promise<T>
}

export const apiClient = {
  get: <T>(path: string) => request<T>(path, { method: 'GET' }),
  post: <T>(path: string, body: unknown) =>
    request<T>(path, { method: 'POST', body: JSON.stringify(body) }),
  put: <T>(path: string, body: unknown) =>
    request<T>(path, { method: 'PUT', body: JSON.stringify(body) }),
  patch: <T>(path: string, body: unknown) =>
    request<T>(path, { method: 'PATCH', body: JSON.stringify(body) }),
  delete: <T>(path: string) => request<T>(path, { method: 'DELETE' }),
  upload: <T>(path: string, formData: FormData) =>
    request<T>(path, { method: 'POST', body: formData }),
}
