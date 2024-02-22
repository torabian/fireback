package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnInterpolateConfigDto extends JsonSerializable {
    public Float[] sources;
    public Float[] targets;
    public static class VM extends ViewModel {
    // upper: Sources sources
    private MutableLiveData< Float[] > sources = new MutableLiveData<>();
    public MutableLiveData< Float[] > getSources() {
        return sources;
    }
    public void setSources( Float[]  v) {
        sources.setValue(v);
    }
    // upper: Targets targets
    private MutableLiveData< Float[] > targets = new MutableLiveData<>();
    public MutableLiveData< Float[] > getTargets() {
        return targets;
    }
    public void setTargets( Float[]  v) {
        targets.setValue(v);
    }
    }
}