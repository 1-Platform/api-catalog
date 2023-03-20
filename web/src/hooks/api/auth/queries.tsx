import { apiRequest } from '@config/apiRequest';
import { useQuery } from '@tanstack/react-query';

export const authQueryKeys = {
  userInfo: ['user-info'] as const,
  userAction: ['user-action'] as const
};

const fetchUserInfo = async () => {
  const { data } = await apiRequest.get('/user/info');
  return data;
};

export const useGetUserInfo = () =>
  useQuery({
    queryKey: authQueryKeys.userInfo,
    queryFn: fetchUserInfo
  });

const fetchUserAction = async (action: string) => {
  const { data } = await apiRequest.get<{ userAction: string }>('/api/v1/user-action', {
    params: {
      action
    }
  });
  return data;
};

export const useGetUserAction = (action: string) =>
  useQuery({
    queryKey: authQueryKeys.userAction,
    queryFn: () => fetchUserAction(action)
  });
