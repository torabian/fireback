import {useCallback} from 'react';
import {ListRenderItemInfo} from 'react-native';
import {useQueryClient} from 'react-query';
import {InfinteList} from '../../../components/infinite-list/InfiniteList';
import {useDatatableFiltering} from '../../../hooks/useDatatableFiltering';
import {useReindexedContent} from '../../../hooks/useReindexedContent';
import {UserEntity} from '../../../sdk/core/react-tools';
import {useGetUsers} from '../../../sdk/modules/workspaces/useGetUsers';
import {UserItemCard} from './UserItemCard';

export const UserArchiveScreen = () => {
  const queryClient = useQueryClient();

  // udf is a mechanism of handling filters, search, pagination
  // of an array over http.
  const udf = useDatatableFiltering({});

  // reindexed content is a way to make it infinit scroll easier,
  // it would accept udf object and append the items
  const {indexedData, reindex} = useReindexedContent(udf);

  const renderItem = useCallback(
    (rProps: ListRenderItemInfo<UserEntity>) => (
      <UserItemCard entity={rProps.item} />
    ),
    [],
  );

  const {query} = useGetUsers({
    queryClient,
    query: {deep: true, startIndex: udf.debouncedFilters.startIndex},
  });

  return (
    <InfinteList<UserEntity>
      keyExtractor={v => v.uniqueId}
      renderItem={renderItem}
      data={[]}
      query={query}
      udf={udf}
      items={indexedData}
    />
  );
};
UserArchiveScreen.Name = 'UserArchiveScreen';
