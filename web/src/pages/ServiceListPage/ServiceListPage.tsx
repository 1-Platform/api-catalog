import {
  Button,
  Divider,
  Dropdown,
  DropdownItem,
  DropdownToggle,
  Label,
  Menu,
  MenuContent,
  MenuItem,
  MenuList,
  PageSection,
  SearchInput,
  Split,
  SplitItem,
  Stack,
  StackItem,
  Title
} from '@patternfly/react-core';
import { CubeIcon, UsersIcon } from '@patternfly/react-icons';
import { css } from '@patternfly/react-styles';

export const ServiceListPage = (): JSX.Element => (
  <>
    <PageSection isWidthLimited isCenterAligned variant="light">
      Something top
    </PageSection>
    <PageSection isWidthLimited isCenterAligned variant="light">
      <Split>
        <SplitItem className="w-56 mr-12 space-y-8">
          <div>
            <Menu isPlain>
              <MenuContent>
                <MenuList className="space-y-2 py-0">
                  <MenuItem
                    className={css('rounded-md menu', 'menu-active')}
                    icon={<CubeIcon className="text-lg" />}
                  >
                    Services
                  </MenuItem>
                  <MenuItem className="rounded-md menu" icon={<UsersIcon className="text-xl" />}>
                    Teams
                  </MenuItem>
                </MenuList>
              </MenuContent>
            </Menu>
          </div>
          <Divider />
          <div>
            <div className="font-bold text-sm mb-4">Related Tags</div>
            <div className="flex gap-2 flex-wrap">
              {[...Array(10)].map((_, i) => (
                <Label key={`tags-${i + 1}`} variant="outline">
                  hydra{i + 1}
                </Label>
              ))}
            </div>
          </div>
        </SplitItem>
        <SplitItem isFilled>
          <Stack hasGutter>
            <StackItem>
              <Split hasGutter>
                <SplitItem isFilled>
                  <SearchInput className="max-w-lg" placeholder="Search for a service by name" />
                </SplitItem>
                <Dropdown
                  isPlain
                  className="mr-0"
                  toggle={<DropdownToggle className="pr-0">Recent</DropdownToggle>}
                  dropdownItems={[<DropdownItem key="order-by">Created At</DropdownItem>]}
                />
                <SplitItem />
              </Split>
            </StackItem>
            <StackItem>
              {[...Array(10)].map((_, i) => (
                <div key={i + 1}>
                  {i !== 0 && <Divider />}
                  <div className="flex justify-center items-center h-24 space-x-4">
                    <div>
                      <div className="w-16 h-16 bg-primary rounded-md" />
                    </div>
                    <div className="flex-grow">
                      <Title headingLevel="h6" className="text-base mb-1">
                        One Platform GraphQL API
                      </Title>
                      <p className="text-sm">The one platform Graphql API</p>
                    </div>
                    <div>
                      <div className="w-8 h-8 bg-blue-400 rounded-md" />
                    </div>
                    <div className="w-44 text-center">One Platform</div>
                    <div className="w-40">Updated 8 days ago</div>
                    <div>
                      <Button className="rounded">Explore</Button>
                    </div>
                  </div>
                </div>
              ))}
            </StackItem>
          </Stack>
        </SplitItem>
      </Split>
    </PageSection>
  </>
);
