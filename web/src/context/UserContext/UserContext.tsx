import { createContext, ReactNode } from 'react';
import { useGetUserInfo } from '@api/auth';

type Props = {
  children: ReactNode;
};

const UserContext = createContext({});

export const UserProvider = ({ children }: Props) => {
  const { data: user, isLoading } = useGetUserInfo();

  return <UserContext.Provider>{children}</UserContext.Provider>;
};
