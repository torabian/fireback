package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class PreferenceEntity extends JsonSerializable {
    public String timezone;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Timezone timezone
    private MutableLiveData< String > timezone = new MutableLiveData<>();
    public MutableLiveData< String > getTimezone() {
        return timezone;
    }
    public void setTimezone( String  v) {
        timezone.setValue(v);
    }
    // Handling error message for each field
    // upper: Timezone timezone
    private MutableLiveData<String> timezoneMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTimezoneMsg() {
        return timezoneMsg;
    }
    public void setTimezoneMsg(String v) {
        timezoneMsg.setValue(v);
    }
  }
}