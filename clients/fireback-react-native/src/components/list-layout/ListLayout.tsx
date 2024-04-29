import { StyleSheet, Text, View } from "react-native";
import { themeEl } from "../../themes/theme";

export function ListLayout({
  title,
  children,
}: {
  title: string;
  children: any;
}) {
  return (
    <View style={styles.wrapper}>
      <Text style={themeEl.h1}>{title}</Text>
      {children}
    </View>
  );
}

const styles = StyleSheet.create({
  wrapper: {
    flex: 1,
  },
});
