import {useCallback, useEffect} from 'react';
import {ListRenderItemInfo, Text, View} from 'react-native';
import {useQueryClient} from 'react-query';
import {InfinteList} from '../../../components/infinite-list/InfiniteList';
import {useDatatableFiltering} from '../../../hooks/useDatatableFiltering';
import {useReindexedContent} from '../../../hooks/useReindexedContent';
import {UserEntity} from '../../../sdk/core/react-tools';
import {useGetUsers} from '../../../sdk/modules/abac/useGetUsers';
import {UserItemCard} from './UserItemCard';
import {ListLayout} from '@/modules/fireback/components/list-layout/ListLayout';

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

  useEffect(() => {
    const rows: any = query.data?.data?.items || [];

    reindex(rows, '');
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [query.data?.data?.items]);

  return (
    <ListLayout title="Users">
      <InfinteList<UserEntity>
        keyExtractor={v => v.uniqueId}
        renderItem={renderItem}
        data={[]}
        query={query}
        udf={udf}
        items={indexedData}
      />
    </ListLayout>
  );
};
UserArchiveScreen.Name = 'UserArchiveScreen';
