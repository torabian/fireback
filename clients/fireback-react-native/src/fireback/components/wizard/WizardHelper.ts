import {FormikHelpers, FormikProps} from 'formik';
import {View} from 'react-native';
import Toast from 'react-native-toast-message';

export interface WizardStep {
  label: string;
  component: Element | (() => JSX.Element);
  formNeedsValidation?: boolean;
  isValid?: (formikProps: FormikProps<any>) => boolean;
}

export interface LayoutData {
  fx?: number;
  fy?: number;
  width: number;
  height: number;
  px: number;
  py: number;
}

export function detectViewDistances(views: View[]): Promise<LayoutData[]> {
  const items: Promise<LayoutData>[] = [];

  views.forEach(item => {
    const p$ = new Promise<LayoutData>(resolve => {
      item.measure((fx, fy, width, height, px, py) => {
        resolve({fx, fy, width, height, px, py});
      });
    });

    items.push(p$);
  });

  return Promise.all(items);
}

/**
 * Gets a set of views, and finds empty distance between them
 */
export function calculateLines(
  items: LayoutData[],
  height = 2,
  distance = 2,
): LayoutData[] {
  const lines: LayoutData[] = [];

  for (let i = 0; i < items.length - 1; i++) {
    const item = items[i];

    if (!item.px && !item.py) {
      continue;
    }

    const line: LayoutData = {
      height,
      width: items[i + 1].px - item.px - item.width - distance * 2,
      px: item.px + item.width + distance,
      py: item.height / 2 - height / 2,
    };

    lines.push(line);
  }

  return lines;
}

const toLowerFirst = (input: string): string => {
  if (!input || typeof input !== 'string') {
    return '';
  }

  const str = input.split('');
  str[0] = str[0].toLowerCase();

  return str.join('');
};
