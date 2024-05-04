package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.ResponseErrorException;
public class GsmSendSmsWithProviderAction {
    public static class Req extends JsonSerializable {
    public GsmProviderEntity gsmProvider;
    public String toNumber;
    public String body;
    // upper: GsmProvider gsmProvider
    private MutableLiveData<String> gsmProviderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getGsmProviderMsg() {
        return gsmProviderMsg;
    }
    public void setGsmProviderMsg(String v) {
        gsmProviderMsg.setValue(v);
    }
    // upper: ToNumber toNumber
    private MutableLiveData<String> toNumberMsg = new MutableLiveData<>();
    public MutableLiveData<String> getToNumberMsg() {
        return toNumberMsg;
    }
    public void setToNumberMsg(String v) {
        toNumberMsg.setValue(v);
    }
    // upper: Body body
    private MutableLiveData<String> bodyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBodyMsg() {
        return bodyMsg;
    }
    public void setBodyMsg(String v) {
        bodyMsg.setValue(v);
    }
    }
    public static class ReqViewModel extends ViewModel {
    // upper: GsmProvider gsmProvider
    private MutableLiveData< GsmProviderEntity > gsmProvider = new MutableLiveData<>();
    public MutableLiveData< GsmProviderEntity > getGsmProvider() {
        return gsmProvider;
    }
    public void setGsmProvider( GsmProviderEntity  v) {
        gsmProvider.setValue(v);
    }
    // upper: ToNumber toNumber
    private MutableLiveData< String > toNumber = new MutableLiveData<>();
    public MutableLiveData< String > getToNumber() {
        return toNumber;
    }
    public void setToNumber( String  v) {
        toNumber.setValue(v);
    }
    // upper: Body body
    private MutableLiveData< String > body = new MutableLiveData<>();
    public MutableLiveData< String > getBody() {
        return body;
    }
    public void setBody( String  v) {
        body.setValue(v);
    }
    // upper: GsmProvider gsmProvider
    private MutableLiveData<String> gsmProviderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getGsmProviderMsg() {
        return gsmProviderMsg;
    }
    public void setGsmProviderMsg(String v) {
        gsmProviderMsg.setValue(v);
    }
    // upper: ToNumber toNumber
    private MutableLiveData<String> toNumberMsg = new MutableLiveData<>();
    public MutableLiveData<String> getToNumberMsg() {
        return toNumberMsg;
    }
    public void setToNumberMsg(String v) {
        toNumberMsg.setValue(v);
    }
    // upper: Body body
    private MutableLiveData<String> bodyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBodyMsg() {
        return bodyMsg;
    }
    public void setBodyMsg(String v) {
        bodyMsg.setValue(v);
    }
public void applyException(Throwable e) {
    if (!(e instanceof ResponseErrorException)) {
        return;
    }
    ResponseErrorException responseError = (ResponseErrorException) e;
    // @todo on fireback: This needs to be recursive.
    responseError.error.errors.forEach(item -> {
        if (item.location != null && item.location.equals("gsmProvider")) {
            this.setGsmProviderMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("toNumber")) {
            this.setToNumberMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("body")) {
            this.setBodyMsg(item.messageTranslated);
        }
    });
}
    }
    public static class Res extends JsonSerializable {
    public String queueId;
    // upper: QueueId queueId
    private MutableLiveData<String> queueIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getQueueIdMsg() {
        return queueIdMsg;
    }
    public void setQueueIdMsg(String v) {
        queueIdMsg.setValue(v);
    }
    }
}