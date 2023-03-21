import { pageLinks } from '@app/router/links';
import {
  Menu,
  PageSection,
  Split,
  SplitItem,
  Title,
  MenuContent,
  MenuItem,
  MenuList,
  Divider,
  Breadcrumb,
  BreadcrumbItem,
  Label,
  MenuGroup,
  Card,
  CardBody,
  CardTitle
} from '@patternfly/react-core';
import {
  CloudIcon,
  EditIcon,
  HomeIcon,
  LightbulbIcon,
  LinkIcon,
  OutlinedEnvelopeIcon,
  OutlinedUserIcon,
  UsersIcon
} from '@patternfly/react-icons';
import { css } from '@patternfly/react-styles';
import { Link } from 'react-router-dom';

import restApiLogo from '@app/assets/rest-logo.svg';
import { TableComposable, Tbody, Td, Th, Thead, Tr } from '@patternfly/react-table';
import { useState } from 'react';

const enum Mode {
  ServiceList = 'service-list',
  MemberList = 'members-list'
}

export const TeamDetailsPage = (): JSX.Element => {
  const [mode, setMode] = useState<Mode>(Mode.ServiceList);

  return (
    <PageSection isWidthLimited isCenterAligned variant="light" className="pr-12">
      <Split>
        <SplitItem className="w-56 mr-12 space-y-8">
          <div>
            <Menu isPlain>
              <MenuContent>
                <MenuList className="space-y-2 py-0">
                  <MenuItem
                    className={css('rounded-md menu', mode === Mode.ServiceList && 'menu-active')}
                    icon={<CloudIcon className="text-lg" />}
                    onClick={() => setMode(Mode.ServiceList)}
                  >
                    Services
                  </MenuItem>
                  <MenuItem
                    className={css('rounded-md menu', mode === Mode.MemberList && 'menu-active')}
                    icon={<UsersIcon className="text-xl" />}
                    onClick={() => setMode(Mode.MemberList)}
                  >
                    Team Members
                  </MenuItem>
                </MenuList>
                <Divider className="mt-4" />
                <MenuGroup label="Settings">
                  <MenuList>
                    <MenuItem className="rounded-md menu" icon={<EditIcon className="text-lg" />}>
                      Edit Team
                    </MenuItem>
                  </MenuList>
                </MenuGroup>
              </MenuContent>
            </Menu>
          </div>
        </SplitItem>
        <SplitItem isFilled>
          <div className="mb-4">
            <Breadcrumb>
              <BreadcrumbItem>
                <Link to={pageLinks.dashboard}>
                  <HomeIcon />
                </Link>
              </BreadcrumbItem>
              <BreadcrumbItem>
                <Link to={pageLinks.teamList}>Teams</Link>
              </BreadcrumbItem>
              <BreadcrumbItem isActive>One Platform</BreadcrumbItem>
            </Breadcrumb>
          </div>
          <div className="mb-4 flex items-center">
            <div className="flex-grow">
              <Title size="4xl" headingLevel="h1">
                One Platform
              </Title>
            </div>
            <div>
              <Label color="green" icon={<LightbulbIcon />}>
                Operational
              </Label>
            </div>
          </div>
          <div className="space-x-2 mb-6">
            <Label color="blue" icon={<OutlinedUserIcon />}>
              Nilesh Patil
            </Label>
            <Label color="blue" icon={<OutlinedEnvelopeIcon />}>
              one-platform-devs@redhat.com
            </Label>
          </div>
          <div className="mb-4">
            <Title size="lg" headingLevel="h3" className="mb-1">
              Description
            </Title>
            <p>Lorem ipsum dolor silicit</p>
          </div>
          <div>
            <Title size="lg" headingLevel="h3" className="mb-2">
              Links
            </Title>
            <div>
              <Label color="blue" variant="outline" icon={<LinkIcon />}>
                Source doc
              </Label>
            </div>
          </div>
          {mode === Mode.ServiceList && (
            <div>
              <div className="mb-4">
                <Divider className="my-8" />
                <Title size="2xl" headingLevel="h2" className="mb-2">
                  Services
                </Title>
              </div>
              <Card isSelectableRaised className="flex items-center flex-row pr-6 w-1/2">
                <div className="flex-grow">
                  <CardTitle>OP GraphQL API</CardTitle>
                  <CardBody className="flex items-center">
                    <div>Some description about the API</div>
                  </CardBody>
                </div>
                <img src={restApiLogo} alt="rest api" className="h-12 rounded-md" />
              </Card>
            </div>
          )}
          {mode === Mode.MemberList && (
            <div>
              <div className="mb-4">
                <Divider className="my-8" />
                <Title size="2xl" headingLevel="h2" className="mb-2">
                  Members
                </Title>
              </div>
              <TableComposable>
                <Thead>
                  <Tr>
                    <Th>Name</Th>
                    <Th>Email</Th>
                    <Th>Role</Th>
                    <Th>Status</Th>
                  </Tr>
                </Thead>
                <Tbody>
                  <Tr>
                    <Td>Akhil Mohan</Td>
                    <Td>akmohan@redhat.com</Td>
                    <Td>Admin</Td>
                    <Td>Accepted</Td>
                  </Tr>
                </Tbody>
              </TableComposable>
            </div>
          )}
        </SplitItem>
      </Split>
    </PageSection>
  );
};
