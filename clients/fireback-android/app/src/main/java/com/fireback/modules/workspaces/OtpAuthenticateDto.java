package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class OtpAuthenticateDto extends JsonSerializable {
    public String value;
    public String otp;
    public String type;
    public String password;
    public static class VM extends ViewModel {
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Otp otp
    private MutableLiveData< String > otp = new MutableLiveData<>();
    public MutableLiveData< String > getOtp() {
        return otp;
    }
    public void setOtp( String  v) {
        otp.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
    }
    }
}