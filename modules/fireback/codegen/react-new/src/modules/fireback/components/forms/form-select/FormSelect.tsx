import { JsonQuery } from "@/modules/fireback/definitions/definitions";
import { IResponseList } from "@/modules/fireback/sdk/core/http-tools";
import { UseRemoteQuery } from "@/modules/fireback/sdk/core/react-tools";
import classNames from "classnames";
import { FormikProps } from "formik";
import { get, isArray, isObject, set } from "lodash";
import { useState } from "react";
import { useQueryClient, UseQueryResult } from "react-query";
import Select from "react-select/async";
import { useT } from "../../../hooks/useT";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";

export interface FormSelectBase<T, ValueIdentifier>
  extends BaseFormElementProps {
  /**
   * @description label is what user will see as the text on top of the input or near
   * it depenging on the design.
   */
  label?: string;

  /**
   * @description placeholder is the same common placeholder element of any input
   */
  placeholder?: string;

  /**
   * @description errorMessage is usually a string message when available it will make the field red
   * and show the message below it to the user. Can be directly read from the formik errors object as well
   */
  errorMessage?: string;

  /**
   * @description keyExtractor extract the value of the input which will be compared to determine
   * which option is selected, as well as will be used upon a new selection
   */
  keyExtractor?: (t: T) => ValueIdentifier;

  /**
   * @description List of the items which will be used to show as options. This is not async,
   * if you want to have dynamic options maybe better to use FormEntitySelect.
   */
  // options?: T[];

  /**
   * @description children object.
   */
  children?: any;

  /**
   * @description fnLabelFormat will be called on each item in the list, to create the string which user
   * will be seeing.
   */
  fnLabelFormat?: (item: T) => string;

  /**
   * @description Triggers each type user types into the auto suggestion list.
   */
  onInputChange?: (t: string) => void;

  /**
   * @description Skips the autocompletion component and renders the html <select... components
   * instead regardless.
   */
  convertToNative?: boolean;

  /**
   * @description You can have different type of the select.
   * @enum auto means automatically decideds for you
   * @enum verbose means it would show the options as a radio list so user can choose.
   */
  type?: "auto" | "verbose";

  /**
   * @description name property of the input will appear on html[name=xxx]
   */
  name?: string;

  /**
   * Fireback Query Result which includes items and react-query query object.
   * This is the only way to provide the form select with options,
   * even static array needs to be converted.
   * @param params
   * @returns
   */
  querySource: (params: UseRemoteQuery) => {
    query: UseQueryResult<IResponseList<T>, any>;
    items: T[];
    keyExtractor?: (item: T) => any;
  };

  /**
   * Call back to create JsonQuery object to filter and search the endpoint
   * @param keyword
   * @returns
   */
  jsonQuery?: (keyword: string) => JsonQuery;

  /**
   * @description withPreloads
   * Goes to the query to left join inner tables (objects) or foreign relations if needed.
   */
  withPreloads?: string;
}

interface FormSelectEffectBase<TargetType, T, ValueIdentifier> {
  form: FormikProps<TargetType>;
  field: string;

  /**
   * When set true, it would skip adding ListId or Id fields suffix for objects
   * and arrays used in Fireback entities
   */
  skipFirebackMetaData?: boolean;
}

interface FormSelectEffect<TargetType, T, ValueIdentifier>
  extends FormSelectEffectBase<TargetType, T, ValueIdentifier> {
  beforeSet?: (item: T) => ValueIdentifier;
}

interface FormSelectMultipleEffect<TargetType, T, ValueIdentifier>
  extends FormSelectEffectBase<TargetType, T, ValueIdentifier> {
  beforeSet?: (items: T[]) => ValueIdentifier[];
}

export interface FormSelectProps<T, ValueIdentifier>
  extends FormSelectBase<T, ValueIdentifier> {
  /**
   * @description value is the form element actual values which will be read from the form object,
   * regardless of the options type
   */
  value?: T | ValueIdentifier;

  /**
   * @description allows the user to have multiple selection
   */
  multiple?: boolean;

  /**
   * @description Will be triggered regardless of the usage when a value has been changed.
   * @returns
   */
  onChange?: (value: T) => void;

  /**
   * @description formEffect
   * Magic option used for applying the value change directly into a formik object,
   * useful for selecting object, array items
   */
  formEffect?: FormSelectEffect<any, T, ValueIdentifier>;
}

export interface FormSelectMultipleProps<T, ValueIdentifier>
  extends FormSelectBase<T, ValueIdentifier> {
  /**
   * @description value is the form element actual values which will be read from the form object,
   * regardless of the options type
   */
  value?: T[];

  /**
   * @description Will be triggered regardless of the usage when a value has been changed.
   * @returns
   */
  onChange?: (value: T[]) => void;

  /**
   * @description formEffect
   * Magic option used for applying the value change directly into a formik object,
   * useful for selecting object, array items
   */
  formEffect?: FormSelectMultipleEffect<any, T, ValueIdentifier>;
}

function resolveJsonQuery(
  keyword: string,
  userInput?: (keyword: string) => JsonQuery
): JsonQuery {
  if (userInput) {
    return userInput(keyword);
  }
  return {
    name: {
      operation: "contains",
      value: keyword,
    },
  };
}

export function FormSelectMultiple<T, V>(props: FormSelectMultipleProps<T, V>) {
  return <FormSelect<T, V> {...(props as any)} multiple={true} />;
}
export function FormSelect<T, V>(props: FormSelectProps<T, V>) {
  const t = useT();

  const queryClient = useQueryClient();
  let [keyword, setKeyword] = useState<string>("");

  if (!props.querySource) {
    return <div>No query source to render</div>;
  }

  const { query, keyExtractor: queryKeyExtractor } = props.querySource({
    queryClient,
    query: {
      itemsPerPage: 20,
      jsonQuery: resolveJsonQuery(keyword, props.jsonQuery),
      withPreloads: props.withPreloads,
    },
    queryOptions: {
      refetchOnWindowFocus: false,
    },
  });

  const keyExtractor: (t: T) => V =
    props.keyExtractor || queryKeyExtractor || ((item) => JSON.stringify(item));

  const options = query?.data?.data?.items;

  const onChange = (value: T | T[]) => {
    // if there are form effect, we need to apply them, depending on the type
    if (props?.formEffect?.form) {
      const { formEffect } = props;
      const newValue = {
        ...formEffect.form.values,
      };

      if (formEffect.beforeSet) {
        value = formEffect.beforeSet(value as T) as any;
      }

      set(newValue, formEffect.field, value);

      // We need to apply to the form effect based on the actual value of the data which
      // has been changed, so it would work outof the box.
      // For the object, we need to add the Id field as well alongside the object itself.
      // This might be unnecessary.
      if (
        isObject(value) &&
        (value as any).uniqueId &&
        formEffect.skipFirebackMetaData !== true
      ) {
        set(newValue, formEffect.field + "Id", (value as any).uniqueId);
      }

      // If array, we need to extract all of the items uniqueId, and send with ListId suffix
      // for fireback to pick them up.
      if (isArray(value) && formEffect.skipFirebackMetaData !== true) {
        const arrayTarget = formEffect.field + "ListId";
        set(
          newValue,
          arrayTarget,
          (value || []).map((t: any) => t.uniqueId)
        );
      }

      formEffect?.form.setValues(newValue);
    }

    // regardless of formEffect, if there is unchange we are going to call onChange, if it's provided.
    if (props.onChange && typeof props.onChange === "function") {
      props.onChange(value as T);
    }
  };

  // Let's pick the value from formEffect.
  let value = props.value;
  if (value === undefined && props.formEffect?.form) {
    const possibleValue = get(
      props.formEffect.form.values,
      props.formEffect.field
    );
    if (possibleValue !== undefined) {
      value = possibleValue;
    }
  }

  if (typeof value !== "object" && keyExtractor && value !== undefined) {
    value = options.find((item) => keyExtractor(item) === value);
  }

  // if (props.type === "verbose") {
  //   return <VerboseSelect {...props} />;
  // }

  const promiseOptions = (inputValue: string) =>
    new Promise<T[]>((resolve) => {
      setTimeout(() => {
        resolve(options);
      }, 100);
    });

  return (
    <BaseFormElement {...props}>
      {props.children}
      {props.convertToNative ? (
        <select
          value={value as any}
          multiple={props.multiple}
          onChange={(e) => {
            const item = options?.find(
              (t: any) => t.uniqueId === e.target.value
            ) as any;

            onChange(item);
          }}
          className={classNames(
            "form-select",
            props.errorMessage && "is-invalid",
            props.validMessage && "is-valid"
          )}
          disabled={props.disabled}
          aria-label="Default select example"
        >
          <option key={undefined} value={""}>
            {t.selectPlaceholder}
          </option>
          {options?.filter(Boolean).map((t) => {
            const itemValue = keyExtractor(t);
            return (
              <option key={itemValue as any} value={itemValue as any}>
                {props.fnLabelFormat(t)}
              </option>
            );
          })}
        </select>
      ) : (
        <>
          <Select
            value={value as any}
            onChange={(newValue) => {
              onChange(newValue as T);
            }}
            isMulti={props.multiple}
            classNames={{
              container(propsx: any) {
                return classNames(
                  props.errorMessage &&
                    " form-control form-control-no-padding is-invalid",
                  props.validMessage && "is-valid"
                );
              },
              control(props2: any) {
                return classNames("form-control form-control-no-padding");
              },
              menu(props) {
                return "react-select-menu-area";
              },
            }}
            isSearchable
            defaultOptions={options}
            placeholder={t.searchplaceholder}
            noOptionsMessage={() => t.noOptions}
            getOptionValue={keyExtractor as any}
            loadOptions={promiseOptions}
            formatOptionLabel={props.fnLabelFormat}
            onInputChange={setKeyword}
          />
        </>
      )}
    </BaseFormElement>
  );
}

// function VerboseSelect<T, ValueIdentifier>(
//   props: FormSelectProps<T, ValueIdentifier>
// ) {
//   return (
//     <BaseFormElement {...props}>
//       <div className="form-select-verbos">
//         {options?.map((item) => {
//           const value = props.keyExtractor(item);

//           return (
//             <label key={`${value}`}>
//               <input
//                 name={props.name}
//                 type="radio"
//                 onClick={(t) => {
//                   props.onChange(value);
//                 }}
//                 value={`${value}`}
//                 checked={value === props.value}
//               />
//               {props.fnLabelFormat(item)}
//             </label>
//           );
//         })}
//       </div>
//     </BaseFormElement>
//   );
// }
