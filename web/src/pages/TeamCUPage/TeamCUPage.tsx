import { Link } from 'react-router-dom';

import {
  ActionGroup,
  Button,
  Form,
  FormGroup,
  PageSection,
  Split,
  SplitItem,
  TextArea,
  TextInput,
  Title} from '@patternfly/react-core';
import { ChevronLeftIcon } from '@patternfly/react-icons';

import { pageLinks } from '@app/router/links';

export const TeamCUPage = (): JSX.Element => (
  <PageSection variant="light" isWidthLimited isCenterAligned>
    <Split className="space-x-16 flex justify-center">
      <SplitItem>
        <Link to={pageLinks.teamList}>
          <Button variant="link" icon={<ChevronLeftIcon />} className="bg-blue-100">
            Back
          </Button>
        </Link>
      </SplitItem>
      <SplitItem>
        <div className="mb-8">
          <Title headingLevel="h1" className="mb-2">
            Create new team
          </Title>
          <p className="text-gray-500">
            A team is a group of users, that handles the development and mainteince of a service
          </p>
        </div>
        <div>
          <Form>
            <FormGroup label="Team Name" isRequired>
              <TextInput />
            </FormGroup>
            <FormGroup label="Team Slug" helperText="A unique identifier for your team" isRequired>
              <TextInput />
            </FormGroup>
            <FormGroup label="Description">
              <TextArea rows={10} />
            </FormGroup>
            <ActionGroup>
              <Button>Submit</Button>
              <Button variant="link">Cancel</Button>
            </ActionGroup>
          </Form>
        </div>
      </SplitItem>
    </Split>
  </PageSection>
);
