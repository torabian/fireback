import { useQueryClient, type UseQueryResult } from "@tanstack/react-query";
import classNames from "classnames";
import { type FormikProps } from "formik";
import { get, isArray, isObject, set } from "lodash";
import { useState } from "react";
import Select from "react-select/async";
import { type UseRemoteQuery } from "../../../types/remoteQuery";
import { useS } from "../../../hooks/useS";
import { strings } from "../../strings/translations";
import {
  BaseFormElement,
  type BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import type { ResponseDto } from "@fireback/js-remote-ctx/envelopes/google-json-style-guide/generated/ResponseDto";

export interface FormSelectBase<
  T,
  ValueIdentifier,
> extends BaseFormElementProps {
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
    query: UseQueryResult<ResponseDto<T>, any>;
    items: T[];
    keyExtractor?: (item: T) => any;
  };

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

interface FormSelectEffect<
  TargetType,
  T,
  ValueIdentifier,
> extends FormSelectEffectBase<TargetType, T, ValueIdentifier> {
  beforeSet?: (item: T) => ValueIdentifier;
}

interface FormSelectMultipleEffect<
  TargetType,
  T,
  ValueIdentifier,
> extends FormSelectEffectBase<TargetType, T, ValueIdentifier> {
  beforeSet?: (items: T[]) => ValueIdentifier[];
}

export interface FormSelectProps<T, ValueIdentifier> extends FormSelectBase<
  T,
  ValueIdentifier
> {
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

export interface FormSelectMultipleProps<
  T,
  ValueIdentifier,
> extends FormSelectBase<T, ValueIdentifier> {
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

export function FormSelectMultiple<T, V>(props: FormSelectMultipleProps<T, V>) {
  return <FormSelect<T, V> {...(props as any)} multiple={true} />;
}
export function FormSelect<T, V>(props: FormSelectProps<T, V>) {
  const s = useS(strings);

  const queryClient = useQueryClient();
  let [keyword, setKeyword] = useState<string>("");

  if (!props.querySource) {
    return <div>{s.components.noQuerySourceToRender}</div>;
  }

  // Bug fix: this used to fetch only 20 items and never actually filter by what
  // was typed at all (see promiseOptions below) - `keyword` was tracked but never
  // read anywhere, so the dropdown always showed the exact same unfiltered first
  // page regardless of the search box's contents. A generous itemsPerPage (these
  // are admin-configuration pickers - providers/templates realistically number in
  // the dozens, not thousands) plus real client-side filtering against `keyword`
  // covers this without depending on server-side search support, which the
  // generic Browse actions these querySources wrap don't have (SearchPhrase is
  // only meaningful to reactivesearch's own hand-written providers - see
  // abac.QueryMenusReact/QueryRolesReact - not the generic entity Browse/Query
  // path). `query.searchPhrase` is still forwarded in case a given querySource
  // does have real server-side search to offer, but nothing currently requires it.
  const { query, keyExtractor: queryKeyExtractor } = props.querySource({
    queryClient,
    query: {
      itemsPerPage: 200,
      withPreloads: props.withPreloads,
      searchPhrase: keyword,
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
          (value || []).map((t: any) => t.uniqueId),
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
      props.formEffect.field,
    );
    if (possibleValue !== undefined) {
      value = possibleValue;
    }
  }

  if (typeof value !== "object" && keyExtractor && value !== undefined) {
    value = (options || []).find((item) => keyExtractor(item) === value);
  }

  // if (props.type === "verbose") {
  //   return <VerboseSelect {...props} />;
  // }

  // Bug fix: this used to always resolve the full, unfiltered `options` list
  // regardless of inputValue - typing into the search box never actually narrowed
  // anything down. Match case-insensitively against whatever's actually rendered
  // for each option (fnLabelFormat, when given - the same string the user is
  // looking at), falling back to the raw item otherwise.
  const promiseOptions = (inputValue: string) =>
    new Promise<T[]>((resolve) => {
      setTimeout(() => {
        if (!inputValue) {
          resolve(options);
          return;
        }
        const needle = inputValue.toLowerCase();
        // Bug fix: falling back to String(item) for an object-shaped option
        // (e.g. the plain { label, value } pairs createQuerySource wraps a
        // static array in, as every EmailProvider/GsmProvider "Type" select
        // does) stringified to the useless "[object Object]" - matching
        // nothing typed, so the dropdown went empty for every keystroke
        // instead of narrowing down. react-select's own default rendering
        // (formatOptionLabel left unset, as none of these callers set it)
        // already falls back to reading option.label - mirror that same
        // convention here so the filter matches what's actually on screen.
        const labelOf = (item: T) =>
          (props.fnLabelFormat
            ? props.fnLabelFormat(item)
            : (item as any)?.label ?? String(item)) ?? "";
        resolve(
          (options || []).filter((item) =>
            labelOf(item).toLowerCase().includes(needle),
          ),
        );
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
              (t: any) => t.uniqueId === e.target.value,
            ) as any;

            onChange(item);
          }}
          className={classNames(
            "form-select",
            props.errorMessage && "is-invalid",
            props.validMessage && "is-valid",
          )}
          disabled={props.disabled}
          aria-label={s.components.defaultSelectExample}
        >
          <option key={undefined} value={""}>
            {s.selectPlaceholder}
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
                  props.validMessage && "is-valid",
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
            placeholder={s.searchplaceholder}
            noOptionsMessage={() => s.noOptions}
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
