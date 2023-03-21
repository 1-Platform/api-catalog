import { Link } from 'react-router-dom';

import { Button, Card, CardBody, PageSection, Title } from '@patternfly/react-core';
import { ChevronLeftIcon } from '@patternfly/react-icons';

import welcomeImg from '@app/assets/hello-there.svg';

export const LoginPage = (): JSX.Element => (
  <div className="flex h-full">
    <PageSection className="w-[44%] bg-gray-200 flex flex-col justify-between">
      <div className="px-16">Logo</div>
      <div className="flex flex-col justify-center space-y-4">
        <img src={welcomeImg} alt="welcome back" />
      </div>
      <div className="px-16">
        <Link to="/">
          <Button icon={<ChevronLeftIcon />} variant="link" className="pl-0">
            Back
          </Button>
        </Link>
      </div>
    </PageSection>
    <PageSection
      className="w-[56%] flex flex-col justify-center items-center"
      padding={{ default: 'noPadding' }}
    >
      <div className="w-full max-w-xl flex flex-col">
        <div>
          <Title headingLevel="h1" size="4xl" className="mb-4">
            Login
          </Title>
          <p className="text-gray-500">Welcome back. Please login to your account</p>
          <div className="w-16 h-2 bg-blue-500 border-none rounded-md mt-4 mb-8" />
        </div>
        <div className="space-y-4 w-full border border-solid border-gray-200 p-2 rounded-sm">
          <Button className="w-full py-3">Sign In With Keycloak</Button>
        </div>
      </div>
    </PageSection>
  </div>
);
