import { RouterProvider } from 'react-router-dom';

import { router } from './router';

import './App.scss';

export const App = () => <RouterProvider router={router} />;
