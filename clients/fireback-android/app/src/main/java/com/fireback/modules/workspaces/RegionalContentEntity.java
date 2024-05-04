package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class RegionalContentEntity extends JsonSerializable {
    public String content;
    public String region;
    public String title;
    public String languageId;
    public String keyGroup;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Content content
    private MutableLiveData< String > content = new MutableLiveData<>();
    public MutableLiveData< String > getContent() {
        return content;
    }
    public void setContent( String  v) {
        content.setValue(v);
    }
    // upper: Region region
    private MutableLiveData< String > region = new MutableLiveData<>();
    public MutableLiveData< String > getRegion() {
        return region;
    }
    public void setRegion( String  v) {
        region.setValue(v);
    }
    // upper: Title title
    private MutableLiveData< String > title = new MutableLiveData<>();
    public MutableLiveData< String > getTitle() {
        return title;
    }
    public void setTitle( String  v) {
        title.setValue(v);
    }
    // upper: LanguageId languageId
    private MutableLiveData< String > languageId = new MutableLiveData<>();
    public MutableLiveData< String > getLanguageId() {
        return languageId;
    }
    public void setLanguageId( String  v) {
        languageId.setValue(v);
    }
    // upper: KeyGroup keyGroup
    private MutableLiveData< String > keyGroup = new MutableLiveData<>();
    public MutableLiveData< String > getKeyGroup() {
        return keyGroup;
    }
    public void setKeyGroup( String  v) {
        keyGroup.setValue(v);
    }
    // Handling error message for each field
    // upper: Content content
    private MutableLiveData<String> contentMsg = new MutableLiveData<>();
    public MutableLiveData<String> getContentMsg() {
        return contentMsg;
    }
    public void setContentMsg(String v) {
        contentMsg.setValue(v);
    }
    // upper: Region region
    private MutableLiveData<String> regionMsg = new MutableLiveData<>();
    public MutableLiveData<String> getRegionMsg() {
        return regionMsg;
    }
    public void setRegionMsg(String v) {
        regionMsg.setValue(v);
    }
    // upper: Title title
    private MutableLiveData<String> titleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTitleMsg() {
        return titleMsg;
    }
    public void setTitleMsg(String v) {
        titleMsg.setValue(v);
    }
    // upper: LanguageId languageId
    private MutableLiveData<String> languageIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLanguageIdMsg() {
        return languageIdMsg;
    }
    public void setLanguageIdMsg(String v) {
        languageIdMsg.setValue(v);
    }
    // upper: KeyGroup keyGroup
    private MutableLiveData<String> keyGroupMsg = new MutableLiveData<>();
    public MutableLiveData<String> getKeyGroupMsg() {
        return keyGroupMsg;
    }
    public void setKeyGroupMsg(String v) {
        keyGroupMsg.setValue(v);
    }
  }
}