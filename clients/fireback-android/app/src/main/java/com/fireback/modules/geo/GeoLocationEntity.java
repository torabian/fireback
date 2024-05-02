package com.fireback.modules.geo;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class GeoLocationEntity extends JsonSerializable {
    public String name;
    public String code;
    public GeoLocationTypeEntity type;
    public String status;
    public String flag;
    public String officialName;
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
    // upper: Code code
    private MutableLiveData< String > code = new MutableLiveData<>();
    public MutableLiveData< String > getCode() {
        return code;
    }
    public void setCode( String  v) {
        code.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< GeoLocationTypeEntity > type = new MutableLiveData<>();
    public MutableLiveData< GeoLocationTypeEntity > getType() {
        return type;
    }
    public void setType( GeoLocationTypeEntity  v) {
        type.setValue(v);
    }
    // upper: Status status
    private MutableLiveData< String > status = new MutableLiveData<>();
    public MutableLiveData< String > getStatus() {
        return status;
    }
    public void setStatus( String  v) {
        status.setValue(v);
    }
    // upper: Flag flag
    private MutableLiveData< String > flag = new MutableLiveData<>();
    public MutableLiveData< String > getFlag() {
        return flag;
    }
    public void setFlag( String  v) {
        flag.setValue(v);
    }
    // upper: OfficialName officialName
    private MutableLiveData< String > officialName = new MutableLiveData<>();
    public MutableLiveData< String > getOfficialName() {
        return officialName;
    }
    public void setOfficialName( String  v) {
        officialName.setValue(v);
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
    // upper: Code code
    private MutableLiveData<String> codeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCodeMsg() {
        return codeMsg;
    }
    public void setCodeMsg(String v) {
        codeMsg.setValue(v);
    }
    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTypeMsg() {
        return typeMsg;
    }
    public void setTypeMsg(String v) {
        typeMsg.setValue(v);
    }
    // upper: Status status
    private MutableLiveData<String> statusMsg = new MutableLiveData<>();
    public MutableLiveData<String> getStatusMsg() {
        return statusMsg;
    }
    public void setStatusMsg(String v) {
        statusMsg.setValue(v);
    }
    // upper: Flag flag
    private MutableLiveData<String> flagMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFlagMsg() {
        return flagMsg;
    }
    public void setFlagMsg(String v) {
        flagMsg.setValue(v);
    }
    // upper: OfficialName officialName
    private MutableLiveData<String> officialNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getOfficialNameMsg() {
        return officialNameMsg;
    }
    public void setOfficialNameMsg(String v) {
        officialNameMsg.setValue(v);
    }
  }
}