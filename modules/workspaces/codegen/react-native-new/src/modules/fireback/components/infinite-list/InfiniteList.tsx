import {DataFilteringResult} from '@/modules/fireback/hooks/useDatatableFiltering';
import {IResponseList} from '@/modules/fireback/sdk/core/http-tools';
import {
  FlatList,
  FlatListProps,
  ListRenderItem,
  RefreshControl,
  StyleSheet,
  ActivityIndicator,
  View,
} from 'react-native';
import {UseQueryResult} from 'react-query';

type InfiniteListProps<T> = {
  items: Array<T>;
  keyExtractor: (item: T, index: number) => string;
  renderItem: ListRenderItem<T>;
  query?: UseQueryResult<IResponseList<T>, any>;
  udf: DataFilteringResult;
} & FlatListProps<T>;

export function InfinteList<T = any>(props: InfiniteListProps<T>) {
  const isRefreshing = props.query?.isRefetching;

  const onRefresh = () => {
    props.query?.refetch();
  };

  return (
    <FlatList
      {...props}
      data={props.items}
      initialNumToRender={50}
      style={styles.wrapper}
      renderItem={props.renderItem}
      keyExtractor={props.keyExtractor}
      ListFooterComponent={
        <View style={styles.footer}>
          <ActivityIndicator
            size="large"
            color="#0000ff"
            style={{width: 5, height: 5, position: 'absolute'}}
          />
        </View>
      }
      onEndReachedThreshold={0.5}
      onEndReached={() => props.udf.increaseIndex(20)}
      refreshControl={
        <RefreshControl
          refreshing={isRefreshing || false}
          onRefresh={onRefresh}
          colors={['#9Bd35A', '#689F38']}
          progressBackgroundColor="#ffffff"
        />
      }
    />
  );
}

const styles = StyleSheet.create({
  footer: {
    alignContent: 'center',
    justifyContent: 'center',
    alignItems: 'center',
    margin: 30,
  },
  wrapper: {
    flex: 1,
  },
});
