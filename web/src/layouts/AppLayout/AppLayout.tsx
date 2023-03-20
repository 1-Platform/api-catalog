import { Outlet } from 'react-router-dom';

import { Page } from '@patternfly/react-core';

import { Navbar } from './components/Navbar';

export const AppLayout = (): JSX.Element => (
  <Page>
    <Outlet />
  </Page>
);
