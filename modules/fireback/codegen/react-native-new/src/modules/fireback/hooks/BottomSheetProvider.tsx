import React, {createContext, useContext, useState, ReactNode} from 'react';
import BottomSheet, {BottomSheetView} from '@gorhom/bottom-sheet';
import {StyleSheet, Text} from 'react-native';

type SheetItem = {
  id: number;
  Component: React.ComponentType<{onConfirm: () => void; onReject: () => void}>;
  resolve: () => void;
  reject: () => void;
};

const BottomSheetContext = createContext<{
  openBottomSheet: (Component: SheetItem['Component']) => Promise<void>;
} | null>(null);

let uniqueId = 0;

export const BottomSheetProvider = ({children}: {children: ReactNode}) => {
  const [sheets, setSheets] = useState<SheetItem[]>([]);

  const openBottomSheet = (Component: SheetItem['Component']) => {
    return new Promise<void>((resolve, reject) => {
      const id = uniqueId++;
      const newSheet: SheetItem = {
        id,
        Component,
        resolve: () => {
          resolve();
          setSheets(prev => prev.filter(s => s.id !== id));
        },
        reject: () => {
          reject();
          setSheets(prev => prev.filter(s => s.id !== id));
        },
      };
      setSheets(prev => [...prev, newSheet]);
    });
  };

  return (
    <BottomSheetContext.Provider value={{openBottomSheet}}>
      {children}
      {sheets.map(({id, Component, resolve, reject}) => (
        <BottomSheet
          key={id}
          index={0}
          snapPoints={['50%']}
          onClose={reject}
          enablePanDownToClose>
          <BottomSheetView style={styles.contentContainer}>
            <Component onConfirm={resolve} onReject={reject} />
          </BottomSheetView>
        </BottomSheet>
      ))}
    </BottomSheetContext.Provider>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'grey',
  },
  contentContainer: {
    flex: 1,
    padding: 36,
    alignItems: 'center',
  },
});

export const useBottomSheet = () => {
  const ctx = useContext(BottomSheetContext);
  if (!ctx)
    throw new Error('useBottomSheet must be inside BottomSheetProvider');
  return ctx;
};
