import {useLocale} from '@/modules/fireback/hooks/useLocale';
import {FormikProps} from 'formik';
import {useEffect, useRef} from 'react';
import {useQueryClient} from 'react-query';

interface DtoEntity<T> {
  data?: T | null;
}

/**
 * Set of hooks we might need for entity manager screens (update/create)
 */
export function useCommonEntityManager<T>(props?: DtoEntity<T> | undefined) {
  const formik = useRef<FormikProps<T>>();
  const queryClient = useQueryClient();
  useEffect(() => {
    if (props?.data) {
      formik.current?.setValues(props.data);
    }
  }, [props?.data]);

  //   const router = useRouter();
  //   const uniqueId = router.query.uniqueId as string;
  //   const linkerId = router.query.linkerId as string;
  //   const isEditing = !!uniqueId;
  const {locale} = useLocale();

  return {
    // router,
    isEditing: false,
    locale,
    queryClient,
    formik,
    // uniqueId,
    // linkerId,
  };
}
