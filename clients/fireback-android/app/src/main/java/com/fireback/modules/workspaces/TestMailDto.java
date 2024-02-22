package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class TestMailDto extends JsonSerializable {
    public String senderId;
    public String toName;
    public String toEmail;
    public String subject;
    public String content;
    public static class VM extends ViewModel {
    // upper: SenderId senderId
    private MutableLiveData< String > senderId = new MutableLiveData<>();
    public MutableLiveData< String > getSenderId() {
        return senderId;
    }
    public void setSenderId( String  v) {
        senderId.setValue(v);
    }
    // upper: ToName toName
    private MutableLiveData< String > toName = new MutableLiveData<>();
    public MutableLiveData< String > getToName() {
        return toName;
    }
    public void setToName( String  v) {
        toName.setValue(v);
    }
    // upper: ToEmail toEmail
    private MutableLiveData< String > toEmail = new MutableLiveData<>();
    public MutableLiveData< String > getToEmail() {
        return toEmail;
    }
    public void setToEmail( String  v) {
        toEmail.setValue(v);
    }
    // upper: Subject subject
    private MutableLiveData< String > subject = new MutableLiveData<>();
    public MutableLiveData< String > getSubject() {
        return subject;
    }
    public void setSubject( String  v) {
        subject.setValue(v);
    }
    // upper: Content content
    private MutableLiveData< String > content = new MutableLiveData<>();
    public MutableLiveData< String > getContent() {
        return content;
    }
    public void setContent( String  v) {
        content.setValue(v);
    }
    }
}