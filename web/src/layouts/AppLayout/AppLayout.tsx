import { Outlet } from 'react-router-dom';

import { Page } from '@patternfly/react-core';

import { Navbar } from './components/Navbar';
import { Sidebar } from './components/Sidebar';

export const AppLayout = (): JSX.Element => (
  <Page header={<Navbar />} sidebar={<Sidebar />}>
    <Outlet />
  </Page>
);
