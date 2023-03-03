import { ReactNode } from 'react';

import { Nav, NavItem, NavList, PageSidebar, Split, SplitItem } from '@patternfly/react-core';
import { HomeIcon, TableIcon, UserIcon } from '@patternfly/react-icons';

type SNProps = {
  title: string;
  icon: ReactNode;
  isActive?: boolean;
};

const SidebarNavItem = ({ title, icon, isActive }: SNProps) => (
  <NavItem isActive={isActive}>
    <Split hasGutter className="pf-u-p-md cursor-pointer">
      <SplitItem>{icon}</SplitItem>
      <SplitItem>{title}</SplitItem>
    </Split>
  </NavItem>
);

const SidebarNav = () => (
  <Nav>
    <NavList>
      <SidebarNavItem title="Home" icon={<HomeIcon />} />
      <SidebarNavItem title="Explore" icon={<TableIcon />} />
      <SidebarNavItem title="My APIs" icon={<UserIcon />} />
    </NavList>
  </Nav>
);

export const Sidebar = () => <PageSidebar nav={<SidebarNav />} />;
