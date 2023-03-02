import { Outlet } from 'react-router-dom';

import { Page } from '@patternfly/react-core';

export const AppLayout = (): JSX.Element => (
  <Page>
    <Outlet />
  </Page>
);
