package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class TemperatureHmiComponentDto extends JsonSerializable {
    public String viewMode;
    public String units;
    public float maximumTemperature;
    public float minimumTemperature;
    public static class VM extends ViewModel {
    // upper: ViewMode viewMode
    private MutableLiveData< String > viewMode = new MutableLiveData<>();
    public MutableLiveData< String > getViewMode() {
        return viewMode;
    }
    public void setViewMode( String  v) {
        viewMode.setValue(v);
    }
    // upper: Units units
    private MutableLiveData< String > units = new MutableLiveData<>();
    public MutableLiveData< String > getUnits() {
        return units;
    }
    public void setUnits( String  v) {
        units.setValue(v);
    }
    // upper: MaximumTemperature maximumTemperature
    private MutableLiveData< Float > maximumTemperature = new MutableLiveData<>();
    public MutableLiveData< Float > getMaximumTemperature() {
        return maximumTemperature;
    }
    public void setMaximumTemperature( Float  v) {
        maximumTemperature.setValue(v);
    }
    // upper: MinimumTemperature minimumTemperature
    private MutableLiveData< Float > minimumTemperature = new MutableLiveData<>();
    public MutableLiveData< Float > getMinimumTemperature() {
        return minimumTemperature;
    }
    public void setMinimumTemperature( Float  v) {
        minimumTemperature.setValue(v);
    }
    }
}