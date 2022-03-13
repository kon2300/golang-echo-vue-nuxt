import type { IncomingMessage, ServerResponse } from 'http'

export default async (req: IncomingMessage, res: ServerResponse) => {
  const result = await $fetch('http://backend:8080/ping')

  return result
}
