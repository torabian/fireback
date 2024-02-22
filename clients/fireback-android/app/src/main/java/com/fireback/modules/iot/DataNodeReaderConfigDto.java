package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DataNodeReaderConfigDto extends JsonSerializable {
    public int interval;
    public static class VM extends ViewModel {
    // upper: Interval interval
    private MutableLiveData< Integer > interval = new MutableLiveData<>();
    public MutableLiveData< Integer > getInterval() {
        return interval;
    }
    public void setInterval( Integer  v) {
        interval.setValue(v);
    }
    }
}