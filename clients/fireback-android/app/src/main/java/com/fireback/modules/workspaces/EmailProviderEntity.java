package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class EmailProviderEntity extends JsonSerializable {
    public String type;
    public String apiKey;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: ApiKey apiKey
    private MutableLiveData< String > apiKey = new MutableLiveData<>();
    public MutableLiveData< String > getApiKey() {
        return apiKey;
    }
    public void setApiKey( String  v) {
        apiKey.setValue(v);
    }
    // Handling error message for each field
    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTypeMsg() {
        return typeMsg;
    }
    public void setTypeMsg(String v) {
        typeMsg.setValue(v);
    }
    // upper: ApiKey apiKey
    private MutableLiveData<String> apiKeyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getApiKeyMsg() {
        return apiKeyMsg;
    }
    public void setApiKeyMsg(String v) {
        apiKeyMsg.setValue(v);
    }
  }
}