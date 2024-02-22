package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class GpioEntity extends JsonSerializable {
    public String name;
    public int index;
    public String analogFunction;
    public String rtcGpio;
    public String comments;
    public GpioModeEntity mode;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Index index
    private MutableLiveData< Integer > index = new MutableLiveData<>();
    public MutableLiveData< Integer > getIndex() {
        return index;
    }
    public void setIndex( Integer  v) {
        index.setValue(v);
    }
    // upper: AnalogFunction analogFunction
    private MutableLiveData< String > analogFunction = new MutableLiveData<>();
    public MutableLiveData< String > getAnalogFunction() {
        return analogFunction;
    }
    public void setAnalogFunction( String  v) {
        analogFunction.setValue(v);
    }
    // upper: RtcGpio rtcGpio
    private MutableLiveData< String > rtcGpio = new MutableLiveData<>();
    public MutableLiveData< String > getRtcGpio() {
        return rtcGpio;
    }
    public void setRtcGpio( String  v) {
        rtcGpio.setValue(v);
    }
    // upper: Comments comments
    private MutableLiveData< String > comments = new MutableLiveData<>();
    public MutableLiveData< String > getComments() {
        return comments;
    }
    public void setComments( String  v) {
        comments.setValue(v);
    }
    // upper: Mode mode
    private MutableLiveData< GpioModeEntity > mode = new MutableLiveData<>();
    public MutableLiveData< GpioModeEntity > getMode() {
        return mode;
    }
    public void setMode( GpioModeEntity  v) {
        mode.setValue(v);
    }
    }
}