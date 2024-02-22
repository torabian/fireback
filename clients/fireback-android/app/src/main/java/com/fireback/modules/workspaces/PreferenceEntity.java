package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PreferenceEntity extends JsonSerializable {
    public String timezone;
    public static class VM extends ViewModel {
    // upper: Timezone timezone
    private MutableLiveData< String > timezone = new MutableLiveData<>();
    public MutableLiveData< String > getTimezone() {
        return timezone;
    }
    public void setTimezone( String  v) {
        timezone.setValue(v);
    }
    }
}