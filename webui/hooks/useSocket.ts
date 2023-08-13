export const useSocket = (url: string) => {
    const socket = new WebSocket(url)

    return {
        socket,
    }
}