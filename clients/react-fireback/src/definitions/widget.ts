import { WidgetEntity } from "src/sdk/fireback";

export enum WidgetAvailableProvider {
  SingleDigitalGpio = "single_digital_gpio",
  NumericValueRealtimeChart = "numeric_value_realtime_chart",
  TemperatureControl = "temperature_control",
}

export enum SupportedWidgetFamily {
  SystemSmall = "systemSmall",
  SystemMedium = "systemMedium",
  SystemLarge = "systemLarge",
}

export interface Widget {
  displayName: string;
  description: string;
  supportedFamilies: SupportedWidgetFamily[];
}

export interface WidgetProps<T> {
  widget: WidgetEntity;
  data: T;
  preview?: boolean;
}
