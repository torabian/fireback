import React, {useState, useEffect} from 'react';
import {Text} from 'react-native';
import io from 'socket.io-client';

const URL = 'http://localhost:4500';

const socket = io(URL, {
  autoConnect: true,
  query: {
    b64: 1,
    access_token:
      '7de54866efd3febfc2ff7e0b2470f1b2fc917877848a8a3955bc3a888c94e254',
  },
  transports: ['websocket'],
});

export function SocketConnection() {
  const [isConnected, setIsConnected] = useState(socket.connected);
  const [lastPong, setLastPong] = useState(null);

  useEffect(() => {
    socket.on('connect', () => {
      console.log('connected');
      setIsConnected(true);

      socket.emit('notice', 'asdasdasds');
    });

    socket.on('disconnect', () => {
      console.log('disconnect');
      console.log(2);
      setIsConnected(false);
    });

    socket.on('reply', data => {
      console.log(data);
    });

    socket.on('BOOK_EVENT_CREATED', data => {
      console.log(55, data);
    });

    socket.on('error', data => {
      if (data === 'Unauthorized access') {
        socket.close();
      }
      console.log(1, data);
    });

    // io.of('/orders').on('connection', socket => {
    //   console.log('inside');
    //   socket.on('order:list', () => {});
    //   socket.on('order:create', () => {});
    // });

    // socket.on('pong', () => {
    //   setLastPong(new Date().toISOString());
    // });

    return () => {
      socket.off('connect');
      socket.off('disconnect');
      socket.off('pong');
    };
  }, []);

  const sendPing = () => {
    socket.emit('ping');
  };

  return <Text>{isConnected ? ' Yes' : 'No'}</Text>;
}
