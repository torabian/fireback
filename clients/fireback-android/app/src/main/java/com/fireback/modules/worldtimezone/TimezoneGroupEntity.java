package com.fireback.modules.worldtimezone;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class TimezoneGroupUtcItems extends JsonSerializable {
    public String name;
  public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
  }
}
public class TimezoneGroupEntity extends JsonSerializable {
    public String value;
    public String abbr;
    public int offset;
    public Boolean isdst;
    public String text;
    public TimezoneGroupUtcItems[] utcItems;
    public static class VM extends ViewModel {
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Abbr abbr
    private MutableLiveData< String > abbr = new MutableLiveData<>();
    public MutableLiveData< String > getAbbr() {
        return abbr;
    }
    public void setAbbr( String  v) {
        abbr.setValue(v);
    }
    // upper: Offset offset
    private MutableLiveData< Integer > offset = new MutableLiveData<>();
    public MutableLiveData< Integer > getOffset() {
        return offset;
    }
    public void setOffset( Integer  v) {
        offset.setValue(v);
    }
    // upper: Isdst isdst
    private MutableLiveData< Boolean > isdst = new MutableLiveData<>();
    public MutableLiveData< Boolean > getIsdst() {
        return isdst;
    }
    public void setIsdst( Boolean  v) {
        isdst.setValue(v);
    }
    // upper: Text text
    private MutableLiveData< String > text = new MutableLiveData<>();
    public MutableLiveData< String > getText() {
        return text;
    }
    public void setText( String  v) {
        text.setValue(v);
    }
    // upper: UtcItems utcItems
    private MutableLiveData< TimezoneGroupUtcItems[] > utcItems = new MutableLiveData<>();
    public MutableLiveData< TimezoneGroupUtcItems[] > getUtcItems() {
        return utcItems;
    }
    public void setUtcItems( TimezoneGroupUtcItems[]  v) {
        utcItems.setValue(v);
    }
    }
}