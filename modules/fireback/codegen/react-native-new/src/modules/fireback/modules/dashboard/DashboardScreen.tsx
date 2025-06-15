import {ScrollView, TouchableOpacity, View} from 'react-native';
import {
  BarChart,
  LineChart,
  PieChart,
  PopulationPyramid,
} from 'react-native-gifted-charts';
import {ConfirmDrawer} from '../../components/confirm/ConfirmDrawer';
import {useBottomSheet} from '../../hooks/BottomSheetProvider';

export const DashboardScreen = () => {
  const {openBottomSheet} = useBottomSheet();

  const data = [{value: 50}, {value: 80}, {value: 90}, {value: 70}];

  return (
    <ScrollView>
      <View style={{flexDirection: 'row'}}>
        <TouchableOpacity
          onPress={() => {
            openBottomSheet(({onConfirm, onReject}) => (
              <ConfirmDrawer
                title="Even gets title"
                onConfirm={onConfirm}
                onReject={onReject}
              />
            )).then(() =>
              openBottomSheet(({onConfirm, onReject}) => (
                <ConfirmDrawer
                  title="Even gets title"
                  onConfirm={onConfirm}
                  onReject={onReject}
                />
              )),
            );
          }}
          style={{flex: 1}}>
          <BarChart data={data} />
        </TouchableOpacity>
        <View style={{flex: 1}}>
          <LineChart data={data} />
        </View>
      </View>
      <View>
        <View>
          <PieChart data={data} />
        </View>
        <View>
          <PopulationPyramid
            data={[
              {left: 10, right: 12},
              {left: 9, right: 8},
            ]}
          />
        </View>
      </View>
    </ScrollView>
  );
};

DashboardScreen.Name = 'DashboardScreen';
