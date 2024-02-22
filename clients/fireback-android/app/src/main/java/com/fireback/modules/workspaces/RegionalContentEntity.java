package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class RegionalContentEntity extends JsonSerializable {
    public String content;
    public String region;
    public String title;
    public String languageId;
    public String keyGroup;
    public static class VM extends ViewModel {
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
    }
}