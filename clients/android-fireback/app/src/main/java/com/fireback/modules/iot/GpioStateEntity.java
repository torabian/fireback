package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class GpioStateEntity extends JsonSerializable {
    public GpioModeEntity gpioMode;
    public int gpioIndexSnapshot;
    public int gpioModeSnapshot;
    public int gpioValueSnapshot;
    public GpioEntity gpio;
    public static class VM extends ViewModel {
    // upper: GpioMode gpioMode
    private MutableLiveData< GpioModeEntity > gpioMode = new MutableLiveData<>();
    public MutableLiveData< GpioModeEntity > getGpioMode() {
        return gpioMode;
    }
    public void setGpioMode( GpioModeEntity  v) {
        gpioMode.setValue(v);
    }
    // upper: GpioIndexSnapshot gpioIndexSnapshot
    private MutableLiveData< Integer > gpioIndexSnapshot = new MutableLiveData<>();
    public MutableLiveData< Integer > getGpioIndexSnapshot() {
        return gpioIndexSnapshot;
    }
    public void setGpioIndexSnapshot( Integer  v) {
        gpioIndexSnapshot.setValue(v);
    }
    // upper: GpioModeSnapshot gpioModeSnapshot
    private MutableLiveData< Integer > gpioModeSnapshot = new MutableLiveData<>();
    public MutableLiveData< Integer > getGpioModeSnapshot() {
        return gpioModeSnapshot;
    }
    public void setGpioModeSnapshot( Integer  v) {
        gpioModeSnapshot.setValue(v);
    }
    // upper: GpioValueSnapshot gpioValueSnapshot
    private MutableLiveData< Integer > gpioValueSnapshot = new MutableLiveData<>();
    public MutableLiveData< Integer > getGpioValueSnapshot() {
        return gpioValueSnapshot;
    }
    public void setGpioValueSnapshot( Integer  v) {
        gpioValueSnapshot.setValue(v);
    }
    // upper: Gpio gpio
    private MutableLiveData< GpioEntity > gpio = new MutableLiveData<>();
    public MutableLiveData< GpioEntity > getGpio() {
        return gpio;
    }
    public void setGpio( GpioEntity  v) {
        gpio.setValue(v);
    }
    }
}