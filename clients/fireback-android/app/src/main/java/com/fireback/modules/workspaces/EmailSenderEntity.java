package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class EmailSenderEntity extends JsonSerializable {
    public String fromName;
    public String fromEmailAddress;
    public String replyTo;
    public String nickName;
    public static class VM extends ViewModel {
    // upper: FromName fromName
    private MutableLiveData< String > fromName = new MutableLiveData<>();
    public MutableLiveData< String > getFromName() {
        return fromName;
    }
    public void setFromName( String  v) {
        fromName.setValue(v);
    }
    // upper: FromEmailAddress fromEmailAddress
    private MutableLiveData< String > fromEmailAddress = new MutableLiveData<>();
    public MutableLiveData< String > getFromEmailAddress() {
        return fromEmailAddress;
    }
    public void setFromEmailAddress( String  v) {
        fromEmailAddress.setValue(v);
    }
    // upper: ReplyTo replyTo
    private MutableLiveData< String > replyTo = new MutableLiveData<>();
    public MutableLiveData< String > getReplyTo() {
        return replyTo;
    }
    public void setReplyTo( String  v) {
        replyTo.setValue(v);
    }
    // upper: NickName nickName
    private MutableLiveData< String > nickName = new MutableLiveData<>();
    public MutableLiveData< String > getNickName() {
        return nickName;
    }
    public void setNickName( String  v) {
        nickName.setValue(v);
    }
    }
}