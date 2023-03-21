import {
  Avatar,
  Brand,
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

import logo from '@app/assets/logo-long.svg';

const AVATAR =
  'https://avataaars.io/?avatarStyle=Circle&topType=ShortHairDreads02&accessoriesType=Blank&hairColor=PastelPink&facialHairType=Blank&clotheType=ShirtCrewNeck&clotheColor=PastelYellow&eyeType=Wink&eyebrowType=RaisedExcited&mouthType=Twinkle&skinColor=DarkBrown';

export const Navbar = () => (
  <Masthead backgroundColor="light" className="border-b border-solid border-gray-200">
    <MastheadMain>
      <MastheadBrand className="flex items-center">
        <Brand src={logo} className="h-8" alt="API Catalog" />
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
                href="#"
                target="_blank"
                rel="noopener noreferrer"
                className="text-black font-bold"
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
                      <Avatar src={AVATAR} alt="Avatar image" size="lg" />
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
                <Avatar
                  src={AVATAR}
                  alt="Avatar image"
                  size="md"
                  className="cursor-pointer h-12 w-12"
                />
              </Popover>
            </ToolbarItem>
          </ToolbarGroup>
        </ToolbarContent>
      </Toolbar>
    </MastheadContent>
  </Masthead>
);
