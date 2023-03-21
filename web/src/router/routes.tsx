import { createBrowserRouter } from 'react-router-dom';

import { AppLayout } from '../layouts';
import { HomePage } from '../pages/HomePage';
import { LoginPage } from '../pages/LoginPage';
import { ServiceListPage } from '../pages/ServiceListPage';
import { TeamCUPage } from '../pages/TeamCUPage';
import { TeamDetailsPage } from '../pages/TeamDetailsPage';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <AppLayout />,
    children: [
      {
        index: true,
        element: <HomePage />
      },
      {
        path: '/login',
        element: <LoginPage />
      },
      {
        path: '/services',
        element: <ServiceListPage />
      },
      {
        path: '/teams/new',
        element: <TeamCUPage />
      },
      {
        path: '/teams/:id',
        element: <TeamDetailsPage />
      }
    ]
  }
]);
