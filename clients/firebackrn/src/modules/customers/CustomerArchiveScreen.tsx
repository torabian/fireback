import React, {useCallback, useEffect, useState} from 'react';
import {ListRenderItemInfo, ViewToken} from 'react-native';
import {InfinteList} from '../../components/infinite-list/InfiniteList';
import {ListLayout} from '../../components/list-layout/ListLayout';

// import {useActivityOverSocket2} from '@/hooks/useActivityOversocket';
import {
  useActivityOverSocket2,
  useActivityOverSocket3,
} from '@/hooks/useActivityOversocket';
import {useDatatableFiltering} from '@/hooks/useDatatableFiltering';
import {useDebouncedEffect} from '@/hooks/useDebouncedEffect';
import {useReindexedContent} from '@/hooks/useReindexedContent';
import {CustomerEntity} from '@/sdk/fireback/modules/demo/CustomerEntity';
import {useGetCustomers} from '@/sdk/fireback/modules/demo/useGetCustomers';
import {useQueryClient} from 'react-query';
import {CustomerCard} from './CustomerCard';
import {useReactiveCustomerActivity} from '@/sdk/fireback/modules/demo/useReactiveCustomerActivity';
import {UserActivityDto} from '@/sdk/fireback/modules/demo/UserActivityDto';

export function CustomerArchiveScreen() {
  const cq = useQueryClient();
  const {withDebounce} = useDebouncedEffect();
  const udf = useDatatableFiltering({});
  const [latestChange, setLatestChange] = useState<UserActivityDto>();

  const [ids, setIds] = useState<string[]>([]);

  // We code use ideally udf mechanism to implement n infinite scroll system
  const {query} = useGetCustomers({
    queryClient: cq,
    query: {deep: true, startIndex: udf.debouncedFilters.startIndex},
  });

  useEffect(() => {
    operate('');
  }, []);

  const {operate, write} = useReactiveCustomerActivity({
    onMessage(msg) {
      const content = JSON.parse(msg as any);

      // Running message parse and setting value on a seprate thread.
      // Seems to much for javascript to parse a json, and update a UI
      setTimeout(() => {
        setLatestChange(content);
      }, 100);
    },
  });

  // On real socket server
  useActivityOverSocket3(cq, latestChange);

  // This is a mock one I have prepared
  useActivityOverSocket2(cq, ids);

  // Reindexed content is a way to have infinite scrolls
  const {indexedData, reindex} = useReindexedContent(udf);

  const onViewableItemsChanged = (info: {
    viewableItems: ViewToken<CustomerEntity>[];
    changed: ViewToken<CustomerEntity>[];
  }) => {
    withDebounce(() => {
      const newIds = info.viewableItems.map(item => `${item.item.uniqueId}`);

      setIds(newIds);
      // Dont' forget
      write(JSON.stringify({ids: newIds}));
    }, 500);
  };

  useEffect(() => {
    const rows: any = query.data?.data?.items || [];

    reindex(rows, '');
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [query.data?.data?.items]);

  const renderItem = useCallback(
    (rProps: ListRenderItemInfo<CustomerEntity>) => (
      <CustomerCard entity={rProps.item} />
    ),
    [],
  );

  return (
    <ListLayout title="Customers">
      <InfinteList<CustomerEntity>
        data={[]}
        query={query}
        udf={udf}
        items={indexedData}
        renderItem={renderItem}
        onViewableItemsChanged={onViewableItemsChanged}
        keyExtractor={entity => `${entity?.uniqueId}`}
      />
    </ListLayout>
  );
}
