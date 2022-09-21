import React from 'react';
import ReactDOM from 'react-dom/client';
import {createHashRouter, RouterProvider} from 'react-router-dom';

import './translations/translations';

import Home from './pages/home';
import Game from './pages/game';

import './assets/scss/main.scss';
import { AuthProvider } from './hooks/auth';
import { DecksProvider } from './hooks/decks';

const router = createHashRouter([
  { path: '/game/:room_id', element: <Game/> },
  { path: '/', element: <Home/> },
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <AuthProvider>
      <DecksProvider>
        <RouterProvider router={router}/>
      </DecksProvider>
    </AuthProvider>
  </React.StrictMode>
)
