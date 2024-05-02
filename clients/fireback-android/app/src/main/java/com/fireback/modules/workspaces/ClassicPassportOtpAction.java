package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ClassicPassportOtpAction {
    public static class Req extends JsonSerializable {
    public String value;
    public String otp;
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Otp otp
    private MutableLiveData<String> otpMsg = new MutableLiveData<>();
    public MutableLiveData<String> getOtpMsg() {
        return otpMsg;
    }
    public void setOtpMsg(String v) {
        otpMsg.setValue(v);
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
    // upper: Otp otp
    private MutableLiveData< String > otp = new MutableLiveData<>();
    public MutableLiveData< String > getOtp() {
        return otp;
    }
    public void setOtp( String  v) {
        otp.setValue(v);
    }
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Otp otp
    private MutableLiveData<String> otpMsg = new MutableLiveData<>();
    public MutableLiveData<String> getOtpMsg() {
        return otpMsg;
    }
    public void setOtpMsg(String v) {
        otpMsg.setValue(v);
    }
    }
    public static class Res extends JsonSerializable {
    public int suspendUntil;
    public UserSessionDto session;
    public int validUntil;
    public int blockedUntil;
    public int secondsToUnblock;
    // upper: SuspendUntil suspendUntil
    private MutableLiveData<String> suspendUntilMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSuspendUntilMsg() {
        return suspendUntilMsg;
    }
    public void setSuspendUntilMsg(String v) {
        suspendUntilMsg.setValue(v);
    }
    // upper: Session session
    private MutableLiveData<String> sessionMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSessionMsg() {
        return sessionMsg;
    }
    public void setSessionMsg(String v) {
        sessionMsg.setValue(v);
    }
    // upper: ValidUntil validUntil
    private MutableLiveData<String> validUntilMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValidUntilMsg() {
        return validUntilMsg;
    }
    public void setValidUntilMsg(String v) {
        validUntilMsg.setValue(v);
    }
    // upper: BlockedUntil blockedUntil
    private MutableLiveData<String> blockedUntilMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBlockedUntilMsg() {
        return blockedUntilMsg;
    }
    public void setBlockedUntilMsg(String v) {
        blockedUntilMsg.setValue(v);
    }
    // upper: SecondsToUnblock secondsToUnblock
    private MutableLiveData<String> secondsToUnblockMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSecondsToUnblockMsg() {
        return secondsToUnblockMsg;
    }
    public void setSecondsToUnblockMsg(String v) {
        secondsToUnblockMsg.setValue(v);
    }
    }
}