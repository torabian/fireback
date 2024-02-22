package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class FakeIotEnvDto extends JsonSerializable {
    public float core1temperature;
    public float core2temperature;
    public static class VM extends ViewModel {
    // upper: Core1temperature core1temperature
    private MutableLiveData< Float > core1temperature = new MutableLiveData<>();
    public MutableLiveData< Float > getCore1temperature() {
        return core1temperature;
    }
    public void setCore1temperature( Float  v) {
        core1temperature.setValue(v);
    }
    // upper: Core2temperature core2temperature
    private MutableLiveData< Float > core2temperature = new MutableLiveData<>();
    public MutableLiveData< Float > getCore2temperature() {
        return core2temperature;
    }
    public void setCore2temperature( Float  v) {
        core2temperature.setValue(v);
    }
    }
}