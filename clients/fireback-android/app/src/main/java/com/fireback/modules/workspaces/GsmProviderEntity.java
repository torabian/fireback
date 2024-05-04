package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class GsmProviderEntity extends JsonSerializable {
    public String apiKey;
    public String mainSenderNumber;
    public String type;
    public String invokeUrl;
    public String invokeBody;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: ApiKey apiKey
    private MutableLiveData< String > apiKey = new MutableLiveData<>();
    public MutableLiveData< String > getApiKey() {
        return apiKey;
    }
    public void setApiKey( String  v) {
        apiKey.setValue(v);
    }
    // upper: MainSenderNumber mainSenderNumber
    private MutableLiveData< String > mainSenderNumber = new MutableLiveData<>();
    public MutableLiveData< String > getMainSenderNumber() {
        return mainSenderNumber;
    }
    public void setMainSenderNumber( String  v) {
        mainSenderNumber.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: InvokeUrl invokeUrl
    private MutableLiveData< String > invokeUrl = new MutableLiveData<>();
    public MutableLiveData< String > getInvokeUrl() {
        return invokeUrl;
    }
    public void setInvokeUrl( String  v) {
        invokeUrl.setValue(v);
    }
    // upper: InvokeBody invokeBody
    private MutableLiveData< String > invokeBody = new MutableLiveData<>();
    public MutableLiveData< String > getInvokeBody() {
        return invokeBody;
    }
    public void setInvokeBody( String  v) {
        invokeBody.setValue(v);
    }
    // Handling error message for each field
    // upper: ApiKey apiKey
    private MutableLiveData<String> apiKeyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getApiKeyMsg() {
        return apiKeyMsg;
    }
    public void setApiKeyMsg(String v) {
        apiKeyMsg.setValue(v);
    }
    // upper: MainSenderNumber mainSenderNumber
    private MutableLiveData<String> mainSenderNumberMsg = new MutableLiveData<>();
    public MutableLiveData<String> getMainSenderNumberMsg() {
        return mainSenderNumberMsg;
    }
    public void setMainSenderNumberMsg(String v) {
        mainSenderNumberMsg.setValue(v);
    }
    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTypeMsg() {
        return typeMsg;
    }
    public void setTypeMsg(String v) {
        typeMsg.setValue(v);
    }
    // upper: InvokeUrl invokeUrl
    private MutableLiveData<String> invokeUrlMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInvokeUrlMsg() {
        return invokeUrlMsg;
    }
    public void setInvokeUrlMsg(String v) {
        invokeUrlMsg.setValue(v);
    }
    // upper: InvokeBody invokeBody
    private MutableLiveData<String> invokeBodyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInvokeBodyMsg() {
        return invokeBodyMsg;
    }
    public void setInvokeBodyMsg(String v) {
        invokeBodyMsg.setValue(v);
    }
  }
}