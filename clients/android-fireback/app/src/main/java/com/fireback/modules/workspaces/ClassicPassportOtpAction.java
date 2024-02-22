package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ClassicPassportOtpAction {
    public static class Req extends JsonSerializable {
    public String value;
    public String otp;
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
    // upper: Otp otp
    private MutableLiveData< String > otp = new MutableLiveData<>();
    public MutableLiveData< String > getOtp() {
        return otp;
    }
    public void setOtp( String  v) {
        otp.setValue(v);
    }
    }
    public static class Res extends JsonSerializable {
    public int suspendUntil;
    public UserSessionDto session;
    public int validUntil;
    public int blockedUntil;
    public int secondsToUnblock;
    }
}