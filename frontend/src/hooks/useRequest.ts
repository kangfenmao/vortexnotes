import useSWR, { SWRConfiguration, SWRResponse } from 'swr'
import { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'

export type RequestType = AxiosRequestConfig | null

interface Return<Data, Error>
  extends Pick<
    SWRResponse<AxiosResponse<Data>, AxiosError<Error>>,
    'isValidating' | 'isLoading' | 'error' | 'mutate'
  > {
  data: Data | undefined
  response: AxiosResponse<Data> | undefined
}

export interface SWRConfig<Data = unknown, Error = unknown>
  extends Omit<SWRConfiguration<AxiosResponse<Data>, AxiosError<Error>>, 'fallbackData'> {
  fallbackData?: Data
}

export default function useRequest<Data = unknown, Error = unknown>(
  request: RequestType,
  { fallbackData, ...config }: SWRConfig<Data, Error> = {}
): Return<Data, Error> {
  const {
    data: response,
    error,
    isValidating,
    isLoading,
    mutate
  } = useSWR<AxiosResponse<Data>, AxiosError<Error>>(
    request,
    /**
     * NOTE: Typescript thinks `request` can be `null` here, but the fetcher
     * function is actually only called by `useSWR` when it isn't.
     */
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    () => window.$http.request<Data>(request!),
    {
      revalidateOnFocus: Boolean(config.revalidateOnFocus),
      fallbackData:
        fallbackData &&
        ({
          status: 200,
          statusText: 'InitialData',
          // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
          config: request!,
          headers: {},
          data: fallbackData
        } as AxiosResponse<Data>),
      ...config
    }
  )

  return {
    data: response && response.data,
    response,
    error,
    isValidating,
    isLoading,
    mutate
  }
}
