package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ClassicSigninAction {
    public static class Req extends JsonSerializable {
    public String value;
    public String password;
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Password password
    private MutableLiveData<String> passwordMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPasswordMsg() {
        return passwordMsg;
    }
    public void setPasswordMsg(String v) {
        passwordMsg.setValue(v);
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
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
    }
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Password password
    private MutableLiveData<String> passwordMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPasswordMsg() {
        return passwordMsg;
    }
    public void setPasswordMsg(String v) {
        passwordMsg.setValue(v);
    }
    }
}