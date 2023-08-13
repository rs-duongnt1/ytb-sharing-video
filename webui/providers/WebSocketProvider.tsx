import { createContext, useContext, useEffect, useRef, useState } from "react";
import React from "react";
import { Config } from "../config/config";
export const WebSocketContext = createContext(
  {} as {
    sendMessage: (event: string, message: string) => void;
    data: any;
  }
);

interface Props {
  children: React.ReactNode;
}

export const WebSocketProvider: React.FC<Props> = ({ children }) => {
  const socket = useRef<WebSocket>();
  const [data, setData] = useState<any>();
  const connect = () => {
    socket.current = new WebSocket(Config.websocketUrl);
    if (socket.current?.OPEN === 1) {
      socket.current.onmessage = (msg) => {
        try {
          const data = JSON.parse(msg.data);
          setData({
            error: false,
            data: data,
            message: null,
          });
        } catch (_) {
          setData({
            error: true,
            data: null,
            message: msg.data,
          });
        }
      };

      socket.current.onerror = (ev) => {
        console.log(ev);
      };

      socket.current.onclose = () => {
        const reconnectInterial = setInterval(() => {
          connect();
          if (socket.current?.OPEN === 1) {
            console.log("Reconnect successful");

            clearInterval(reconnectInterial);
          }
        }, 3000);
      };
    }
  };

  useEffect(() => {
    connect();
  }, [socket]);
  const sendMessage = (event: string, message: string) => {
    socket.current?.send(
      JSON.stringify({
        event,
        message,
        token: localStorage.getItem("t"),
      })
    );
  };

  return (
    <WebSocketContext.Provider value={{ sendMessage, data }}>
      {children}
    </WebSocketContext.Provider>
  );
};

export const useWebSocket = () => {
  const context = useContext(WebSocketContext);
  if (!context) {
    throw new Error("useCounter must be used within a CounterProvider");
  }
  return context;
};
