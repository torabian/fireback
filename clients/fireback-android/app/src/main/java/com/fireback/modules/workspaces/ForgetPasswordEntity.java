package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ForgetPasswordEntity extends JsonSerializable {
    public UserEntity user;
    public PassportEntity passport;
    public String status;
    public String validUntil;
    public String blockedUntil;
    public int secondsToUnblock;
    public String otp;
    public String recoveryAbsoluteUrl;
    public static class VM extends ViewModel {
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: Passport passport
    private MutableLiveData< PassportEntity > passport = new MutableLiveData<>();
    public MutableLiveData< PassportEntity > getPassport() {
        return passport;
    }
    public void setPassport( PassportEntity  v) {
        passport.setValue(v);
    }
    // upper: Status status
    private MutableLiveData< String > status = new MutableLiveData<>();
    public MutableLiveData< String > getStatus() {
        return status;
    }
    public void setStatus( String  v) {
        status.setValue(v);
    }
    // upper: ValidUntil validUntil
    private MutableLiveData< String > validUntil = new MutableLiveData<>();
    public MutableLiveData< String > getValidUntil() {
        return validUntil;
    }
    public void setValidUntil( String  v) {
        validUntil.setValue(v);
    }
    // upper: BlockedUntil blockedUntil
    private MutableLiveData< String > blockedUntil = new MutableLiveData<>();
    public MutableLiveData< String > getBlockedUntil() {
        return blockedUntil;
    }
    public void setBlockedUntil( String  v) {
        blockedUntil.setValue(v);
    }
    // upper: SecondsToUnblock secondsToUnblock
    private MutableLiveData< Integer > secondsToUnblock = new MutableLiveData<>();
    public MutableLiveData< Integer > getSecondsToUnblock() {
        return secondsToUnblock;
    }
    public void setSecondsToUnblock( Integer  v) {
        secondsToUnblock.setValue(v);
    }
    // upper: Otp otp
    private MutableLiveData< String > otp = new MutableLiveData<>();
    public MutableLiveData< String > getOtp() {
        return otp;
    }
    public void setOtp( String  v) {
        otp.setValue(v);
    }
    // upper: RecoveryAbsoluteUrl recoveryAbsoluteUrl
    private MutableLiveData< String > recoveryAbsoluteUrl = new MutableLiveData<>();
    public MutableLiveData< String > getRecoveryAbsoluteUrl() {
        return recoveryAbsoluteUrl;
    }
    public void setRecoveryAbsoluteUrl( String  v) {
        recoveryAbsoluteUrl.setValue(v);
    }
    }
}