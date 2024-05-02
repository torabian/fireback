package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class CheckClassicPassportAction {
    public static class Req extends JsonSerializable {
    public String value;
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    }
    public static class ReqViewModel extends ViewModel {
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    }
    public static class Res extends JsonSerializable {
    public Boolean exists;
    // upper: Exists exists
    private MutableLiveData<String> existsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getExistsMsg() {
        return existsMsg;
    }
    public void setExistsMsg(String v) {
        existsMsg.setValue(v);
    }
    }
}