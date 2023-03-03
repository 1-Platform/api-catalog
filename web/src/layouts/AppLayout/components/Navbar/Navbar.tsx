import {
  Avatar,
  Button,
  ButtonVariant,
  Flex,
  FlexItem,
  Masthead,
  MastheadBrand,
  MastheadContent,
  MastheadMain,
  Popover,
  Text,
  Title,
  TitleSizes,
  Toolbar,
  ToolbarContent,
  ToolbarGroup,
  ToolbarItem
} from '@patternfly/react-core';
import imgAvatar from '@patternfly/react-core/src/components/Avatar/examples/avatarImg.svg';

export const Navbar = () => (
  <Masthead>
    <MastheadMain>
      <MastheadBrand>
        <Title headingLevel="h6" className="font-mono">
          API CATALOG
        </Title>
      </MastheadBrand>
    </MastheadMain>
    <MastheadContent>
      <Toolbar isFullHeight isStatic>
        <ToolbarContent>
          <ToolbarGroup
            variant="icon-button-group"
            alignment={{ default: 'alignRight' }}
            spacer={{ default: 'spacerNone', md: 'spacerMd' }}
            spaceItems={{ default: 'spaceItemsSm' }}
          >
            <ToolbarItem>
              <Button
                component="a"
                aria-label="DOC URL"
                variant={ButtonVariant.link}
                style={{ color: '#fff' }}
                href="#"
                target="_blank"
                rel="noopener noreferrer"
              >
                Go to Docs
              </Button>
            </ToolbarItem>
            <ToolbarItem>
              <Popover
                flipBehavior={['bottom-end']}
                hasAutoWidth
                showClose={false}
                bodyContent={
                  <Flex
                    style={{ width: '200px' }}
                    direction={{ default: 'column' }}
                    alignItems={{ default: 'alignItemsCenter' }}
                    spaceItems={{ default: 'spaceItemsSm' }}
                  >
                    <FlexItem>
                      <Avatar src={imgAvatar} alt="Avatar image" size="lg" />
                    </FlexItem>
                    <FlexItem>
                      <Title headingLevel="h6" size={TitleSizes.lg}>
                        Akhil Mohan
                      </Title>
                    </FlexItem>
                    <FlexItem>
                      <Text component="small">akmohan@redhat.com</Text>
                    </FlexItem>
                    <FlexItem className="pf-u-w-100">
                      <Button isBlock>Logout</Button>
                    </FlexItem>
                  </Flex>
                }
              >
                <Avatar src={imgAvatar} alt="Avatar image" size="md" className="cursor-pointer" />
              </Popover>
            </ToolbarItem>
          </ToolbarGroup>
        </ToolbarContent>
      </Toolbar>
    </MastheadContent>
  </Masthead>
);
