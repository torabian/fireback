import {UserEntity} from '@/modules/fireback/sdk/modules/abac/UserEntity';
import {useGetUserByUniqueId} from '../../../sdk/modules/abac/useGetUserByUniqueId';
import {useS} from '@/modules/fireback/hooks/useS';
import {strings} from './strings/translations';

export const UserSingleScreen = () => {
  const {uniqueId, queryClient} = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetUserByUniqueId({query: {uniqueId}});
  var d: UserEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);

  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({locale, router}) => {
          router.push(UserEntity.Navigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}></CommonSingleManager>
    </>
  );
};
