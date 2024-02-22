package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class UserEntity extends JsonSerializable {
    public PersonEntity person;
    public static class VM extends ViewModel {
    // upper: Person person
    private MutableLiveData< PersonEntity > person = new MutableLiveData<>();
    public MutableLiveData< PersonEntity > getPerson() {
        return person;
    }
    public void setPerson( PersonEntity  v) {
        person.setValue(v);
    }
    }
}