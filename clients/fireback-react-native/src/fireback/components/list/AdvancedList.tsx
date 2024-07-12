import React, {useEffect, useState} from 'react';

import {
  Alert,
  LayoutAnimation,
  RefreshControl,
  StyleSheet,
  TouchableOpacity,
  View,
} from 'react-native';
import {SwipeListView} from 'react-native-swipe-list-view';
import {useMutation, useQuery, useQueryClient} from 'react-query';
import DeleteIcon from '~/assets/icons/delete-icon.svg';
import {CommonFlatListEmptyComponent} from '~/components/list/CommonFlatListEmptyComponent';
import {useListFiltering} from '~/components/list/ListFilteringCommon';
import {openDialg} from '~/components/modal/Modal';
import {PageTitle} from '~/components/page-title/PageTitle';
import {TagsList} from '~/components/tags-list/TagsList';
import colors from '~/constants/colors';
import {IResponse, IResponseList} from '~/interfaces/JSONStyle';
import {AdvancedListInteractionPool} from '~/interfaces/Lists';
import {BookDto} from '~/modules/books/BookDto';

export interface AdvancedListProps {
  title?: string;
  description?: string;
  EditForm: any;
  EntityManager: any;
  RenderPlaceHolder: any;
  FilterForm: any;
  interactionPool?: AdvancedListInteractionPool<any>;
  RenderItem: any;
  filtersValidationSchema?: any;
  queryKey: string;
  keyExtractor: (t: any) => any;
}

export const AdvancedList = ({
  EditForm,
  EntityManager,
  title,
  description,
  RenderPlaceHolder,
  RenderItem,
  filtersValidationSchema,
  queryKey,
  interactionPool,
  FilterForm,
  keyExtractor,
}: AdvancedListProps) => {
  const [userRefreshed, setUserRefreshed] = useState<boolean>(false);
  const [selection, setSelection] = useState<string[]>([]);

  const {FilterBtn, filtersList, filters} = useListFiltering({
    Form: FilterForm,
    initialFilters: {},
    validationSchema: filtersValidationSchema,
  });

  const mutation = useMutation<any, unknown, string>(content => {
    return interactionPool?.remove(content);
  });

  const {data, isLoading, refetch, isRefetching} = useQuery(
    queryKey,
    () => interactionPool?.query(filters),
    {
      cacheTime: 1000,
    },
  );

  useEffect(() => {
    refetch();
  }, [filters]);

  const onLongPress = () => {
    Alert.alert('Long');
  };

  const queryClient = useQueryClient();
  const layoutAnimConfig = {
    duration: 300,
    update: {
      type: LayoutAnimation.Types.easeInEaseOut,
    },
    delete: {
      duration: 100,
      type: LayoutAnimation.Types.easeInEaseOut,
      property: LayoutAnimation.Properties.opacity,
    },
  };

  const onDelete = (key: string) => {
    mutation.mutate(key, {
      onSuccess(response: IResponse<{affectedRows: number}>) {
        LayoutAnimation.configureNext(layoutAnimConfig);
        queryClient.setQueryData<IResponseList<any>>(queryKey, data => {
          if (!data) {
            return {data: {}};
          }

          if (data?.data?.items && key) {
            data.data.items = data.data.items.filter(
              t => BookDto.getPrimaryKey(t) !== key,
            );
          }

          return data;
        });
      },

      onError(error: any) {
        // formikProps.setErrors(mutationErrorsToFormik(error));
      },
    });
  };

  const onRefresh = () => {
    setUserRefreshed(true);
    refetch();
  };

  useEffect(() => {
    if (isRefetching === false) {
      setUserRefreshed(false);
    }
  }, [isRefetching]);

  const onPress = (item: any) => {
    openDialg({
      title: item.title || '',
      data: item,
      Component: (props: any) => <EntityManager {...props} />,
    });
  };

  const showSkeleton = (isRefetching && !data) || isLoading;

  return (
    <>
      <PageTitle
        SideAction={FilterForm && FilterBtn}
        title={title}
        description={description}>
        <TagsList items={filtersList} setValue={() => {}} value={'bye'} />
      </PageTitle>

      <SwipeListView
        refreshControl={
          <RefreshControl refreshing={userRefreshed} onRefresh={onRefresh} />
        }
        contentContainerStyle={styles.contentContainerStyle}
        keyExtractor={keyExtractor}
        style={styles.flatList}
        ListEmptyComponent={<CommonFlatListEmptyComponent response={data} />}
        data={data?.data?.items}
        leftOpenValue={40}
        rightOpenValue={-75}
        renderHiddenItem={(data, rowMap) => (
          <TouchableOpacity
            onPress={() => onDelete(BookDto.getPrimaryKey(data.item))}
            style={styles.deleteButton}>
            <DeleteIcon />
          </TouchableOpacity>
        )}
        renderItem={
          showSkeleton
            ? RenderPlaceHolder
            : props => (
                <RenderItem
                  {...props}
                  onPress={onPress}
                  onLongPress={onLongPress}
                />
              )
        }
        ListFooterComponent={() => <View style={{paddingBottom: 80}} />}
      />
    </>
  );
};

const styles = StyleSheet.create({
  wrapper: {flex: 1, padding: 15},
  contentContainerStyle: {
    paddingBottom: 10,
  },
  flatList: {flex: 1, padding: 10, backgroundColor: colors.gray},
  row: {flexDirection: 'row', justifyContent: 'space-between'},
  label: {fontWeight: 'bold'},
  deleteButton: {
    top: 25,
    width: 30,
    height: 30,
    right: 0,
  },
});
