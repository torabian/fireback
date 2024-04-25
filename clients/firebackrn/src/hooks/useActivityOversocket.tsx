import {UserActivityDto} from '@/sdk/fireback/modules/demo/UserActivityDto';
import {useEffect} from 'react';
import {QueryClient} from 'react-query';

interface Subscription {
  unsubscribe: () => void;
}

interface Subscription {
  unsubscribe: () => void;
}

function intervalWithCancellation(
  callback: (value: any) => void,
  interval: number,
): Subscription {
  let aborted = false;

  callback(-1);

  const intervalId = setInterval(() => {
    if (!aborted) {
      callback(intervalId);
    } else {
      clearInterval(intervalId);
    }
  }, interval);

  return {
    unsubscribe: () => {
      aborted = true;
      clearInterval(intervalId);
    },
  };
}

export const useActivityOverSocket2 = (q: QueryClient, ids: string[]) => {
  useEffect(() => {
    const subscription = intervalWithCancellation(async () => {
      q.setQueriesData('activity', function () {
        return {
          data: {
            items: (ids || []).map(m => {
              return {
                uniqueId: m,
                status: Math.floor(Math.random() * 3) + 1,
              };
            }),
          },
        };
      });
    }, 1000);

    return () => {
      subscription.unsubscribe();
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [ids]);
};

export const useActivityOverSocket3 = (
  q: QueryClient,
  data?: UserActivityDto,
) => {
  useEffect(() => {
    setTimeout(() => {
      q.setQueriesData('activity', function () {
        const items = (data?.activities || []).map(m => {
          return {
            uniqueId: m.uniqueId,
            status: m.activity,
          };
        });

        return {
          data: {
            items,
          },
        };
      });
    }, 150);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data]);
};
