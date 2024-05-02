package com.fireback.modules.licenses;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class LicensableProductEntity extends JsonSerializable {
    public String name;
    public String privateKey;
    public String publicKey;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: PrivateKey privateKey
    private MutableLiveData< String > privateKey = new MutableLiveData<>();
    public MutableLiveData< String > getPrivateKey() {
        return privateKey;
    }
    public void setPrivateKey( String  v) {
        privateKey.setValue(v);
    }
    // upper: PublicKey publicKey
    private MutableLiveData< String > publicKey = new MutableLiveData<>();
    public MutableLiveData< String > getPublicKey() {
        return publicKey;
    }
    public void setPublicKey( String  v) {
        publicKey.setValue(v);
    }
    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: PrivateKey privateKey
    private MutableLiveData<String> privateKeyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPrivateKeyMsg() {
        return privateKeyMsg;
    }
    public void setPrivateKeyMsg(String v) {
        privateKeyMsg.setValue(v);
    }
    // upper: PublicKey publicKey
    private MutableLiveData<String> publicKeyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPublicKeyMsg() {
        return publicKeyMsg;
    }
    public void setPublicKeyMsg(String v) {
        publicKeyMsg.setValue(v);
    }
  }
}