import {Formik, FormikProps} from 'formik';
import React, {useState} from 'react';
import {ScrollView, StyleSheet, TouchableOpacity, View} from 'react-native';
import {KeyboardAwareScrollView} from 'react-native-keyboard-aware-scroll-view';
import FilterIcon from '~/assets/icons/filter-icon.svg';
import colors from '~/constants/colors';
import t from '~/constants/t';
import {ListItem} from '~/interfaces/UI';
import {FormButton} from '../form-button/FormButton';
import {openDialg} from '../modal/Modal';

export const FilteringFooter = ({form}: {form: FormikProps<any>}) => {
  return (
    <View style={styles.modalFooter}>
      <FormButton label={'Apply'} onPress={() => form.submitForm()} />
      <FormButton
        label={'Reset'}
        type="secondary"
        onPress={() => form.resetForm()}
      />
    </View>
  );
};

function asArray(data: any) {
  const keys = Object.keys(data);
  const result: ListItem[] = [];

  for (const key of keys) {
    const label = `${key}: ${data[key]}`;
    const value = data[key];
    if (!value) {
      continue;
    }
    result.push({
      label,
      value,
    });
  }

  return result;
}

export function useListFiltering<T>({
  initialFilters,
  Form,
  beforeSubmit,
  validationSchema,
}: {
  initialFilters: T;
  Form: any;
  beforeSubmit?: (t: T) => T;
  validationSchema?: any;
}) {
  const [filters, setFilters] = useState<T>(initialFilters);

  const FilterModal = ({fnClose}: {fnClose: () => void}) => (
    <KeyboardAwareScrollView>
      <Formik
        initialValues={filters}
        validationSchema={validationSchema}
        onSubmit={data => {
          let dataNew: any = {...data};
          if (beforeSubmit) {
            dataNew = beforeSubmit(dataNew);
          }
          setFilters(dataNew);
          fnClose();
        }}
        onReset={() => {
          setFilters(initialFilters);
          fnClose();
        }}>
        {(form: FormikProps<T>) => {
          return (
            <>
              <Form form={form} />
              <FilteringFooter form={form} />
            </>
          );
        }}
      </Formik>
    </KeyboardAwareScrollView>
  );

  const onPress = () => {
    openDialg({
      title: 'Filters',
      Component: (props: any) => <FilterModal {...props} />,
    });
  };

  const FilterBtn = ({testID}: {testID?: string}) => (
    <TouchableOpacity onPress={onPress}>
      <FilterIcon width={40} testID={testID} />
    </TouchableOpacity>
  );

  const filtersList = asArray(filters);

  return {
    FilterBtn,
    filters,
    filtersList,
    setFilters,
  };
}

export const FilterView = ({
  form,
  children,
  title,
}: {
  title: string;
  form: FormikProps<any>;
  children: any;
}) => {
  return (
    <View style={styles.wrapper}>
      <ScrollView style={{paddingHorizontal: 20, flex: 1}}>
        {children}
      </ScrollView>

      <FilteringFooter form={form} />
    </View>
  );
};

const styles = StyleSheet.create({
  modalFooter: {
    // flexDirection: 'row',
    minHeight: 100,
    paddingVertical: 10,
  },
  wrapper: {flex: 1},
});
